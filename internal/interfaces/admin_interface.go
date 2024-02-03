package interfaces

import (
	"github.com/gin-gonic/gin"
	"github.com/iki-rumondor/sips/internal/http/request"
	"github.com/iki-rumondor/sips/internal/models"
)

type AdminHandlerInterface interface {
	SignIn(*gin.Context)
}

type AdminServiceInterface interface {
	VerifyAdmin(*request.SignIn) (string, error)
}

type AdminRepoInterface interface {
	FindAdminBy(column string, value interface{}) (*models.Admin, error)
}
