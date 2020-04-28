package models

import (
	log "github.com/sirupsen/logrus"
	"time"
)

type ServiceAgreement struct {
	Id               int       `json:"id" xorm:"not null pk autoincr comment('id') INT(11)"`
	RecommendationId int       `json:"recommendation_id" xorm:"not null default 0 comment('推荐id') index INT(11)"`
	Number           string    `json:"number" xorm:"not null default '' comment('协议编号') VARCHAR(20)"`
	AgentId          int       `json:"agent_id" xorm:"not null default 0 comment('agent id') INT(11)"`
	AgentInfo        string    `json:"agent_info" xorm:"not null default '' comment('经纪人信息') VARCHAR(60)"`
	AgentServices    string    `json:"agent_services" xorm:"not null default '0' comment('经纪人的服务') VARCHAR(500)"`
	TalentId         int       `json:"talent_id" xorm:"not null default 0 comment('talent id') INT(11)"`
	TalentInfo       string    `json:"talent_info" xorm:"not null default '' comment('人才信息') VARCHAR(60)"`
	JobId            int       `json:"job_id" xorm:"not null default 0 comment('job id') INT(11)"`
	JobNo            string    `json:"job_no" xorm:"not null default '' comment('职位编号') VARCHAR(20)"`
	CompanyName      string    `json:"company_name" xorm:"not null default '' comment('单位名称') VARCHAR(255)"`
	Department       string    `json:"department" xorm:"not null default '' comment('部门') VARCHAR(200)"`
	JobName          string    `json:"job_name" xorm:"not null default '' comment('职位名称') VARCHAR(200)"`
	JobLevel         int       `json:"job_level" xorm:"not null default 0 comment('岗位职级') INT(11)"`
	JobCity          int       `json:"job_city" xorm:"not null default 0 comment('工作城市') INT(11)"`
	JobDistrict      int       `json:"job_district" xorm:"not null default 0 comment('工作城市区域') INT(11)"`
	JobLocation      string    `json:"job_location" xorm:"not null default '' comment('工作地点') VARCHAR(200)"`
	ServiceFee       int       `json:"service_fee" xorm:"not null default 0 comment('服务费') INT(11)"`
	CreatedAt        time.Time `json:"created_at" xorm:"not null default 'CURRENT_TIMESTAMP' created comment('创建时间') TIMESTAMP"`
	UpdatedAt        time.Time `json:"updated_at" xorm:"not null default 'CURRENT_TIMESTAMP' updated comment('更新时间') TIMESTAMP"`
	DeletedAt        time.Time `json:"deleted_at" xorm:"deleted comment('删除时间') TIMESTAMP"`
}

func InsertServiceAgreement(m *ServiceAgreement) error {
	var err error
	if _, err = X.Insert(m); err != nil {
		log.WithFields(log.Fields{"err": err, "ServiceAgreement": m}).Error("models.InsertServiceAgreement has an error")
		err = DBErr
	}
	return err
}

func GetServiceAgreement(id int) (m *ServiceAgreement, err error) {
	var has bool
	m = &ServiceAgreement{}
	if has, err = X.ID(id).Get(m); err != nil {
		log.WithFields(log.Fields{"err": err, "id": id}).Error("models.GetServiceAgreement has an error")
		err = DBErr
	}
	if !has {
		m = nil
	}
	return
}

func UpdateServiceAgreement(m *ServiceAgreement) error {
	var err error
	if _, err = X.Id(m.Id).AllCols().Update(m); err != nil {
		log.WithFields(log.Fields{"err": err, "ServiceAgreement": m}).Error("models.UpdateServiceAgreement has an error")
		err = DBErr
	}
	return err
}

func DeleteServiceAgreement(id int) error {
	var err error
	if _, err := X.Id(id).Delete(new(ServiceAgreement)); err != nil {
		log.WithFields(log.Fields{"err": err, "id": id}).Error("models.DeleteServiceAgreement has an error")
		err = DBErr
	}
	return err
}
