package objptr

type Student struct {
	Name string
	Age  int

	Card *Card
}

type Card struct {
	ID int
}


func StringLen(s string) int

func String(s string) string

func NewStudent(s Student) Student

func NewStudentPtr(s *Student) *Student

func UpStudentPtr(s *Student, name string, age int)

func StudentName(s *Student) string
