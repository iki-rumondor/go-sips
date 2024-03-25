package response

type Pembimbing struct {
	Uuid string `json:"uuid"`
	Nama string `json:"nama"`
	Nip  string `json:"nip"`
}

type Prodi struct {
	Uuid     string `json:"uuid"`
	Name     string `json:"name"`
	Kaprodi  string `json:"kaprodi"`
	Username string `json:"username"`
}

type Pengaturan struct {
	Uuid  string `json:"uuid"`
	Name  string `json:"name"`
	Value string `json:"value"`
}

type User struct {
	Uuid      string         `json:"uuid"`
	Username  string         `json:"username"`
	Role      string         `json:"role"`
	Mahasiswa *DataMahasiswa `json:"mahasiswa"`
	Penasihat *Pembimbing    `json:"penasihat"`
}
