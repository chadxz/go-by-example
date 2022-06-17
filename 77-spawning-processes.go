package main

import (
	"fmt"
	"io"
	"os/exec"
)

func main() {
	dateOut, err := exec.Command("date").Output()
	if err != nil {
		panic(err)
	}

	fmt.Println("> date")
	fmt.Println(string(dateOut))

	_, err = exec.Command("date", "-x").Output()
	if err != nil {
		switch e := err.(type) {
		case *exec.Error:
			fmt.Println("failed executing:", err)
		case *exec.ExitError:
			fmt.Println("command exit rc =", e.ExitCode())
		default:
			panic(err)
		}
	}

	grepCmd := exec.Command("grep", "hello")
	grepIn, _ := grepCmd.StdinPipe()
	grepOut, _ := grepCmd.StdoutPipe()
	_ = grepCmd.Start()
	_, _ = grepIn.Write([]byte("hello grep\ngoodbye grep"))
	_ = grepIn.Close()
	grepBytes, _ := io.ReadAll(grepOut)
	_ = grepCmd.Wait()

	fmt.Println("> grep hello")
	fmt.Println(string(grepBytes))

	lsOut, err := exec.Command("bash", "-c", "ls -a -l -h").Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("> ls -a -l -h")
	fmt.Println(string(lsOut))
}
