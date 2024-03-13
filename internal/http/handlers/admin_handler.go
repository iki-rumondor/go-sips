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

	jwt, err := h.Service.VerifyPengguna(&body)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": jwt,
	})
}

func (h *AdminHandler) GetUser(c *gin.Context) {

	uuid := c.Param("uuid")

	resp, err := h.Service.GetUser(uuid)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	
	c.JSON(http.StatusOK, response.DATA_RES(resp))
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

func (h *AdminHandler) SetMahasiswaPeringatan(c *gin.Context) {

	if err := h.Service.SetMahasiswaPeringatan(); err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.SUCCESS_RES("Status Mahasiswa Berhasil Diperbarui"))
}

func (h *AdminHandler) GetMahasiswaPeringatan(c *gin.Context) {

	mahasiswa, err := h.Service.GetMahasiswaPeringatan()
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	var res []*response.StatusMahasiswa
	for _, item := range *mahasiswa {
		res = append(res, &response.StatusMahasiswa{
			Status: fmt.Sprintf("%d", item.Status),
			Mahasiswa: &response.Mahasiswa{
				Uuid:        item.Mahasiswa.Uuid,
				Nim:         item.Mahasiswa.Nim,
				Nama:        item.Mahasiswa.Nama,
				Angkatan:    fmt.Sprintf("%d", item.Mahasiswa.Angkatan),
				Ipk:         fmt.Sprintf("%.2f", item.Mahasiswa.Ipk),
				TotalSks:    fmt.Sprintf("%d", item.Mahasiswa.TotalSks),
				JumlahError: fmt.Sprintf("%d", item.Mahasiswa.JumlahError),
			},
			CreatedAt: item.CreatedAt,
			UpdatedAt: item.UpdatedAt,
		})
	}

	c.JSON(http.StatusOK, response.DATA_RES(res))
}

func (h *AdminHandler) CreatePembimbing(c *gin.Context) {
	var body request.Pembimbing
	if err := c.BindJSON(&body); err != nil {
		utils.HandleError(c, response.BADREQ_ERR(err.Error()))
		return
	}

	if _, err := govalidator.ValidateStruct(&body); err != nil {
		utils.HandleError(c, response.BADREQ_ERR(err.Error()))
		return
	}

	if err := h.Service.CreatePembimbing(&body); err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.SUCCESS_RES("Data Pembimbing Akademik Berhasil Ditambahkan"))
}

func (h *AdminHandler) GetAllPembimbing(c *gin.Context) {

	resp, err := h.Service.FindAllPembimbing()
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.DATA_RES(resp))
}

func (h *AdminHandler) GetPembimbing(c *gin.Context) {
	uuid := c.Param("uuid")
	resp, err := h.Service.FindPembimbing(uuid)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.DATA_RES(resp))
}

func (h *AdminHandler) UpdatePembimbing(c *gin.Context) {
	var body request.Pembimbing
	if err := c.BindJSON(&body); err != nil {
		utils.HandleError(c, response.BADREQ_ERR(err.Error()))
		return
	}

	if _, err := govalidator.ValidateStruct(&body); err != nil {
		utils.HandleError(c, response.BADREQ_ERR(err.Error()))
		return
	}

	uuid := c.Param("uuid")

	if err := h.Service.UpdatePembimbing(uuid, &body); err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.SUCCESS_RES("Berhasil Memperbarui Data Pembimbing Akademik"))
}

func (h *AdminHandler) DeletePembimbing(c *gin.Context) {
	uuid := c.Param("uuid")
	if err := h.Service.DeletePembimbing(uuid); err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.SUCCESS_RES("Berhasil Menghapus Data Pembimbing Akademik"))
}

func (h *AdminHandler) UpdateKelas(c *gin.Context) {
	var body request.KelasRule
	if err := c.BindJSON(&body); err != nil {
		utils.HandleError(c, response.BADREQ_ERR(err.Error()))
		return
	}

	if _, err := govalidator.ValidateStruct(&body); err != nil {
		utils.HandleError(c, response.BADREQ_ERR(err.Error()))
		return
	}

	if err := h.Service.UpdateKelas(&body); err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.SUCCESS_RES("Berhasil Sinkronisasi Kelas Mahasiswa"))
}

func (h *AdminHandler) GetClasses(c *gin.Context) {

	resp, err := h.Service.GetClasses()
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.DATA_RES(resp))
}
