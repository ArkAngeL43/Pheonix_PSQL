package Phoenix_Data

import (
	"database/sql"
	"fmt"
	Phoenix_Errors "main/modules/Errors"
)

func (Config *Sql_Configuration) Alter_User(user, new_pass string) {
	user1 := user
	pass1 := new_pass
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

	defer db.Close()
	query := "ALTER USER %s WITH password '%s';"
	query = fmt.Sprintf(query, user1, pass1)
	if _, x := db.Exec(string(query)); x != nil {
		fmt.Println("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m|> \033[31m Error when trying to run the query [ ", query, " ]")
		fmt.Println("", x)
		fmt.Println("\033[38;5;55m|\033[31m-\033[38;5;55m|> \033[31m Configuration [Pass]   : ", Config.Password)
		fmt.Println("\033[38;5;55m|\033[31m-\033[38;5;55m|> \033[31m Configuration [User]   : ", Config.Username)
		fmt.Println("\033[38;5;55m|\033[31m-\033[38;5;55m|> \033[31m Configuration [Port]   : ", Config.Port)
		fmt.Println("\033[38;5;55m|\033[31m-\033[38;5;55m|> \033[31m Configuration [Host]   : ", Config.Hostname)
		fmt.Println("\033[38;5;55m|\033[31m-\033[38;5;55m|> \033[31m Configuration [SSL ]   : ", Config.SSL)
		fmt.Println("\033[38;5;55m|\033[31m-\033[38;5;55m|> \033[31m Configuration [DBMS]   : ", Config.Database_Name)
		fmt.Println("\033[38;5;55m|\033[31m-\033[38;5;55m|> \033[31m Configuration Altering : ", user1)
		fmt.Println("\033[38;5;55m|\033[31m-\033[38;5;55m|> \033[31m Configuration With pass: ", pass1)
	} else {
		fmt.Println("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m|> \033[37m User -> ", user1, " Has been altered with password -> ", pass1)
	}

}
