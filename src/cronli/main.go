package main

import (
	"github.com/robfig/cron"
	"os"
	"os/exec"
	"strings"
	"time"
	"fmt"
)

func main() {
	c := cron.New()
	if len(os.Args) != 3 {
		println("Please provide 2 args: schedule and bash command")
		os.Exit(1)
	}
	schedule := os.Args[1]
	command := []string{"/usr/bin/env", "sh", "-c", os.Args[2]}
	c.AddFunc(schedule, func() {
		cmd := exec.Command(command[0], command[1:]...)
		cmd.Stderr = os.Stderr
		cmd.Stdout = os.Stdout
		cmd.Run()
	})
	fmt.Println("Schedule:", schedule)
	fmt.Println("Command:", strings.Join(command, " "))
	firstRuns := []string{}
	next := time.Now().Local()
	fmt.Println("Current time:", next.String())
	for i := 0; i < 5; i++ {
		next = c.Entries()[0].Schedule.Next(next)
		firstRuns = append(firstRuns, next.String())
	}
	fmt.Println("First Runs:", strings.Join(firstRuns, ", "))
	fmt.Println("Started.")
	fmt.Println("********")
	c.Start()
	select {}
}
