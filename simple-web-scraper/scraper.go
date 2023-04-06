package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/gocolly/colly"
)

var names []string

func main() {
	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting: ", r.URL)
	})

	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong: ", err)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Page visited: ", r.Request.URL)
	})

	// print online players
	c.OnHTML("tr td a", func(e *colly.HTMLElement) {

		name := e.Text
		names = append(names, name)
		//fmt.Println(names)
	})

	c.OnScraped(func(r *colly.Response) {
		fmt.Println(r.Request.URL, " scraped!")
		for _, v := range names {
			if strings.Contains(v, "Carlos Corleone") {
				fmt.Println(v)
			}
			if strings.Contains(v, "Eddie Corleone") {
				fmt.Println(v)
			}
			if strings.Contains(v, "Jack Corleone") {
				fmt.Println(v)
			}
			if strings.Contains(v, "Glen Urra") {
				fmt.Println(v)
			}
			if strings.Contains(v, "Novak Dakovic") {
				fmt.Println(v)
			}
			if strings.Contains(v, "Xander Falcone") {
				fmt.Println(v)
			}
			if strings.Contains(v, "Joshua Falcone") {
				fmt.Println(v)
			}
			if strings.Contains(v, "Kobe Mamba") {
				fmt.Println(v)
			}
			if strings.Contains(v, "Levi Falcone") {
				fmt.Println(v)
			}
			if strings.Contains(v, "Narmy Falcone") {
				fmt.Println(v)
			}
			if strings.Contains(v, "Donnie Pazzini") {
				fmt.Println(v)
			}
			if strings.Contains(v, "Aidan Falcone") {
				fmt.Println(v)
			}
		}
	})

	c.Visit("https://sa-mp.im/server/players")

}
