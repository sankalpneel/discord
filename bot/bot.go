package bot

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/sankalpneel/discord/config"
	"github.com/sankalpneel/discord/temperature"
)

var BotID string
var goBot *discordgo.Session

var MSG string

func Start() {
	MSG = "HI I am a Bot. I can tell you the the temperature. To get the temperature msg me in below formart:\n**temp palce_name**\n\n-Developed by Sankalp Kumar."
	goBot, err := discordgo.New("Bot " + config.Token)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	u, err := goBot.User("@me")
	if err != nil {
		fmt.Println(err.Error())
	}

	BotID = u.ID

	goBot.AddHandler(messageHandler)

	err = goBot.Open()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Bot is Running!")
}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == BotID {
		return
	}

	msg := strings.Split(m.Content, " ")
	//fmt.Println(msg)
	if msg[0] == "ping" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "pong")
	} else if msg[0] == "temp" {
		_, err := s.ChannelMessageSend(m.ChannelID, "Temperature of "+msg[1]+" is = "+temperature.Query(msg[1]))
		if err != nil {
			s.ChannelMessageSend(m.ChannelID, "Oops Something went wrong...Try again!!")
		}
	} else {
		_, _ = s.ChannelMessageSend(m.ChannelID, MSG)
	}

}
