package response

type PracticalTool struct {
	Uuid         string        `json:"uuid"`
	Available    *bool         `json:"available"`
	Condition    string        `json:"condition"`
	Note         string        `json:"note"`
	Subject      *Subject      `json:"subject"`
	AcademicYear *AcademicYear `json:"academic_year"`
	CreatedAt    int64         `json:"created_at"`
	UpdatedAt    int64         `json:"updated_at"`
}
