package Phoenix_Data

import (
	"fmt"
	Phoenix_Errors "main/modules/Errors"
)

var (
	NULL_Counter   = 0
	Column_Counter = 0
	Row_Counter    = 0
)

func (Config *Sql_Configuration) Run(table string) {
	// thank you so fucking much, whoever you are `https://groups.google.com/g/golang-nuts/c/Fvwz10kvLLE`
	var Rows_ []interface{}
	var Cols_ []string
	var BB_Row []string
	Auth = Config.Execute_Auth()
	row, x := Auth.Query(fmt.Sprintf("SELECT * FROM %s", table))
	if Phoenix_Errors.EM(x) {
		fmt.Println("\033[31m[-] Error: Could not run the given query for `SELECT * FROM ....` i got the error -> ", x)
	} else {
		columns, x2 := row.Columns()
		if Phoenix_Errors.EM(x2) {
			fmt.Println("\033[31m[-] Error: Could not get the column count or rows or even scan it got the error -> ", x2)
		} else {
			Column_Counter = len(columns)
			sl := make([]interface{}, len(columns))
			cs := make([]interface{}, len(sl))
			for i := range sl {
				cs[i] = &sl[i]
			}
			Cols_ = append(Cols_, columns...)
			for row.Next() {
				x = row.Scan(cs...)
				if Phoenix_Errors.EM(x) {
					fmt.Println("\033[31m[-] Error: Could not scan the rows with the given scan and run arguments, got the error -> ", x)
				} else {

					for _, value := range sl {
						Row_Counter++
						switch value.(type) {
						case nil:
							NULL_Counter++
							fmt.Print(" | \t NULL")

						case []byte:
							BB_Row = append(BB_Row, string(value.([]byte)))
						default:
							Rows_ = append(Rows_, value)
						}
					}
				}
			}
			fmt.Println("\033[31m---------------------------------------------------------------------------------------------")
			fmt.Println("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m|> \033[37m Column Names  \033[38;5;43m=>  \033[37m", Cols_)
			fmt.Println("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m|> \033[37m Total Columns \033[38;5;43m=>  \033[37m", Column_Counter)
			fmt.Println("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m|> \033[37m Total Rows    \033[38;5;43m=>  \033[37m", Row_Counter)
			fmt.Println("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m|> \033[37m Total NULLS   \033[38;5;43m=>  \033[37m", NULL_Counter)
			fmt.Println("\033[31m---------------------------------------------------------------------------------------------")
			fmt.Println("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m|> \033[37m Row data organized by column  \033[38;5;43m=> ", Rows_)
		}
	}
}
