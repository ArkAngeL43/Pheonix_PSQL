package Pheonix_Constants_MAIN

import Phoenix_SQL "main/modules/Sql"

var Structure = Phoenix_SQL.Sql_Configuration{}

var (
	Connect_and_authenticate_string = "host=%s port=%d user=%s password=%s dbname=%s sslmode=%s"
)
