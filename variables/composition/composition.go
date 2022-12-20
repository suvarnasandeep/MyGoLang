package composition

import "fmt"

type Vehicle struct {
	numberOfWheel      int
	numberOfPassangers int
}

type Car struct {
	make    string
	model   string
	year    int
	vehicle Vehicle
}

func (v Vehicle) showVehicle() {
	fmt.Println("No of passengers : ", v.numberOfPassangers)
	fmt.Println("No of wheels", v.numberOfWheel)
}

func (c Car) showCar() {
	fmt.Println("Make : ", c.make)
	fmt.Println("Model : ", c.model)
	fmt.Println("Year : ", c.year)
	c.vehicle.showVehicle()

}

func TestComposition() {

	suv := Vehicle{
		numberOfPassangers: 7,
		numberOfWheel:      4,
	}

	car := Car{
		make:    "volvo",
		model:   "xc90",
		year:    2022,
		vehicle: suv,
	}

	car.showCar()

	carTesla := Car{
		make:    "Tesla",
		model:   "V9",
		year:    2021,
		vehicle: suv,
	}

	carTesla.year = 2000
	carTesla.showCar()

}
