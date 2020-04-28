package models

import (
	log "github.com/sirupsen/logrus"
	"time"
)

type AgentInfo struct {
	Id                int       `json:"id" xorm:"not null pk autoincr comment('id') INT(11)"`
	Name              string    `json:"name" xorm:"not null default '' comment('姓名') VARCHAR(30)"`
	Mobile            string    `json:"mobile" xorm:"not null default '' comment('手机号') unique VARCHAR(20)"`
	IdNumber          string    `json:"id_number" xorm:"not null default '' comment('身份证号') VARCHAR(18)"`
	Gender            int       `json:"gender" xorm:"not null default 0 comment('性别') TINYINT(4)"`
	Email             string    `json:"email" xorm:"not null default '' comment('邮箱') VARCHAR(100)"`
	AlipayAccount     string    `json:"alipay_account" xorm:"not null default '' comment('支付宝帐号') VARCHAR(100)"`
	Avatar            string    `json:"avatar" xorm:"not null default '' comment('头像URL') VARCHAR(500)"`
	ProvideServices   string    `json:"provide_services" xorm:"not null default '' comment('可提供的服务标签') VARCHAR(50)"`
	ServiceCity       string    `json:"service_city" xorm:"not null default '' comment('服务范围-城市') VARCHAR(20)"`
	ServiceIndustry   string    `json:"service_industry" xorm:"not null default '' comment('专注行业标签') VARCHAR(20)"`
	CustomTags        string    `json:"custom_tags" xorm:"not null default '' comment('更多标签') VARCHAR(200)"`
	UserFrom          string    `json:"user_from" xorm:"not null default '' comment('用户来源') index VARCHAR(20)"`
	ProfileRegistered int       `json:"profile_registered" xorm:"not null default 0 comment('资料是否已经注册 0 未注册 1 已注册') TINYINT(1)"`
	ProfileFilled     int       `json:"profile_filled" xorm:"not null default 0 comment('资料是否已经全部完善 0 未全部完善 1 已全部完善') TINYINT(1)"`
	IdentityVerified  int       `json:"identity_verified" xorm:"not null default 0 comment('身份认证是否完成 0 未完成 1已完成') TINYINT(1)"`
	Status            int       `json:"status" xorm:"not null default 1 comment('状态 1 正常 2 禁用') TINYINT(1)"`
	Token             string    `json:"token" xorm:"not null default '' comment('token') VARCHAR(500)"`
	CreatedAt         time.Time `json:"created_at" xorm:"not null default 'CURRENT_TIMESTAMP' created comment('创建时间') TIMESTAMP"`
	UpdatedAt         time.Time `json:"updated_at" xorm:"not null default 'CURRENT_TIMESTAMP' updated comment('更新时间') TIMESTAMP"`
	DeletedAt         time.Time `json:"deleted_at" xorm:"deleted comment('删除时间') TIMESTAMP"`
}

func InsertAgentInfo(m *AgentInfo) error {
	var err error
	if _, err = X.Insert(m); err != nil {
		log.WithFields(log.Fields{"err": err, "AgentInfo": m}).Error("models.InsertAgentInfo has an error")
		err = DBErr
	}
	return err
}

func GetAgentInfo(id int) (m *AgentInfo, err error) {
	var has bool
	m = &AgentInfo{}
	if has, err = X.ID(id).Get(m); err != nil {
		log.WithFields(log.Fields{"err": err, "id": id}).Error("models.GetAgentInfo has an error")
		err = DBErr
	}
	if !has {
		m = nil
	}
	return
}

func UpdateAgentInfo(m *AgentInfo) error {
	var err error
	if _, err = X.Id(m.Id).AllCols().Update(m); err != nil {
		log.WithFields(log.Fields{"err": err, "AgentInfo": m}).Error("models.UpdateAgentInfo has an error")
		err = DBErr
	}
	return err
}

func DeleteAgentInfo(id int) error {
	var err error
	if _, err := X.Id(id).Delete(new(AgentInfo)); err != nil {
		log.WithFields(log.Fields{"err": err, "id": id}).Error("models.DeleteAgentInfo has an error")
		err = DBErr
	}
	return err
}
