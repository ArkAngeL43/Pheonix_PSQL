package Phoenix_Command

import (
	"fmt"
	Phonix_Main_Constants "main/modules"
)

// this is not a requirement, current alternative to the mapping
func Clear(dt string) {
	fmt.Println("\x1b[H\x1b[2J\x1b[3J")
}

func Auxilary_Set_Pass(password string) {
	Phonix_Main_Constants.Structure.Password = password
	fmt.Println(Color_Bar("\033[38;5;43m+"), "Config  : Pass -> ", password)

}

func Auxilary_Set_User(username string) {
	Phonix_Main_Constants.Structure.Username = username
	fmt.Println(Color_Bar("\033[38;5;43m+"), "Config  : User -> ", username)

}

func Auxilary_Set_Host(hostname string) {
	Phonix_Main_Constants.Structure.Hostname = hostname
	fmt.Println(Color_Bar("\033[38;5;43m+"), "Config  : Host -> ", hostname)
}

func Auxilary_Set_Dbname(database_name string) {
	Phonix_Main_Constants.Structure.Database_Name = database_name
	fmt.Println(Color_Bar("\033[38;5;43m+"), "Config  : DB Name -> ", database_name)
}

func Auxilary_Set_Port(port string) {
	Phonix_Main_Constants.Structure.Port = port
	fmt.Println(Color_Bar("\033[38;5;43m+"), "Config  : Port -> ", port)
}

func Auxilary_Set_SSL(ssl string) {
	switch ssl {
	case "verify-ca":
		Phonix_Main_Constants.Structure.SSL = "verify-ca"
	case "require":
		Phonix_Main_Constants.Structure.SSL = "require"
	case "verify-full":
		Phonix_Main_Constants.Structure.SSL = "verify-full"
	case "deactivate":
		Phonix_Main_Constants.Structure.SSL = "disable"
	case "off":
		Phonix_Main_Constants.Structure.SSL = "disable"
	}
	fmt.Println(Color_Bar("\033[38;5;43m+"), "Config  : SSL Mode -> ", ssl)
}

func Settings(filler string) {
	fmt.Println(filler)
	fmt.Println(Color_Bar("\033[38;5;43m+"), "Configuration  : Port  -> ", Phonix_Main_Constants.Structure.Port)
	fmt.Println(Color_Bar("\033[38;5;43m+"), "Configuration  : Host  -> ", Phonix_Main_Constants.Structure.Hostname)
	fmt.Println(Color_Bar("\033[38;5;43m+"), "Configuration  : DBMS  -> ", Phonix_Main_Constants.Structure.Database_Name)
	fmt.Println(Color_Bar("\033[38;5;43m+"), "Configuration  : Pass  -> ", Phonix_Main_Constants.Structure.Password)
	fmt.Println(Color_Bar("\033[38;5;43m+"), "Configuration  : User  -> ", Phonix_Main_Constants.Structure.Username)
	fmt.Println(Color_Bar("\033[38;5;43m+"), "Configuration  : SSL   -> ", Phonix_Main_Constants.Structure.SSL)
}

func Color_Bar(filler string) string {
	a := fmt.Sprintf("\033[38;5;55m|%s\033[38;5;55m|> ", filler)
	return a
}
