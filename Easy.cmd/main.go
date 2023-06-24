package main

import (
	"fmt"
	"lesson/loyihalar"
)

func main() {
	account := loyihalar.BankAccount{
		AccountNumber: "132534536",
		HolderName:    "Nodir",
		Balance:       100.0,
	}
	account.Deposit(500.0)
	account.Withdraw(200.0)
	account.Display()
	///////////////////////////////////////////////////////////////////////////////////////////////////////////
		book:=loyihalar.Book {
			Title: "Nodir",
			Author: "John",
			Price: 500.0,
		}
		book.Chegirma(40)
	
		book.Disp()
	/*////////////////////////////////////////////////////////////////////////////////////////////////*/
	Car:=loyihalar.RentCar{
		Make: "Kia",
		Model: "k5",
		RentDay: 30.0,

	}	
	fmt.Println(Car.Count(14))
	Car.Display()
	////////////////////////////////////////////////////////////////////////////////////////////////////
	Maosh:=loyihalar.Employee {
		Name: "Nodir",
		Age: 18,
		Salary: 455.0,
	}
	Maosh.CountSalary(5)
	Maosh.Display()
	/////////////////////////////////////////////////////////////////////////////////////////////////////
	ticket:=loyihalar.OnlineTicket {
		TicketNumber: 123,
		EventName: "one belt",
		Price: 100.0,
		Status: "Aviable",
	}
	ticket.Check()
	ticket.Display()
	//////////////////////////////////////////////////////////////////////////////////////////////
	Mahsulot:=loyihalar.Product{
		Name: "Olma",
		Price: 5000.0,
		Quantity: 3,
	}
	Mahsulot.OldNew(1)
		Mahsulot.Display()
	Mahsulot.CalcProduct(3)
	Mahsulot.Display()
	////////////////////////////////////////////////////////////////////////////////////////////////
	// Hisoblash:=loyihalar.Geometriya{length: 5.0, width: 4.0 }
	// Area:=Hisoblash.Yuzi()
	// Sum:=Hisoblash.Perimetr()
	// fmt.Println(Area)
	// fmt.Println(Sum)
	///////////////////////////////////////////////////////////////////////////////////////////////////
	Talaba:=loyihalar.Student {
		Name: "Nodir",
		Grades: []float64{50,70,85},
	}
	Talaba.Checkst()
}