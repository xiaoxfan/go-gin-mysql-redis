package models

import (
	log "github.com/sirupsen/logrus"
	"time"
)

type OrderInfo struct {
	Id        int       `json:"id" xorm:"not null pk autoincr comment('id') INT(11)"`
	UserId    int       `json:"user_id" xorm:"not null comment('用户id') index INT(11)"`
	UserType  int       `json:"user_type" xorm:"not null comment('用户类型') index TINYINT(4)"`
	OrderType int       `json:"order_type" xorm:"not null comment('订单类型') index TINYINT(4)"`
	Amount    int       `json:"amount" xorm:"not null default 0 comment('总金额/分') INT(11)"`
	OrderNo   string    `json:"order_no" xorm:"comment('支付订单号') index VARCHAR(45)"`
	OutNo     string    `json:"out_no" xorm:"comment('第三方支付订单号') VARCHAR(45)"`
	PrepayId  string    `json:"prepay_id" xorm:"comment('prepay_id 微信') VARCHAR(45)"`
	PayUrl    string    `json:"pay_url" xorm:"not null default '' comment('支付url') VARCHAR(4096)"`
	PayWay    int       `json:"pay_way" xorm:"not null default 0 comment('支付方式') TINYINT(4)"`
	PayTime   time.Time `json:"pay_time" xorm:"comment('支付时间') TIMESTAMP"`
	PayNotify string    `json:"pay_notify" xorm:"not null default '' comment('支付系统通知') VARCHAR(4096)"`
	PayStatus int       `json:"pay_status" xorm:"not null default 0 comment('支付状态') TINYINT(4)"`
	CreatedAt time.Time `json:"created_at" xorm:"not null default 'CURRENT_TIMESTAMP' created comment('创建时间') TIMESTAMP"`
	UpdatedAt time.Time `json:"updated_at" xorm:"not null default 'CURRENT_TIMESTAMP' updated comment('更新时间') TIMESTAMP"`
	DeletedAt time.Time `json:"deleted_at" xorm:"deleted comment('删除时间') TIMESTAMP"`
}

func InsertOrderInfo(m *OrderInfo) error {
	var err error
	if _, err = X.Insert(m); err != nil {
		log.WithFields(log.Fields{"err": err, "OrderInfo": m}).Error("models.InsertOrderInfo has an error")
		err = DBErr
	}
	return err
}

func GetOrderInfo(id int) (m *OrderInfo, err error) {
	var has bool
	m = &OrderInfo{}
	if has, err = X.ID(id).Get(m); err != nil {
		log.WithFields(log.Fields{"err": err, "id": id}).Error("models.GetOrderInfo has an error")
		err = DBErr
	}
	if !has {
		m = nil
	}
	return
}

func UpdateOrderInfo(m *OrderInfo) error {
	var err error
	if _, err = X.Id(m.Id).AllCols().Update(m); err != nil {
		log.WithFields(log.Fields{"err": err, "OrderInfo": m}).Error("models.UpdateOrderInfo has an error")
		err = DBErr
	}
	return err
}

func DeleteOrderInfo(id int) error {
	var err error
	if _, err := X.Id(id).Delete(new(OrderInfo)); err != nil {
		log.WithFields(log.Fields{"err": err, "id": id}).Error("models.DeleteOrderInfo has an error")
		err = DBErr
	}
	return err
}
