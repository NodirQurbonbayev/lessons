package Hard

import "fmt"

type Passenger struct{
	Name string
	Age int
}

type Flight struct{
	FlightNumber int 
	Source string
	Destination string
	Passengers []Passenger
}
func (f *Flight) AddPassenger(p Passenger) {
	f.Passengers=append(f.Passengers,p)
}
func (f *Flight) RemovePassenger(name string) {
	for i, v := range f.Passengers {
		if name==v.Name {
			f.Passengers=append(f.Passengers[:i],f.Passengers[i+1:]...)
			break
		}
	}
}
func (f Flight) DisplayPassenger() {
	fmt.Printf("Birinchi Yo'lovchi: %v\n",f.FlightNumber)
	for _, v := range f.Passengers {
		fmt.Printf("-%v",v)
	}

}
func main()  {
	flight:=Flight{
		FlightNumber: "A213",
		Source: "New-york",
		Destination: "America",
	}
	Hard.AddPassenger("John")
	Hard.AddPassenger("Jammy")
	Hard.AddPassenger("Jek")
	Hard.DisplayPassenger()
	Hard.RemovePassenger("Jek")
	Hard.DisplayPassenger()


}