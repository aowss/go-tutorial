package main

import (
    "fmt"
    "example.com/greetings"
)

func main() {
    message := greetings.Hello("Aowss")
    fmt.Println(message)
}