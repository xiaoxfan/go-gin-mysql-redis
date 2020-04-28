package models

import (
	log "github.com/sirupsen/logrus"
	"time"
)

type TalentInfo struct {
	Id                int       `json:"id" xorm:"not null pk autoincr comment('id') INT(11)"`
	Name              string    `json:"name" xorm:"not null default '' comment('姓名') VARCHAR(30)"`
	Age               int       `json:"age" xorm:"not null default 0 comment('年龄') TINYINT(4)"`
	IdNumber          string    `json:"id_number" xorm:"not null default '' comment('身份证号') VARCHAR(18)"`
	Mobile            string    `json:"mobile" xorm:"not null default '' comment('手机号') unique VARCHAR(20)"`
	Gender            int       `json:"gender" xorm:"not null default 0 comment('性别') TINYINT(4)"`
	Email             string    `json:"email" xorm:"not null default '' comment('邮箱') VARCHAR(100)"`
	Avatar            string    `json:"avatar" xorm:"not null default '' comment('头像URL') VARCHAR(500)"`
	IntentionJobType  int       `json:"intention_job_type" xorm:"not null default 0 comment('求职意向 1 全职 2 实习 3 管培生') TINYINT(1)"`
	IntentionJobs     string    `json:"intention_jobs" xorm:"not null default '' comment('意向求职岗位') VARCHAR(256)"`
	IntentionJobLevel int       `json:"intention_job_level" xorm:"not null default 0 comment('求职岗位职级') TINYINT(1)"`
	ExpectedSalary    int       `json:"expected_salary" xorm:"not null default 0 comment('期望薪资') TINYINT(1)"`
	ExpectedCity      int       `json:"expected_city" xorm:"not null default 0 comment('期望城市') INT(11)"`
	ExpectedDistrict  int       `json:"expected_district" xorm:"not null default 0 comment('期望城市行政区') INT(11)"`
	ExpectedIndustry  string    `json:"expected_industry" xorm:"not null default '' comment('求职行业') VARCHAR(20)"`
	TargetUnitType    int       `json:"target_unit_type" xorm:"not null default 0 comment('定向求职') TINYINT(1)"`
	TargetUnit        string    `json:"target_unit" xorm:"not null default '' comment('定向求职单位') VARCHAR(256)"`
	WillingServiceFee int       `json:"willing_service_fee" xorm:"not null default 0 comment('愿意支付的服务费') INT(11)"`
	EmploymentStatus  int       `json:"employment_status" xorm:"not null default 0 comment('当前就业状态') TINYINT(1)"`
	SchoolLevel       int       `json:"school_level" xorm:"not null default 0 comment('学校类型') TINYINT(1)"`
	Degree            int       `json:"degree" xorm:"not null default 0 comment('学历层次') TINYINT(1)"`
	StudyType         int       `json:"study_type" xorm:"not null default 0 comment('学习形式') TINYINT(1)"`
	UniversityName    string    `json:"university_name" xorm:"not null default '' comment('学校全称') VARCHAR(200)"`
	SchoolName        string    `json:"school_name" xorm:"not null default '' comment('学院名称') VARCHAR(200)"`
	MajorName         string    `json:"major_name" xorm:"not null default '' comment('专业名称') VARCHAR(200)"`
	GraduateTime      string    `json:"graduate_time" xorm:"not null default '' comment('毕业时间') VARCHAR(30)"`
	LatestCompanyName string    `json:"latest_company_name" xorm:"not null default '' comment('最近工作单位名称') VARCHAR(200)"`
	LatestJobLevel    int       `json:"latest_job_level" xorm:"not null default 0 comment('最近工作的岗位职级') TINYINT(1)"`
	LatestPosition    string    `json:"latest_position" xorm:"not null default '' comment('最近工作岗位') VARCHAR(200)"`
	LatestDepartment  string    `json:"latest_department" xorm:"not null default '' comment('最近工作部门') VARCHAR(200)"`
	WorkYear          int       `json:"work_year" xorm:"not null default 0 comment('工作年限') TINYINT(1)"`
	CustomTags        string    `json:"custom_tags" xorm:"not null default '' comment('自定义身份标签') VARCHAR(200)"`
	UserFrom          string    `json:"user_from" xorm:"not null default '' comment('用户来源') index VARCHAR(20)"`
	BiaAccessToken    string    `json:"bia_access_token" xorm:"not null default '' comment('微背调报告access token') VARCHAR(256)"`
	BiaReportId       int       `json:"bia_report_id" xorm:"not null default 0 comment('微背调报告id') INT(11)"`
	ProfileRegistered int       `json:"profile_registered" xorm:"not null default 0 comment('资料是否已经注册 0 未注册 1 已注册') TINYINT(1)"`
	ProfileFilled     int       `json:"profile_filled" xorm:"not null default 0 comment('资料是否已经全部完善 0 未全部完善 1 已全部完善') TINYINT(1)"`
	IdentityVerified  int       `json:"identity_verified" xorm:"not null default 0 comment('身份认证是否完成 0 未完成 1已完成') TINYINT(1)"`
	Status            int       `json:"status" xorm:"not null default 1 comment('状态 1 正常 2 禁用') TINYINT(1)"`
	Token             string    `json:"token" xorm:"not null default '' comment('token') VARCHAR(500)"`
	CreatedAt         time.Time `json:"created_at" xorm:"not null default 'CURRENT_TIMESTAMP' created comment('创建时间') TIMESTAMP"`
	UpdatedAt         time.Time `json:"updated_at" xorm:"not null default 'CURRENT_TIMESTAMP' updated comment('更新时间') TIMESTAMP"`
	DeletedAt         time.Time `json:"deleted_at" xorm:"deleted comment('删除时间') TIMESTAMP"`
}

func InsertTalentInfo(m *TalentInfo) error {
	var err error
	if _, err = X.Insert(m); err != nil {
		log.WithFields(log.Fields{"err": err, "TalentInfo": m}).Error("models.InsertTalentInfo has an error")
		err = DBErr
	}
	return err
}

func GetTalentInfo(id int) (m *TalentInfo, err error) {
	var has bool
	m = &TalentInfo{}
	if has, err = X.ID(id).Get(m); err != nil {
		log.WithFields(log.Fields{"err": err, "id": id}).Error("models.GetTalentInfo has an error")
		err = DBErr
	}
	if !has {
		m = nil
	}
	return
}

func UpdateTalentInfo(m *TalentInfo) error {
	var err error
	if _, err = X.Id(m.Id).AllCols().Update(m); err != nil {
		log.WithFields(log.Fields{"err": err, "TalentInfo": m}).Error("models.UpdateTalentInfo has an error")
		err = DBErr
	}
	return err
}

func DeleteTalentInfo(id int) error {
	var err error
	if _, err := X.Id(id).Delete(new(TalentInfo)); err != nil {
		log.WithFields(log.Fields{"err": err, "id": id}).Error("models.DeleteTalentInfo has an error")
		err = DBErr
	}
	return err
}
