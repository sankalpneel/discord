package main

import (
	"fmt"

	"github.com/sankalpneel/discord/bot"
	"github.com/sankalpneel/discord/config"
)

func main() {
	err := config.ReadConfig()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	bot.Start()
	<-make(chan struct{})
	return
}