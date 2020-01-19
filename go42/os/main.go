package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	go signalListen()
	for {
		time.Sleep(10 * time.Second)
	}

	env := os.Environ()
	procAttr := &os.ProcAttr{
		Env:   env,
		Files: []*os.File{
			os.Stdin,
			os.Stdout,
			os.Stderr,
		},
	}
	Pid, err := os.StartProcess("/bin/ls", []string{"ls", "-l"}, procAttr)
	if err != nil {
		fmt.Sprintf("Error %v starting process!", err)
		os.Exit(1)
	}
	fmt.Printf("The process id is %v", Pid)
}

func signalListen()  {
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGUSR2)
	for {
		s := <- c
		fmt.Printf("get signal:", s)
	}
}
