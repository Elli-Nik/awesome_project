package family

import "fmt"

type Family struct {
	Name          string
	Age           int
	Qualification string
	Hobby         string
}

func FillFamily() []Family {
	member1 := Family{"Лиза", 20, "Студент", "tt"}
	member2 := Family{"Валентина", 76, "Пенсионер", "рисование"}
	cat := Family{"Фея", 2, "кошка", "фанатично спать"}
	return []Family{member1, member2, cat}
}

func ShowFamily(Family []Family) {
	fmt.Printf("Кто обитает в доме:\n")
	for _, members := range Family {
		fmt.Printf("Имя - %s\n Возраст - %d\n Кем является - %s\n Хобби - %s\n", members.Name, members.Age, members.Qualification, members.Hobby)
	}
	fmt.Print("\n")
}
