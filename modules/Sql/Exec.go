package Phoenix_Data

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	Phoenix_Errors "main/modules/Errors"
	Phoenix_Data "main/modules/Options"

	_ "github.com/lib/pq"
)

func (Config *Sql_Configuration) Execute_SQL_File(filename string) {
	var v Phoenix_Data.Options
	if v.Database_file == "" {
		v.Database_file = filename
	}
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
	qu, x := ioutil.ReadFile(v.Database_file)
	if Phoenix_Errors.EM(x) {
		fmt.Println("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m|> \033[31m Error when loading SQL file (", v.Database_file, ") into a runable query ", x)
	} else {
		if _, x := db.Exec(string(qu)); x != nil {
			fmt.Println("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m|> \033[31m Error when loading SQL file (", v.Database_file, ") into a runable query ", x)
		} else {
			fmt.Println("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m|> \033[38;5;43m Query has been run, status was okay")
			fmt.Println("-----------------------------------")
		}
	}
}
