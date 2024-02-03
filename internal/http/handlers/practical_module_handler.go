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

type PracticalModuleHandler struct {
	Service interfaces.PracticalModuleServiceInterface
}

func NewPracticalModuleHandler(service interfaces.PracticalModuleServiceInterface) interfaces.PracticalModuleHandlerInterface {
	return &PracticalModuleHandler{
		Service: service,
	}
}

func (h *PracticalModuleHandler) CreatePracticalModule(c *gin.Context) {
	var body request.PracticalModule
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

	if err := h.Service.CreatePracticalModule(userUuid, &body); err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, response.SUCCESS_RES("Modul Praktikum Berhasil Ditambahkan"))
}

func (h *PracticalModuleHandler) GetAllPracticalModules(c *gin.Context) {
	userUuid := c.GetString("uuid")
	if userUuid == "" {
		utils.HandleError(c, response.HANDLER_INTERR)
		return
	}

	result, err := h.Service.GetAllPracticalModules(userUuid)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	var resp []*response.PracticalModule
	for _, item := range *result {
		resp = append(resp, &response.PracticalModule{
			Uuid:      item.Uuid,
			Available: item.Available,
			Note:      item.Note,
			AcademicYear: &response.AcademicYear{
				Uuid: item.AcademicYear.Uuid,
				Name: item.AcademicYear.Name,
			},
			Subject: &response.Subject{
				Uuid: item.Subject.Uuid,
				Name: item.Subject.Name,
				Code: item.Subject.Code,
			},
			Laboratory: &response.Laboratory{
				Uuid: item.Laboratory.Uuid,
				Name: item.Laboratory.Name,
			},
			CreatedAt: item.CreatedAt,
			UpdatedAt: item.UpdatedAt,
		})
	}

	c.JSON(http.StatusOK, response.DATA_RES(resp))
}

func (h *PracticalModuleHandler) GetPracticalModule(c *gin.Context) {
	uuid := c.Param("uuid")
	userUuid := c.GetString("uuid")
	if userUuid == "" {
		utils.HandleError(c, response.HANDLER_INTERR)
		return
	}
	result, err := h.Service.GetPracticalModule(userUuid, uuid)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	resp := &response.PracticalModule{
		Uuid:      result.Uuid,
		Available: result.Available,
		Note:      result.Note,
		AcademicYear: &response.AcademicYear{
			Uuid: result.AcademicYear.Uuid,
			Name: result.AcademicYear.Name,
		},
		Subject: &response.Subject{
			Uuid: result.Subject.Uuid,
			Name: result.Subject.Name,
			Code: result.Subject.Code,
		},
		Laboratory: &response.Laboratory{
			Uuid: result.Laboratory.Uuid,
			Name: result.Laboratory.Name,
		},
		CreatedAt: result.CreatedAt,
		UpdatedAt: result.UpdatedAt,
	}

	c.JSON(http.StatusOK, response.DATA_RES(resp))
}

func (h *PracticalModuleHandler) UpdatePracticalModule(c *gin.Context) {
	var body request.UpdatePracticalModule
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
	if err := h.Service.UpdatePracticalModule(userUuid, uuid, &body); err != nil {
		utils.HandleError(c, err)
		return
	}
	c.JSON(http.StatusOK, response.SUCCESS_RES("Modul Praktikum Berhasil Diperbarui"))
}

func (h *PracticalModuleHandler) DeletePracticalModule(c *gin.Context) {
	userUuid := c.GetString("uuid")
	if userUuid == "" {
		utils.HandleError(c, response.HANDLER_INTERR)
		return
	}

	uuid := c.Param("uuid")
	if err := h.Service.DeletePracticalModule(userUuid, uuid); err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.SUCCESS_RES("Modul Praktikum Berhasil Dihapus"))
}
