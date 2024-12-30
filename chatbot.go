package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	responses := map[string][]string{
		"hello":   {"Hi there!", "Hello!", "Hey! How can I help you?"},
		"how are you": {"I'm just a program, but I'm feeling great!", "Doing well, thanks for asking!", "I'm here to assist you."},
		"help":    {"Sure, what do you need help with?", "I'm here to assist you. Ask me anything!", "How can I be of service?"},
		"bye":     {"Goodbye!", "See you later!", "Take care!"},
		"default": {"I'm not sure I understand.", "Can you rephrase that?", "Interesting... tell me more."},
	}

	fmt.Println("Basic Chatbot Simulator")
	fmt.Println("========================")
	fmt.Println("Start chatting! Type 'bye' to end the conversation.\n")

	for {
		fmt.Print("You: ")
		var input string
		fmt.Scanln(&input)

		lowerInput := strings.ToLower(input)
		if lowerInput == "bye" {
			fmt.Println("Chatbot: Goodbye! Have a nice day!")
			break
		}

		// Find a matching response
		matched := false
		for key, replies := range responses {
			if strings.Contains(lowerInput, key) {
				randomIndex := rand.Intn(len(replies))
				fmt.Printf("Chatbot: %s\n", replies[randomIndex])
				matched = true
				break
			}
		}

		// Default response if no keyword matches
		if !matched {
			defaultReplies := responses["default"]
			randomIndex := rand.Intn(len(defaultReplies))
			fmt.Printf("Chatbot: %s\n", defaultReplies[randomIndex])
		}
	}
}
