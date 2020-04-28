package models

import (
	log "github.com/sirupsen/logrus"
	"time"
)

type Favorite struct {
	Id           int       `json:"id" xorm:"not null pk autoincr comment('id') INT(11)"`
	UserId       int       `json:"user_id" xorm:"not null default 0 comment('用户id') unique(uniq_user_favorite_item) INT(11)"`
	UserType     int       `json:"user_type" xorm:"not null default 0 comment('用户类型') unique(uniq_user_favorite_item) TINYINT(4)"`
	ItemId       int       `json:"item_id" xorm:"not null default 0 comment('收藏条目id') unique(uniq_user_favorite_item) INT(11)"`
	FavoriteType int       `json:"favorite_type" xorm:"not null default 0 comment('收藏类型') unique(uniq_user_favorite_item) TINYINT(4)"`
	CreatedAt    time.Time `json:"created_at" xorm:"not null default 'CURRENT_TIMESTAMP' created comment('创建时间') TIMESTAMP"`
	UpdatedAt    time.Time `json:"updated_at" xorm:"not null default 'CURRENT_TIMESTAMP' updated comment('更新时间') TIMESTAMP"`
	DeletedAt    time.Time `json:"deleted_at" xorm:"deleted comment('删除时间') unique(uniq_user_favorite_item) TIMESTAMP"`
}

func InsertFavorite(m *Favorite) error {
	var err error
	if _, err = X.Insert(m); err != nil {
		log.WithFields(log.Fields{"err": err, "Favorite": m}).Error("models.InsertFavorite has an error")
		err = DBErr
	}
	return err
}

func GetFavorite(id int) (m *Favorite, err error) {
	var has bool
	m = &Favorite{}
	if has, err = X.ID(id).Get(m); err != nil {
		log.WithFields(log.Fields{"err": err, "id": id}).Error("models.GetFavorite has an error")
		err = DBErr
	}
	if !has {
		m = nil
	}
	return
}

func UpdateFavorite(m *Favorite) error {
	var err error
	if _, err = X.Id(m.Id).AllCols().Update(m); err != nil {
		log.WithFields(log.Fields{"err": err, "Favorite": m}).Error("models.UpdateFavorite has an error")
		err = DBErr
	}
	return err
}

func DeleteFavorite(id int) error {
	var err error
	if _, err := X.Id(id).Delete(new(Favorite)); err != nil {
		log.WithFields(log.Fields{"err": err, "id": id}).Error("models.DeleteFavorite has an error")
		err = DBErr
	}
	return err
}
