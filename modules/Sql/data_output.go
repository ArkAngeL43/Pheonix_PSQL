package Phoenix_Data

import (
	"fmt"
	"io/ioutil"
	Phoenix_Errors "main/modules/Errors"

	"gopkg.in/yaml.v2"
)

type Database_Information struct {
	DB_Name string `yaml:"Database_Name"`
	DB_Host string `yaml:"Database_Host"`
	DB_Port string `yaml:"Database_Port"`
	DB_Pass string `yaml:"Database_Pass"`
	DB_SSLM string `yaml:"Database_SSL"`
	DB_User string `yaml:"Database_User"`
}

func (Config *Sql_Configuration) Reader_Database_Config(filename string) {
	f, x := ioutil.ReadFile(filename)
	if Phoenix_Errors.EM(x) {
		fmt.Println("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m|> \033[31m Error when trying to read the YAML file -> ", x)
	}
	d := make(map[string]Database_Information)
	x = yaml.Unmarshal(f, &d)
	if Phoenix_Errors.EM(x) {
		fmt.Println("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m|> \033[31m Error when trying to parse the YAML structure -> ", x)
	}
	for _, v := range d {
		Config.Database_Name = v.DB_Name
		fmt.Println("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m|> \033[38;5;43m Configuration [New Name]   : \033[37m", Config.Database_Name)
		Config.Hostname = v.DB_Host
		fmt.Println("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m|> \033[38;5;43m Configuration [New Host]   : \033[37m", Config.Hostname)
		Config.Port = v.DB_Port
		fmt.Println("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m|> \033[38;5;43m Configuration [New Port]   : \033[37m", Config.Port)
		Config.SSL = v.DB_SSLM
		fmt.Println("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m|> \033[38;5;43m Configuration [New SSLM]   : \033[37m", Config.SSL)
		Config.Password = v.DB_Pass
		fmt.Println("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m|> \033[38;5;43m Configuration [New Pass]   : \033[37m", Config.Password)
		Config.Username = v.DB_User
		fmt.Println("\033[38;5;55m|\033[38;5;43m+\033[38;5;55m|> \033[38;5;43m Configuration [New User]   : \033[37m", Config.Username)
	}
	Config.Auth("")
	Config.Ping("tcp")
}
