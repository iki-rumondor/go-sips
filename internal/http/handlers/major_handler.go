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

type MajorHandler struct {
	Service interfaces.MajorServiceInterface
}

func NewMajorHandler(service interfaces.MajorServiceInterface) interfaces.MajorHandlerInterface {
	return &MajorHandler{
		Service: service,
	}
}

func (h *MajorHandler) CreateMajor(c *gin.Context) {
	var body request.Major
	if err := c.BindJSON(&body); err != nil {
		utils.HandleError(c, response.BADREQ_ERR(err.Error()))
		return
	}

	if _, err := govalidator.ValidateStruct(&body); err != nil {
		utils.HandleError(c, response.BADREQ_ERR(err.Error()))
		return
	}

	if err := h.Service.CreateMajor(&body); err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, response.SUCCESS_RES("Jurusan Berhasil Ditambahkan"))
}

func (h *MajorHandler) GetAllMajors(c *gin.Context) {
	majors, err := h.Service.GetAllMajors()
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	var resp []*response.Major
	for _, item := range *majors {
		resp = append(resp, &response.Major{
			Uuid:      item.Uuid,
			Name:      item.Name,
			CreatedAt: item.CreatedAt,
			UpdatedAt: item.UpdatedAt,
		})
	}

	c.JSON(http.StatusOK, response.DATA_RES(resp))
}

func (h *MajorHandler) GetMajor(c *gin.Context) {
	uuid := c.Param("uuid")
	major, err := h.Service.GetMajor(uuid)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	resp := &response.Major{
		Uuid:      major.Uuid,
		Name:      major.Name,
		CreatedAt: major.CreatedAt,
		UpdatedAt: major.UpdatedAt,
	}

	c.JSON(http.StatusOK, response.DATA_RES(resp))
}

func (h *MajorHandler) UpdateMajor(c *gin.Context) {
	var body request.Major
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
	if err := h.Service.UpdateMajor(uuid, &body); err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.SUCCESS_RES("Jurusan Berhasil Diperbarui"))
}

func (h *MajorHandler) DeleteMajor(c *gin.Context) {
	uuid := c.Param("uuid")
	if err := h.Service.DeleteMajor(uuid); err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.SUCCESS_RES("Jurusan Berhasil Dihapus"))
}
