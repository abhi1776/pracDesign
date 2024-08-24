package model

import "github.com/beego/beego/v2/client/orm"

type ParkingLotTable struct {
	Id     int    `orm:"column(id)" json:"id"`
	Number string `orm:"column(number)" json:"number"`
	Color  string `orm:"column(color)" json:"color"`
}

func (t *ParkingLotTable) TableName() string {
	return "parking_lot_table"
}

func init() {
	orm.RegisterModel(new(ParkingLotTable))
}

func (t *ParkingLotTable) AddParkingLotTable() (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(t)
	return
}

func (t *ParkingLotTable) GetParkingLotTableById(id int) (vd *ParkingLotTable, err error) {
	o := orm.NewOrm()
	t.Id = id
	if err = o.Read(t); err == nil {
		return t, nil
	}
	return nil, err
}
