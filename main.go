package main

import (
	tmpestbotserver "tmpest-bot-server/pkg/tmpest-bot-server"
)

func main() {
	server, error := tmpestbotserver.NewTmpestBotServer()
	if error != nil {
		return
	}

	server.Start()
}
