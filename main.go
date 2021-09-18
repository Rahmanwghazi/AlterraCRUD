package main

import (
	"github.com/Rahmanwghazi/AlterraCRUD/config"
	"github.com/Rahmanwghazi/AlterraCRUD/routes"
)

func main() {
	config.InitDB()
	e := routes.New()
	e.Start(":8000")
}
