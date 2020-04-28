package models

import (
	log "github.com/sirupsen/logrus"
	"time"
)

type Job struct {
	Id             int       `json:"id" xorm:"not null pk autoincr comment('id') INT(11)"`
	Name           string    `json:"name" xorm:"not null default '' comment('岗位名称') index VARCHAR(100)"`
	AgentId        int       `json:"agent_id" xorm:"not null default 0 comment('顾问ID') index INT(11)"`
	CompanyName    string    `json:"company_name" xorm:"not null default '' comment('单位名称') index VARCHAR(255)"`
	Industry       int       `json:"industry" xorm:"not null default 0 comment('所属行业') TINYINT(4)"`
	Department     string    `json:"department" xorm:"not null default '' comment('部门') VARCHAR(200)"`
	SalaryLow      int       `json:"salary_low" xorm:"not null default 0 comment('薪资low') INT(11)"`
	SalaryHigh     int       `json:"salary_high" xorm:"not null default 0 comment('薪资high') INT(11)"`
	City           int       `json:"city" xorm:"not null default 0 comment('工作城市') index INT(11)"`
	District       int       `json:"district" xorm:"not null default 0 comment('工作城市区域') INT(11)"`
	Location       string    `json:"location" xorm:"not null default '' comment('工作地点') VARCHAR(200)"`
	Degree         int       `json:"degree" xorm:"not null default 0 comment('学历要求') TINYINT(4)"`
	WorkExperience int       `json:"work_experience" xorm:"not null default 0 comment('工作经验') TINYINT(4)"`
	JobDescription string    `json:"job_description" xorm:"not null default '' comment('工作内容') VARCHAR(4096)"`
	Status         int       `json:"status" xorm:"not null default 0 comment('是否下架') index TINYINT(4)"`
	CreatedAt      time.Time `json:"created_at" xorm:"not null default 'CURRENT_TIMESTAMP' created comment('创建时间') TIMESTAMP"`
	UpdatedAt      time.Time `json:"updated_at" xorm:"not null default 'CURRENT_TIMESTAMP' updated comment('更新时间') TIMESTAMP"`
	DeletedAt      time.Time `json:"deleted_at" xorm:"deleted comment('删除时间') TIMESTAMP"`
}

func InsertJob(m *Job) error {
	var err error
	if _, err = X.Insert(m); err != nil {
		log.WithFields(log.Fields{"err": err, "Job": m}).Error("models.InsertJob has an error")
		err = DBErr
	}
	return err
}

func GetJob(id int) (m *Job, err error) {
	var has bool
	m = &Job{}
	if has, err = X.ID(id).Get(m); err != nil {
		log.WithFields(log.Fields{"err": err, "id": id}).Error("models.GetJob has an error")
		err = DBErr
	}
	if !has {
		m = nil
	}
	return
}

func UpdateJob(m *Job) error {
	var err error
	if _, err = X.Id(m.Id).AllCols().Update(m); err != nil {
		log.WithFields(log.Fields{"err": err, "Job": m}).Error("models.UpdateJob has an error")
		err = DBErr
	}
	return err
}

func DeleteJob(id int) error {
	var err error
	if _, err := X.Id(id).Delete(new(Job)); err != nil {
		log.WithFields(log.Fields{"err": err, "id": id}).Error("models.DeleteJob has an error")
		err = DBErr
	}
	return err
}
