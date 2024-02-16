package response

type TahunAjaran struct {
	Uuid      string `json:"uuid"`
	Tahun     string `json:"tahun"`
	Semester  string `json:"semester"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}
