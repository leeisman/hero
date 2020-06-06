package main

import (
	"hero/cmd"
)

func main() {
	//建立table
	cmd.AutoMigrate()
	//建立預設資料
	cmd.SeedRun()
	//command初始
	cmd.Execute()
}
