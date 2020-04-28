package models

import (
	log "github.com/sirupsen/logrus"
	"time"
)

type College struct {
	Id        int       `json:"id" xorm:"INT(11)"`
	Name      string    `json:"name" xorm:"not null index VARCHAR(100)"`
	CreatedAt time.Time `json:"created_at" xorm:"created TIMESTAMP"`
	UpdatedAt time.Time `json:"updated_at" xorm:"updated TIMESTAMP"`
	DeletedAt time.Time `json:"deleted_at" xorm:"deleted TIMESTAMP"`
}

func InsertCollege(m *College) error {
	var err error
	if _, err = X.Insert(m); err != nil {
		log.WithFields(log.Fields{"err": err, "College": m}).Error("models.InsertCollege has an error")
		err = DBErr
	}
	return err
}

func GetCollege(id int) (m *College, err error) {
	var has bool
	m = &College{}
	if has, err = X.ID(id).Get(m); err != nil {
		log.WithFields(log.Fields{"err": err, "id": id}).Error("models.GetCollege has an error")
		err = DBErr
	}
	if !has {
		m = nil
	}
	return
}

func UpdateCollege(m *College) error {
	var err error
	if _, err = X.Id(m.Id).AllCols().Update(m); err != nil {
		log.WithFields(log.Fields{"err": err, "College": m}).Error("models.UpdateCollege has an error")
		err = DBErr
	}
	return err
}

func DeleteCollege(id int) error {
	var err error
	if _, err := X.Id(id).Delete(new(College)); err != nil {
		log.WithFields(log.Fields{"err": err, "id": id}).Error("models.DeleteCollege has an error")
		err = DBErr
	}
	return err
}
