package loyihalar

type Geometriya struct{
	length float64
	width float64
}
func (d *Geometriya) Yuzi() float64 {
	return d.length*d.width
}
func (d *Geometriya) Perimetr() float64{
	return 2*(d.length+d.width)
}