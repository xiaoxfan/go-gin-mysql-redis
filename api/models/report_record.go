package models

import (
	log "github.com/sirupsen/logrus"
	"time"
)

type ReportRecord struct {
	Id        int       `json:"id" xorm:"not null pk autoincr comment('id') INT(11)"`
	Title     string    `json:"title" xorm:"not null default '' comment('投诉事项') VARCHAR(200)"`
	Content   string    `json:"content" xorm:"not null comment('投诉内容') VARCHAR(400)"`
	Creator   int       `json:"creator" xorm:"not null default 0 comment('投诉人') INT(11)"`
	CreatedAt time.Time `json:"created_at" xorm:"not null default 'CURRENT_TIMESTAMP' created comment('创建时间') TIMESTAMP"`
	UpdatedAt time.Time `json:"updated_at" xorm:"not null default 'CURRENT_TIMESTAMP' updated comment('更新时间') TIMESTAMP"`
	DeletedAt time.Time `json:"deleted_at" xorm:"deleted comment('删除时间') TIMESTAMP"`
}

func InsertReportRecord(m *ReportRecord) error {
	var err error
	if _, err = X.Insert(m); err != nil {
		log.WithFields(log.Fields{"err": err, "ReportRecord": m}).Error("models.InsertReportRecord has an error")
		err = DBErr
	}
	return err
}

func GetReportRecord(id int) (m *ReportRecord, err error) {
	var has bool
	m = &ReportRecord{}
	if has, err = X.ID(id).Get(m); err != nil {
		log.WithFields(log.Fields{"err": err, "id": id}).Error("models.GetReportRecord has an error")
		err = DBErr
	}
	if !has {
		m = nil
	}
	return
}

func UpdateReportRecord(m *ReportRecord) error {
	var err error
	if _, err = X.Id(m.Id).AllCols().Update(m); err != nil {
		log.WithFields(log.Fields{"err": err, "ReportRecord": m}).Error("models.UpdateReportRecord has an error")
		err = DBErr
	}
	return err
}

func DeleteReportRecord(id int) error {
	var err error
	if _, err := X.Id(id).Delete(new(ReportRecord)); err != nil {
		log.WithFields(log.Fields{"err": err, "id": id}).Error("models.DeleteReportRecord has an error")
		err = DBErr
	}
	return err
}
