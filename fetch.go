package main

import (
	"encoding/xml"
	"net/http"
)

func fetchRSSFeed(url string) (RSSFeed, error) {
	resp, err := http.Get(url)
	if err != nil {
		return RSSFeed{}, err
	}
	params := RSSFeed{}

	if err = xml.NewDecoder(resp.Body).Decode(&params); err != nil {
		return RSSFeed{}, err
	}
	return params, nil
}
