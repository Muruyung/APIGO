package model

type Tb_user struct {
	Id       int64  `xorm:"not null 'id'" json:"id"`
	Nama     string `xorm:"not null 'nama'" json:"nama"`
	Username string `xorm:"not null 'username'" json:"username"`
}

type Tb_admin struct {
	Id       int64  `xorm:"not null 'id'" json:"id"`
	Nama     string `xorm:"not null 'nama'" json:"nama"`
	Username string `xorm:"not null 'username'" json:"username"`
}
