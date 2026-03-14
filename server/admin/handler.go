package admin

import (
	"fmt"
	"server/client"
	"strings"
)

func Handler() {
	var input string
	for {
		fmt.Printf("> ")
		fmt.Scanln(&input)

		inputs := strings.Split(input, " ")
		switch inputs[0] {
		case "groups":
			client.ClientsMu.Lock()
			for name, groups := range client.Clients {
				fmt.Printf("%s: %d\n", name, len(groups))
			}
			client.ClientsMu.Unlock()
		case "print":
			if len(inputs) < 3 {
				fmt.Println("print [group/client] [name/id]")
				continue
			}

			if inputs[1] == "group" {

			}
		}
	}
}
