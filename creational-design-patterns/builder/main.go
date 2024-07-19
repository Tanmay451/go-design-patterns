package builder

import "fmt"

func Print() {
	car := NewCar("Tata")
	car.SetYear(2020)
	car.SetMiles(1000.4)
	fmt.Println(car)
}
