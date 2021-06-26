package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string {
		"https://users.roblox.com",
		"https://games.roblox.com",
		"https://inventory.roblox.com",
		"https://avatar.roblox.com",
		"https://economy.roblox.com",
		"https://catalog.roblox.com",
		"https://groups.roblox.com",
		"https://trades.roblox.com",
		"https://api.roblox.com",
	}

	for _, link := range links {
		go checkLink(link)
	}
}

func checkLink(link string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "Is down")
		return
	}

	fmt.Println(link, "Is up")
}