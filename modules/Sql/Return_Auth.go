package Phoenix_Data

import (
	"database/sql"
	"fmt"
	Phoenix_Errors "main/modules/Errors"
)

func (Config *Sql_Configuration) Execute_Auth() *sql.DB {
	db, err := sql.Open("postgres", fmt.Sprintf(Authentiation_Str,
		Config.Hostname,
		Config.Port,
		Config.Username,
		Config.Password,
		Config.Database_Name,
		Config.SSL))
	if Phoenix_Errors.EM(err) {
		fmt.Println("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m|> \033[31m Error when trying to authenticate to the database -> ", err)
		fmt.Println("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m|> \033[31m Configuration [Pass]   : ", Config.Password)
		fmt.Println("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m|> \033[31m Configuration [User]   : ", Config.Username)
		fmt.Println("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m|> \033[31m Configuration [Port]   : ", Config.Port)
		fmt.Println("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m|> \033[31m Configuration [Host]   : ", Config.Hostname)
		fmt.Println("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m|> \033[31m Configuration [SSL ]   : ", Config.SSL)
		fmt.Println("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m|> \033[31m Configuration [DBMS]   : ", Config.Database_Name)
	}
	err = db.Ping()
	if Phoenix_Errors.EM(err) {
		fmt.Println("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m|> \033[31m Error when trying to authenticate to the database -> ", err)
		fmt.Println("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m|> \033[31m Configuration [Pass]   : ", Config.Password)
		fmt.Println("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m|> \033[31m Configuration [User]   : ", Config.Username)
		fmt.Println("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m|> \033[31m Configuration [Port]   : ", Config.Port)
		fmt.Println("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m|> \033[31m Configuration [Host]   : ", Config.Hostname)
		fmt.Println("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m|> \033[31m Configuration [SSL ]   : ", Config.SSL)
		fmt.Println("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m|> \033[31m Configuration [DBMS]   : ", Config.Database_Name)
	}
	return db
}
