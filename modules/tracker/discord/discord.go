package discord

import (
	"log"

	"github.com/bwmarrin/discordgo"
	"github.com/spf13/viper"
)

var dg *discordgo.Session
var webookid string
var webhooktoken string
var bottoken string

func init() {
	bottoken = viper.GetString("discord.bottoken")
	webookid = viper.GetString("discord.webhookid")
	webhooktoken = viper.GetString("discord.webhooktoken")
}

// send embeded message to discord channel
// spezified by webhook token and webhook id
// return the id of the send message
func SendEmbedMessageWithWebhook(title string,
	reference string,
	embeds []*discordgo.MessageEmbed) (id string, err error) {

	if dg == nil {
		createSession()
	}

	webhook, err := dg.WebhookWithToken(webookid, webhooktoken)
	if err != nil {
		return "", err
	}

	msgSend := discordgo.MessageSend{
		Content: title,
		Embeds:  embeds,
	}

	st, err := dg.ChannelMessageSendComplex(webhook.ChannelID, &msgSend)

	if err != nil {
		return "", err
	}

	return st.ID, nil
}

// create discord session and end with log fatal if it was not successful
func createSession() {
	log.Printf("starting discord session")

	if bottoken == "" {
		log.Fatalf("no discord bot token set in config, abort ...")
	}

	var err error
	dg, err = discordgo.New("Bot " + bottoken)
	if err != nil {
		log.Fatal("error creating discord session,", err)
	}
}
