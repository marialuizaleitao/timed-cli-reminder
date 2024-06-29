package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

const reminderFile = "reminders.txt"

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: reminder [add|list]")
		return
	}

	command := os.Args[1]
	switch command {
	case "add":
		addReminder()
	case "list":
		listReminders()
	default:
		fmt.Println("Unknown command:", command)
	}
}

func addReminder() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: reminder add [your reminder]")
		return
	}

	reminder := strings.Join(os.Args[2:], " ")
	file, err := os.OpenFile(reminderFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	timestamp := time.Now().Format("2006-01-02 15:04:05")
	entry := fmt.Sprintf("%s - %s\n", timestamp, reminder)
	if _, err := file.WriteString(entry); err != nil {
		fmt.Println("Error writing to file:", err)
	}
	fmt.Println("Reminder added!")
}

func listReminders() {
	file, err := os.Open(reminderFile)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}
