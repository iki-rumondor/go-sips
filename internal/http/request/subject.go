package request

type Subject struct {
	Name      string `json:"name" valid:"required~field name tidak ditemukan"`
	Code      string `json:"code" valid:"required~field code tidak ditemukan"`
	Practical *bool   `json:"practical"`
}
