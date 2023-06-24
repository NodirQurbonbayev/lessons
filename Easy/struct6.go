package loyihalar
import "fmt"
type Product struct{
	Name string
	Price float64
	Quantity int
}
func (n *Product) CalcProduct(amount int) {
	n.Quantity +=amount
}
func (n *Product) OldNew(newQuantity int) {
	if n.Quantity>=newQuantity {
		n.Quantity-=newQuantity
	}else {
		fmt.Println("Mahsulot kam!")
	}
}
func (n Product) Jami() float64{
	return float64(n.Quantity)*n.Price
}

func (n Product) Display() {
	fmt.Printf("Name: %s\n",n.Name)
	fmt.Printf("Price: $%.2f\n",n.Price)
	fmt.Printf("Quantity: %d\n",n.Quantity)
	fmt.Println("Narxi: ",n.Jami())
}
