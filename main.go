package main

import (
	"fmt"
	Phoenix_Graphics "main/modules/Files"
	"os"

	Phoenix_Structure "main/modules"
	Phoenix_Colors "main/modules/Coloring"
	Phoenix_Console "main/modules/Console"
	Phoenix_Options "main/modules/Options"

	"github.com/spf13/pflag"
)

func main() {
	Phoenix_Graphics.Open_File(
		Phoenix_Graphics.Banner_file1,
		Phoenix_Colors.GSIVQPMV,
		true)
	fmt.Println(Phoenix_Colors.XUAXFHWO, "─────────────────────────────────────────────────────────────────────────────────────────")
	Phoenix_Console.Console()
}

var Options = Phoenix_Options.Options{}
var Options_SQL = Phoenix_Structure.Structure
var flags = pflag.FlagSet{SortFlags: false}
var Version = "{ 1.0 (BETA) } "

func init() {
	flags.StringVar(&Options.Screen_Resolution, "res", "Verticle", "Set the resolution for desired output |<Verticle>-<Landscape>|")
	flags.StringVar(&Phoenix_Structure.Structure.Database_Name, "dbname", "empty", "Set the database name to log into")
	flags.StringVar(&Phoenix_Structure.Structure.Hostname, "dbhost", "127.0.0.1", "Set the database host to log into")
	flags.StringVar(&Phoenix_Structure.Structure.Port, "dbport", "5432", "Set the database port number to connect to")
	flags.StringVar(&Phoenix_Structure.Structure.Username, "dbuser", "empty", "Set the database username to log into")
	flags.StringVar(&Phoenix_Structure.Structure.Password, "dbpass", "empty", "Set the database password to use to login")
	flags.StringVar(&Phoenix_Structure.Structure.Database_Query, "query", "empty", "Execute a Database query, print results and exit")
	flags.StringVar(&Options.Database_file, "sqlf", "examples/example.sql", "Execute a .SQL file, print results and exit")
	flags.StringVar(&Options.Output_File, "output", "log.txt", "Execute SQL queries and direct the output to a file (defualt=log.txt)")
	flags.StringVar(&Phoenix_Structure.Structure.SSL, "ssl", "disable", "Set the SSL mode (supported sslmode only 'require' (default), 'verify-full', 'verify-ca', and 'disable')")
	flags.BoolVar(&Options.Generate_HTML, "html", false, "Upon output, generate a HTML template as a table of all the data from execution")
	flags.BoolVar(&Options.Generate_YAML, "yaml", false, "Upon output, generate a YAML template of all of the data from a query")
	flags.BoolVar(&Options.Generate_JSON, "json", false, "Upon output, generate a JSON template of all of the data from a given query")
	flags.BoolVar(&Options.Databases, "dbs", false, "Upon being true, list all useable databases and exit")
	flags.BoolVar(&Options.Help_Internal, "helpi", false, "Upon being true, output all information about the modules in the CLI and exit")
	flags.BoolVar(&Options.Help_External, "helpe", false, "Upon being true, output all information about flags in the CLI and exit")
	flags.BoolVar(&Options.Banner_File, "banner", true, "Upon being tue, output the standard banner for your screen resolution, if false no banner will be shown")
	flags.BoolVar(&Options.Ping, "ping", false, "Upon being true, ping the database to see if its alive and exit (Defualt is a host of 127.0.0.1 and port of 5432)")
	flags.BoolVar(&Options.Auth, "auth", false, "Upon being true, attempt to authenticate the database to see if credentials are correct")
	flags.BoolVar(&Options.Version, "version", false, "Upon being true print script version and exit")
	flags.Parse(os.Args[1:])
	if Options.Version {
		Phoenix_Graphics.Open_File(
			Phoenix_Graphics.Banner_file1,
			Phoenix_Colors.GSIVQPMV,
			true)
		fmt.Println("Currnet Version -> ", Version)
		os.Exit(0)
	}
	if Options.Ping == true {
		Phoenix_Console.Parser("ping", Phoenix_Console.No_Args_Commands)
		os.Exit(0)
	}
	if Options.Auth == true {
		Phoenix_Console.Parser("ping", Phoenix_Console.No_Args_Commands)
		os.Exit(0)
	}
	if Options.Help_Internal == true {
		Phoenix_Graphics.Open_File(
			Phoenix_Graphics.Banner_file1,
			Phoenix_Colors.GSIVQPMV,
			true)
		fmt.Printf("\n\n\n\n")
		Phoenix_Console.Output_Commands("")
		os.Exit(0)
	}
	if Options.Help_External == true {
		Phoenix_Graphics.Open_File(
			Phoenix_Graphics.Banner_file1,
			Phoenix_Colors.GSIVQPMV,
			true)
		fmt.Printf("\n\n\n\n")

		Phoenix_Console.Output_Flags("")
		os.Exit(0)
	}
}
