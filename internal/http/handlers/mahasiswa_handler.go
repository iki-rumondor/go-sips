package handlers

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/iki-rumondor/sips/internal/http/request"
	"github.com/iki-rumondor/sips/internal/http/response"
	"github.com/iki-rumondor/sips/internal/interfaces"
	"github.com/iki-rumondor/sips/internal/utils"
)

type MahasiswaHandler struct {
	Service interfaces.MahasiswaServiceInterface
}

func NewMahasiswaHandler(service interfaces.MahasiswaServiceInterface) interfaces.MahasiswaHandlerInterface {
	return &MahasiswaHandler{
		Service: service,
	}
}

func (h *MahasiswaHandler) Import(c *gin.Context) {
	file, err := c.FormFile("mahasiswa")
	if err != nil {
		utils.HandleError(c, response.NOTFOUND_ERR("File Tidak Ditemukan"))
		return
	}

	if err := utils.IsCsvFile(file); err != nil {
		utils.HandleError(c, err)
		return
	}

	tempFolder := "internal/temp"
	pathFile := filepath.Join(tempFolder, file.Filename)

	if err := c.SaveUploadedFile(file, pathFile); err != nil {
		utils.HandleError(c, response.HANDLER_INTERR)
	}

	defer func() {
		if err := os.Remove(pathFile); err != nil {
			fmt.Println(err.Error())
		}
	}()

	uuid := c.Param("userUuid")
	failedImport, err := h.Service.CreateMahasiswaCSV(uuid, pathFile)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.DATA_RES(failedImport))
}

func (h *MahasiswaHandler) GetAll(c *gin.Context) {
	class := c.DefaultQuery("kelas", "")
	angkatan := c.DefaultQuery("angkatan", "")
	options := map[string]string{
		"class":    class,
		"angkatan": angkatan,
	}

	resp, err := h.Service.GetAllMahasiswa(options)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.DATA_RES(resp))
}

func (h *MahasiswaHandler) Get(c *gin.Context) {
	uuid := c.Param("uuid")
	result, err := h.Service.GetMahasiswa(uuid)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	resp := response.Mahasiswa{
		Uuid:        result.Uuid,
		Nim:         result.Nim,
		Nama:        result.Nama,
		Kelas:       result.Class,
		Percepatan:  result.Percepatan,
		Angkatan:    fmt.Sprintf("%d", result.Angkatan),
		Ipk:         fmt.Sprintf("%.2f", result.Ipk),
		TotalSks:    fmt.Sprintf("%d", result.TotalSks),
		JumlahError: fmt.Sprintf("%d", result.JumlahError),
		Pembimbing: &response.Pembimbing{
			Uuid: result.PembimbingAkademik.Uuid,
			Nama: result.PembimbingAkademik.Nama,
		},
		CreatedAt: result.CreatedAt,
		UpdatedAt: result.UpdatedAt,
	}

	c.JSON(http.StatusOK, response.DATA_RES(resp))
}

func (h *MahasiswaHandler) Update(c *gin.Context) {
	var body request.Mahasiswa
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
	if err := h.Service.UpdateMahasiswa(uuid, &body); err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.SUCCESS_RES("Mahasiswa Berhasil Diperbarui"))
}

func (h *MahasiswaHandler) Delete(c *gin.Context) {
	uuid := c.Param("uuid")
	if err := h.Service.DeleteMahasiswa(uuid); err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.SUCCESS_RES("Mahasiswa Berhasil Dihapus"))
}

func (h *MahasiswaHandler) DeleteAll(c *gin.Context) {
	userUuid := c.Param("userUuid")
	if err := h.Service.DeleteAllMahasiswa(userUuid); err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.SUCCESS_RES("Mahasiswa Berhasil Dihapus"))
}

func (h *MahasiswaHandler) GetData(c *gin.Context) {
	nim := c.Param("nim")
	if nim == "" {
		utils.HandleError(c, response.BADREQ_ERR("Nim Tidak Ditemukan"))
		return
	}

	resp, err := h.Service.GetDataMahasiswa(nim)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.DATA_RES(resp))
}

func (h *MahasiswaHandler) GetMahasiswaByUserUuid(c *gin.Context) {
	userUuid := c.Param("userUuid")
	if userUuid == "" {
		utils.HandleError(c, response.BADREQ_ERR("Uuid Tidak Ditemukan"))
		return
	}

	resp, err := h.Service.GetMahasiswaByUserUuid(userUuid)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.DATA_RES(resp))
}

func (h *MahasiswaHandler) GetMahasiswaProdi(c *gin.Context) {
	userUuid := c.Param("userUuid")
	resp, err := h.Service.GetMahasiswaProdi(userUuid)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.DATA_RES(resp))
}

func (h *MahasiswaHandler) GetMahasiswaByPenasihat(c *gin.Context) {
	class := c.DefaultQuery("kelas", "")
	angkatan := c.DefaultQuery("angkatan", "")
	min_angkatan := c.DefaultQuery("min_angkatan", "")
	options := map[string]string{
		"class":        class,
		"angkatan":     angkatan,
		"min_angkatan": min_angkatan,
	}

	userUuid := c.Param("userUuid")
	if userUuid == "" {
		utils.HandleError(c, response.BADREQ_ERR("Uuid Tidak Ditemukan"))
		return
	}

	resp, err := h.Service.GetAllMahasiswaByPenasihat(userUuid, options)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.DATA_RES(resp))
}

func (h *MahasiswaHandler) UpdatePengaturan(c *gin.Context) {
	var body request.Pengaturan
	if err := c.BindJSON(&body); err != nil {
		utils.HandleError(c, response.BADREQ_ERR(err.Error()))
		return
	}

	if _, err := govalidator.ValidateStruct(&body); err != nil {
		utils.HandleError(c, response.BADREQ_ERR(err.Error()))
		return
	}

	if err := h.Service.UpdatePengaturan(&body); err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.SUCCESS_RES("Pengaturan Berhasil Diperbarui"))
}

func (h *MahasiswaHandler) GetMahasiswaPercepatan(c *gin.Context) {
	resp, err := h.Service.GetMahasiswaPercepatan()
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.DATA_RES(resp))
}

func (h *MahasiswaHandler) GetProdiPercepatan(c *gin.Context) {
	prodiUuid := c.Param("uuid")
	resp, err := h.Service.GetProdiPercepatan(prodiUuid)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.DATA_RES(resp))
}
