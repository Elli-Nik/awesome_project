package main

import (
	"awesome_project/house_information/family"
	"awesome_project/house_information/furniture"
	"awesome_project/house_information/roomsAndsizes"
	"awesome_project/house_information/tech"
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

func DbShowFurniture() []furniture.Furniture {
	db, err := sql.Open("postgres", "postgres://awesome_project:123@localhost:5436/test_db?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query("SELECT name, color, material FROM furniture")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	furs := make([]*furniture.Furniture, 0)
	for rows.Next() {
		fur := new(furniture.Furniture)
		err := rows.Scan(&fur.Name, &fur.Color, &fur.Material)
		if err != nil {
			log.Fatal(err)
		}
		furs = append(furs, fur)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
	convertedFurniture := make([]furniture.Furniture, len(furs))
	for i, fur := range furs {
		convertedFurniture[i] = furniture.Furniture{
			Name:     fur.Name,
			Color:    fur.Color,
			Material: fur.Material,
		}
	}
	return convertedFurniture
}

func DbShowFamily() []family.Family {
	db, err := sql.Open("postgres", "postgres://awesome_project:123@localhost:5436/test_db?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query("SELECT name, age, qualification, hobby FROM family")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	fams := make([]*family.Family, 0)
	for rows.Next() {
		fam := new(family.Family)
		err := rows.Scan(&fam.Name, &fam.Age, &fam.Qualification, &fam.Hobby)
		if err != nil {
			log.Fatal(err)
		}
		fams = append(fams, fam)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
	convertedFamily := make([]family.Family, len(fams))
	for i, fam := range fams {
		convertedFamily[i] = family.Family{
			Name:          fam.Name,
			Age:           fam.Age,
			Qualification: fam.Qualification,
			Hobby:         fam.Hobby,
		}
	}
	return convertedFamily
}

func DbShowTech() []tech.Tech {
	db, err := sql.Open("postgres", "postgres://awesome_project:123@localhost:5436/test_db?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query("SELECT name, model, color FROM tech")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	techs := make([]*tech.Tech, 0)
	for rows.Next() {
		t := new(tech.Tech)
		err := rows.Scan(&t.Name, &t.Model, &t.Color)
		if err != nil {
			log.Fatal(err)
		}
		techs = append(techs, t)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
	convertedTech := make([]tech.Tech, len(techs))
	for i, t := range techs {
		convertedTech[i] = tech.Tech{
			Name:  t.Name,
			Model: t.Model,
			Color: t.Color,
		}
	}
	return convertedTech
}

func DbShowRoomsSizes() []roomsAndsizes.RoomsSizes {
	db, err := sql.Open("postgres", "postgres://awesome_project:123@localhost:5436/test_db?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query("SELECT name, value FROM roomsandsizes")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	sizes := make([]*roomsAndsizes.RoomsSizes, 0)
	for rows.Next() {
		size := new(roomsAndsizes.RoomsSizes)
		var err = rows.Scan(&size.Name, &size.Value)
		if err != nil {
			log.Fatal(err)
		}
		sizes = append(sizes, size)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
	convertedRoomsAndsizes := make([]roomsAndsizes.RoomsSizes, len(sizes))
	for i, size := range sizes {
		convertedRoomsAndsizes[i] = roomsAndsizes.RoomsSizes{
			Name:  size.Name,
			Value: size.Value,
		}
	}
	return convertedRoomsAndsizes
}
