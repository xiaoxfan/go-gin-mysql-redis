package models

import (
	log "github.com/sirupsen/logrus"
	"time"
)

type Operator struct {
	Id         int       `json:"id" xorm:"not null pk autoincr comment('主键') INT(11)"`
	Name       string    `json:"name" xorm:"not null default '' comment('姓名') VARCHAR(20)"`
	Mobile     string    `json:"mobile" xorm:"not null default '' comment('手机号') unique VARCHAR(11)"`
	Department string    `json:"department" xorm:"not null default '' comment('部门') VARCHAR(20)"`
	IsAdmin    int       `json:"is_admin" xorm:"not null default 0 comment('是否管理员') TINYINT(1)"`
	Position   string    `json:"position" xorm:"not null default '' comment('岗位') VARCHAR(20)"`
	Email      string    `json:"email" xorm:"not null default '' comment('邮箱') VARCHAR(50)"`
	Password   string    `json:"password" xorm:"not null default '' comment('密码md5值') VARCHAR(32)"`
	Token      string    `json:"token" xorm:"not null default '' comment('token') VARCHAR(400)"`
	Status     int       `json:"status" xorm:"not null default 1 comment('状态 1 启用 2 停用') TINYINT(2)"`
	CreatedAt  time.Time `json:"created_at" xorm:"not null default 'CURRENT_TIMESTAMP' created comment('创建时间') TIMESTAMP"`
	UpdatedAt  time.Time `json:"updated_at" xorm:"not null default 'CURRENT_TIMESTAMP' updated comment('更新时间') TIMESTAMP"`
	DeletedAt  time.Time `json:"deleted_at" xorm:"deleted comment('删除时间') TIMESTAMP"`
}

func InsertOperator(m *Operator) error {
	var err error
	if _, err = X.Insert(m); err != nil {
		log.WithFields(log.Fields{"err": err, "Operator": m}).Error("models.InsertOperator has an error")
		err = DBErr
	}
	return err
}

func GetOperator(id int) (m *Operator, err error) {
	var has bool
	m = &Operator{}
	if has, err = X.ID(id).Get(m); err != nil {
		log.WithFields(log.Fields{"err": err, "id": id}).Error("models.GetOperator has an error")
		err = DBErr
	}
	if !has {
		m = nil
	}
	return
}

func UpdateOperator(m *Operator) error {
	var err error
	if _, err = X.Id(m.Id).AllCols().Update(m); err != nil {
		log.WithFields(log.Fields{"err": err, "Operator": m}).Error("models.UpdateOperator has an error")
		err = DBErr
	}
	return err
}

func DeleteOperator(id int) error {
	var err error
	if _, err := X.Id(id).Delete(new(Operator)); err != nil {
		log.WithFields(log.Fields{"err": err, "id": id}).Error("models.DeleteOperator has an error")
		err = DBErr
	}
	return err
}
