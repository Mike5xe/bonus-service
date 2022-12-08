package draft

type Car struct {
	Label  string
	Engine Engine
	Color  Color
}
type Color struct {
	blue string
	red  string
}
type Engine struct {
	Power int
}

func NewCar(label string, engine Engine) Car {
	return Car{
		Label:  label,
		Engine: engine,
	}
}

func NewBMW() Car {
	return Car{
		Label: "BMW",
		Engine: Engine{
			Power: 10000000000000,
		},
	}
}
