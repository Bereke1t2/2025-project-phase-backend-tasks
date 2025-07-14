package main
import (
	"fmt"
)

type Student struct {
	Name  string
	Grades map[string]float64
	marks map[string]float64
}
func (s *Student) CalculateAverage() float64 {
	total := 0.0
	numSubjects := len(s.Grades)
	for _, grade := range s.Grades {
		total += grade
	}
	if numSubjects == 0 {
		return 0.0
	}
	return total / float64(numSubjects)
}

func (s *Student) getAllGrades() {
	for subject, grade := range s.Grades {
		fmt.Printf("Subject: %s, Grade: %s\n", subject, getGradeLetter(grade))
	}
}

func getGradeLetter(average float64) string {
	if average >= 90 {
		return "A+"
	} else if average >= 85 {
		return "A"
	} else if average >= 80 {
		return "A-"
	} else if average >= 75 {
		return "B+"
	} else if average >= 70 {
		return "B"
	} else if average >= 65 {
		return "B-"
	} else if average >= 60 {
		return "C+"
	} else if average >= 50 {
		return "C"
	} else if average >= 45 {
		return "C-"
	} else if average >= 40 {
		return "D"
	} else {
		return "F"
	}
}
func (s *Student) AddGrade(subject string, grade float64) {
	if s.Grades == nil {
		s.Grades = make(map[string]float64)
	}
	s.Grades[subject] = grade
}

func start_service() {
	fmt.Println("Student Grade Calculator Service Started")
	fmt.Println("Enter student name:")
	var name string
	fmt.Scanln(&name)
	student := Student{
		Name:   name,
		Grades: make(map[string]float64),
	}
	var numSubjects int
	fmt.Println("Enter number of subjects:")
	fmt.Scanln(&numSubjects)
	for i := 0; i < numSubjects; i++ {
		fmt.Printf("Enter subject name for subject %d: ", i+1)
		var subject string
		fmt.Scanln(&subject)
		fmt.Printf("Enter mark for %s: ", subject)
		var grade float64
		fmt.Scanln(&grade)
		student.AddGrade(subject, grade)
	}
	student.getAllGrades()
	average := student.CalculateAverage()
	fmt.Printf("Average grade for %s: %.2f\n", student.Name, average)
	fmt.Printf("Letter grade: %s\n", getGradeLetter(average))
}