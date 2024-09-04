package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/urfave/cli/v2"
)

var (
	conversationsDir string
	screenshotsDir   string
)

func main() {
	godotenv.Load()
	app := &cli.App{
		Name: "whatsapp-assistant",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "conversations-dir",
				Aliases:     []string{"convs-dir"},
				EnvVars:     []string{"CONVERSATIONS_DIR"},
				Value:       "./conversations",
				Destination: &conversationsDir,
			},
			&cli.StringFlag{
				Name:        "screenshots-dir",
				Aliases:     []string{"sc-dir"},
				EnvVars:     []string{"SCREENSHOTS_DIR"},
				Value:       "./conversations",
				Destination: &screenshotsDir,
			},
		},
		Action: chat,
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
