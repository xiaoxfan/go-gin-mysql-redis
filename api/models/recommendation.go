package models

import (
	log "github.com/sirupsen/logrus"
	"time"
)

type Recommendation struct {
	Id           int       `json:"id" xorm:"not null pk autoincr comment('id') INT(11)"`
	AgentId      int       `json:"agent_id" xorm:"not null default 0 comment('顾问id') unique(uniq_talent_job_agent) INT(11)"`
	TalentId     int       `json:"talent_id" xorm:"not null default 0 comment('人才id') unique(uniq_talent_job_agent) INT(11)"`
	JobId        int       `json:"job_id" xorm:"not null default 0 comment('职位id') unique(uniq_talent_job_agent) INT(11)"`
	ServiceFee   int       `json:"service_fee" xorm:"not null default 0 comment('服务费') INT(11)"`
	OrderId      int       `json:"order_id" xorm:"not null default 0 comment('订单id') index INT(11)"`
	HireCompany  string    `json:"hire_company" xorm:"not null default '' comment('录用单位') VARCHAR(256)"`
	HireDate     string    `json:"hire_date" xorm:"not null default '' comment('录用日期') VARCHAR(20)"`
	UnitRelation int       `json:"unit_relation" xorm:"not null default 0 comment('入职公司与实际公司关系') TINYINT(4)"`
	Status       int       `json:"status" xorm:"not null default 0 comment('推荐状态') index TINYINT(4)"`
	CreatedAt    time.Time `json:"created_at" xorm:"not null default 'CURRENT_TIMESTAMP' created comment('创建时间') TIMESTAMP"`
	UpdatedAt    time.Time `json:"updated_at" xorm:"not null default 'CURRENT_TIMESTAMP' updated comment('更新时间') TIMESTAMP"`
	DeletedAt    time.Time `json:"deleted_at" xorm:"deleted comment('删除时间') unique(uniq_talent_job_agent) TIMESTAMP"`
}

func InsertRecommendation(m *Recommendation) error {
	var err error
	if _, err = X.Insert(m); err != nil {
		log.WithFields(log.Fields{"err": err, "Recommendation": m}).Error("models.InsertRecommendation has an error")
		err = DBErr
	}
	return err
}

func GetRecommendation(id int) (m *Recommendation, err error) {
	var has bool
	m = &Recommendation{}
	if has, err = X.ID(id).Get(m); err != nil {
		log.WithFields(log.Fields{"err": err, "id": id}).Error("models.GetRecommendation has an error")
		err = DBErr
	}
	if !has {
		m = nil
	}
	return
}

func UpdateRecommendation(m *Recommendation) error {
	var err error
	if _, err = X.Id(m.Id).AllCols().Update(m); err != nil {
		log.WithFields(log.Fields{"err": err, "Recommendation": m}).Error("models.UpdateRecommendation has an error")
		err = DBErr
	}
	return err
}

func DeleteRecommendation(id int) error {
	var err error
	if _, err := X.Id(id).Delete(new(Recommendation)); err != nil {
		log.WithFields(log.Fields{"err": err, "id": id}).Error("models.DeleteRecommendation has an error")
		err = DBErr
	}
	return err
}
