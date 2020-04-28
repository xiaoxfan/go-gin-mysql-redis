package models

import (
	log "github.com/sirupsen/logrus"
	"time"
)

type Position struct {
	Id        int64     `json:"id" xorm:"BIGINT(20)"`
	Name      string    `json:"name" xorm:"TINYTEXT"`
	CreatedAt time.Time `json:"created_at" xorm:"created TIMESTAMP"`
	UpdatedAt time.Time `json:"updated_at" xorm:"updated TIMESTAMP"`
	DeletedAt time.Time `json:"deleted_at" xorm:"deleted TIMESTAMP"`
}

func InsertPosition(m *Position) error {
	var err error
	if _, err = X.Insert(m); err != nil {
		log.WithFields(log.Fields{"err": err, "Position": m}).Error("models.InsertPosition has an error")
		err = DBErr
	}
	return err
}

func GetPosition(id int) (m *Position, err error) {
	var has bool
	m = &Position{}
	if has, err = X.ID(id).Get(m); err != nil {
		log.WithFields(log.Fields{"err": err, "id": id}).Error("models.GetPosition has an error")
		err = DBErr
	}
	if !has {
		m = nil
	}
	return
}

func UpdatePosition(m *Position) error {
	var err error
	if _, err = X.Id(m.Id).AllCols().Update(m); err != nil {
		log.WithFields(log.Fields{"err": err, "Position": m}).Error("models.UpdatePosition has an error")
		err = DBErr
	}
	return err
}

func DeletePosition(id int) error {
	var err error
	if _, err := X.Id(id).Delete(new(Position)); err != nil {
		log.WithFields(log.Fields{"err": err, "id": id}).Error("models.DeletePosition has an error")
		err = DBErr
	}
	return err
}
