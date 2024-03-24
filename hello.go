package main

import "fmt"

func main() {
	fmt.Println(Hello("world"))
}
func Hello(s string) string {
	return "Hello" + s
}
