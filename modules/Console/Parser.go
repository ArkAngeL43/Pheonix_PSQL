package Phoenix_Command

import (
	"fmt"
	"strings"
)

func Parser(I string, Map_of_choice map[string]interface{}) {
	I = strings.TrimSpace(I)
	if Map_of_choice[I] == nil {
		fmt.Println("\033[31m[-] Error: Command not found")
	} else {
		dt := Map_of_choice[I]
		dt.(func(string))("")
	}
}

func Parser_With_Args(Module string, Map_of_choice map[string]interface{}, Module_Arg string) {
	Module = strings.TrimSpace(Module)
	if Map_of_choice[Module] == nil || Map_of_choice[Module] == "" {
		fmt.Println("\033[31m[-] Error: Command not found, or value came empty")
	} else {
		dt := Map_of_choice[Module]
		dt.(func(string))(Module_Arg)
	}
}
