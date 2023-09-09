package repositories

type StudentRepository interface {
	GetStudentName() string
}

type StudentManager struct {
	Id     int
	Name   string
	Age    string
	Gender string
}

func NewStudentManager() StudentRepository {
	return &StudentManager{}
}

func (repo *StudentManager) GetStudentName() string {
	student := &StudentManager{
		Id:     1,
		Name:   "iMooc",
		Age:    "23",
		Gender: "male",
	}
	return student.Name
}
