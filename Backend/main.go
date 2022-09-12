package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "os/exec"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(34.133.39.78:3306)/sopesDB")
	defer db.Close()

	if err != nil {
		log.Fatal(err)
	}

	var version string

	err2 := db.QueryRow("SELECT VERSION()").Scan(&version)

	if err2 != nil {
		log.Fatal(err2)
	}

	fmt.Println(version)

	/*
		fmt.Println("Datos obtenidos desde el Modulo: ")
		fmt.Println("")

		cmd := exec.Command("sh", "-c", "cat /proc/ram_201903850")
		out, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Println(err)
		}
		output := string(out[:])
		fmt.Println(output)
	*/
}
