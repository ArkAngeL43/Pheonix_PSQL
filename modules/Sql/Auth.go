package Phoenix_Data

import (
	"database/sql"
	"fmt"
	Phoenix_Errors "main/modules/Errors"
)

// Contains authentiation

func (Config *Sql_Configuration) Auth(filler string) {
	c, x := sql.Open("postgres", fmt.Sprintf(Authentiation_Str,
		Config.Hostname,
		Config.Port,
		Config.Username,
		Config.Password,
		Config.Database_Name,
		Config.SSL))
	if Phoenix_Errors.EM(x) {
		fmt.Println("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m|> \033[31m Got an error when trying to authenticate the database ~~> ", x)
		fmt.Println("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m|> \033[31m Configuration [Pass]   : ", Config.Password)
		fmt.Println("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m|> \033[31m Configuration [User]   : ", Config.Username)
		fmt.Println("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m|> \033[31m Configuration [Port]   : ", Config.Port)
		fmt.Println("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m|> \033[31m Configuration [Host]   : ", Config.Hostname)
		fmt.Println("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m|> \033[31m Configuration [SSL ]   : ", Config.SSL)
		fmt.Println("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m|> \033[31m Configuration [DBMS]   : ", Config.Database_Name)
	}
	x = c.Ping()
	if Phoenix_Errors.EM(x) {
		fmt.Println("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m|> \033[31m Got an error when trying to authenticate the database ~~> ", x)
		fmt.Println("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m|> \033[31m Configuration [Pass]   : ", Config.Password)
		fmt.Println("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m|> \033[31m Configuration [User]   : ", Config.Username)
		fmt.Println("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m|> \033[31m Configuration [Port]   : ", Config.Port)
		fmt.Println("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m|> \033[31m Configuration [Host]   : ", Config.Hostname)
		fmt.Println("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m|> \033[31m Configuration [SSL ]   : ", Config.SSL)
		fmt.Println("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m|> \033[31m Configuration [DBMS]   : ", Config.Database_Name)
	} else {
		fmt.Println("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m|> \033[32m Authentication has been made, login correct")

		fmt.Println("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m|> \033[38;5;43m Configuration [Pass]   : \033[37m", Config.Password)
		fmt.Println("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m|> \033[38;5;43m Configuration [User]   : \033[37m", Config.Username)
		fmt.Println("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m|> \033[38;5;43m Configuration [Port]   : \033[37m", Config.Port)
		fmt.Println("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m|> \033[38;5;43m Configuration [Host]   : \033[37m", Config.Hostname)
		fmt.Println("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m|> \033[38;5;43m Configuration [SSL ]   : \033[37m", Config.SSL)
		fmt.Println("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m|> \033[38;5;43m Configuration [DBMS]   : \033[37m", Config.Database_Name)
	}
	defer c.Close()
}
