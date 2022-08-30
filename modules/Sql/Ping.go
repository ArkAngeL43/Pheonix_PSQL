package Phoenix_Data

import (
	"fmt"
	Phoenix_Errors "main/modules/Errors"
	"net"
)

// package contains
/*
	: CONTAINS
		- Function to ping databases
*/

func (Config *Sql_Configuration) Ping(method string) {
	con, x := net.Dial("tcp", net.JoinHostPort(Config.Hostname, Config.Port))
	if Phoenix_Errors.EM(x) {
		fmt.Println("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m|> Got a error when trying to connect -> ", x)
	} else {
		fmt.Println("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m|>  Configuration   : DBMS -> \033[38;5;43m Up and alive")
		fmt.Println("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m|>  Configuration   : Port -> \033[38;5;43m", Config.Port)
		fmt.Println("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m|>  Configuration   : Host -> \033[38;5;43m", Config.Hostname)
	}
	defer con.Close()
}
