package main

import _ "github.com/denisenkom/go-mssqldb"
import "database/sql"
import "log"
import "fmt"
import "flag"

var debug = flag.Bool("debug", false, "enable debugging")
var password = flag.String("password", "[ibdkifu", "the database password")
var port *int = flag.Int("port", 1433, "the database port")
var server = flag.String("server", "192.168.0.194", "the database server")
var user = flag.String("user", "sa", "the database user")

func main() {
	flag.Parse() // parse the command line args

	if *debug {
		fmt.Printf(" password:%s\n", *password)
		fmt.Printf(" port:%d\n", *port)
		fmt.Printf(" server:%s\n", *server)
		fmt.Printf(" user:%s\n", *user)
	}

	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;", *server, *user, *password, *port)
	if *debug {
		fmt.Printf(" connString:%s\n", connString)
	}
	conn, err := sql.Open("mssql", connString)
	if err != nil {
		log.Fatal("Open connection failed:", err.Error())
	}
	defer conn.Close()

	stmt, err := conn.Query("select 1")
	if err != nil {
		log.Fatal("Prepare failed:", err.Error())
	}
	defer stmt.Close()
	var somenumber int64
	var somechars string
	for stmt.Next() {
		err = stmt.Scan(&somenumber, &somechars)
	}


	if err != nil {
		log.Fatal("Scan failed:", err.Error())
	}
	fmt.Printf("somenumber:%d\n", somenumber)
	fmt.Printf("somechars:%s\n", somechars)

	fmt.Printf("bye\n")

}
