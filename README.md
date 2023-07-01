# ChatGPT-on-terminal
A simple project in Golang that brings chat GPT right onto the terminal.

## How to use 
* You should have [go compiler](https://go.dev/doc/install) on your system
* Clone this repository and navigate to its root directory

### For Linux and MacOS
```
echo "export GOBIN=~/go/bin/" >> ~/.bashrc
echo "export PATH=$PATH:$GOBIN" >> ~/.bashrc
```
* Source the .bashrc, then build and run main.go
  ```
  source ~/.bashrc
  go run main.go
  ```

### For Windows
```
set GOBIN=%USERPROFILE%\go\bin
set PATH=%PATH%;%GOBIN%
```
* Source the .profile, then build and run main.go
  ```
  source ~/.profile
  go run main.go
  ```


**You're all set** 🚀, start using the tool by typing 
  ```
  ask-gpt askgpt <enter password> <enter a prompt>
  ``` 

