package response

type FailedImport struct {
	Nim   string `json:"nim"`
	Nama  string `json:"nama"`
	Pesan string `json:"pesan"`
}

type Mahasiswa struct {
	Uuid        string `json:"uuid"`
	Nim         string `json:"nim"`
	Nama        string `json:"nama"`
	Angkatan    string `json:"angkatan"`
	TotalSks    string `json:"total_sks"`
	Ipk         string `json:"ipk"`
	JumlahError string `json:"jumlah_error"`
	CreatedAt   int64  `json:"created_at"`
	UpdatedAt   int64  `json:"updated_at"`
}
