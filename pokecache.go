package main

import (
	"errors"
	"fmt"
)

type Cache struct {
	data []map[string]Response
	urls []map[int]string
	len  int
}

func (c *Cache) get_data(url string) (Response, bool) {
	for _, data := range c.data {
		if response, exists := data[url]; exists {
			return response, true
		}
	}
	return Response{}, false
}
func (c *Cache) set_data(url string, res Response) (Response, bool) {
	if c.len <= 0 {
		c.data = append(c.data, map[string]Response{url: res})
		c.urls = append(c.urls, map[int]string{c.len + 1: url})

	} else {
		for _, data := range c.data {

			if _, exists := data[url]; !exists {
				data[url] = res
				fmt.Println(url)
				c.urls = append(c.urls, map[int]string{c.len + 1: url})

				return res, true
			} else {
				fmt.Println("Already Exists")
				return Response{}, false
			}
		}
	}

	return Response{}, false
}
func create_cache() Cache {
	cache := Cache{data: nil, urls: nil, len: -1}
	return cache
}
func (c *Cache) get_data_by_num(num int) (Response, error) {
	var search_url string
	for index, url := range c.urls {
		if num == index {
			search_url = url[num]
		}
	}
	if search_url != "" {
		res, _ := c.get_data(search_url)
		return res, nil
	}

	return Response{}, errors.New("Error")
}

var cache = create_cache()
