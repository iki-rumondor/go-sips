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

type LaboratoryHandler struct {
	Service interfaces.LaboratoryServiceInterface
}

func NewLaboratoryHandler(service interfaces.LaboratoryServiceInterface) interfaces.LaboratoryHandlerInterface {
	return &LaboratoryHandler{
		Service: service,
	}
}

func (h *LaboratoryHandler) CreateLaboratory(c *gin.Context) {
	var body request.Laboratory
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

	if err := h.Service.CreateLaboratory(userUuid, &body); err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, response.SUCCESS_RES("Laboratorium Berhasil Ditambahkan"))
}

func (h *LaboratoryHandler) GetAllLaboratories(c *gin.Context) {
	userUuid := c.GetString("uuid")
	if userUuid == "" {
		utils.HandleError(c, response.HANDLER_INTERR)
		return
	}

	result, err := h.Service.GetAllLaboratories(userUuid)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	var resp []*response.Laboratory
	for _, item := range *result {
		resp = append(resp, &response.Laboratory{
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

func (h *LaboratoryHandler) GetLaboratory(c *gin.Context) {
	uuid := c.Param("uuid")
	userUuid := c.GetString("uuid")
	if userUuid == "" {
		utils.HandleError(c, response.HANDLER_INTERR)
		return
	}

	result, err := h.Service.GetLaboratory(userUuid, uuid)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	resp := &response.Laboratory{
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

func (h *LaboratoryHandler) UpdateLaboratory(c *gin.Context) {
	var body request.Laboratory
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
	if err := h.Service.UpdateLaboratory(userUuid, uuid, &body); err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.SUCCESS_RES("Laboratorium Berhasil Diperbarui"))
}

func (h *LaboratoryHandler) DeleteLaboratory(c *gin.Context) {
	userUuid := c.GetString("uuid")
	if userUuid == "" {
		utils.HandleError(c, response.HANDLER_INTERR)
		return
	}

	uuid := c.Param("uuid")
	if err := h.Service.DeleteLaboratory(userUuid, uuid); err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.SUCCESS_RES("Laboratorium Berhasil Dihapus"))
}
