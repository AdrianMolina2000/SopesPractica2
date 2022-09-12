package main

import (
	"database/sql"
	"fmt"
	"log"
	"os/exec"
	_ "os/exec"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(34.133.39.78:3306)/sopesDB")
	defer db.Close()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Datos obtenidos desde el Modulo: ")
	fmt.Println("")

	for {
		cmd := exec.Command("sh", "-c", "cat /proc/ram_201903850")
		cmd2 := exec.Command("sh", "-c", "cat /proc/cpu_201903850")

		out, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Println(err)
		}

		out2, err := cmd2.CombinedOutput()
		if err != nil {
			fmt.Println(err)
		}

		output := string(out[:])
		output2 := string(out2[:])

		query := `INSERT INTO ram(valor) VALUES (?);`
		query2 := `INSERT INTO cpu(valor) VALUES (?);`

		result, err := db.Exec(query, output)
		if err != nil {
			fmt.Println(err)
		}

		result2, err := db.Exec(query2, output2)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(result)
		fmt.Println(result2)

		time.Sleep(5 * time.Second)
	}

}
