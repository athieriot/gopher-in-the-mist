package main

import (
	"net/http"
	"io/ioutil"
	"fmt"
)

func GetBody() ([]byte, error) {
	client := &http.Client{}
	// Our friend the Underscore ;) // HL
	req, _ := http.NewRequest("GET", "http://api.com/", nil)

	// A Function can return more than one value // HL
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	// Defer the execution of a function after the enclosing block // HL
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	return body, nil
}

func main() {
	body, _ := GetBody()
	fmt.Println(string(body))
}
