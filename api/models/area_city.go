package models

type AreaCity struct {
	Id           int    `json:"id" xorm:"not null pk comment('id') INT(11)"`
	Pid          int    `json:"pid" xorm:"not null comment('pid') INT(11)"`
	Deep         int    `json:"deep" xorm:"not null comment('deep') INT(11)"`
	Name         string `json:"name" xorm:"not null comment('name') VARCHAR(50)"`
	PinyinPrefix string `json:"pinyin_prefix" xorm:"not null comment('pingyin prefix') VARCHAR(5)"`
	Pinyin       string `json:"pinyin" xorm:"not null comment('pinyin') VARCHAR(50)"`
	ExtId        string `json:"ext_id" xorm:"not null comment('ext_id') VARCHAR(20)"`
	ExtName      string `json:"ext_name" xorm:"not null comment('ext_name') VARCHAR(50)"`
}

func InsertAreaCity(m *AreaCity) error {
	var err error
	if _, err = X.Insert(m); err != nil {
		log.WithFields(log.Fields{"err": err, "AreaCity": m}).Error("models.InsertAreaCity has an error")
		err = DBErr
	}
	return err
}

func GetAreaCity(id int) (m *AreaCity, err error) {
	var has bool
	m = &AreaCity{}
	if has, err = X.ID(id).Get(m); err != nil {
		log.WithFields(log.Fields{"err": err, "id": id}).Error("models.GetAreaCity has an error")
		err = DBErr
	}
	if !has {
		m = nil
	}
	return
}

func UpdateAreaCity(m *AreaCity) error {
	var err error
	if _, err = X.Id(m.Id).AllCols().Update(m); err != nil {
		log.WithFields(log.Fields{"err": err, "AreaCity": m}).Error("models.UpdateAreaCity has an error")
		err = DBErr
	}
	return err
}

func DeleteAreaCity(id int) error {
	var err error
	if _, err := X.Id(id).Delete(new(AreaCity)); err != nil {
		log.WithFields(log.Fields{"err": err, "id": id}).Error("models.DeleteAreaCity has an error")
		err = DBErr
	}
	return err
}
