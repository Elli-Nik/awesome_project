package roomsAndsizes

import (
	"fmt"
)

type RoomsSizes struct {
	Name  string
	Value int
}

func FillRoomsSizes() []RoomsSizes {
	ceiling := RoomsSizes{"Высота потолков, см: ", 260}
	square := RoomsSizes{"Площадь, кв.м: ", 50}
	numRooms := RoomsSizes{"Количество комнат: ", 3}

	return []RoomsSizes{ceiling, square, numRooms}
}

func ShowRoomsSizes(RoomsSizes []RoomsSizes) {
	for _, roomsSizes1 := range RoomsSizes {
		fmt.Printf("%s\n %d\n", roomsSizes1.Name, roomsSizes1.Value)
	}
	fmt.Print("\n")
}
