package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/iki-rumondor/sips/internal/config"
	"github.com/iki-rumondor/sips/internal/http/middleware"
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

	admin := router.Group("api").Use(middleware.IsValidJWT())
	{
		admin.POST("mahasiswa/import", handlers.MahasiswaHandler.Import)
		admin.GET("mahasiswa", handlers.MahasiswaHandler.GetAll)
		admin.GET("mahasiswa/:uuid", handlers.MahasiswaHandler.Get)
		admin.PUT("mahasiswa/:uuid", handlers.MahasiswaHandler.Update)
		admin.DELETE("mahasiswa/:uuid", handlers.MahasiswaHandler.Delete)

		admin.POST("tahun_ajaran", handlers.TahunAjaranHandler.Create)
		admin.GET("tahun_ajaran", handlers.TahunAjaranHandler.GetAll)
		admin.GET("tahun_ajaran/:uuid", handlers.TahunAjaranHandler.Get)
		admin.PUT("tahun_ajaran/:uuid", handlers.TahunAjaranHandler.Update)
		admin.DELETE("tahun_ajaran/:uuid", handlers.TahunAjaranHandler.Delete)

		admin.GET("percepatan", handlers.AdminHandler.GetMahasiswaPercepatan)
		admin.POST("percepatan", handlers.AdminHandler.SetMahasiswaPercepatan)
	}

	return router
}
