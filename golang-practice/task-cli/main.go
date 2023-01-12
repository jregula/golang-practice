/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"path/filepath"
	"task-cli/cmd"
	"task-cli/db"

	homedir "github.com/mitchellh/go-homedir"
)

func main() {
	home, _ := homedir.Dir()
	dbPath := filepath.Join(home, "tasks.db")
	err := db.Init(dbPath)
	if err != nil {
		panic(err)
	}
	cmd.Execute()
}
