package controllers

import (
	"net/http"

	"github.com/PedroVMB/API-GO-CHECKLIST-CONDOMINIO/database"
	"github.com/PedroVMB/API-GO-CHECKLIST-CONDOMINIO/models"
	"github.com/gin-gonic/gin"
)

func GetAllAdm(c *gin.Context){
	var adms []models.Administrador
	database.DB.Find(&adms)
	c.JSON(200, adms)
}

// func GetAdmByInativeStatus(c *gin.Context){
// 	var adms []models.Administrador
// 	database.DB.Raw("SELECT * from administradores WHERE ").Find(&adms)
// 	c.JSON(http.StatusOK, adms)
// }

func GetAdmById(c *gin.Context){
	var adm models.Administrador
	id := c.Params.ByName("id")
	database.DB.First(&adm, id)
	if adm.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not Found": "Administrador not found"})
		return
	}

	c.JSON(http.StatusOK, adm)
}

func CreateAdm(c *gin.Context){
	var adm models.Administrador
	if err := c.ShouldBindJSON(&adm); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"erro": err.Error()})
			return
	}
	if err := models.ValidatesDataAdm(&adm); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"erro": err.Error()})
		return
	}

	database.DB.Create(&adm)
	c.JSON(http.StatusOK, adm)
}

func UpdateAdm(c *gin.Context){
	var adm models.Administrador
	id := c.Params.ByName("id")
	database.DB.First(&adm, id)
	if err := c.ShouldBindJSON(&adm); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"erro": err.Error()})
		return
	}
	if err := models.ValidatesDataAdm(&adm); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"erro": err.Error()})
		return
	}

	database.DB.Save(&adm)
	c.JSON(http.StatusOK, adm)

}

func DeleteAdm(c *gin.Context){
	var adm models.Administrador
	id := c.Params.ByName("id")
	database.DB.Delete(&adm, id)
	c.JSON(http.StatusOK, gin.H{"data": "Admnistrador successfully deleted"})
}
