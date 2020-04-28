package models

import (
	log "github.com/sirupsen/logrus"
	"time"
)

type Captcha struct {
	Id            int       `json:"id" xorm:"not null pk autoincr comment('id') INT(11)"`
	Receiver      string    `json:"receiver" xorm:"not null comment('接收方') index(idx_receiver_user_type_type) VARCHAR(11)"`
	UserType      int       `json:"user_type" xorm:"not null default 0 comment('用户类型') index(idx_receiver_user_type_type) TINYINT(4)"`
	Type          int       `json:"type" xorm:"not null comment('验证码类型') index(idx_receiver_user_type_type) TINYINT(4)"`
	Code          string    `json:"code" xorm:"not null default '' comment('code') VARCHAR(6)"`
	ExpiredSecond int       `json:"expired_second" xorm:"not null default 0 comment('过期时间/秒') INT(11)"`
	CreatedAt     time.Time `json:"created_at" xorm:"not null default 'CURRENT_TIMESTAMP' created comment('创建时间') TIMESTAMP"`
	UpdatedAt     time.Time `json:"updated_at" xorm:"not null default 'CURRENT_TIMESTAMP' updated comment('更新时间') TIMESTAMP"`
	DeletedAt     time.Time `json:"deleted_at" xorm:"deleted comment('删除时间') TIMESTAMP"`
}

func InsertCaptcha(m *Captcha) error {
	var err error
	if _, err = X.Insert(m); err != nil {
		log.WithFields(log.Fields{"err": err, "Captcha": m}).Error("models.InsertCaptcha has an error")
		err = DBErr
	}
	return err
}

func GetCaptcha(id int) (m *Captcha, err error) {
	var has bool
	m = &Captcha{}
	if has, err = X.ID(id).Get(m); err != nil {
		log.WithFields(log.Fields{"err": err, "id": id}).Error("models.GetCaptcha has an error")
		err = DBErr
	}
	if !has {
		m = nil
	}
	return
}

func UpdateCaptcha(m *Captcha) error {
	var err error
	if _, err = X.Id(m.Id).AllCols().Update(m); err != nil {
		log.WithFields(log.Fields{"err": err, "Captcha": m}).Error("models.UpdateCaptcha has an error")
		err = DBErr
	}
	return err
}

func DeleteCaptcha(id int) error {
	var err error
	if _, err := X.Id(id).Delete(new(Captcha)); err != nil {
		log.WithFields(log.Fields{"err": err, "id": id}).Error("models.DeleteCaptcha has an error")
		err = DBErr
	}
	return err
}
