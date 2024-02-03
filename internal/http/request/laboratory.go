package request

type Laboratory struct {
	Name string `json:"name" valid:"required~field name tidak ditemukan"`
}
