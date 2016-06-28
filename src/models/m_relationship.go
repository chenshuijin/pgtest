package models
import (
	"conf"
    "gopkg.in/pg.v4"
)

type Relationship struct {
	Id     string
	User_id   string
	State   string
	Type string
}

func GetRshipById(id string) ([]Relationship, error) {
	db := pg.Connect(conf.GetDbConf())
	defer db.Close()

	var rships []Relationship
	//err := db.Model(&rships).Select(rship)
	err := db.Model(&rships).Where("id = ?", id).Order("user_id").Select()
	return rships, err
}

func GetRshipByRid(id, oid string) (*Relationship, error) {
	db := pg.Connect(conf.GetDbConf())
	defer db.Close()
	// Select all users.
	rship := &Relationship{
		Id:id,
		User_id:oid,
	}
	err := db.Select(rship)
	return rship, err
}

func CreateOrUpdteRelationship(rship *Relationship) error {
	db := pg.Connect(conf.GetDbConf())
	defer db.Close()
	_, err := db.Model(&rship).OnConflict("(id, user_id) DO UPDATE").Set("state = ?state").Create()
	return err
}

func UpdateRelationship(rship *Relationship) error {
	db := pg.Connect(conf.GetDbConf())
	defer db.Close()
	err := db.Update(rship)
	return err
}