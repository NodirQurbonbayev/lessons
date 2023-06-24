package loyihalar

import "fmt"

type Book struct {
	Title string
	Author string
	Price float64
}
func (n *Book) Chegirma(amount float64) {
	n.Price = n.Price-n.Price*amount/100
}
func (n Book) Disp() {
	fmt.Println("Title:",n.Title)
	fmt.Println("Author:",n.Author)
	fmt.Println("Price:",n.Price)
}