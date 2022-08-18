package main

import "fmt"

type House struct {
	SquareMeters int
	Material     string
	Rooms        int
	Roof         string
}

type HouseBuilderI interface {
	SquareMeters(v int) HouseBuilderI
	Material(v string) HouseBuilderI
	Rooms(v int) HouseBuilderI
	Roof(v string) HouseBuilderI
	Build() House
}

type houseBuilder struct {
	squareMeters int
	material     string
	rooms        int
	roof         string
}

func (h houseBuilder) SquareMeters(v int) HouseBuilderI {
	h.squareMeters = v
	return h
}
func (h houseBuilder) Material(v string) HouseBuilderI {
	h.material = v
	return h
}
func (h houseBuilder) Rooms(v int) HouseBuilderI {
	h.rooms = v
	return h
}
func (h houseBuilder) Roof(v string) HouseBuilderI {
	h.roof = v
	return h
}
func (h houseBuilder) Build() House {
	return House{
		SquareMeters: h.squareMeters,
		Material:     h.material,
		Rooms:        h.rooms,
		Roof:         h.roof,
	}
}

func NewHouseBuilder() HouseBuilderI {
	return houseBuilder{}
}

type preparedHouse struct {
	houseBuilder
}

func (p preparedHouse) Build() House {
	return House{
		SquareMeters: 150,
		Material:     "Дерево",
		Rooms:        6,
		Roof:         "черепица",
	}
}

func NewPreparedHouseBuilder() HouseBuilderI {
	return preparedHouse{}
}
func main() {
	houseBuilder := NewHouseBuilder()
	house := houseBuilder.SquareMeters(100).Material("Кирпич").Rooms(5).Roof("шифер").Build()
	fmt.Println(house)
	preparedHouseBuilder := NewPreparedHouseBuilder()
	preparedHouse := preparedHouseBuilder.Build()
	fmt.Println(preparedHouse)
}
