package pokeapirequest

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"sync"

	"github.com/Qu-Ack/pokedexcli/pokecache"
)



type conf struct {
		start int
		end int
}

type Pokestats struct {
		Name string `json:"name"`
		BaseExperience int `json:"base_experience"`
		Height int `json:"height"`
		Stats []struct {
				BaseStat int `json:"base_stat"`
				Effort int `json:"effort"`
		} `json:"stats"`
		Weight int `json:"weight"`
}



type pokelocation struct {
		Name  string `json:"name"`
		Region struct {
				Name string `json:"name"`
		} `json:"region"`

}

type pokename struct {
		PokemonEncounters []struct {
				Pokemon struct {
						Name string `json:"name"`
				} `json:"pokemon"`
		} `json:"pokemon_encounters"`
}

var Pokedex = make(map[string]Pokestats)

var location = conf{
		start: -9,
		end: 0,
}



func apiRequest(url string) ([]byte, error) {
		res , err := http.Get(url)
		if err != nil {
				return nil, errors.New("unexpected fetch error")
		}
		body , err := io.ReadAll(res.Body)
		res.Body.Close()

		if res.StatusCode == 404 {
				return nil, errors.New("404")
		}

		if res.StatusCode > 299 {
				return nil,errors.New(fmt.Sprintf("response failed with status code : %v ", res.StatusCode))
		}

		if err != nil {
				return nil,errors.New("error while reading the body")
		}
		return body,nil


}

func jsonConvert[D any](data []byte, target *D) error {
		err:= json.Unmarshal(data, target)
		return err
}

func GetPokeStats(pokemon string) Pokestats {
		return Pokedex[pokemon]
}


func PokePokemonGet(city string) error {
		value , ok := pokecache.Cache[city]
		mu := &sync.Mutex{}
		if ok {
				pokemon := pokename{}
				err:= jsonConvert(value.Val, &pokemon)
				if err != nil {
						return err 
				}
				for _, elem := range pokemon.PokemonEncounters {
						fmt.Printf(" - %v\n", elem.Pokemon.Name)
				}
				return nil
		}  else {
				response , err  := apiRequest(fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%v", city))
				if err != nil {
						if err.Error() == "404" {
								return errors.New("no such area exists")
						} else {
								return err
						}
				}
				pokecache.Add(city, response, mu)
				pokemon := pokename{}
				jerr := jsonConvert(response, &pokemon)	
				if jerr != nil {
						return jerr
				}
				for _, elem := range pokemon.PokemonEncounters {
						fmt.Printf(" - %v\n", elem.Pokemon.Name)
				}
				return nil
		}
}


func canCatch(baseexp int) bool {
		chance := rand.Intn(baseexp)
		if chance < 50 {
				return true
		}
		return false
}


func PokeCatch(pokemon string) error {
		response , err := apiRequest(fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%v", pokemon))
		if err != nil {
				return err
		} 
		pokestats := Pokestats{}
		jerr := jsonConvert(response, &pokestats)
		if jerr != nil {
				return jerr
		}
		value , ok := Pokedex[pokestats.Name]
		if ok {
				fmt.Printf("you have already caught %v pokemon\n", value.Name)
		} else {
				catchChance := canCatch(pokestats.BaseExperience)
				if catchChance {
						Pokedex[pokestats.Name] = pokestats
						fmt.Printf("%v was caught\n" , pokestats.Name)
				} else {
						fmt.Printf("%v escaped\n", pokestats.Name)
				}
				return nil
		}
		return nil

}




func PokeLocationGet() error {
		var wg sync.WaitGroup
		errChan := make(chan error)
		pokecache.ReapLoop(30000000000)
		mu := &sync.Mutex{}
		location.start = location.start + 10
		location.end = location.end + 10
		for i:= location.start ; i <= location.end; i++ {

				if i == 21 {
						continue
				}
				cacheEntry , ok := pokecache.Cache[string(i)]
				if ok {
						result := pokelocation{}
						err := jsonConvert(cacheEntry.Val, &result)
						if err != nil {
								return err
						}

						fmt.Println(result.Name)
				} else {
						wg.Add(1)
						go func(i int) {
								defer wg.Done()
								resLocation, err :=  apiRequest(fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%v", i))
								if err != nil {
										errChan <- err	
										return
								}
								pokecache.Add(string(i), resLocation, mu)
								result := pokelocation{}
								jerr := jsonConvert(resLocation, &result)
								if jerr != nil {
										errChan <- err
										return
								}
								fmt.Println(result.Name)
						}(i)
				}



		}
		wg.Wait()
		close(errChan)

		for err := range errChan {
				if err != nil {
						return err
				}
		}



		return nil



}



func PokePrevLocationGet() error {
		if location.start <= 1 {
				return errors.New("can't go back")
		}
		mu := &sync.Mutex{}

		var wg sync.WaitGroup
		errChan := make(chan error)

		location.start = location.start - 10
		location.end = location.end - 10

		for i:= location.start ; i <= location.end; i++ {
				if i == 21 {
						continue
				}
				cacheEntry , ok := pokecache.Cache[string(i)]
				if ok {
						result := pokelocation{}
						err := jsonConvert(cacheEntry.Val , &result) 
						if err != nil {
								return err
						}
						fmt.Println(result.Name)

				} else {

						wg.Add(1)
						go func(i int) {
								defer wg.Done()
								resLocation, err := apiRequest(fmt.Sprintf("https://pokeapi.co/api/v2/location/%v", i))
								if err != nil {
										errChan <- err
								}
								pokecache.Add(string(i), resLocation, mu)
								result := pokelocation{}
								jerr := jsonConvert(resLocation, &result)
								if jerr != nil {
										errChan <- err
								}
								fmt.Println(result.Name)
						}(i)
				}

		}
		wg.Wait()
		close(errChan)

		for err := range errChan {
				if err != nil {
						return err
				}
		}
		return nil

}
