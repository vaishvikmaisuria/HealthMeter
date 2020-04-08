package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vaishvikmaisuria/CoronaVision/db"
)

// DataController <controller>
// is used for describing controller actions for data.
type DataController struct{}

// GetData <function>
// is used to handle get action of data controller which will return <count> number of coronavirus data.
// url: /v1/data/GetData?count=80 , by default <count> = 50
func (nc DataController) GetData(c *gin.Context) {
	countStr := c.DefaultQuery("count", "1")
	count, err := strconv.Atoi(countStr)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "Please, enter corrent number of count parameter.",
		})
		return
	}

	dbClient := db.GetClient()
	data, err := dbClient.DataGet(int64(count))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error":       "Problem with fetching data from DB",
			"description": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, data)
}

// GetTable <function>
// is used to handle get action of data controller which will return <count> number of coronavirus data.
// url: /v1/data/GetTable , by default <count> = 1
func (nc DataController) GetTable(c *gin.Context) {

	dbClient := db.GetClient()
	data, err := dbClient.CTableGet()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error":       "Problem with fetching data from DB",
			"description": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, data)
}

// GetTest <function>
// is used to handle get action of news controller which will return <count> number of news.
// url: /v1/news?count=80 , by default <count> = 50
func (nc DataController) GetTest(c *gin.Context) {
	c.JSON(200, gin.H{
		"method":  "v1/news",
		"message": "Hello from Get function!",
	})
}
