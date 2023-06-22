package loyihalar

import "fmt"

type Employee struct{
	Name string
	Age int 
	Salary float64
}
func (f *Employee) CountSalary(amount float64) {
	f.Salary+=amount
}
func (f Employee) Display() {
	fmt.Println("Name:",f.Name)
	fmt.Println("Age:",f.Age)
	fmt.Println("Salary:",f.Salary)
}