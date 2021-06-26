package main

import (
	"fmt"
	"net/http"
	"time"
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
		"https://privatemessages.roblox.com",
		"https://api.roblox.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	// for i := 0; i < len(links); i++ {
	// 	go checkLink(<-c, c)
	// }

	// for {
	// 	go checkLink(<-c, c)
	// }

	for l := range c {
		time.Sleep(5 * time.Second)
		go checkLink(l, c)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "Is down")
		c <- link
		return
	}

	fmt.Println(link, "Is up")
	c <- link
}