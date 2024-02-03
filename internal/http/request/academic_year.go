package request

type AcademicYear struct {
	Name      string `json:"name" valid:"required~field name tidak ditemukan"`
}
