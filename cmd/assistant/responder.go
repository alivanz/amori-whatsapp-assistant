package main

import (
	"fmt"
	"strings"
)

// NOTE: in this file, I'm thinking to implement the text processor (AI-powered) here.
// we can use things like OpenAI.
// also, I found https://github.com/go-nlp repositories, but some repos look "old" already.

func GenerateResponse(context *Context, question string) string {
	question = strings.ToLower(question)

	// check for keywords
	if strings.Contains(question, "relationship") {
		for _, msg := range context.Messages {
			if strings.Contains(strings.ToLower(msg.Content), "relationship") {
				return fmt.Sprintf("Regarding your relationship, here's what was mentioned: \"%s\" - %s", msg.Content, msg.Sender)
			}
		}
		return "I couldn't find any messages related to your relationship."
	} else if strings.Contains(question, "first message") {
		if len(context.Messages) > 0 {
			msg := context.Messages[0]
			return fmt.Sprintf("The first message was: \"%s\" - %s", msg.Content, msg.Sender)
		}
		return "There are no messages in the conversation."
	} else if strings.Contains(question, "last message") {
		if len(context.Messages) > 0 {
			msg := context.Messages[len(context.Messages)-1]
			return fmt.Sprintf("The last message was: \"%s\" - %s", msg.Content, msg.Sender)
		}
		return "There are no messages in the conversation."
	}
	return "I don't have a specific answer for that, for now..."
}

func GenerateResponseWithScreenshot(context *Context, question, screenshot string) string {
	// can be expanded to analyze the context, screenshot, and generate more intelligent responses.
	return fmt.Sprintf("This is a response to your question: %s (with screenshot %s)", question, screenshot)
}
