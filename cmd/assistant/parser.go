package main

import (
	"bufio"
	"os"
	"strings"
)

type Message struct {
	Sender  string
	Content string
}

func ReadConversationFile(filePath string) ([]Message, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var messages []Message
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		parts := strings.SplitN(line, ":", 2)
		if len(parts) != 2 {
			continue
		}

		messages = append(messages, Message{
			Sender:  strings.TrimSpace(parts[0]),
			Content: strings.TrimSpace(parts[1]),
		})
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return messages, nil
}
