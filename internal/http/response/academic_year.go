package response

type AcademicYear struct {
	Uuid            string             `json:"uuid"`
	Name            string             `json:"name"`
	AcademicPlan    *[]AcademicPlan    `json:"academic_plans"`
	PracticalTool   *[]PracticalTool   `json:"practical_tools"`
	PracticalModule *[]PracticalModule `json:"practical_modules"`
	CreatedAt       int64              `json:"created_at"`
	UpdatedAt       int64              `json:"updated_at"`
}
