package Phoenix_Data

import "fmt"

func (Config *Sql_Configuration) Create_Database(dbname, owner string) {
	var exec = "CREATE DATABASE %s OWNER %s"
	DT := fmt.Sprintf(exec, dbname, owner)
	Auth = Config.Execute_Auth()
	if _, x := Auth.Exec(string(DT)); x != nil {
		fmt.Println("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m|> \033[31m Could not execute query -> ", x)
	} else {
		fmt.Println("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m|> \033[38;5;43m Query has been run, Database [", dbname, "] Has been successfuly created")
	}
}
