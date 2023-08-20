package main

import (
	"encoding/json"
	"fmt"

	welcome "example.com/greetings/wel"
)

func main() {
	fmt.Println(welcome.Welcome("szx"))

	R1 := "router1"
	R2 := "router2"

	net1 := make([]string, 0)

	net1 = append(net1, R1)
	net1 = append(net1, R2)

	fmt.Println(net1)

	j, _ := json.Marshal(net1)
	fmt.Println(string(j))
	// log.SetPrefix("greetings: ")
	// log.SetFlags(0)
	// mes, err := greetings.Hello("szx")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(mes)

	// msgs, err := greetings.Hellos([]string{"szx", "xiaoming", "xiaolei"})
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(msgs)

}
