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

	if err := utils.IsExcelFile(file); err != nil {
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

	failedImport, err := h.Service.ImportMahasiswa(pathFile)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.DATA_RES(failedImport))
}

func (h *MahasiswaHandler) GetAll(c *gin.Context) {
	result, err := h.Service.GetAllMahasiswa()
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	var resp []*response.Mahasiswa
	for _, item := range *result {
		resp = append(resp, &response.Mahasiswa{
			Uuid:        item.Uuid,
			Nim:         item.Nim,
			Nama:        item.Nama,
			Angkatan:    fmt.Sprintf("%d", item.Angkatan),
			Ipk:         fmt.Sprintf("%.2f", item.Ipk),
			TotalSks:    fmt.Sprintf("%d", item.TotalSks),
			JumlahError: fmt.Sprintf("%d", item.JumlahError),
			CreatedAt:   item.CreatedAt,
			UpdatedAt:   item.UpdatedAt,
		})
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
		Angkatan:    fmt.Sprintf("%d", result.Angkatan),
		Ipk:         fmt.Sprintf("%.2f", result.Ipk),
		TotalSks:    fmt.Sprintf("%d", result.TotalSks),
		JumlahError: fmt.Sprintf("%d", result.JumlahError),
		CreatedAt:   result.CreatedAt,
		UpdatedAt:   result.UpdatedAt,
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

	c.JSON(http.StatusOK, response.SUCCESS_RES("Jurusan Berhasil Diperbarui"))
}

func (h *MahasiswaHandler) Delete(c *gin.Context) {
	uuid := c.Param("uuid")
	if err := h.Service.DeleteMahasiswa(uuid); err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.SUCCESS_RES("Jurusan Berhasil Dihapus"))
}
