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
	// r.GET("/administradores/inativos", controllers.GetAdmByInativeStatus)

	//Sindico routes
	r.GET("/sindicos", controllers.GetAllSindico)
	r.GET("/sindicos/:id", controllers.GetSindicoById)
	r.POST("/sindicos", controllers.CreateSindico)
	r.DELETE("/sindicos/:id", controllers.DeleteSindico)
	r.PATCH("/sindicos/:id", controllers.UpdateSindico)

	r.Run(":5000")
}
