package pokeapirequest

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"reflect"
)


type conf struct {
		start int
		end int
}


func ApiRequest(url string) error {
		res , err := http.Get(url)
		if err != nil {
				return errors.New("unexpected fetch error")
		}
		body , err := io.ReadAll(res.Body)
		res.Body.Close()

		if res.StatusCode > 299 {
				return errors.New(fmt.Sprintf("response failed with status code : %v ", res.StatusCode))
		}

		if err != nil {
				return errors.New("error while reading the body")
		}
		fmt.Println(reflect.TypeOf(body))
		fmt.Println(body)
		return nil


}

func PokeGet()  {

}
