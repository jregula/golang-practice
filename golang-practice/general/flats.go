package main
import (
"fmt"
)

type flat struct {
	size float64
	price float64
	location string
}

type house struct {
	size float64
	price float64
	location string
}


func(f *flat) calculatePrice() float64{
	return f.price *1.2
}

func(h *house) calculatePrice() float64{
	return h.price
}

type Pricer interface {
	calculatePrice() float64
}

func main() {
	flatA := &flat{
		12.0,
		4555.55,
		"camden",
	}
	costOfFlatA := Pricer(flatA)

	fmt.Println(costOfFlatA.calculatePrice())
}