package scraper

import (
	"fmt"
	"time"

	"github.com/vaishvikmaisuria/CoronaVision/db"
	"github.com/vaishvikmaisuria/CoronaVision/models"
)

// CoronaScraper <interface>
// is used to describe CoronaScraper class instance
type CoronaScraper interface {
	Run() []models.CoronaData
}

// CoronaFunc <type>
// is used to simplify Corona func type signature
type CoronaFunc func() []models.CoronaData

// Start <function>
// is used to start process of web resources crawling every 3 minutes
func Start() {
	go startScraper()
}

func startScraper() {
	// array of scraper for different coronavirus sources which implement CoronaScraper interface.
	scraper := ScrapWorldometer{}

	// other scrapers

	// duration of each crawling process
	duration := time.Minute * 3
	fmt.Println("TOTAl timee")
	for range time.Tick(duration) {
		// all states collected from each crawler
		// totalData := scraper.Run()
		fmt.Println("TOTAL work")
		// next step: adding totalData to db
		// all news collected from each crawler
		var totalData []models.CoronaStates
		var totalTable []models.CountryTable

		data := scraper.Run()

		totalData = append(totalData, data.StateData)
		totalTable = append(totalTable, data.TableData)

		dbClient := db.GetClient()
		dbClient.DataInsert(totalData)   // here it should show error but it will be ignored by mongo and it will continue to write
		dbClient.TableInsert(totalTable) // here it should show error but it will be ignored by mongo and it will continue to write
	}
}
