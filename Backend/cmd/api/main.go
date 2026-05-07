package main

import (
	"enrollapp/internal/db"
	"fmt"
)

func main() {
	db.ConnectDB()
	//Vamos a probar poner datos en la tabla de careers_tb en la base de datos StudentsDB

	fmt.Println("Hello, World!")
}
