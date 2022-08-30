package Phoenix_Command

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"
	"strings"

	Phoenix_Helpers "main/modules/Helpers"
)

func Handelreturncon(c chan os.Signal) {
	signal.Notify(c, os.Interrupt)
	for s := <-c; ; s = <-c {
		switch s {
		case os.Interrupt:
			fmt.Println("[-] Detected interruption")
			os.Exit(0)
		case os.Kill:
			fmt.Println("GOT OS KILL")
			os.Exit(0)
		}
	}
}

/*

User must input data like the following

module <module number> <command>

an example is if all modules are set

*/

func Console() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(Color_Bar("\033[38;5;43m+"), "Console  :>> \033[37m")
		t, _ := reader.ReadString('\n')
		t = strings.Replace(t, "\n", "", -1)
		go Handelreturncon(make(chan os.Signal, 1))
		var Modules = map[string]string{
			"1": "No argument map",
		}
		if strings.Contains(t, "module") {
			Mods := Modules[string(t[7:8])]
			switch Mods {
			case "No argument map":
				if strings.Compare(t, "module 1 authenticate") == 0 {

				} else {
					Parser(string(t[8:]), No_Args_Commands)
				}
			}
		}
		if strings.Contains(t, "set") {
			Parser_With_Args(string(t[4:8]), Auxilary, string(t[9:]))
		}
		if strings.Contains(t, "run") {
			I := t[4:14]
			I = strings.TrimSpace(I)
			if strings.Contains(t, "alteruserp") {
				type TG struct {
					User string
					Pass string
				}
				I = strings.TrimSpace(I)
				if Aux_DB_2_args == nil {
					fmt.Println("\033[31m[-] Error: Command not found")
				} else {
					data := strings.FieldsFunc(t[15:], Phoenix_Helpers.Split)
					dt := Aux_DB_2_args[I]
					dt.(func(string, string))(data[0], data[1])
				}
			}
			if strings.Contains(t, "createdbms") {
				dt := Aux_DB[I]
				var a string
				fmt.Print("Enter the owner> ")
				fmt.Scanf("%s", &a)
				dt.(func(string, string))(t[15:], a)
			} else {
				Parser_With_Args(I, Aux_DB, string(t[15:]))
			}
		}
		if strings.Contains(t, "load") {
			I := t[4:14]
			I = strings.TrimSpace(I)
			Parser_With_Args(I, Load, strings.TrimSpace(string(t[14:])))
		}
	}
}
