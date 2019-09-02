package main

import (
	"anonbot/configuration"
	"anonbot/database"
	"anonbot/handlers"
	"anonbot/repositories"
	"anonbot/routing"
	"context"
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
	"log"
	"net/http"
)

func registerHandlers(db *sql.DB, router routing.Router, parent context.Context) {
	messageRepository := repositories.NewMessageRepository(db)

	handlers.RegisterMessageHandlers(messageRepository, router, parent)
}

func main() {
	config, err := configuration.LoadYamlConfig("config.yaml")
	if err != nil {
		log.Fatalln("Error opening config at path ./config/config.yaml")
	}

	router := routing.NewRouter("/api/v1")
	ctx, _ := context.WithCancel(context.Background())

	cs, err := config.GetConnectionString("postgres")
	if err != nil {
		log.Fatal(errors.Wrap(errors.WithStack(err), "Resolving postgres connection string"))
	}

	db := database.OpenSql(database.SqlTypePostgres, cs)
	defer func() {
		if err := db.Close(); err != nil {
			log.Fatalln("Closing db connection")
		}
	}()

	registerHandlers(db, router, ctx)


	log.Println("Starting server on port 8000")

	if port, err := config.GetAppSettingInt("port"); err != nil {
		log.Fatal(errors.Wrap(errors.WithStack(err), "Resolving app setting port"))
	} else {
		log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), router.AsHandler()))
	}
}
