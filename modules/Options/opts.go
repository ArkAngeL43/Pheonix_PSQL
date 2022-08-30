package Options

type Options struct {
	Screen_Resolution string // Enable and pre set screen resolution verticle or landscape only for output
	Database_Name     string // Enable and pre set the database name
	Databse_Host      string // Enable and pre set the database Hostname
	Database_Port     string // Enable and pre set the database port
	Database_User     string // Enable and pre set the database username
	Dastabase_Pass    string // Enable and pre set the database password
	Database_Query    string // Database query, if they want to execute one command and exit
	Database_file     string // Database file, aka SQL file .SQL files to load queries
	Output_File       string // Location of the output
	Output_Filepath   string // Location of the output to a filepath
	SSL               bool   // Enable SSL mode
	Generate_HTML     bool   // Generate a HTML file of output from query
	Generate_YAML     bool   // Generate a YAML file of output from query
	Generate_JSON     bool   // Generate a JSON file of output from query
	Databases         bool   // List all databases and exit
	Output            bool   // if user wants to output they need this flag
	Help_Internal     bool   // Help with internal modules [Commands]
	Help_External     bool   // Help with external modules [flags]
	Banner_File       bool   // Optional banner file for logo
	Ping              bool   // Try ping and exit
	Auth              bool   // Try authentication and exit
	Version           bool   // Version of the script
}
