package main

type Command struct {
	name        string
	instruction string
	callback    func()
}
type Response struct {
	Count    int     `json:"count"`
	Previous *string `json:"previous"`
	Next     *string `json:"next"`
	Results  []Data  `json:"results"`
}
type Data struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}
