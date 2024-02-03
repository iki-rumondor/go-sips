package response

type PracticalModule struct {
	Uuid         string        `json:"uuid"`
	Available    *bool          `json:"available"`
	Note         string        `json:"note"`
	Subject      *Subject      `json:"subject"`
	Laboratory   *Laboratory   `json:"laboratory"`
	AcademicYear *AcademicYear `json:"academic_year"`
	CreatedAt    int64         `json:"created_at"`
	UpdatedAt    int64         `json:"updated_at"`
}
