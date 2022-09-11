package main

import (
	"fmt"
	"os/exec"
)

func main() {
	fmt.Println("Datos obtenidos desde el Modulo: ")
	fmt.Println("")

	cmd := exec.Command("sh", "-c", "cat /proc/ram_201903850")
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
	}

	output := string(out[:])
	fmt.Println(output)
	PTMADRE
}
