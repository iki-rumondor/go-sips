package request

type TeacherSkill struct {
	Skill       string `json:"skill" valid:"required~field skill tidak ditemukan"`
	TeacherUuid string `json:"teacher_uuid" valid:"required~field teacher_uuid_uuid tidak ditemukan"`
}
