package main

import (
	"fmt"
	"io/ioutil"
	"os/exec"
)

func main() {
	// date command
	dateCmd := exec.Command("date")

	dateOut, err := dateCmd.Output()
	if err != nil {
		panic(err)
	}

	fmt.Println("> date")
	fmt.Println(string(dateOut))

	// grep command
	grepCmd := exec.Command("grep", "hello")
	grepIn, _ := grepCmd.StdinPipe()
	grepOut, _ := grepCmd.StdoutPipe()
	grepCmd.Start()
	grepIn.Write([]byte("hello grep\ngoodbye grep"))
	grepIn.Close()
	grepBytes, _ := ioutil.ReadAll(grepOut)
	grepCmd.Wait()

	fmt.Println("> grep hello")
	fmt.Println(string(grepBytes))

	// bash command
	lsCmd := exec.Command("bash", "-c", "ls -a -l -h")
	lsOut, err := lsCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("> ls -a -l -h")
	fmt.Println(string(lsOut))

}
