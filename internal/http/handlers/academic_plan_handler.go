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

type AcademicPlanHandler struct {
	Service interfaces.AcademicPlanServiceInterface
}

func NewAcademicPlanHandler(service interfaces.AcademicPlanServiceInterface) interfaces.AcademicPlanHandlerInterface {
	return &AcademicPlanHandler{
		Service: service,
	}
}

func (h *AcademicPlanHandler) CreateAcademicPlan(c *gin.Context) {
	var body request.AcademicPlan
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

	if err := h.Service.CreateAcademicPlan(userUuid, &body); err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, response.SUCCESS_RES("RPS Berhasil Ditambahkan"))
}

func (h *AcademicPlanHandler) GetAllAcademicPlans(c *gin.Context) {
	userUuid := c.GetString("uuid")
	if userUuid == "" {
		utils.HandleError(c, response.HANDLER_INTERR)
		return
	}

	result, err := h.Service.GetAllAcademicPlans(userUuid)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	var resp []*response.AcademicPlan
	for _, item := range *result {
		resp = append(resp, &response.AcademicPlan{
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
			CreatedAt: item.CreatedAt,
			UpdatedAt: item.UpdatedAt,
		})
	}

	c.JSON(http.StatusOK, response.DATA_RES(resp))
}

func (h *AcademicPlanHandler) GetAcademicPlan(c *gin.Context) {
	uuid := c.Param("uuid")
	userUuid := c.GetString("uuid")
	if userUuid == "" {
		utils.HandleError(c, response.HANDLER_INTERR)
		return
	}
	result, err := h.Service.GetAcademicPlan(userUuid, uuid)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	resp := &response.AcademicPlan{
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
		CreatedAt: result.CreatedAt,
		UpdatedAt: result.UpdatedAt,
	}

	c.JSON(http.StatusOK, response.DATA_RES(resp))
}

func (h *AcademicPlanHandler) UpdateAcademicPlan(c *gin.Context) {
	var body request.UpdateAcademicPlan
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
	if err := h.Service.UpdateAcademicPlan(userUuid, uuid, &body); err != nil {
		utils.HandleError(c, err)
		return
	}
	c.JSON(http.StatusOK, response.SUCCESS_RES("RPS Berhasil Diperbarui"))
}

func (h *AcademicPlanHandler) DeleteAcademicPlan(c *gin.Context) {
	userUuid := c.GetString("uuid")
	if userUuid == "" {
		utils.HandleError(c, response.HANDLER_INTERR)
		return
	}
	
	uuid := c.Param("uuid")
	if err := h.Service.DeleteAcademicPlan(userUuid, uuid); err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.SUCCESS_RES("RPS Berhasil Dihapus"))
}
