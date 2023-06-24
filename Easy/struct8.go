package loyihalar

import "fmt"
type Student struct{
	Name string 
	Grades []float64
}
func (s *Student) Checkst() {
	sum:=0.0
	for _, grade := range s.Grades {
		sum+=grade
	}
	str:=sum/float64(len(s.Grades))
	fmt.Println("Ism: ",s.Name)
	fmt.Println("O'rtacha baho: ",str)
}