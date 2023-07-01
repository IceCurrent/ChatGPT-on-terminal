# ChatGPT-on-terminal
A simple project in Golang that brings chat GPT right onto the terminal.

## How to use 
* Clone this repository and navigate to its root directory

### For MacOS
```
echo "export GOBIN=~/go/bin/" >> ~/.bashrc
echo "export PATH=$PATH:$GOBIN" >> ~/.bashrc
```
* Source the .bashrc, then build and run main.go
  ```
  source ~/.bashrc
  go run main.go
  ```
* You're all set, start using the tool by typing
  ```
  ask-gpt askgpt <enter password> <enter a prompt>
  ``` 

