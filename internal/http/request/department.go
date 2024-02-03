package request

type Department struct {
	MajorUuid string `json:"major_uuid" valid:"required~field major_uuid tidak ditemukan"`
	Name      string `json:"name" valid:"required~field name tidak ditemukan"`
	Head      string `json:"head" valid:"required~field head tidak ditemukan"`
}
