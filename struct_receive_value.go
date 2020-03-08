package main
import "fmt"

const usixteenbitmax float64 = 65535
const kmh_multiple float64 = 1.60934

type car struct {
	gas_pedal uint64
	brake_pedal uint64
	steering uint64
	top_speed float64
}

func (c car) kmh() float64{
    return float64(c.gas_pedal) * (c.top_speed/usixteenbitmax)
}

func (c car) mph() float64 {
	return float64(c.gas_pedal) * (c.top_speed/usixteenbitmax/kmh_multiple)
}

func main() {
   a_car := car {
		 gas_pedal: 34556,
		 brake_pedal: 0,
		 steering: 235,
		 top_speed: 235.990,
	 }
	 fmt.Println("Top Speed of my Car:: ", a_car.top_speed)
	 fmt.Println(a_car.kmh())
	 fmt.Println(a_car.mph())
}