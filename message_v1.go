package main

import (
	"bufio"
	"fmt"
	"motd"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Your Greeting: ")
	phrase, _ := reader.ReadString('\n')
	phrase = strings.TrimSpace(phrase)

	fmt.Print("Your Name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	m := motd.Greeting(name, phrase)
	fmt.Println(m)
}
