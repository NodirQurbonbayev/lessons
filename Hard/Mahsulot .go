package Hard


type MahsulotTuri struct{
	Name string
	Price float64
}
type User struct{
	Ism string
	Email string
	Savatcha []MahsulotTuri
}
func (f *User) AddMahsulot(newmahsulot MahsulotTuri) {
	f.Savatcha=append(f.Savatcha,newmahsulot)
}
func(f *User) OchirMahsulot(newmahsulot MahsulotTuri) {
	Old:=[]MahsulotTuri{}
		for _, v := range f.Savatcha {
			if v==newmahsulot {
				continue
			}
			Old=append(Old,v)
		}
		f.Savatcha=Old
}		
		

