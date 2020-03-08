package main
import "fmt"
type car struct{
	gas_pedal uint64
	brake_pedal uint64
	steering uint64
	top_speed float64
}

func main() {
	a_car:= car{
							gas_pedal: 3456,
							brake_pedal: 0,
							steering: 12356,
							top_speed: 350.08,
	}
	fmt.Println(a_car.top_speed)
}