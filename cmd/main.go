package main

import (
	"fmt"
	"log"
	"os"

	"github.com/iki-rumondor/go-monev/internal/config"
	"github.com/iki-rumondor/go-monev/internal/http/routes"
	"github.com/iki-rumondor/go-monev/internal/migrate"
)

func main() {
	gormDB, err := config.NewMysqlDB()
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	if len(os.Args)-1 > 0 {
		migrate.ReadTerminal(gormDB, os.Args)
		return
	}

	handlers := config.GetAppHandlers(gormDB)

	var PORT = os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	routes.StartServer(handlers).Run(fmt.Sprintf(":%s", PORT))
}
