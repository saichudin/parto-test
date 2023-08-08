package main

import (
	"saichudin/parto-test/config"
	"saichudin/parto-test/migration"
	"saichudin/parto-test/router"

	"github.com/gin-gonic/gin"
)

func main() {
	config.InitDb()
	migration.Migrate()

	r := gin.Default()
	router.Route(r)
	r.Run()
}
