package pokeapirequest

import (
		"encoding/json"
		"errors"
		"fmt"
		"io"
		"net/http"
		"sync"
		"github.com/Qu-Ack/pokedexcli/pokecache"

)



type conf struct {
		start int
		end int
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
