package handlers

import (
	"io"
	"net/http"
)

// Contains Helpers to fetch robots.txt file
// parse it and read Site_map url
// Robots.txt file is followed only when the necessary flag is
// passed in the web crawler config

func fetchRobotsTxt(domain string) (string, error) {
	url := domain + "/robots.txt"
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
