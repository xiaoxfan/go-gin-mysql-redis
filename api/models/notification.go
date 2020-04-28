package models

import (
	log "github.com/sirupsen/logrus"
	"time"
)

type Notification struct {
	Id         int       `json:"id" xorm:"not null pk autoincr comment('id') INT(11)"`
	UserId     int       `json:"user_id" xorm:"not null default 0 comment('用户id') index(idx_user_id_type) INT(11)"`
	UserType   int       `json:"user_type" xorm:"not null default 0 comment('用户类型') index(idx_user_id_type) TINYINT(4)"`
	Type       int       `json:"type" xorm:"not null default 0 comment('通知类型') index TINYINT(4)"`
	TargetId   int       `json:"target_id" xorm:"not null default 0 comment('目标对象id') INT(11)"`
	ReadStatus int       `json:"read_status" xorm:"not null default 0 comment('是否已读 1 已读') index TINYINT(4)"`
	Content    string    `json:"content" xorm:"not null default '' comment('通知内容json') VARCHAR(500)"`
	CreatedAt  time.Time `json:"created_at" xorm:"not null default 'CURRENT_TIMESTAMP' created comment('创建时间') TIMESTAMP"`
	UpdatedAt  time.Time `json:"updated_at" xorm:"not null default 'CURRENT_TIMESTAMP' updated comment('更新时间') TIMESTAMP"`
	DeletedAt  time.Time `json:"deleted_at" xorm:"deleted comment('删除时间') TIMESTAMP"`
}

func InsertNotification(m *Notification) error {
	var err error
	if _, err = X.Insert(m); err != nil {
		log.WithFields(log.Fields{"err": err, "Notification": m}).Error("models.InsertNotification has an error")
		err = DBErr
	}
	return err
}

func GetNotification(id int) (m *Notification, err error) {
	var has bool
	m = &Notification{}
	if has, err = X.ID(id).Get(m); err != nil {
		log.WithFields(log.Fields{"err": err, "id": id}).Error("models.GetNotification has an error")
		err = DBErr
	}
	if !has {
		m = nil
	}
	return
}

func UpdateNotification(m *Notification) error {
	var err error
	if _, err = X.Id(m.Id).AllCols().Update(m); err != nil {
		log.WithFields(log.Fields{"err": err, "Notification": m}).Error("models.UpdateNotification has an error")
		err = DBErr
	}
	return err
}

func DeleteNotification(id int) error {
	var err error
	if _, err := X.Id(id).Delete(new(Notification)); err != nil {
		log.WithFields(log.Fields{"err": err, "id": id}).Error("models.DeleteNotification has an error")
		err = DBErr
	}
	return err
}
