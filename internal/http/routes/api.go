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
		public.GET("mahasiswa/nim/:nim", handlers.MahasiswaHandler.GetData)
		public.GET("user/:uuid", handlers.AdminHandler.GetUser)
		public.GET("mahasiswa", handlers.MahasiswaHandler.GetAll)
		public.GET("percepatan", handlers.MahasiswaHandler.GetMahasiswaPercepatan)
		public.GET("percepatan/prodi/:uuid", handlers.MahasiswaHandler.GetProdiPercepatan)
		public.GET("prodi", handlers.AdminHandler.GetAllProdi)

	}

	admin := router.Group("api").Use(middleware.IsValidJWT())
	{
		admin.GET("mahasiswa/user/:userUuid", handlers.MahasiswaHandler.GetMahasiswaByUserUuid)
		admin.GET("mahasiswa/prodi/:userUuid", handlers.MahasiswaHandler.GetMahasiswaProdi)
		admin.GET("mahasiswa/penasihat/:userUuid", handlers.MahasiswaHandler.GetMahasiswaByPenasihat)
		admin.GET("dashboard/penasihat/:userUuid", handlers.AdminHandler.GetPenasihatDashboard)
		admin.GET("dashboard/kaprodi/:userUuid", handlers.AdminHandler.GetKaprodiDashboard)
		admin.GET("dashboard/kajur", handlers.AdminHandler.GetKajurDashboard)

		admin.POST("mahasiswa/import/:userUuid", handlers.MahasiswaHandler.Import)
		admin.GET("mahasiswa/:uuid", handlers.MahasiswaHandler.Get)
		admin.PUT("mahasiswa/:uuid", handlers.MahasiswaHandler.Update)
		admin.DELETE("mahasiswa/:uuid", handlers.MahasiswaHandler.Delete)
		admin.DELETE("mahasiswa/prodi/:userUuid", handlers.MahasiswaHandler.DeleteAll)

		admin.POST("pembimbing", handlers.AdminHandler.CreatePembimbing)
		admin.GET("pembimbing", handlers.AdminHandler.GetAllPembimbing)
		admin.GET("pembimbing/prodi/:userUuid", handlers.AdminHandler.GetPembimbingProdi)
		admin.GET("pembimbing/:uuid", handlers.AdminHandler.GetPembimbing)
		admin.PUT("pembimbing/:uuid", handlers.AdminHandler.UpdatePembimbing)
		admin.DELETE("pembimbing/:uuid", handlers.AdminHandler.DeletePembimbing)

		admin.POST("prodi", handlers.AdminHandler.CreateProdi)
		admin.GET("prodi/:uuid", handlers.AdminHandler.GetProdi)
		admin.PUT("prodi/:uuid", handlers.AdminHandler.UpdateProdi)
		admin.DELETE("prodi/:uuid", handlers.AdminHandler.DeleteProdi)

		admin.POST("mahasiswa/kelas", handlers.AdminHandler.UpdateKelas)
		admin.GET("classes", handlers.AdminHandler.GetClasses)
		admin.GET("pengaturan", handlers.AdminHandler.GetPengaturan)
		admin.GET("pengaturan/:name", handlers.AdminHandler.GetPengaturanByName)
		admin.PUT("pengaturan", handlers.MahasiswaHandler.UpdatePengaturan)

		admin.GET("users", handlers.AdminHandler.GetAllUsers)
		admin.POST("kajur", handlers.AdminHandler.CreateKajur)
		admin.PATCH("user/:uuid/username", handlers.AdminHandler.UpdateUsername)
		admin.PATCH("user/:uuid/password", handlers.AdminHandler.UpdatePassword)
		admin.PATCH("mahasiswa/rekomendasi", handlers.AdminHandler.RekomendasiMahasiswa)
	}

	return router
}
