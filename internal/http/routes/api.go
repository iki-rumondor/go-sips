package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/iki-rumondor/sips/internal/config"
)

func StartServer(handlers *config.Handlers) *gin.Engine {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12,
	}))

	public := router.Group("api")
	{
		public.POST("signin", handlers.AdminHandler.SignIn)

	}

	// admin := router.Group("api").Use(middleware.IsValidJWT(), middleware.IsRole("ADMIN"))
	// {
	// 	admin.POST("majors", handlers.MajorHandler.CreateMajor)
	// 	admin.GET("majors", handlers.MajorHandler.GetAllMajors)
	// 	admin.GET("majors/:uuid", handlers.MajorHandler.GetMajor)
	// 	admin.PUT("majors/:uuid", handlers.MajorHandler.UpdateMajor)
	// 	admin.DELETE("majors/:uuid", handlers.MajorHandler.DeleteMajor)

	// 	admin.POST("departments", handlers.DepartmentHandler.CreateDepartment)
	// 	admin.GET("departments", handlers.DepartmentHandler.GetAllDepartments)
	// 	admin.GET("departments/:uuid", handlers.DepartmentHandler.GetDepartment)
	// 	admin.PUT("departments/:uuid", handlers.DepartmentHandler.UpdateDepartment)
	// 	admin.DELETE("departments/:uuid", handlers.DepartmentHandler.DeleteDepartment)

	// 	admin.POST("academic-years", handlers.AcademicYearHandler.CreateAcademicYear)
	// 	admin.PUT("academic-years/:uuid", handlers.AcademicYearHandler.UpdateAcademicYear)
	// 	admin.DELETE("academic-years/:uuid", handlers.AcademicYearHandler.DeleteAcademicYear)

	// }

	return router
}
