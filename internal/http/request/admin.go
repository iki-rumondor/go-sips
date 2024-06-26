package request

type SignIn struct {
	Username string `json:"username" valid:"required~field username tidak ditemukan"`
	Password string `json:"password" valid:"required~field password tidak ditemukan"`
}

type Pembimbing struct {
	UserUuid string `json:"user_uuid" valid:"required~field user uuid tidak ditemukan"`
	Nama     string `json:"nama" valid:"required~field nama tidak ditemukan"`
}

type UpdatePembimbing struct {
	Nama string `json:"nama" valid:"required~field nama tidak ditemukan"`
}

type Prodi struct {
	Name     string `json:"name" valid:"required~field nama tidak ditemukan"`
	Kaprodi  string `json:"kaprodi" valid:"required~field kaprodi tidak ditemukan"`
	Username string `json:"username" valid:"required~field username tidak ditemukan"`
}

type Kajur struct {
	Username        string `json:"username" valid:"required~field username tidak ditemukan"`
	Password        string `json:"password" valid:"required~field password tidak ditemukan"`
	ConfirmPassword string `json:"confirm_password" valid:"required~field konfirmasi password tidak ditemukan"`
	RoleID          uint   `json:"role_id" valid:"required~field role id tidak ditemukan"`
}

type KelasRule struct {
	JumlahMahasiswa string `json:"amount" valid:"required~field Jumlah Mahasiswa tidak ditemukan, int~field Jumlah Mahasiswa harus berupa bilangan bulat, range(1|200)~field Jumlah Mahasiswa tidak valid"`
}

type PercepatanCond struct {
	TotalSks    string `json:"total_sks" valid:"required~field total_sks tidak ditemukan, int~field total_sks harus berupa bilangan bulat, range(0|200)~field total_sks tidak valid"`
	Ipk         string `json:"ipk" valid:"required~field ipk tidak ditemukan, float~field ipk harus berupa bilangan desimal, range(0|4)~field ipk tidak valid"`
	JumlahError string `json:"jumlah_error" valid:"required~field jumlah_error tidak ditemukan, int~field jumlah_error harus berupa bilangan bulat, range(0|200)~field jumlah_error tidak valid"`
}

type Pengaturan struct {
	TotalSks           string `json:"total_sks" valid:"required~field total_sks tidak ditemukan, int~field total_sks harus berupa bilangan bulat, range(0|200)~field total_sks tidak valid"`
	Ipk                string `json:"ipk" valid:"required~field ipk tidak ditemukan, float~field ipk harus berupa bilangan desimal, range(0|4)~field ipk tidak valid"`
	AngkatanPercepatan string `json:"angkatan_percepatan" valid:"required~field angkatan percepatan tidak ditemukan, int~field angkatan percepatan tidak valid, range(2000|3000)~field angkatan percepatan tidak valid"`
	AngkatanKelas      string `json:"angkatan_kelas" valid:"required~field angkatan kelas tidak ditemukan, int~field angkatan kelas tidak valid, range(2000|3000)~field angkatan kelas tidak valid"`
	JumlahMahasiswa    string `json:"jumlah_mahasiswa" valid:"required~field jumlah mahasiswa per kelas tidak ditemukan, int~field Jumlah Mahasiswa harus berupa bilangan bulat, range(1|200)~field Jumlah Mahasiswa tidak valid"`
	MaksimalPercepatan string `json:"maksimal_percepatan" valid:"required~field maksimal percepatan tidak ditemukan, int~field Maksimal Mahasiswa Percepatan harus berupa bilangan bulat, range(1|200)~field Maksimal Mahasiswa Percepatan tidak valid"`
}

type UpdateUsername struct {
	Username string `json:"username" valid:"required~field username tidak ditemukan"`
}

type UpdatePassword struct {
	CurrentPassword string `json:"current_password" valid:"required~field password lama tidak ditemukan"`
	NewPassword     string `json:"new_password" valid:"required~field password baru tidak ditemukan"`
	ConfirmPassword string `json:"confirm_password" valid:"required~field konfirmasi password tidak ditemukan"`
}

type RekomendasiMahasiswa struct {
	UuidMahasiswa  string `json:"uuid_mahasiswa" valid:"required~field uuid mahasiswa tidak ditemukan"`
	UuidPembimbing string `json:"uuid_pembimbing" valid:"required~field uuid pembimbing tidak ditemukan"`
}
