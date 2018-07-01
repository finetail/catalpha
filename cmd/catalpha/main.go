package main

import (
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"strings"

	"catalphafine"
	"github.com/bwmarrin/discordgo"
)

var (
	token string
	name  = "catalpha2"
)

func main() {
	token = readToken()

	session, err := discordgo.New()
	if err != nil {
		log.Printf("\t%v\n", err)
	}
	session.Token = token


	session.AddHandler(handler)

	err = session.Open()
	if err != nil {
		log.Printf("\t%v\n", err)
	}

	ch := make(chan os.Signal, 1)
	signal.Notify(ch)
	_ = <-ch
	err = session.Close()
	if err != nil {
		log.Printf("\t%v\n", err)
	}
}

func handler(s *discordgo.Session, m *discordgo.MessageCreate) {
	c, err := s.State.Channel(m.ChannelID)
	if err != nil {
		log.Printf("\t%v\n", err)
	}

	var msg string
	if m.Author.Username != name {
		log.Printf("\tMessage from %s: %s\n", m.Author.Username, m.Content)
		msg = catalpha.Catalpha(m)
	}

	if msg != "" {
		_, err = s.ChannelMessageSend(c.ID, msg)
		if err != nil {
			log.Printf("\t%v\n", err)
		}
		log.Printf("\tMessage sent: %s\n", msg)
	}
}

func readToken() string {
	file, err := os.Open("token.txt")
	if err != nil {
		log.Printf("\t%v\n", err)
	}
	defer file.Close()

	b, err := ioutil.ReadAll(file)
	if err != nil {
		log.Printf("\t%v\n", err)
	}

	token = strings.TrimSpace(string(b))

	return token
}
