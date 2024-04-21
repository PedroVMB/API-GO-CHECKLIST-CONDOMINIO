package routes

import (
	"github.com/PedroVMB/API-GO-CHECKLIST-CONDOMINIO/controllers"
	"github.com/gin-gonic/gin"
)


func HandleRequests() {
	r := gin.Default()
	///Administrador routes
	r.GET("/administradores", controllers.GetAllAdm)
	r.GET("/administradores/:id", controllers.GetAdmById)
	r.POST("/administradores", controllers.CreateAdm)
	r.DELETE("/administradores/:id", controllers.DeleteAdm)
	r.PATCH("/administradores/:id", controllers.UpdateAdm)

	//Sindico routes

	r.Run(":8000")
}