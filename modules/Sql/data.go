package Phoenix_Data

// package contains
/*
	: CONTAINS
		- Types
		- Variables
		- Contents
		- Structurs
		- Constants
		- Extra Data

Development team [ Scare security aka +The legacy hackers+ ]
*/

var ()

const (
	Authentiation_Str = "host=%s port=%s user=%s password=%s dbname=%s sslmode=%v"
)

type Sql_Configuration struct {
	Hostname       string
	Port           string
	Connect        string
	Username       string
	Password       string
	Database_Name  string
	Database_Query string
	Query_File     string
	SSL            string
}
