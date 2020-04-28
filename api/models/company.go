package models

import (
	log "github.com/sirupsen/logrus"
	"time"
)

type Company struct {
	Id                   int64     `json:"id" xorm:"BIGINT(20)"`
	OutId                int64     `json:"out_id" xorm:"BIGINT(20)"`
	PercentileScore      int       `json:"percentile_score" xorm:"INT(11)"`
	CategoryScore        int       `json:"category_score" xorm:"INT(11)"`
	StaffNumRange        string    `json:"staff_num_range" xorm:"TINYTEXT"`
	Updatetime           int64     `json:"updatetime" xorm:"BIGINT(20)"`
	FromTime             int64     `json:"from_time" xorm:"BIGINT(20)"`
	Type                 int       `json:"type" xorm:"TINYINT(4)"`
	IsMicroEnt           int       `json:"is_micro_ent" xorm:"TINYINT(4)"`
	RegNumber            string    `json:"reg_number" xorm:"TINYTEXT"`
	RegCapital           string    `json:"reg_capital" xorm:"TINYTEXT"`
	Name                 string    `json:"name" xorm:"TINYTEXT"`
	RegInstitute         string    `json:"reg_institute" xorm:"TINYTEXT"`
	RegLocation          string    `json:"reg_location" xorm:"TINYTEXT"`
	Industry             string    `json:"industry" xorm:"TINYTEXT"`
	ApprovedTime         int64     `json:"approved_time" xorm:"BIGINT(20)"`
	SocialStaffNum       int       `json:"social_staff_num" xorm:"INT(11)"`
	Logo                 string    `json:"logo" xorm:"TINYTEXT"`
	TaxNumber            string    `json:"tax_number" xorm:"TINYTEXT"`
	BusinessScope        string    `json:"business_scope" xorm:"VARCHAR(4091)"`
	Property3            string    `json:"property3" xorm:"TINYTEXT"`
	Alias                string    `json:"alias" xorm:"TINYTEXT"`
	OrgNumber            string    `json:"org_number" xorm:"TINYTEXT"`
	RegStatus            string    `json:"reg_status" xorm:"TINYTEXT"`
	EstiblishTime        int64     `json:"estiblish_time" xorm:"BIGINT(20)"`
	LegalPersonName      string    `json:"legal_person_name" xorm:"TINYTEXT"`
	ToTime               int64     `json:"to_time" xorm:"BIGINT(20)"`
	ActualCapital        string    `json:"actual_capital" xorm:"TINYTEXT"`
	CompanyOrgType       string    `json:"company_org_type" xorm:"TINYTEXT"`
	Base                 string    `json:"base" xorm:"TINYTEXT"`
	CreditCode           string    `json:"credit_code" xorm:"TINYTEXT"`
	Email                string    `json:"email" xorm:"TINYTEXT"`
	WebsiteList          string    `json:"website_list" xorm:"TINYTEXT"`
	PhoneNumber          string    `json:"phone_number" xorm:"TINYTEXT"`
	UpdateTimes          int64     `json:"update_times" xorm:"BIGINT(20)"`
	LegalPersonId        int64     `json:"legal_person_id" xorm:"BIGINT(20)"`
	OrgApprovedInstitute string    `json:"org_approved_institute" xorm:"TINYTEXT"`
	SourceFlag           string    `json:"source_flag" xorm:"TINYTEXT"`
	IsClaimed            int       `json:"is_claimed" xorm:"INT(11)"`
	CreatedAt            time.Time `json:"created_at" xorm:"created TIMESTAMP"`
	UpdatedAt            time.Time `json:"updated_at" xorm:"updated TIMESTAMP"`
	DeletedAt            time.Time `json:"deleted_at" xorm:"deleted TIMESTAMP"`
}

func InsertCompany(m *Company) error {
	var err error
	if _, err = X.Insert(m); err != nil {
		log.WithFields(log.Fields{"err": err, "Company": m}).Error("models.InsertCompany has an error")
		err = DBErr
	}
	return err
}

func GetCompany(id int) (m *Company, err error) {
	var has bool
	m = &Company{}
	if has, err = X.ID(id).Get(m); err != nil {
		log.WithFields(log.Fields{"err": err, "id": id}).Error("models.GetCompany has an error")
		err = DBErr
	}
	if !has {
		m = nil
	}
	return
}

func UpdateCompany(m *Company) error {
	var err error
	if _, err = X.Id(m.Id).AllCols().Update(m); err != nil {
		log.WithFields(log.Fields{"err": err, "Company": m}).Error("models.UpdateCompany has an error")
		err = DBErr
	}
	return err
}

func DeleteCompany(id int) error {
	var err error
	if _, err := X.Id(id).Delete(new(Company)); err != nil {
		log.WithFields(log.Fields{"err": err, "id": id}).Error("models.DeleteCompany has an error")
		err = DBErr
	}
	return err
}
