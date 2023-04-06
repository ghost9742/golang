package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/gocolly/colly"
)

// array for scrapped names
var names []string

var namesString string

func main() {

	sess, err := discordgo.New("Bot insert-private-key-here")
	if err != nil {
		log.Fatal(err)
	}

	sess.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		if m.Author.ID == s.State.User.ID {
			return
		}

		if m.Content == "!online" {

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
				namesString = strings.Join(names, " ")
				//fmt.Println(namesString)
				s.ChannelMessageSend(m.ChannelID, namesString)
				// reset names slice
				names = []string{}

			})

			c.Visit("https://sa-mp.im/server/players")
			//s.ChannelMessageSend(m.ChannelID, "world!")

		}
	})

	sess.Identify.Intents = discordgo.IntentsAllWithoutPrivileged

	err = sess.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer sess.Close()

	fmt.Println("The bot is online!")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}
