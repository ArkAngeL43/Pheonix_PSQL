package Phoenix_Command

import (
	"fmt"
)

func Output_Flags(filler string) {
	fmt.Println("   \033[38;5;55m--res       \033[37m| Set the resolution for desired output, Verticle and Landscape are the only supported modes")
	fmt.Println("	\033[38;5;55m--dbname    \033[37m| Set the database name to log into")
	fmt.Println("	\033[38;5;55m--dbhost    \033[37m| Set the database host to connect to")
	fmt.Println("	\033[38;5;55m--dbport    \033[37m| Set the database port to connect to")
	fmt.Println("	\033[38;5;55m--dbuser    \033[37m| Set the database username to log on as ")
	fmt.Println("	\033[38;5;55m--dbpass    \033[37m| Set the database password to log in with")
	fmt.Println("	\033[38;5;55m--sqlf      \033[37m| Set a .SQL file to run")
	fmt.Println("	\033[38;5;55m--output    \033[37m| Set a output filename, example would be log.txt")
	fmt.Println("	\033[38;5;55m--html      \033[37m| Set wether or not you want to output data from certian commands in a HTML template")
	fmt.Println("	\033[38;5;55m--yam       \033[37m| Set wether or not you want to output data from certian commands in a YAML template")
	fmt.Println("	\033[38;5;55m--dbs       \033[37m| Upon being true or activated, print all database names possible and exit")
	fmt.Println("	\033[38;5;55m--json      \033[37m| Set wether or not you want to output data from certian commands in a JSON template")
	fmt.Println("	\033[38;5;55m--ping      \033[37m| Upon being true or activated, ping if a database is alive and exit (--dbhost, and --dbport must be used)")
	fmt.Println("	\033[38;5;55m--auth      \033[37m| Upon being true or activated, attempt to authenticate a database and exit (--dbname, --dbhost, --dbpass, --dbuser, --dbport must be used)")
	fmt.Println("	\033[38;5;55m--version   \033[37m| Upon being true or activated, print script or suite version and exit")
}

func Output_Commands(filler string) {
	fmt.Println("========= Modules")
	fmt.Println("	\033[38;5;55mset   \033[37m| when using set, you set the session variables such as user/dbname/host etc...")
	fmt.Println("                                \033[38;5;43mset \033[31msyntax    \033[37m: set <pass/user/host/port/name/sslm> <value>")
	fmt.Println("                                \033[38;5;43mset \033[31mexample   \033[37m: set user username1234")
	fmt.Println("                                \033[38;5;43mset \033[31mdictionary\033[37m: set user -> set the sessions username to log into")
	fmt.Println("                                \033[38;5;43mset \033[31mdictionary\033[37m: set pass -> set the sessions password to log into")
	fmt.Println("                                \033[38;5;43mset \033[31mdictionary\033[37m: set name -> set the sessions DBMS     to log into")
	fmt.Println("                                \033[38;5;43mset \033[31mdictionary\033[37m: set port -> set the sessions port     to log into")
	fmt.Println("                                \033[38;5;43mset \033[31mdictionary\033[37m: set host -> set the sessions host     to log into")
	fmt.Println("                                \033[38;5;43mset \033[31mdictionary\033[37m: set sslm -> set the sessions SSL mode ")
	fmt.Println("	\033[38;5;55mrun   \033[37m| when using run, you grab/edit data from the database and output it such as tables, or alter users")
	fmt.Println("                                \033[38;5;43mrun \033[31msyntax        \033[37m: run <createuser/createdbms/selectrowc/selectallr/selectallt/alteruserp> <value>")
	fmt.Println("                                \033[38;5;43mrun \033[31mexample       \033[37m: run createuser username1234")
	fmt.Println("                                \033[38;5;43mrun \033[31mexample       \033[37m: run createdbms database_name123")
	fmt.Println("                                \033[38;5;43mrun alter \033[31msyntax  \033[37m: when running the command alteruserp you are going to need to set values as so ")
	fmt.Println("                                \033[38;5;43mrun alter \033[31msyntax  \033[37m: username:new_password, note you must have admin access to do this, an example below")
	fmt.Println("                                \033[38;5;43mrun alter \033[31msyntax  \033[37m: will modify and change the password to a username")
	fmt.Println("                                \033[38;5;43mrun alter \033[31mexample \033[37m: run alteruserp username1234:admin123")
}
