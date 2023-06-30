# ChatGPT-on-terminal
A simple project in Golang that brings chat GPT right onto the terminal.

## Requirements
* You should have go installed on your system (preferably go 1.20)
* install go-client for Open AI
  ```
  go get github.com/sashabaranov/go-openai
  ```
* install cobra
  ```
  go get -u github.com/spf13/cobra@latest
  ```

## How to use 
* Clone this repository and navigate to it's root directory
* Build and run main.go
  ```
  go run main.go
  ```
* You're all set, start using the tool by typing
  ```
  ask-gpt askgpt <enter password> <enter a prompt>
  ``` 

