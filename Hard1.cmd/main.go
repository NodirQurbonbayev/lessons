package main


import (
	//"fmt"
	"lesson/Hard"
)
func main()  {
	// mijoz1:=Hard.User {
	// 	Ism: "Nodir",
	// 	Email: "Nodirqurbonbayev3@gmail.com",
	// 	Savatcha: []Hard.MahsulotTuri{},
	// }
	// mijoz2:=Hard.User{
	// 	Ism: "John",
	// 	Email: "John@gmail.com",
	// 	Savatcha: []Hard.MahsulotTuri{},
	// }
	// sut:=Hard.MahsulotTuri{
	// 	Name: "Milk",
	// 	Price: 15000.0,
	// }
	// non:=Hard.MahsulotTuri{
	// 	Name: "Non",
	// 	Price: 12000.0,
	// }
	// Chears:=Hard.MahsulotTuri{
	// 	Name: "Lays",
	// 	Price: 18000.0,
	// }
	// mijoz1.AddMahsulot(sut)
	// mijoz1.AddMahsulot(non)
	// mijoz1.AddMahsulot(Chears)
	// fmt.Println(mijoz1)
	// mijoz2.AddMahsulot(Chears)
	// mijoz2.AddMahsulot(sut)
	// mijoz2.AddMahsulot(non)
	// mijoz1.OchirMahsulot(Chears)
	// fmt.Println(mijoz2)
	nodir:=Hard.Users{
		Name: "Nodir",
		Email: "Nodir@gmail.com",
		Friend: []Hard.Users{},
	}
	john:=Hard.Users{
		Name: "John",
		Email: "John@gmail.com",
		Friend: []Hard.Users{},
	}
	Parker:=Hard.Users{
		Name: "Parker",
		Email: "parker@gmail.com",
		Friend: []Hard.Users{},
	}
	nodir.AddFriend(john)
	Parker.AddFriend(nodir)
	john.AddFriend(Parker)
	john.AddFriend(nodir)
	john.AddFriend(john)
	john.Display()
}