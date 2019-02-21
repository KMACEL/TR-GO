package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
)

func main() {

	//#################################################################################
	//Örnek1
	//#################################################################################
	t, err := exec.Command("ls", "-la").CombinedOutput()

	if err != nil {
		os.Stderr.WriteString(err.Error())
	}
	fmt.Println(string(t))

	//#################################################################################
	//Örnek2
	//#################################################################################

	cmd := exec.Command("bash", "-c",
		"cd xx;"+
			"source yyy;"+
			"cd ..;"+
			"echo $PATH ;")
	cmd.Dir = "project/test"
	f, _ := os.Create("Test.log")
	multiWriter := io.MultiWriter(os.Stdout, f)
	cmd.Stdout = multiWriter
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Run()
}
