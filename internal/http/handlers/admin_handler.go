package handlers

import (
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

func (h *AdminHandler) GetPembimbingProdi(c *gin.Context) {
	userUuid := c.Param("userUuid")
	resp, err := h.Service.FindPembimbingProdi(userUuid)
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
	var body request.UpdatePembimbing
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

func (h *AdminHandler) CreateProdi(c *gin.Context) {
	var body request.Prodi
	if err := c.BindJSON(&body); err != nil {
		utils.HandleError(c, response.BADREQ_ERR(err.Error()))
		return
	}

	if _, err := govalidator.ValidateStruct(&body); err != nil {
		utils.HandleError(c, response.BADREQ_ERR(err.Error()))
		return
	}

	if err := h.Service.CreateProdi(&body); err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.SUCCESS_RES("Data Program Studi Berhasil Ditambahkan"))
}

func (h *AdminHandler) GetAllProdi(c *gin.Context) {

	resp, err := h.Service.FindAllProdi()
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.DATA_RES(resp))
}

func (h *AdminHandler) GetProdi(c *gin.Context) {
	uuid := c.Param("uuid")
	resp, err := h.Service.FindProdi(uuid)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.DATA_RES(resp))
}

func (h *AdminHandler) UpdateProdi(c *gin.Context) {
	var body request.Prodi
	if err := c.BindJSON(&body); err != nil {
		utils.HandleError(c, response.BADREQ_ERR(err.Error()))
		return
	}

	if _, err := govalidator.ValidateStruct(&body); err != nil {
		utils.HandleError(c, response.BADREQ_ERR(err.Error()))
		return
	}

	uuid := c.Param("uuid")

	if err := h.Service.UpdateProdi(uuid, &body); err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.SUCCESS_RES("Berhasil Memperbarui Data Program Studi"))
}

func (h *AdminHandler) DeleteProdi(c *gin.Context) {
	uuid := c.Param("uuid")
	if err := h.Service.DeleteProdi(uuid); err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.SUCCESS_RES("Berhasil Menghapus Data Program Studi"))
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

func (h *AdminHandler) GetPenasihatDashboard(c *gin.Context) {

	userUuid := c.Param("userUuid")
	resp, err := h.Service.GetPenasihatDashboard(userUuid)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.DATA_RES(resp))
}

func (h *AdminHandler) GetKaprodiDashboard(c *gin.Context) {
	userUuid := c.Param("userUuid")
	resp, err := h.Service.GetKaprodiDashboard(userUuid)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.DATA_RES(resp))
}

func (h *AdminHandler) GetKajurDashboard(c *gin.Context) {
	resp, err := h.Service.GetKajurDashboard()
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.DATA_RES(resp))
}

func (h *AdminHandler) GetPengaturan(c *gin.Context) {

	resp, err := h.Service.GetPengaturan()
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.DATA_RES(resp))
}

func (h *AdminHandler) GetPengaturanByName(c *gin.Context) {
	name := c.Param("name")
	resp, err := h.Service.GetPengaturanByName(name)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.DATA_RES(resp))
}

func (h *AdminHandler) GetAllUsers(c *gin.Context) {

	resp, err := h.Service.FindAllUsers()
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.DATA_RES(resp))
}

func (h *AdminHandler) CreateKajur(c *gin.Context) {
	var body request.Kajur
	if err := c.BindJSON(&body); err != nil {
		utils.HandleError(c, response.BADREQ_ERR(err.Error()))
		return
	}

	if _, err := govalidator.ValidateStruct(&body); err != nil {
		utils.HandleError(c, response.BADREQ_ERR(err.Error()))
		return
	}

	if err := h.Service.CreateKajur(&body); err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.SUCCESS_RES("Ketua Jurusan Berhasil Ditambahkan"))
}

func (h *AdminHandler) UpdateUsername(c *gin.Context) {
	var body request.UpdateUsername
	if err := c.BindJSON(&body); err != nil {
		utils.HandleError(c, response.BADREQ_ERR(err.Error()))
		return
	}

	if _, err := govalidator.ValidateStruct(&body); err != nil {
		utils.HandleError(c, response.BADREQ_ERR(err.Error()))
		return
	}

	uuid := c.Param("uuid")

	if err := h.Service.UpdateUsername(uuid, &body); err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.SUCCESS_RES("Username Berhasil Diperbarui"))
}

func (h *AdminHandler) UpdatePassword(c *gin.Context) {
	var body request.UpdatePassword
	if err := c.BindJSON(&body); err != nil {
		utils.HandleError(c, response.BADREQ_ERR(err.Error()))
		return
	}

	if _, err := govalidator.ValidateStruct(&body); err != nil {
		utils.HandleError(c, response.BADREQ_ERR(err.Error()))
		return
	}

	uuid := c.Param("uuid")

	if err := h.Service.UpdatePassword(uuid, &body); err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.SUCCESS_RES("Password Berhasil Diperbarui"))
}

func (h *AdminHandler) RekomendasiMahasiswa(c *gin.Context) {
	var body request.RekomendasiMahasiswa
	if err := c.BindJSON(&body); err != nil {
		utils.HandleError(c, response.BADREQ_ERR(err.Error()))
		return
	}

	if _, err := govalidator.ValidateStruct(&body); err != nil {
		utils.HandleError(c, response.BADREQ_ERR(err.Error()))
		return
	}

	if err := h.Service.RekomendasiMahasiswa(&body); err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.SUCCESS_RES("Mahasiswa Berhasil Direkomendasikan"))
}
