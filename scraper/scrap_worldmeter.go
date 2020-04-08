package scraper

import (
	"regexp"
	"strings"

	"github.com/gocolly/colly"

	// "github.com/vaishvikmaisuria/CoronaVision/db"
	"github.com/vaishvikmaisuria/CoronaVision/models"
	"github.com/vaishvikmaisuria/CoronaVision/utils"
)

// ScrapWorldometer <struct>
// is used to present Secret Magazine crawler.
type ScrapWorldometer struct{}

const (
	scrapURLWM = "https://www.worldometers.info/coronavirus/"
)

// Run <function>
// is used to start crawling process.
func (wm ScrapWorldometer) Run() models.CoronaData {
	var totalData models.CoronaData
	coronaFuncs := wm.runCoronaScraping

	tmpData := coronaFuncs()
	totalData = tmpData

	return totalData
}

func (wm ScrapWorldometer) runCoronaScraping() models.CoronaData {

	// Header
	var totalCases string
	var totalDeaths string
	var totalRecovered string
	// Active Cases
	var currentlyCases string
	var mildCases string
	var criticalCases string
	// Closed Cases
	var outcomeCases string
	var RecoveredDischarged string

	// array of coronastates that will be returned
	var data models.CoronaData

	// table array to hold data
	var table []models.TableRow

	// Instantiate default collector colly instance without any options
	c := colly.NewCollector()

	// On every a element which has href attribute call callback
	c.OnHTML("Body", func(e *colly.HTMLElement) {
		e.ForEach("#maincounter-wrap", func(_ int, e *colly.HTMLElement) {
			link := e.Text
			if strings.Contains(link, "Coronavirus Cases:") {
				re := regexp.MustCompile("[0-9]+")
				totalCases = re.FindAllString(strings.Replace(link, ",", "", -1), -1)[0]
			}
			if strings.Contains(link, "Deaths:") {
				re := regexp.MustCompile("[0-9]+")
				totalDeaths = re.FindAllString(strings.Replace(link, ",", "", -1), -1)[0]
			}
			if strings.Contains(link, "Recovered:") {
				re := regexp.MustCompile("[0-9]+")
				totalRecovered = re.FindAllString(strings.Replace(link, ",", "", -1), -1)[0]
			}
		})

		e.ForEach(".panel_front", func(_ int, e *colly.HTMLElement) {
			link := e.Text
			res1 := strings.Split(link, "\n")

			if strings.Contains(res1[2], "Currently Infected Patients") {
				re := regexp.MustCompile("[0-9]+")
				currentlyCases = re.FindAllString(strings.Replace(res1[1], ",", "", -1), -1)[0]
			}
			if strings.Contains(res1[6], "Mild Condition") {
				// re := regexp.MustCompile("[0-9]+")
				// mildCases = re.FindAllString(strings.Replace(res1[5], ",", "", -1), -1)[0]
				mildCases = res1[5]
			}
			if strings.Contains(res1[9], "Serious or Critical") {
				// re := regexp.MustCompile("[0-9]+")
				// criticalCases = re.FindAllString(strings.Replace(res1[2], ",", "", -1), -1)[0]
				criticalCases = res1[8]
			}
			if strings.Contains(res1[2], "Cases which had an outcome") {
				re := regexp.MustCompile("[0-9]+")
				outcomeCases = re.FindAllString(strings.Replace(res1[1], ",", "", -1), -1)[0]
			}
			if strings.Contains(res1[7], "Recovered / Discharged") {
				// re := regexp.MustCompile("[0-9]+")
				// RecoveredDischarged = re.FindAllString(strings.Replace(res1[0], ",", "", -1), -1)[0]
				RecoveredDischarged = res1[5] + " " + res1[6]
			}
		})

		e.ForEach("#main_table_countries_today", func(_ int, maintable *colly.HTMLElement) {
			maintable.ForEach("tr", func(_ int, row *colly.HTMLElement) {
				newRow := models.TableRow{}
				row.ForEach("td", func(_ int, el *colly.HTMLElement) {
					switch el.Index {
					case 0:
						// Country
						newRow.Country = el.Text
					case 1:
						// Total Cases
						newRow.TotalCase = el.Text
					case 2:
						// New Cases
						newRow.NewCase = el.Text
					case 3:
						// Total Deaths
						newRow.TotalDeaths = el.Text
					case 4:
						// New Deaths
						newRow.NewDeaths = el.Text
					case 5:
						// Total Recovered
						newRow.TotalRecovered = el.Text
					case 6:
						// Active Case
						newRow.ActiveCase = el.Text
					case 7:
						// Critical
						newRow.Critical = el.Text
					case 8:
						// total Case per 1 Million population
						newRow.TotCase = el.Text
					case 9:
						// total Death per 1 Million population
						newRow.TotDeath = el.Text
					case 10:
						// Date of First Case
						newRow.FirstCase = el.Text
					}
				})
				if newRow.Country != "" {
					// append every row to table
					table = append(table, newRow)
				}
			})
		})

		_id := utils.MakeHash(scrapURLWM) // here we are going to create hash from full link in order to set ID of a data to hash value, so mongo won't add in case of duplicates

		mainStates := models.CoronaStates{
			ID:                  _id,
			TotalCases:          totalCases,
			TotalDeaths:         totalDeaths,
			TotalRecovered:      totalRecovered,
			CurrentlyCases:      currentlyCases,
			MildCases:           mildCases,
			CriticalCases:       criticalCases,
			OutcomeCases:        outcomeCases,
			RecoveredDischarged: RecoveredDischarged,
		}

		fullTable := models.CountryTable{
			ID:        _id,
			Fulltable: table,
		}

		data = models.CoronaData{
			StateData: mainStates,
			TableData: fullTable,
		}

	})

	c.Visit(scrapURLWM)
	c.Wait()

	return data
}
