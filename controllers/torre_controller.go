package controllers

import (
	"net/http"

	"github.com/PedroVMB/API-GO-CHECKLIST-CONDOMINIO/database"
	"github.com/PedroVMB/API-GO-CHECKLIST-CONDOMINIO/models"
	"github.com/gin-gonic/gin"
)

func GetAllTorres(c *gin.Context) {
	var torres []models.Torre
	database.DB.Raw("SELECT * FROM torres")
	c.JSON(http.StatusOK, torres)
}

func GetTorreById(c *gin.Context) {
	var torre models.Torre
	id := c.Params.ByName("id")
	database.DB.First(&torre, id)
	if torre.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not Found": "torre not found"})
		return
	}

	c.JSON(http.StatusOK, torre)
}

func CreateTorre(c *gin.Context) {
	var torre models.Torre
	if err := c.ShouldBindJSON(&torre); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"erro": err.Error()})
		return
	}
	if err := models.ValidatesDataTorres(&torre); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"erro": err.Error()})
		return
	}

	database.DB.Create(&torre)
	c.JSON(http.StatusOK, torre)
}

func UpdateTorre(c *gin.Context) {
	var torre models.Torre
	id := c.Params.ByName("id")
	database.DB.First(&torre, id)
	if err := c.ShouldBindJSON(&torre); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"erro": err.Error()})
		return
	}
	if err := models.ValidatesDataTorres(&torre); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"erro": err.Error()})
		return
	}

	database.DB.Save(&torre)
	c.JSON(http.StatusOK, torre)

}

func DeleteTorre(c *gin.Context) {
	var torre models.Torre
	id := c.Params.ByName("id")
	database.DB.Delete(&torre, id)
	c.JSON(http.StatusOK, gin.H{"data": "Torre successfully deleted"})
}
