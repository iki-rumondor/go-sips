package request

type TahunAjaran struct {
	Tahun    string `json:"tahun" valid:"required~field tahun tidak ditemukan; int~field tahun harus berupa angka"`
	Semester string `json:"semester" valid:"required~field semester tidak ditemukan; int~field semester harus berupa angka"`
}
