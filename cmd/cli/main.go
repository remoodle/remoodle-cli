package main

import (
	"os"

	"github.com/ocean-planet/ReMoodle/internal/app/commands/command"
	"github.com/ocean-planet/ReMoodle/internal/app/commands/deadlines"
	"github.com/ocean-planet/ReMoodle/internal/app/commands/help"
	"github.com/ocean-planet/ReMoodle/internal/app/commands/login"
	"github.com/ocean-planet/ReMoodle/internal/app/commands/whoami"
	"github.com/ocean-planet/ReMoodle/internal/app/core"
)

func main() {
	commandRepository := command.NewCommandRepository()
	commandService := command.NewCommandService(commandRepository)
	myApp := core.NewApp(commandService)

	commandRepository.RegisterCommand("help", &help.HelpCommand{CommandService: commandService})
	commandRepository.RegisterCommand("login", &login.LoginCommand{CommandService: commandService})
	commandRepository.RegisterCommand("whoami", &whoami.WhoamiCommand{CommandService: commandService})
	commandRepository.RegisterCommand("deadlines", &deadlines.DeadlineCommand{CommandService: commandService})

	if err := myApp.Run(os.Args[1:]); err != nil {
		os.Exit(1)
	}
}
