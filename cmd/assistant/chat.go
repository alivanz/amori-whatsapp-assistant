package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/urfave/cli/v2"
)

func chat(cli *cli.Context) error {
	// intro
	AssistantWrite("Hi, I'm Christie your dating assistant. Let's start by uploading your WhatsApp conversation.")
	AssistantWrite("Please type the name of the conversation file (The conversation file should be in the \"conversations\" folder):")

	// get the conversation file
	convFile := GetUserInput("> ")
	convPath := filepath.Join(conversationsDir, convFile)
	msgs, err := ReadConversationFile(convPath)
	if err != nil {
		return err
	}

	// process
	var (
		context Context
	)
	AssistantWrite("Thank you for providing the conversation file. I'm currently processing the file. Please wait...")
	for _, msg := range msgs {
		context.AddMessage(msg)
	}
	// simulate process
	time.Sleep(3 * time.Second)
	AssistantWrite("Your conversation file is processed successfully. Now you can ask questions.")

	// interaction loop
	for {
		AssistantWrite(`Please choose from the following options:
	1. Ask a question
	2. Ask a question with an attached screenshot
	99. Quit`)
		option := GetUserInput("> ")
		switch option {
		default:
			AssistantWrite("Invalid option, please try again.")
		case "1":
			// Handle question without screenshot
			question := GetUserInput("Type your question:\n> ")
			response := GenerateResponse(&context, question)
			AssistantWrite(response)
			context.AddQuestionAndResponse(question, response)
		case "2":
			// Handle question with screenshot
			screenshot := GetUserInput("Please input the screenshot file name (The screenshot should be in the \"screenshots\" folder):\n> ")
			screenshotPath := filepath.Join(screenshotsDir, screenshot)
			context.SetScreenshot(screenshot)
			question := GetUserInput("Type your question:\n> ")
			response := GenerateResponseWithScreenshot(&context, question, screenshotPath)
			AssistantWrite(response)
			context.AddQuestionAndResponse(question, response)
		case "99":
			AssistantWrite("Good bye, have a good day!")
			return nil
		}
	}

}

func AssistantWrite(response string) {
	fmt.Printf("Assistant: %s\n", response)
}

func GetUserInput(prompt string) string {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	fmt.Println("") // to make it slightly more readable
	return input
}
