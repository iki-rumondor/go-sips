package request

type PracticalTool struct {
	Available        *bool  `json:"available"`
	Condition        string `json:"condition"`
	Note             string `json:"note" valid:"required~field note tidak ditemukan"`
	SubjectUuid      string `json:"subject_uuid" valid:"required~field subject_uuid tidak ditemukan"`
	AcademicYearUuid string `json:"academic_year_uuid" valid:"required~field academic_year_uuid tidak ditemukan"`
}

type UpdatePracticalTool struct {
	Available *bool  `json:"available"`
	Condition string `json:"condition"`
	Note      string `json:"note" valid:"required~field note tidak ditemukan"`
}
