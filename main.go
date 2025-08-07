package main

import (
	"flag"
)

func main() {
    var folder string
    var email string
    flag.StringVar(&folder, "folder", "", "add a new folder to scan for Git repositories")
    flag.StringVar(&email, "email", "your@email.com", "the email to scan")
    flag.Parse()

    var repositories []string
    if folder != "" {
        repositories = scan(folder)
    }

    stats(email, repositories)
}