package request

type SignIn struct {
	Username string `json:"username" valid:"required~field username tidak ditemukan"`
	Password string `json:"password" valid:"required~field password tidak ditemukan"`
}

type PercepatanCond struct {
	Angkatan    string `json:"angkatan" valid:"required~field angkatan tidak ditemukan, int~field angkatan harus berupa bilangan bulat, range(0|9999)~field angkatan tidak valid"`
	TotalSks    string `json:"total_sks" valid:"required~field total_sks tidak ditemukan, int~field total_sks harus berupa bilangan bulat, range(0|200)~field total_sks tidak valid"`
	Ipk         string `json:"ipk" valid:"required~field ipk tidak ditemukan, float~field ipk harus berupa bilangan desimal, range(0|4)~field ipk tidak valid"`
	JumlahError string `json:"jumlah_error" valid:"required~field jumlah_error tidak ditemukan, int~field jumlah_error harus berupa bilangan bulat, range(0|200)~field jumlah_error tidak valid"`
}
