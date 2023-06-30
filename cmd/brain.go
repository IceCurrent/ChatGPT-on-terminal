/*
Copyright © 2023 SHREYANSH <shreyansh.yashi@gmail.com>
*/
package cmd

import (
	"context"
	"fmt"
	"log"

	openai "github.com/sashabaranov/go-openai"
	//"github.com/joho/godotenv"
)

func getResponse(API_KEY string, input string) {
	//var err = godotenv.Load()

	// if err != nil {
	// 	log.Fatal("Error in loading .env file")
	// }

	// var input string
	// fmt.Scan("Ask Chat-GPT: %v", input)

	//var API_KEY = os.Getenv("API_KEY")

	// fmt.Println(API_KEY)
	// fmt.Println(input)

	var client = openai.NewClient(API_KEY)
	var response, errorGettingMsg = client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT4,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: input,
				},
			},
		},
	)

	if errorGettingMsg != nil {
		log.Fatal("Chat Completion error: ", errorGettingMsg)
	}

	fmt.Println(response.Choices[0].Message.Content)

}
