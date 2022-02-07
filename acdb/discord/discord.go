package discord

import (
	"aircraftTracker/config"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/bwmarrin/discordgo"
)

var botId string
var Dbot *discordgo.Session
var updateHook string
var addRegChannel chan<- string

func Start(myConfig config.Config, addReg chan<- string) {

	log.Println("starting discord bot")

	addRegChannel = addReg

	updateHook = myConfig.DiscordWebHook

	Dbot, err := discordgo.New("Bot " + myConfig.DiscordToken)
	if err != nil {
		log.Println(err.Error())
		return
	}

	u, err := Dbot.User("@me")
	if err != nil {
		log.Println(err.Error())
		return
	}

	botId = u.ID
	Dbot.AddHandler(messageHandler)
	Dbot.AddHandler(disconnectHandler)
	Dbot.AddHandler(readyHandler)

	err = Dbot.Open()
	if err != nil {
		log.Println(err.Error())
		return
	}

}

func SendMessage(msg string) {

	//curl -i -H "Accept: application/json"
	// -H "Content-Type:application/json" -X POST
	// --data "{\"content\": \"Posted Via Command line\"}"
	// discord-webhook-link

	rb, err := json.Marshal(map[string]string{"content": msg})
	if err != nil {
		log.Panicf("error, can not marshal webhook message")
	}
	http.Post(updateHook, "application/json", bytes.NewBuffer(rb))

	//Dbot.ChannelMessageSend()
	// //Dbot..find("")
	// c, err := Dbot.Channel(botId)
	// if err != nil {
	// 	log.Printf("can not create channel %v", err)
	// }
	//log.Printf("send message from bot to discords")
	//Dbot.ChannelMessageSend(readyM.SessionID, "lucki lucki")
}

func readyHandler(s *discordgo.Session, m *discordgo.Ready) {
	log.Printf("discord bot is up and running")

}

func disconnectHandler(s *discordgo.Session, m *discordgo.Disconnect) {
	log.Printf("discord bot is disconnected")
}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {

	// ignore all messages created by the bot itself
	if m.Author.ID == botId {
		return
	}

	if m.Content == "!ping" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "pong")
	}

	if m.Content == "!help" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "some fancy help message")
	}

	if strings.HasPrefix(m.Content, "!add") {
		rc := strings.Split(m.Content, " ")
		if len(rc) > 2 {
			_, _ = s.ChannelMessageSend(m.ChannelID, "wrong add command format, you can only add a single aircraft")
			return
		}
		if len(rc[1]) > 8 {
			_, _ = s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("received registration '%v' is not valid", rc[1]))
			return
		}

		//err := observer.Add(rc[1])
		//if err != nil {
		//	_, _ = s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("aircraft registration '%v' not found in database (%v)", rc[1], err))
		//	return
		//}
		_, _ = s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("add aircraft '%v' to observation list", rc[1]))
		addRegChannel <- rc[1]
	}

	////If we message ping to our bot in our discord it will return us pong .
	//if m.Content == "ping" {
	//	_, _ = s.ChannelMessageSend(m.ChannelID, "pong")
	//}
}
