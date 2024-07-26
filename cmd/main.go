package main

import (
	"fmt"

	"github.com/YerzatCode/learing_kz_backend/pkg/config"
	"github.com/YerzatCode/learing_kz_backend/pkg/controller"
	"github.com/YerzatCode/learing_kz_backend/pkg/database/sqlite"
	"github.com/gin-gonic/gin"
)

func init() {
	config.ConfigPathInit()
}
func main() {
	// INIT: Config
	cfg := config.InitConfig()

	// INIT: Database
	db := sqlite.InitSqlite(cfg.Database)

	// INIT: Controller
	r := gin.Default()
	controller.InitContoller(db, r)

	r.Run(fmt.Sprintf("%s:%s", cfg.Host, cfg.Port))
}
