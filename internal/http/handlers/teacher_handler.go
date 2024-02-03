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

type TeacherHandler struct {
	Service interfaces.TeacherServiceInterface
}

func NewTeacherHandler(service interfaces.TeacherServiceInterface) interfaces.TeacherHandlerInterface {
	return &TeacherHandler{
		Service: service,
	}
}

func (h *TeacherHandler) CreateTeacher(c *gin.Context) {
	var body request.Teacher
	if err := c.BindJSON(&body); err != nil {
		utils.HandleError(c, response.BADREQ_ERR(err.Error()))
		return
	}

	if _, err := govalidator.ValidateStruct(&body); err != nil {
		utils.HandleError(c, response.BADREQ_ERR(err.Error()))
		return
	}

	userUuid := c.GetString("uuid")
	if userUuid == "" {
		utils.HandleError(c, response.HANDLER_INTERR)
		return
	}

	if err := h.Service.CreateTeacher(userUuid, &body); err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, response.SUCCESS_RES("Laboratorium Berhasil Ditambahkan"))
}

func (h *TeacherHandler) GetAllTeachers(c *gin.Context) {
	userUuid := c.GetString("uuid")
	if userUuid == "" {
		utils.HandleError(c, response.HANDLER_INTERR)
		return
	}

	result, err := h.Service.GetAllTeachers(userUuid)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	var resp []*response.Teacher
	for _, item := range *result {
		resp = append(resp, &response.Teacher{
			Uuid: item.Uuid,
			Name: item.Name,
			Department: &response.Department{
				Uuid: item.Department.Uuid,
				Name: item.Department.Name,
				Head: item.Department.Head,
			},
			CreatedAt: item.CreatedAt,
			UpdatedAt: item.UpdatedAt,
		})
	}

	c.JSON(http.StatusOK, response.DATA_RES(resp))
}

func (h *TeacherHandler) GetTeacher(c *gin.Context) {
	uuid := c.Param("uuid")
	userUuid := c.GetString("uuid")
	if userUuid == "" {
		utils.HandleError(c, response.HANDLER_INTERR)
		return
	}

	result, err := h.Service.GetTeacher(userUuid, uuid)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	resp := &response.Teacher{
		Uuid: result.Uuid,
		Name: result.Name,
		Department: &response.Department{
			Uuid: result.Department.Uuid,
			Name: result.Department.Name,
			Head: result.Department.Head,
		},
		CreatedAt: result.CreatedAt,
		UpdatedAt: result.UpdatedAt,
	}

	c.JSON(http.StatusOK, response.DATA_RES(resp))
}

func (h *TeacherHandler) UpdateTeacher(c *gin.Context) {
	var body request.Teacher
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

	userUuid := c.GetString("uuid")
	if userUuid == "" {
		utils.HandleError(c, response.HANDLER_INTERR)
		return
	}

	uuid := c.Param("uuid")
	if err := h.Service.UpdateTeacher(userUuid, uuid, &body); err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.SUCCESS_RES("Laboratorium Berhasil Diperbarui"))
}

func (h *TeacherHandler) DeleteTeacher(c *gin.Context) {
	userUuid := c.GetString("uuid")
	if userUuid == "" {
		utils.HandleError(c, response.HANDLER_INTERR)
		return
	}

	uuid := c.Param("uuid")
	if err := h.Service.DeleteTeacher(userUuid, uuid); err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.SUCCESS_RES("Laboratorium Berhasil Dihapus"))
}
