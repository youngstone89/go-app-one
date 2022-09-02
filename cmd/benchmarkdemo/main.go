package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	body, err := GetProducts()
	if err != nil {
		return
	}
	fmt.Printf("body: %s", string(body))
}

func GetProducts() ([]byte, error) {

	response, error := http.Get("https://reqres.in/api/products")
	if error != nil {
		fmt.Println(error)
		return nil, error
	}
	defer response.Body.Close()
	responseBody, _ := ioutil.ReadAll(response.Body)
	if error != nil {
		fmt.Println(error)
		return nil, error
	}
	return responseBody, nil

}
