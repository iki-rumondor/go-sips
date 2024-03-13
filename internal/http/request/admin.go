package request

type SignIn struct {
	Username string `json:"username" valid:"required~field username tidak ditemukan"`
	Password string `json:"password" valid:"required~field password tidak ditemukan"`
}

type Pembimbing struct {
	Nama string `json:"nama" valid:"required~field nama tidak ditemukan"`
	Nip  string `json:"nip" valid:"required~field nip tidak ditemukan"`
}

type KelasRule struct {
	JumlahMahasiswa string `json:"amount" valid:"required~field Jumlah Mahasiswa tidak ditemukan, int~field Jumlah Mahasiswa harus berupa bilangan bulat, range(1|200)~field Jumlah Mahasiswa tidak valid"`
}

type PercepatanCond struct {
	TotalSks    string `json:"total_sks" valid:"required~field total_sks tidak ditemukan, int~field total_sks harus berupa bilangan bulat, range(0|200)~field total_sks tidak valid"`
	Ipk         string `json:"ipk" valid:"required~field ipk tidak ditemukan, float~field ipk harus berupa bilangan desimal, range(0|4)~field ipk tidak valid"`
	JumlahError string `json:"jumlah_error" valid:"required~field jumlah_error tidak ditemukan, int~field jumlah_error harus berupa bilangan bulat, range(0|200)~field jumlah_error tidak valid"`
}
