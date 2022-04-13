package main

import (
	"log"
	"net/url"
	"time"

	"github.com/sonzqn/pact-consumer-go/src"
)

var token = time.Now().Format("2006-01-02T15:04")

func main() {
	u, _ := url.Parse("http://localhost:8088")
	client := &client.Client{
		BaseURL: u,
	}

	users, err := client.WithToken(token).GetUser(10)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(users)
}
