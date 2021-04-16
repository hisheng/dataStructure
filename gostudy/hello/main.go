/**
@author 张海生<zhanghaisheng@qimao.com>
@dateTime   : 2021/4/16 9:44 上午
@Desc
*/
package main

import (
	"fmt"
	"han.com/age"
	"hisheng.com/greetings"
	"log"
)

func main() {
	log.SetPrefix("greetings:")
	log.SetFlags(0)
	message, err := greetings.Hello("hisheng")
	if err != nil {
		log.Println(err)
	}
	fmt.Println(message)

	names := []string{"Gladys", "Samantha", "Darrin"}
	// Request greeting messages for the names.
	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}
	// If no error was returned, print the returned map of
	// messages to the console.
	fmt.Println(messages)

	age := age.Age(11)
	fmt.Println(age)
}
