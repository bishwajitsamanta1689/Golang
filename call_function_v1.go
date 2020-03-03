package main

import "fmt"

func main() {
	message := greeting("Bishwajit", "hello")
	fmt.Println(message)
}

func greeting(name, message string) string {
	return fmt.Sprintf("%s", "%s", message, name)
}
