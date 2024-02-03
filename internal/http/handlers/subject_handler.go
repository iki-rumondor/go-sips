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

type SubjectHandler struct {
	Service interfaces.SubjectServiceInterface
}

func NewSubjectHandler(service interfaces.SubjectServiceInterface) interfaces.SubjectHandlerInterface {
	return &SubjectHandler{
		Service: service,
	}
}

func (h *SubjectHandler) CreateSubject(c *gin.Context) {
	var body request.Subject
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

	if err := h.Service.CreateSubject(userUuid, &body); err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, response.SUCCESS_RES("Mata Kuliah Berhasil Ditambahkan"))
}

func (h *SubjectHandler) GetAllSubjects(c *gin.Context) {
	userUuid := c.GetString("uuid")
	if userUuid == "" {
		utils.HandleError(c, response.HANDLER_INTERR)
		return
	}

	subjects, err := h.Service.GetAllSubjects(userUuid)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	var resp []*response.Subject
	for _, item := range *subjects {
		var academicPlan *response.AcademicPlan
		if item.AcademicPlan != nil {
			academicPlan = &response.AcademicPlan{
				AcademicYear: &response.AcademicYear{
					Uuid: item.AcademicPlan.AcademicYear.Uuid,
				},
				Uuid:      item.AcademicPlan.Uuid,
				Available: item.AcademicPlan.Available,
				Note:      item.AcademicPlan.Note,
			}
		}

		resp = append(resp, &response.Subject{
			Uuid:      item.Uuid,
			Name:      item.Name,
			Code:      item.Code,
			Practical: item.Practical,
			Department: &response.Department{
				Uuid: item.Department.Uuid,
				Name: item.Department.Name,
				Head: item.Department.Head,
			},
			AcademicPlan: academicPlan,
			CreatedAt:    item.CreatedAt,
			UpdatedAt:    item.UpdatedAt,
		})
	}

	c.JSON(http.StatusOK, response.DATA_RES(resp))
}

func (h *SubjectHandler) GetAllPracticalSubjects(c *gin.Context) {
	userUuid := c.GetString("uuid")
	if userUuid == "" {
		utils.HandleError(c, response.HANDLER_INTERR)
		return
	}

	subjects, err := h.Service.GetPracticalSubjects(userUuid)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	var resp []*response.PracticalSubject

	for _, item := range *subjects {
		var tools *response.PracticalTool
		if item.PracticalTool != nil {
			tools = &response.PracticalTool{
				AcademicYear: &response.AcademicYear{
					Uuid: item.PracticalTool.AcademicYear.Uuid,
				},
				Uuid:      item.PracticalTool.Uuid,
				Available: item.PracticalTool.Available,
				Note:      item.PracticalTool.Note,
			}
		}

		var modules *response.PracticalModule
		if item.PracticalModule != nil {
			modules = &response.PracticalModule{
				AcademicYear: &response.AcademicYear{
					Uuid: item.PracticalModule.AcademicYear.Uuid,
				},
				Uuid:      item.PracticalModule.Uuid,
				Available: item.PracticalModule.Available,
				Note:      item.PracticalModule.Note,
			}
		}

		resp = append(resp, &response.PracticalSubject{
			Uuid:            item.Uuid,
			Name:            item.Name,
			Code:            item.Code,
			PracticalTool:   tools,
			PracticalModule: modules,
			CreatedAt:       item.CreatedAt,
			UpdatedAt:       item.UpdatedAt,
		})
	}

	c.JSON(http.StatusOK, response.DATA_RES(resp))
}

func (h *SubjectHandler) GetSubject(c *gin.Context) {
	uuid := c.Param("uuid")
	userUuid := c.GetString("uuid")
	if userUuid == "" {
		utils.HandleError(c, response.HANDLER_INTERR)
		return
	}

	subject, err := h.Service.GetSubject(userUuid, uuid)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	var academicPlan *response.AcademicPlan
	if subject.AcademicPlan != nil {
		academicPlan = &response.AcademicPlan{
			AcademicYear: &response.AcademicYear{
				Uuid: subject.AcademicPlan.AcademicYear.Uuid,
			},
			Uuid:      subject.AcademicPlan.Uuid,
			Available: subject.AcademicPlan.Available,
			Note:      subject.AcademicPlan.Note,
		}
	}

	resp := &response.Subject{
		Uuid:      subject.Uuid,
		Name:      subject.Name,
		Code:      subject.Code,
		Practical: subject.Practical,
		Department: &response.Department{
			Uuid: subject.Department.Uuid,
			Name: subject.Department.Name,
			Head: subject.Department.Head,
		},
		AcademicPlan: academicPlan,
		CreatedAt:    subject.CreatedAt,
		UpdatedAt:    subject.UpdatedAt,
	}

	c.JSON(http.StatusOK, response.DATA_RES(resp))
}

func (h *SubjectHandler) UpdateSubject(c *gin.Context) {
	var body request.Subject
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
	if err := h.Service.UpdateSubject(userUuid, uuid, &body); err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.SUCCESS_RES("Mata Kuliah Berhasil Diperbarui"))
}

func (h *SubjectHandler) DeleteSubject(c *gin.Context) {
	userUuid := c.GetString("uuid")
	if userUuid == "" {
		utils.HandleError(c, response.HANDLER_INTERR)
		return
	}

	uuid := c.Param("uuid")
	if err := h.Service.DeleteSubject(userUuid, uuid); err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.SUCCESS_RES("Mata Kuliah Berhasil Dihapus"))
}
