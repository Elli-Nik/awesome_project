package main

import (
	"awesome_project/house_information/family"
	"awesome_project/house_information/furniture"
	"awesome_project/house_information/roomsAndsizes"
	"awesome_project/house_information/tech"
)

func FullHouseDb() {
	furs := DbShowFurniture()
	fams := DbShowFamily()
	techs := DbShowTech()
	sizes := DbShowRoomsSizes()

	furniture.ShowFurniture(furs)
	family.ShowFamily(fams)
	tech.ShowTech(techs)
	roomsAndsizes.ShowRoomsSizes(sizes)
}
