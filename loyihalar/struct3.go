package loyihalar
import "fmt"
type RentCar struct{
	Make  string
	Model string
	RentDay float64
}
func (g RentCar) Count(day int ) float64{
	return g.RentDay*float64(day)
}
func (g RentCar) Display() {
	fmt.Println("Make:",g.Make)
	fmt.Println("Model:",g.Model)
	fmt.Println("Rentday:",g.RentDay)
}