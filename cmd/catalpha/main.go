package main

import (
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"strings"

	"github.com/aimof/catalpha"
	"github.com/bwmarrin/discordgo"
)

var (
	token string
	name  = "天飛"
)

func main() {
	token = readToken()

	session, err := discordgo.New()
	if err != nil {
		log.Fatalln(err)
	}
	session.Token = token

	session.AddHandler(handler)

	err = session.Open()
	if err != nil {
		log.Fatalln(err)
	}

	ch := make(chan os.Signal, 1)
	signal.Notify(ch)
	_ = <-ch
	err = session.Close()
	if err != nil {
		log.Fatalln(err)
	}
}

func handler(s *discordgo.Session, m *discordgo.MessageCreate) {
	c, err := s.State.Channel(m.ChannelID)
	if err != nil {
		log.Println(err)
	}

	var msg string
	if m.Author.Username != name {
		msg = catalpha.Catalpha(m)
	}

	if msg != "" {
		_, err = s.ChannelMessageSend(c.ID, msg)
		if err != nil {
			log.Println(err)
		}
	}
}

func readToken() string {
	file, err := os.Open("token.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	b, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatalln(err)
	}

	token = strings.TrimSpace(string(b))

	return token
}
