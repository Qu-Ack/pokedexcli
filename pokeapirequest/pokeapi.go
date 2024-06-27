package pokeapirequest

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
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

var location = conf{
		start: 1,
		end: 1,
}



func apiRequest(url string) ([]byte, error) {
		res , err := http.Get(url)
		if err != nil {
				return nil, errors.New("unexpected fetch error")
		}
		body , err := io.ReadAll(res.Body)
		res.Body.Close()
		
		

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


func PokeLocationGet() error {
		// can use go routines to maybe speed it up ? ??? 
		for i:= location.start ; i <= location.end; i++ {

				if i == 21 {
						continue
				}
				resLocation, err := apiRequest(fmt.Sprintf("https://pokeapi.co/api/v2/location/%v", i))
				if err != nil {
						return err
				}
				result := pokelocation{}
				jerr := jsonConvert(resLocation, &result)
				if jerr != nil {
						return err
				}
				fmt.Println(result.Name)

		}
		location.start = location.start + 10
		location.end = location.start + 10
		return nil


		
}



func PokePrevLocationGet() error {
		if location.start == 1 {
				return errors.New("can't go back")
		}

		location.start = location.start - 10
		location.end = location.end - 10

		for i:= location.start ; i <= location.end; i++ {
				if i == 21 {
						continue
				}
				resLocation, err := apiRequest(fmt.Sprintf("https://pokeapi.co/api/v2/location/%v", i))
				if err != nil {
						return err
				}
				result := pokelocation{}
				jerr := jsonConvert(resLocation, &result)
				if jerr != nil {
						return err
				}
				fmt.Println(result.Name)

		}
		return nil

}
