package handlers

import (
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/iki-rumondor/go-monev/internal/http/request"
	"github.com/iki-rumondor/go-monev/internal/http/response"
	"github.com/iki-rumondor/go-monev/internal/interfaces"
	"github.com/iki-rumondor/go-monev/internal/utils"
)

type DepartmentHandler struct {
	Service interfaces.DepartmentServiceInterface
}

func NewDepartmentHandler(service interfaces.DepartmentServiceInterface) interfaces.DepartmentHandlerInterface {
	return &DepartmentHandler{
		Service: service,
	}
}

func (h *DepartmentHandler) CreateDepartment(c *gin.Context) {
	var body request.Department
	if err := c.BindJSON(&body); err != nil {
		utils.HandleError(c, response.BADREQ_ERR(err.Error()))
		return
	}

	if _, err := govalidator.ValidateStruct(&body); err != nil {
		utils.HandleError(c, response.BADREQ_ERR(err.Error()))
		return
	}

	if err := h.Service.CreateDepartment(&body); err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, response.SUCCESS_RES("Program Studi Berhasil Ditambahkan"))
}

func (h *DepartmentHandler) GetAllDepartments(c *gin.Context) {
	departments, err := h.Service.GetAllDepartments()
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	var resp []*response.Department
	for _, item := range *departments {
		resp = append(resp, &response.Department{
			Uuid: item.Uuid,
			Name: item.Name,
			Head: item.Head,
			User: &response.User{
				Uuid:     item.User.Uuid,
				Username: item.User.Username,
				Role:     item.User.Role.Name,
			},
			Major: &response.Major{
				Uuid: item.Major.Uuid,
				Name: item.Major.Name,
			},
			CreatedAt: item.CreatedAt,
			UpdatedAt: item.UpdatedAt,
		})
	}

	c.JSON(http.StatusOK, response.DATA_RES(resp))
}

func (h *DepartmentHandler) GetDepartment(c *gin.Context) {
	uuid := c.Param("uuid")
	department, err := h.Service.GetDepartment(uuid)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	resp := &response.Department{
		Uuid: department.Uuid,
		Name: department.Name,
		Head: department.Head,
		User: &response.User{
			Uuid:     department.User.Uuid,
			Username: department.User.Username,
			Role:     department.User.Role.Name,
		},
		Major: &response.Major{
			Uuid: department.Major.Uuid,
			Name: department.Major.Name,
		},
		CreatedAt: department.CreatedAt,
		UpdatedAt: department.UpdatedAt,
	}

	c.JSON(http.StatusOK, response.DATA_RES(resp))
}

func (h *DepartmentHandler) UpdateDepartment(c *gin.Context) {
	var body request.Department
	if err := c.BindJSON(&body); err != nil {
		utils.HandleError(c, &response.Error{
			Code:    400,
			Message: err.Error(),
		})
		return
	}

	if _, err := govalidator.ValidateStruct(&body); err != nil {
		utils.HandleError(c, response.BADREQ_ERR(err.Error()))
		return
	}

	uuid := c.Param("uuid")
	if err := h.Service.UpdateDepartment(uuid, &body); err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.SUCCESS_RES("Program Studi Berhasil Diperbarui"))
}

func (h *DepartmentHandler) DeleteDepartment(c *gin.Context) {
	uuid := c.Param("uuid")
	if err := h.Service.DeleteDepartment(uuid); err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.SUCCESS_RES("Program Studi Berhasil Dihapus"))
}
