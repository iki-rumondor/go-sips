package handlers

import (
	"fmt"
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/iki-rumondor/sips/internal/http/request"
	"github.com/iki-rumondor/sips/internal/http/response"
	"github.com/iki-rumondor/sips/internal/interfaces"
	"github.com/iki-rumondor/sips/internal/utils"
)

type TahunAjaranHandler struct {
	Service interfaces.TahunAjaranServiceInterface
}

func NewTahunAjaranHandler(service interfaces.TahunAjaranServiceInterface) interfaces.TahunAjaranHandlerInterface {
	return &TahunAjaranHandler{
		Service: service,
	}
}

func (h *TahunAjaranHandler) Create(c *gin.Context) {
	var body request.TahunAjaran
	if err := c.BindJSON(&body); err != nil {
		utils.HandleError(c, response.BADREQ_ERR(err.Error()))
		return
	}

	if _, err := govalidator.ValidateStruct(&body); err != nil {
		utils.HandleError(c, response.BADREQ_ERR(err.Error()))
		return
	}

	if err := h.Service.CreateTahunAjaran(&body); err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, response.SUCCESS_RES("Tahun Ajaran Berhasil Ditambahkan"))
}

func (h *TahunAjaranHandler) GetAll(c *gin.Context) {
	result, err := h.Service.GetAllTahunAjaran()
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	var resp []*response.TahunAjaran
	for _, item := range *result {
		resp = append(resp, &response.TahunAjaran{
			Uuid:      item.Uuid,
			Tahun:     fmt.Sprintf("%d", item.Tahun),
			Semester:  item.Semester,
			CreatedAt: item.CreatedAt,
			UpdatedAt: item.UpdatedAt,
		})
	}

	c.JSON(http.StatusOK, response.DATA_RES(resp))
}

func (h *TahunAjaranHandler) Get(c *gin.Context) {
	uuid := c.Param("uuid")
	result, err := h.Service.GetTahunAjaran(uuid)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	resp := &response.TahunAjaran{
		Uuid:      result.Uuid,
		Tahun:     fmt.Sprintf("%d", result.Tahun),
		Semester:  result.Semester,
		CreatedAt: result.CreatedAt,
		UpdatedAt: result.UpdatedAt,
	}

	c.JSON(http.StatusOK, response.DATA_RES(resp))
}

func (h *TahunAjaranHandler) Update(c *gin.Context) {
	var body request.TahunAjaran
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
	if err := h.Service.UpdateTahunAjaran(uuid, &body); err != nil {
		utils.HandleError(c, err)
		return
	}
	c.JSON(http.StatusOK, response.SUCCESS_RES("Tahun Ajaran Berhasil Diperbarui"))
}

func (h *TahunAjaranHandler) Delete(c *gin.Context) {
	uuid := c.Param("uuid")
	if err := h.Service.DeleteTahunAjaran(uuid); err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.SUCCESS_RES("Tahun Ajaran Berhasil Dihapus"))
}
