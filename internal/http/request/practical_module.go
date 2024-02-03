package request

type PracticalModule struct {
	SubjectUuid      string `json:"subject_uuid" valid:"required~field subject_uuid tidak ditemukan"`
	LaboratoryUuid   string `json:"laboratory_uuid" valid:"required~field laboratory_uuid tidak ditemukan"`
	AcademicYearUuid string `json:"academic_year_uuid" valid:"required~field academic_year_uuid tidak ditemukan"`
	Available        *bool  `json:"available"`
	Note             string `json:"note" valid:"required~field note tidak ditemukan"`
}

type UpdatePracticalModule struct {
	LaboratoryUuid string `json:"laboratory_uuid" valid:"required~field laboratory_uuid tidak ditemukan"`
	Available      *bool  `json:"available"`
	Note           string `json:"note" valid:"required~field note tidak ditemukan"`
}
