package furniture

import "fmt"

type Furniture struct {
	Name     string
	Color    string
	Material string
}

func FillFurniture() []Furniture {
	sofa1 := Furniture{"Диван", "оранжевый", "текстиль"}
	sofa2 := Furniture{"Диван", "зеленый", "текстиль"}
	armchair := Furniture{"Кресло", "зеленое", "текстиль"}
	table := Furniture{"Обеденный стол", "коричневый", "дерево"}
	desk := Furniture{"Стол", "коричневый", "дерево"}
	wardrobe := Furniture{"Шкаф", "коричневый", "дерево"}
	commode := Furniture{"Комод", "коричневый", "дерево"}
	lamp := Furniture{"Лампа", "серебристо-черная", "металл, пластик"}
	return []Furniture{sofa1, sofa2, armchair, table, desk, wardrobe, commode, lamp}
}

func ShowFurniture(Furniture []Furniture) {
	fmt.Printf("Мебель в  доме:\n")
	for _, furs := range Furniture {
		fmt.Printf("Название - %s\n Цвет - %s\n  Материал - %s\n", furs.Name, furs.Color, furs.Material)
	}
	fmt.Print("\n")
}
