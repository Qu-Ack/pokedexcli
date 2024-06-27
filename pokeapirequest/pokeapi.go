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

func PokeLocationGet()  {
		location := pokelocation{}
		data, _ := apiRequest("https://pokeapi.co/api/v2/location/1")
		err := json.Unmarshal(data, &location)
		if err != nil {
				fmt.Println("something happened")
		}
		fmt.Println(location)
		
}
