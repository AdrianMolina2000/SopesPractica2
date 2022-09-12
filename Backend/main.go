package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os/exec"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Datos obtenidos desde el Modulo: ")
	fmt.Println("")

	bd, err := getDB()
	if err != nil {
		log.Printf("Error en la BDD: " + err.Error())
		return
	} else{
		err = bd.Ping()
		if err != nil {
			log.Printf("Error al conectar con la BDD: " + err.Error())
			return
		}
	}

	router := mux.NewRouter()
	
	
	
	
		/*cmd := exec.Command("sh", "-c", "cat /proc/ram_201903850")
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
	}

	output := string(out[:])
	fmt.Println(output)
	*/
}
