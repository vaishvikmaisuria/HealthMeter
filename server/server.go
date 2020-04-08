package server

import (
	"github.com/vaishvikmaisuria/CoronaVision/db"
	"github.com/vaishvikmaisuria/CoronaVision/scraper"
	"github.com/vaishvikmaisuria/CoronaVision/utils"
)

// Init <function>
// is used to initialize server and all the corresponding services such as DB, Utils, Workers
func Init() {
	// utils
	utils.InitEnvVars()

	// services
	db.InitService()
	// db.GetClient().FillSeedsInformation() // a bit ugly but everything is going to work, only needed to fill seeds information

	// workers
	scraper.Start()

	r := NewRouter()
	// r.Run(":6969")
	// For Docker
	r.Run(":80")
}
