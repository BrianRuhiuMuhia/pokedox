package main

import (
	"encoding/json"
	"net/http"
)

func fetchData(url string) (Response, error) {
	resp, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Response{}, err
	}
	client := http.Client{}
	data, err := client.Do(resp)
	if err != nil {
		return Response{}, err
	}
	defer data.Body.Close()
	var responseData Response
	decoder := json.NewDecoder(data.Body)
	err = decoder.Decode(&responseData)
	if err != nil {
		return Response{}, err
	}
	return responseData, nil

}
