package utils

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

//ReadResponse uses ioutil to read an http response and return the body as a string
func ReadResponse(response *http.Response) (string, error) {
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println("Error: ", err)
	}
	return string(body), err
}

//StructToJson converts a struct to a json string
func StructToJson(data interface{}) string {
	json, err := json.Marshal(data)
	if err != nil {
		log.Println("Error: ", err)
	}
	return string(json)
}
