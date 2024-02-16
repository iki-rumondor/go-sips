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

type AdminHandler struct {
	Service interfaces.AdminServiceInterface
}

func NewAdminHandler(service interfaces.AdminServiceInterface) interfaces.AdminHandlerInterface {
	return &AdminHandler{
		Service: service,
	}
}

func (h *AdminHandler) SignIn(c *gin.Context) {
	var body request.SignIn
	if err := c.BindJSON(&body); err != nil {
		utils.HandleError(c, response.BADREQ_ERR(err.Error()))
		return
	}

	if _, err := govalidator.ValidateStruct(&body); err != nil {
		utils.HandleError(c, response.BADREQ_ERR(err.Error()))
		return
	}

	jwt, err := h.Service.VerifyAdmin(&body)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": jwt,
	})
}

func (h *AdminHandler) SetMahasiswaPercepatan(c *gin.Context) {

	var body request.PercepatanCond
	if err := c.BindJSON(&body); err != nil {
		utils.HandleError(c, response.BADREQ_ERR(err.Error()))
		return
	}

	if _, err := govalidator.ValidateStruct(&body); err != nil {
		utils.HandleError(c, response.BADREQ_ERR(err.Error()))
		return
	}

	if err := h.Service.SetMahasiswaPercepatan(&body); err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.SUCCESS_RES("Mahasiswa Percepatan Berhasil Ditambahkan"))
}

func (h *AdminHandler) GetMahasiswaPercepatan(c *gin.Context) {

	mahasiswa, err := h.Service.GetMahasiswaPercepatan()
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	var res []*response.Mahasiswa
	for _, item := range *mahasiswa {
		res = append(res, &response.Mahasiswa{
			Uuid:        item.Mahasiswa.Uuid,
			Nim:         item.Mahasiswa.Nim,
			Nama:        item.Mahasiswa.Nama,
			Angkatan:    fmt.Sprintf("%d", item.Mahasiswa.Angkatan),
			Ipk:         fmt.Sprintf("%.2f", item.Mahasiswa.Ipk),
			TotalSks:    fmt.Sprintf("%d", item.Mahasiswa.TotalSks),
			JumlahError: fmt.Sprintf("%d", item.Mahasiswa.JumlahError),
			CreatedAt:   item.Mahasiswa.CreatedAt,
			UpdatedAt:   item.Mahasiswa.UpdatedAt,
		})
	}

	c.JSON(http.StatusOK, response.DATA_RES(res))
}
