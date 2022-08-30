package Phoenix_Data

import (
	"database/sql"
	"fmt"
)

var Auth *sql.DB

func (Config *Sql_Configuration) Create_User(user string) {
	var exec = "CREATE USER %s"
	DT := fmt.Sprintf(exec, user)
	Auth = Config.Execute_Auth()
	if _, x := Auth.Exec(string(DT)); x != nil {
		fmt.Println("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m|> \033[31m Could not execute query -> ", x)
	} else {
		fmt.Println("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m|> \033[38;5;43m Query has been run, USER [", user, "] Has been successfuly created")
	}
}
