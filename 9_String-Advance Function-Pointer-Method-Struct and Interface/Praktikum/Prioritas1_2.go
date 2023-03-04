package main

import "fmt"

type Student struct {
	Name  string
	Score []int
}

func (s *Student) AverageScore() float64 {
	totalScore := 0
	for _, score := range s.Score {
		totalScore += score
	}
	return float64(totalScore) / float64(len(s.Score))
}

func (s *Student) MinScore() int {
	minScore := s.Score[0]
	for _, score := range s.Score[1:] {
		if score < minScore {
			minScore = score
		}
	}
	return minScore
}

func (s *Student) MaxScore() int {
	maxScore := s.Score[0]
	for _, score := range s.Score[1:] {
		if score > maxScore {
			maxScore = score
		}
	}
	return maxScore
}

func main() {
	students := []Student{
		{"Rizky", []int{80}},
		{"Kobar", []int{75}},
		{"Ismail", []int{70}},
		{"Umam", []int{75}},
		{"Yopan", []int{60}},
	}

	// Hitung skor rata-rata
	totalAverage := 0.0
	for _, student := range students {
		average := student.AverageScore()
		fmt.Printf("Student's Name %s: %.2f\n", student.Name, average)
		totalAverage += average
	}
	fmt.Printf("Average Score : %.2f\n", totalAverage/float64(len(students)))

	// Temukan siswa dengan skor minimum dan maksimum
	minStudent := students[0]
	maxStudent := students[0]
	for _, student := range students[1:] {
		if student.MinScore() < minStudent.MinScore() {
			minStudent = student
		}
		if student.MaxScore() > maxStudent.MaxScore() {
			maxStudent = student
		}
	}
	fmt.Printf("Min Score of Students : %s (%d)\n", minStudent.Name, minStudent.MinScore())
	fmt.Printf("Max Score of Students: %s (%d)\n", maxStudent.Name, maxStudent.MaxScore())
}
