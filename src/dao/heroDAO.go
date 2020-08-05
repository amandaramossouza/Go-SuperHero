package dao

import (
	"log"
	"fmt"
	"models"
)

//GetAllHeroes retrives all users
func GetAllHeroes() ([]models.Hero, error) {
	heroes := []models.Hero{}
	fmt.Println("heroDAO: line 12")
	rows, err := db.Query(`SELECT uuid, id, name, fullname, intelligence, power, occupation, image, groups, relatives, numrelatives  FROM heroes order by uuid`)
	fmt.Println("heroDAO: line 14")
	fmt.Println(err)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	fmt.Println("heroDAO: line 20")
	for rows.Next() {
		var Uuid			int
		var Id				int
		var Name 			string
		var Fullname		string
		var Intelligence	int
		var Power			int
		var Occupation		int
		var Image			string
		var Groups			string
		var Relatives		string
		var Numrelatives	int
		fmt.Println("heroDAO: line 31")
		err = rows.Scan(&Uuid, &Id, &Name, &Fullname, &Intelligence, &Power, &Occupation, &Image, &Groups, &Relatives, &Numrelatives)
		if err == nil {
			currentHero := models.Hero{Uuid: Uuid, Id: Id, Name: Name, Fullname: Fullname, 
							Intelligence: Intelligence, Power: Power, Occupation: Occupation,
							Image: Image, Groups: Groups, Relatives: Relatives, Numrelatives: Numrelatives}
			heroes = append(heroes, currentHero)
		} else {
			log.Println(err)
			return nil, err
		}
	}
	fmt.Println("heroDAO: line 43")
	return heroes, err
}