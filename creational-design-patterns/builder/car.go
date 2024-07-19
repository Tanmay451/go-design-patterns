package builder

type Car struct {
	Brand string
	Year  int
	Miles float64
}

func NewCar(brand string) *Car {
	return &Car{
		Brand: brand,
	}
}

func (c *Car) SetYear(year int) *Car {
	c.Year = year
	return c
}

func (c *Car) SetMiles(miles float64) *Car {
	c.Miles = miles
	return c
}
