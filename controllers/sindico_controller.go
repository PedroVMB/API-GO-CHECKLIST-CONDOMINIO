package controllers

import (
	"net/http"

	"github.com/PedroVMB/API-GO-CHECKLIST-CONDOMINIO/database"
	"github.com/PedroVMB/API-GO-CHECKLIST-CONDOMINIO/models"
	"github.com/gin-gonic/gin"
)

func GetAllSindico(c *gin.Context) {
	var sindicos []models.Sindico
	database.DB.Raw("SELECT name, email, cpf from sindicos WHERE deleted_at IS NULL").Find(&sindicos)
	c.JSON(http.StatusOK, sindicos)
}

// func GetSindicoByInativeStatus(c *gin.Context){
// 	var adms []models.Administrador
// 	database.DB.Raw("SELECT * from administradors WHERE deleted_at != null").Find(&adms)
// 	c.JSON(http.StatusOK, adms)
// }

func GetSindicoById(c *gin.Context) {
	var sindico models.Sindico
	id := c.Params.ByName("id")
	database.DB.First(&sindico, id)
	if sindico.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not Found": "Sindico not found"})
		return
	}

	c.JSON(http.StatusOK, sindico)
}

func CreateSindico(c *gin.Context) {
	var sindico models.Sindico
	if err := c.ShouldBindJSON(&sindico); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"erro": err.Error()})
		return
	}
	if err := models.ValidatesDataSindico(&sindico); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"erro": err.Error()})
		return
	}

	database.DB.Create(&sindico)
	c.JSON(http.StatusOK, sindico)
}

func UpdateSindico(c *gin.Context) {
	var sindico models.Sindico
	id := c.Params.ByName("id")
	database.DB.First(&sindico, id)
	if err := c.ShouldBindJSON(&sindico); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"erro": err.Error()})
		return
	}
	if err := models.ValidatesDataSindico(&sindico); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"erro": err.Error()})
		return
	}

	database.DB.Save(&sindico)
	c.JSON(http.StatusOK, sindico)

}

func DeleteSindico(c *gin.Context) {
	var sindico models.Sindico
	id := c.Params.ByName("id")
	database.DB.Delete(&sindico, id)
	c.JSON(http.StatusOK, gin.H{"data": "Sindico successfully deleted"})
}
