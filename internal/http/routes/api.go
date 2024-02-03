package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/iki-rumondor/go-monev/internal/config"
	"github.com/iki-rumondor/go-monev/internal/http/middleware"
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
		public.POST("signin", handlers.UserHandler.SignIn)

	}

	user := router.Group("api").Use(middleware.IsValidJWT())
	{
		user.GET("academic-years", handlers.AcademicYearHandler.GetAllAcademicYears)
		user.GET("academic-years/:uuid", handlers.AcademicYearHandler.GetAcademicYear)
		user.GET("dashboards/subjects", handlers.UserHandler.GetCountSubjects)
	}

	department := router.Group("api").Use(middleware.IsValidJWT(), middleware.IsRole("DEPARTMENT"), middleware.SetUserUuid())
	{
		department.POST("subjects", handlers.SubjectHandler.CreateSubject)
		department.GET("subjects", handlers.SubjectHandler.GetAllSubjects)
		department.GET("subjects/:uuid", handlers.SubjectHandler.GetSubject)
		department.PUT("subjects/:uuid", handlers.SubjectHandler.UpdateSubject)
		department.DELETE("subjects/:uuid", handlers.SubjectHandler.DeleteSubject)

		department.POST("laboratories", handlers.LaboratoryHandler.CreateLaboratory)
		department.GET("laboratories", handlers.LaboratoryHandler.GetAllLaboratories)
		department.GET("laboratories/:uuid", handlers.LaboratoryHandler.GetLaboratory)
		department.PUT("laboratories/:uuid", handlers.LaboratoryHandler.UpdateLaboratory)
		department.DELETE("laboratories/:uuid", handlers.LaboratoryHandler.DeleteLaboratory)

		department.POST("teachers", handlers.TeacherHandler.CreateTeacher)
		department.GET("teachers", handlers.TeacherHandler.GetAllTeachers)
		department.GET("teachers/:uuid", handlers.TeacherHandler.GetTeacher)
		department.PUT("teachers/:uuid", handlers.TeacherHandler.UpdateTeacher)
		department.DELETE("teachers/:uuid", handlers.TeacherHandler.DeleteTeacher)

		department.POST("academic-plans", handlers.AcademicPlanHandler.CreateAcademicPlan)
		department.GET("academic-plans", handlers.AcademicPlanHandler.GetAllAcademicPlans)
		department.GET("academic-plans/:uuid", handlers.AcademicPlanHandler.GetAcademicPlan)
		department.PUT("academic-plans/:uuid", handlers.AcademicPlanHandler.UpdateAcademicPlan)
		department.DELETE("academic-plans/:uuid", handlers.AcademicPlanHandler.DeleteAcademicPlan)

		department.POST("practical-modules", handlers.PracticalModuleHandler.CreatePracticalModule)
		department.GET("practical-modules", handlers.PracticalModuleHandler.GetAllPracticalModules)
		department.GET("practical-modules/:uuid", handlers.PracticalModuleHandler.GetPracticalModule)
		department.PUT("practical-modules/:uuid", handlers.PracticalModuleHandler.UpdatePracticalModule)
		department.DELETE("practical-modules/:uuid", handlers.PracticalModuleHandler.DeletePracticalModule)

		department.POST("practical-tools", handlers.PracticalToolHandler.CreatePracticalTool)
		department.GET("practical-tools", handlers.PracticalToolHandler.GetAllPracticalTools)
		department.GET("practical-tools/:uuid", handlers.PracticalToolHandler.GetPracticalTool)
		department.PUT("practical-tools/:uuid", handlers.PracticalToolHandler.UpdatePracticalTool)
		department.DELETE("practical-tools/:uuid", handlers.PracticalToolHandler.DeletePracticalTool)

		department.GET("subjects/practical", handlers.SubjectHandler.GetAllPracticalSubjects)

	}

	admin := router.Group("api").Use(middleware.IsValidJWT(), middleware.IsRole("ADMIN"))
	{
		admin.POST("majors", handlers.MajorHandler.CreateMajor)
		admin.GET("majors", handlers.MajorHandler.GetAllMajors)
		admin.GET("majors/:uuid", handlers.MajorHandler.GetMajor)
		admin.PUT("majors/:uuid", handlers.MajorHandler.UpdateMajor)
		admin.DELETE("majors/:uuid", handlers.MajorHandler.DeleteMajor)

		admin.POST("departments", handlers.DepartmentHandler.CreateDepartment)
		admin.GET("departments", handlers.DepartmentHandler.GetAllDepartments)
		admin.GET("departments/:uuid", handlers.DepartmentHandler.GetDepartment)
		admin.PUT("departments/:uuid", handlers.DepartmentHandler.UpdateDepartment)
		admin.DELETE("departments/:uuid", handlers.DepartmentHandler.DeleteDepartment)

		admin.POST("academic-years", handlers.AcademicYearHandler.CreateAcademicYear)
		admin.PUT("academic-years/:uuid", handlers.AcademicYearHandler.UpdateAcademicYear)
		admin.DELETE("academic-years/:uuid", handlers.AcademicYearHandler.DeleteAcademicYear)

	}

	return router
}
