package full_house

import (
	"awesome_project/house_information/family"
	"awesome_project/house_information/furniture"
	"awesome_project/house_information/roomsAndsizes"
	"awesome_project/house_information/tech"
)

func FullHouse() {
	Family := family.FillFamily()
	RoomsSizes := roomsAndsizes.FillRoomsSizes()
	Furniture := furniture.FillFurniture()
	Tech := tech.FillTech()

	family.ShowFamily(Family)
	roomsAndsizes.ShowRoomsSizes(RoomsSizes)
	furniture.ShowFurniture(Furniture)
	tech.ShowTech(Tech)
}
