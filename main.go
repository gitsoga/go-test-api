package main

import (
  "flag"
  
	_ "github.com/jinzhu/gorm/dialects/mysql"
  "github.com/gitsoga/go-test-api/config"
  "github.com/gitsoga/go-test-api/database"
  "github.com/gitsoga/go-test-api/server"
)

func main() {
  env := flag.String("e", "development", "")
  flag.Parse()

  config.Init(*env)
  database.Init(false)
  defer database.Close()
  if err := server.Init(); err != nil {
      panic(err)
  }
}