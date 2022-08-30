package Phoenix_Data

import (
	"fmt"
	"log"
	"os/exec"
)

/*


Right so i could be missing something but for some reason it was too much of a

hastle and time waster to try and get rows of the db, append, save, and output them

all in a table format in go, so i wrote a mini perl script to do so for us, which uses

the DBI driver to properly get data and output it

	fmt.Println(
		prg,
		arg1,
		prg2,
		arg2,
		prg3,
		arg3,
		prg4,
		arg4,
		prg5,
		arg5,
		prg6,
		arg6,
		prg7,
		arg7)

*/
func (Config *Sql_Configuration) Perl_Rows(Table string) {
	prg := "perl"
	arg1 := "modules/Lang-based/Tables.pl"
	prg2 := "-n"
	arg2 := Config.Database_Name
	prg3 := "-p"
	arg3 := Config.Port
	prg4 := "-h"
	arg4 := Config.Hostname
	prg5 := "-q"
	arg5 := "SELECT * FROM " + Table
	prg6 := "-a"
	arg6 := Config.Password
	prg7 := "-u"
	arg7 := Config.Username
	cmd := exec.Command(
		prg,
		arg1,
		prg2,
		arg2,
		prg3,
		arg3,
		prg4,
		arg4,
		prg5,
		arg5,
		prg6,
		arg6,
		prg7,
		arg7)
	stdout, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(string(stdout))
}
