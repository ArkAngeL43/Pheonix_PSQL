package Phoenix_Data

import (
	"fmt"
)

func (Config *Sql_Configuration) Col_Count(tb_name string) {
	var exec = "SELECT count(*) FROM %s"
	DT := fmt.Sprintf(exec, tb_name)
	Auth = Config.Execute_Auth()
	if _, x := Auth.Exec(string(DT)); x != nil {
		fmt.Println("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m|> \033[31m Could not authenticate SQL DB -> ", x)
	} else {
		var c int
		fmt.Println("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m|> \033[38;5;43m Query has been run")
		row := Auth.QueryRow(DT)
		x := row.Scan(&c)
		if x != nil {
			fmt.Println("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m|> \033[31m Could not execute query -> ", x)
		} else {
			fmt.Println("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m|> ROW COUNT OF TABLE [ ", tb_name, " ] ~~> ", c)
		}
	}
}
