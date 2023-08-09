package main

import (
	"saichudin/parto-test/config"
	"saichudin/parto-test/migration"
	"saichudin/parto-test/router"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	config.InitDb()
	migration.Migrate()

	r := gin.Default()
	router.Route(r)
	r.Run()
}
