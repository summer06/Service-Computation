package main

import (
	"fmt"
	"os"
	orm "webServer_database/orm/services"
	sql "webServer_database/sql/services"

	flag "github.com/spf13/pflag"
)

const (
	PORT string = "8080"
)

func main() {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = PORT
	}

	method := flag.StringP("method", "m", "sql", "method for implementation")
	pPort := flag.StringP("port", "p", PORT, "PORT for httpd listening")
	flag.Parse()
	if len(*pPort) != 0 {
		port = *pPort
	}

	if *method == "sql" {
		fmt.Println("using sql implementation version")
		server := sql.NewServer()
		server.Run(":" + port)
	} else if *method == "orm" {
		fmt.Println("using orm implementation version")
		server := orm.NewServer()
		server.Run(":" + port)
	}

}
