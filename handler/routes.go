package handler

import (
	"dbaas/auth"
	"dbaas/db"
	"fmt"
	// "fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

var Table_Prefix = "gopgx_schema."

func GetHandler(c *gin.Context) {
	tableName := c.Param("table_name")
	// cndn := c.Param("cndn") // Optional: If you plan to use it later
	path := c.Request.URL.Path
	// fmt.Println(queries)

	queries := c.Request.URL.Query()
	result, err := db.Read(Table_Prefix+tableName, queries, path)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}
func Hi(c *gin.Context) {
	c.String(http.StatusOK, "hi")
}

func PostHandler(c *gin.Context) {
	var body map[string]interface{}

	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}
	table := c.Param("table_name")
	db.Insert(Table_Prefix+table, body)
	c.JSON(http.StatusCreated, gin.H{"status": "inserted"})
	return
}

func Create_Table(c *gin.Context) {
	var table_details map[string]string
	if err := c.BindJSON(&table_details); err != nil {
		c.JSON(http.StatusBadRequest, err)
	}
	fmt.Println(table_details)
	err := db.Create_Table(Table_Prefix+c.Param("table_name"), table_details)
	if err != nil {
		c.JSON(400, err)
		return
	}
}
func Delete_table(c *gin.Context) {
	var table_name string
	table_name = c.Param("table_name")
	err := db.Delete_table(table_name)
	if err != nil {
		c.JSON(400, err)
	}
}

func DeleteRowHandler(c *gin.Context) {
	table_name := c.Param("table_name")
	condition := c.Request.URL.Query()
	err := db.DeleteRow(Table_Prefix+table_name, condition)
	if err != nil {
		c.JSON(400, "Bad Request")
		return
	}

}

func UpdateTable(c *gin.Context) {
	table_name := c.Param("table_name")
	condition := c.Request.URL.Query()
	var body map[string]interface{}

	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}
	err := db.UpdateRow(Table_Prefix+table_name, condition, body)
	if err != nil {
		c.JSON(400, err)
	}

}

func NewAPIKey(c *gin.Context) {
	cookies, ok := auth.CheckAndVerifyCookies(c)
	if ok == false {
		c.String(400, "Baddd")

	}
	fmt.Println(cookies)
	apiKey, err := db.InsertAPI(cookies)
	if err != nil {
		c.JSON(400, fmt.Errorf("Bad Request"))
	}
	c.String(200, apiKey)

	// c.JSON(200, apiKey)
	// if err != nil {
	// 	c.AbortWithStatusJSON(400, "Baddd")
	// }
}
