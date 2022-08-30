package Phoenix_Command

import (
	Phoenix_Main_Constants "main/modules"
)

// nag = No argument needed
// this will trigger maps that do not need agruments
var No_Args_Commands = map[string]interface{}{
	"ping":     Phoenix_Main_Constants.Structure.Ping, // ping database and stat it
	"cls":      Clear,                                 // clear terminal
	"clear":    Clear,                                 // clear terminal
	"auth":     Phoenix_Main_Constants.Structure.Auth, // Authenticate
	"settings": Settings,
	"help":     Output_Commands,
}

var Auxilary = map[string]interface{}{
	"pass": Auxilary_Set_Pass,
	"user": Auxilary_Set_User,
	"host": Auxilary_Set_Host,
	"port": Auxilary_Set_Port,
	"name": Auxilary_Set_Dbname,
	"sslm": Auxilary_Set_SSL,
}

var Aux_DB = map[string]interface{}{
	"createuser": Phoenix_Main_Constants.Structure.Create_User,
	"createdbms": Phoenix_Main_Constants.Structure.Create_Database,
	"selectrowc": Phoenix_Main_Constants.Structure.Col_Count,
	"selectallr": Phoenix_Main_Constants.Structure.Perl_Rows,
	"selectallt": Phoenix_Main_Constants.Structure.Run,
}

var Aux_DB_2_args = map[string]interface{}{
	"alteruserp": Phoenix_Main_Constants.Structure.Alter_User,
}

var Load = map[string]interface{}{
	"settingsf": Phoenix_Main_Constants.Structure.Reader_Database_Config,
	"dbqueryef": Phoenix_Main_Constants.Structure.Execute_SQL_File,
}
