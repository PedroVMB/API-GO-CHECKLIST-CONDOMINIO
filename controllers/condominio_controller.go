package controllers

import (
	"net/http"

	"github.com/PedroVMB/API-GO-CHECKLIST-CONDOMINIO/database"
	"github.com/PedroVMB/API-GO-CHECKLIST-CONDOMINIO/models"
	"github.com/gin-gonic/gin"
)

func GetAllCondominios(c *gin.Context) {
	var condominios []models.Condominio
	database.DB.Raw("SELECT DISTINCT nome, cnpj, estado, cidade, bairro, numero, cep, quantidade_torres from condominios WHERE deleted_at IS NULL").Find(&condominios)
	c.JSON(http.StatusOK, condominios)
}

func GetCondominioById(c *gin.Context) {
	var condominio models.Condominio
	id := c.Params.ByName("id")
	database.DB.First(&condominio, id)
	if condominio.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not Found": "Condominio not found"})
		return
	}

	c.JSON(http.StatusOK, condominio)
}

func CreateCondominio(c *gin.Context) {
	var condominio models.Condominio
	if err := c.ShouldBindJSON(&condominio); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"erro": err.Error()})
		return
	}
	if err := models.ValidatesDataCondominio(&condominio); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"erro": err.Error()})
		return
	}

	database.DB.Create(&condominio)
	c.JSON(http.StatusOK, condominio)
}

func UpdateCondominio(c *gin.Context) {
	var condominio models.Condominio
	id := c.Params.ByName("id")
	database.DB.First(&condominio, id)
	if err := c.ShouldBindJSON(&condominio); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"erro": err.Error()})
		return
	}
	if err := models.ValidatesDataCondominio(&condominio); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"erro": err.Error()})
		return
	}

	database.DB.Save(&condominio)
	c.JSON(http.StatusOK, condominio)

}

func DeleteCondominio(c *gin.Context) {
	var adm models.Condominio
	id := c.Params.ByName("id")
	database.DB.Delete(&adm, id)
	c.JSON(http.StatusOK, gin.H{"data": "Condominio successfully deleted"})
}
