package admin

import (
	"APIGO/model"
	"fmt"

	_ "github.com/lib/pq"
	"xorm.io/xorm"
)

// SELECT semua data dari tb_admin
func GetAll(db *xorm.Engine) []model.Tb_admin {
	var tb_admin []model.Tb_admin
	err := db.Find(&tb_admin)

	if err != nil {
		fmt.Println(err)
	}

	return tb_admin
}

// INSERT data ke tb_admin
func Post(db *xorm.Engine, tb_admin model.Tb_admin) {
	_, err := db.Insert(
		model.Tb_admin{
			Id:       model.CreateID(),
			Nama:     tb_admin.Nama,
			Username: tb_admin.Username,
		},
	)
	// fmt.Println(tb_admin)

	if err != nil {
		fmt.Println(err)
		return
	}
}

// UPDATE data di tb_admin
func Put(db *xorm.Engine, tb_admin model.Tb_admin) {
	_, err := db.Where(fmt.Sprintf("id=%d", tb_admin.Id)).Update(
		model.Tb_admin{
			Nama:     tb_admin.Nama,
			Username: tb_admin.Username,
		},
	)
	// fmt.Println(tb_admin)

	if err != nil {
		fmt.Println(err)
		return
	}
}

// DELETE data di tb_admin
func Delete(db *xorm.Engine, id int64) {
	_, err := db.Delete(&model.Tb_admin{Id: id})

	if err != nil {
		fmt.Println(err)
		return
	}
}
