package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "os/exec"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(34.133.39.78:3306)/sopesDB")
	defer db.Close()

	if err != nil {
		log.Fatal(err)
	}

	salida := 0
	for {
		query := `INSERT INTO tarea2(nombre, carnet) VALUES ();`
		result, err := db.Exec(query, "Tiempo", salida)

		if err != nil {
			fmt.Println(err)
		}

		fmt.Println("Tiempo", salida)
		fmt.Println(result)
		if salida == 4 {
			break
		}
		salida += 1
		time.Sleep(1 * time.Second)
	}

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
