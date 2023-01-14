package convert

import (
	"fmt"
	"log"
	"os/exec"
)

func CheckBinExist(bin string) bool {
	path, err := exec.LookPath(bin)
	if err != nil {
		log.Panicf("%s DOES NOT EXIST ON THIS NODE!", bin)
		return false
	} else {
		fmt.Printf("%s executable is in '%s'\n", bin, path)
		return true
	}
}
