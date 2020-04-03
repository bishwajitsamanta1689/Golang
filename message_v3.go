package main

import (
	"bufio"
	"fmt"
	"log"
	"motd"
	"os"
	"strings"
)

func main() {

	//f, err := os.OpenFile("app.txt", os.O_WRONLY|os.O_CREATE, 0644)
	f, err := os.OpenFile("app.txt", os.O_WRONLY, 0644)
	lf, _ := os.OpenFile("info.log", os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		log.SetOutput(lf)
		log.Fatalln("File Doesn't Exist !!")
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Your Greeting: ")
	phrase, _ := reader.ReadString('\n')
	phrase = strings.TrimSpace(phrase)

	fmt.Print("Your Name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	m := motd.Greeting(name, phrase)
	//err := ioutil.WriteFile("/home/cloud_user/go/src/test", []byte(m), 0644)
	_, err = f.Write([]byte(m))

	if err != nil {
		fmt.Println("Unable to Write")
		os.Exit(1)
	}
}
