package models

import (
	log "github.com/sirupsen/logrus"
	"time"
)

type Major struct {
	Id        int       `json:"id" xorm:"INT(11)"`
	Name      string    `json:"name" xorm:"TINYTEXT"`
	CreatedAt time.Time `json:"created_at" xorm:"created TIMESTAMP"`
	UpdatedAt time.Time `json:"updated_at" xorm:"updated TIMESTAMP"`
	DeletedAt time.Time `json:"deleted_at" xorm:"deleted TIMESTAMP"`
}

func InsertMajor(m *Major) error {
	var err error
	if _, err = X.Insert(m); err != nil {
		log.WithFields(log.Fields{"err": err, "Major": m}).Error("models.InsertMajor has an error")
		err = DBErr
	}
	return err
}

func GetMajor(id int) (m *Major, err error) {
	var has bool
	m = &Major{}
	if has, err = X.ID(id).Get(m); err != nil {
		log.WithFields(log.Fields{"err": err, "id": id}).Error("models.GetMajor has an error")
		err = DBErr
	}
	if !has {
		m = nil
	}
	return
}

func UpdateMajor(m *Major) error {
	var err error
	if _, err = X.Id(m.Id).AllCols().Update(m); err != nil {
		log.WithFields(log.Fields{"err": err, "Major": m}).Error("models.UpdateMajor has an error")
		err = DBErr
	}
	return err
}

func DeleteMajor(id int) error {
	var err error
	if _, err := X.Id(id).Delete(new(Major)); err != nil {
		log.WithFields(log.Fields{"err": err, "id": id}).Error("models.DeleteMajor has an error")
		err = DBErr
	}
	return err
}
