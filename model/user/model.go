package user

import (
	"APIGO/model"
	"fmt"

	_ "github.com/lib/pq"
	"xorm.io/xorm"
)

// SELECT semua data dari tb_user
func GetAll(db *xorm.Engine) []model.Tb_user {
	var tb_user []model.Tb_user
	err := db.Find(&tb_user)

	if err != nil {
		fmt.Println(err)
	}

	return tb_user
}

// INSERT data ke tb_user
func Post(db *xorm.Engine, tb_user model.Tb_user) {
	cek, err := db.Insert(
		model.Tb_user{
			Id:       model.CreateID(),
			Nama:     tb_user.Nama,
			Username: tb_user.Username,
		},
	)
	fmt.Println(cek)

	if err != nil {
		fmt.Println(err)
		return
	}
}

// UPDATE data di tb_user
func Put(db *xorm.Engine, tb_user model.Tb_user) {
	_, err := db.Where(fmt.Sprintf("id=%d", tb_user.Id)).Update(
		model.Tb_user{
			Nama:     tb_user.Nama,
			Username: tb_user.Username,
		},
	)
	// fmt.Println(tb_user)

	if err != nil {
		fmt.Println(err)
		return
	}
}

// DELETE data di tb_user
func Delete(db *xorm.Engine, id int64) {
	_, err := db.Delete(&model.Tb_user{Id: id})

	if err != nil {
		fmt.Println(err)
		return
	}
}
