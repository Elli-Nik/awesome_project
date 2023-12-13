package tech

import "fmt"

type Tech struct {
	Name  string
	Model string
	Color string
}

func FillTech() []Tech {
	washingMachine := Tech{"Стиральная машина", "Zanussi", "Белый"}
	fridge := Tech{"Холодильник", "LG", "Серый"}
	microwave := Tech{"Микроволновая печь", "Rolsen", "Серебристый"}
	lamp := Tech{"Лампа", "Camelion", "Белый"}
	tv := Tech{"Телевизор", "LG", "Серый"}
	return []Tech{washingMachine, fridge, microwave, lamp, tv}
}

func ShowTech(Tech []Tech) {
	fmt.Printf("Технические приборы:\n")
	for _, tech := range Tech {
		fmt.Printf("Название - %s\n Модель - %s\n Цвет - %s\n ", tech.Name, tech.Model, tech.Color)
	}
	fmt.Print("\n")
}
