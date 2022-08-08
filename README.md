```
 ______________                   _____            ________                  _____                       
 ___  __ \__  /______________________(_)___  __    __  ___/______________   ____(_)__________________    
 __  /_/ /_  __ \  __ \  _ \_  __ \_  /__  |/_/    _____ \_  _ \_  ___/_ | / /_  /_  ___/  _ \_  ___/    
 _  ____/_  / / / /_/ /  __/  / / /  / __>  <      ____/ //  __/  /   __ |/ /_  / / /__ /  __/(__  )     
 /_/     /_/ /_/\____/\___//_/ /_//_/  /_/|_|      /____/ \___//_/    _____/ /_/  \___/ \___//____/      
 ─────────────────────────────────────────────────────────────────────────────────────────
```

# What is the PostGreSQL tool in the suite 

This tool allows you to connect, authenticate, run commands and even run SQL files for PostGreSQL databases, the idea behind this tool later on is to not just add support for MySQL, SQL Server but to also add support for solutions like MongoDB and eventually make a fully reliable cross platform tool for managing and extracting data from databases. This tool uses Go and Perl to work 

# Setup

run the following command 

`git clone https://www.github.com/ArkAngeL43/Pheonix_PSQL ; cd Pheonix_PSQL ; sudo cpan install DBI ; sudo cpan install Text::Table ; sudo cpan install Getopt::Std`


# Navigation

This program is easy to navigate once you get used to it. you can run the program as follows

> To enter the console 

`go run main.go`

> To get internal/console commands 

`go run main.go --helpi`

> To get external/argument flags 

`go run main.go --helpe`

> To ping a database and exit 

`go run main.go --dbhost 127.0.0.1 --dbport 5432 --ping`

> To auth a database and exit 

`go run main.go --dbhost 127.0.0.1 --dbport 5432 --dbname TestDB --dbuser user123 --dbpass password123 --auth`


# Navigation around the console 

#### Internal commands and usage ####

```
	set   | when using set, you set the session variables such as user/dbname/host etc...
                                set syntax    : set <pass/user/host/port/name/sslm> <value>
                                set example   : set user username1234
                                set dictionary: set user -> set the sessions username to log into
                                set dictionary: set pass -> set the sessions password to log into
                                set dictionary: set name -> set the sessions DBMS     to log into
                                set dictionary: set port -> set the sessions port     to log into
                                set dictionary: set host -> set the sessions host     to log into
                                set dictionary: set sslm -> set the sessions SSL mode 
	run   | when using run, you grab/edit data from the database and output it such as tables, or alter users
                                run syntax        : run <createuser/createdbms/selectrowc/selectallr/selectallt/alteruserp> <value>
                                run example       : run createuser username1234
                                run example       : run createdbms database_name123
                                run alter syntax  : when running the command alteruserp you are going to need to set values as so 
                                run alter syntax  : username:new_password, note you must have admin access to do this, an example below
                                run alter syntax  : will modify and change the password to a username
                                run alter example : run alteruserp username1234:admin123
```

> To set SSL mode require

`set sslm require`

> To set SSL mode disable 

`set sslm disable`

> To set SSL mode verify full 

`set sslm verify-full`

> To set SSL mode verify ca

`set sslm verify-ca`

> To set host 

`set host 127.0.0.1`

> To set port

`set port 5432`

> to set user

`set user user123`

> to set database name to use 

`set name database_name`

> to create a user 

`run createuser username123`

> to create a dbms

`run createdbms database_1`

> to select all row columns 

`run selectrowc db_row_1`

> to select all rows

`run selectallr row_name`

> to select all from the table

`run selectallt table_name`

> to alter a users password 

`run alteruserp username1234:admin123`

> to ping a database 

`module 1 ping`

> to clear the terminal

`module 1 cls` or `module 1 clear`

> to authenticate a database 

`module 1 auth`

> to view all settings

`module 1 settings`

> to view console help 

`module 1 help`

> Loading custom configuration files

take the config file as database settings


**config.yaml**

```yaml
Database_Information:
         Database_Name: "postgres"
         Database_Host: "127.0.0.1"
         Database_Port: "5432"
         Database_Pass: "admin123"
         Database_SSL: "disable"
         Database_User: "postgres"
```

we can load this file into the settings like so 

```
load settingsf config.yaml 
```

> loading `.SQL` files for execution

```
load dbqueryef filename.SQL
```
