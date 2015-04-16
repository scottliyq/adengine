package main

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" // import your used driver
	"log"
)

type MdTargeting struct {
	Id         int
	CampaignId int
	TargetType string
	TargetCode string
}

type TargetGroup struct {
	CampaignId int
	TargetType string
	Number     int
}

type UniqueArray struct {
	List    []string
	ListSet map[string]bool
}

func init() {
	// set default database
	orm.RegisterDataBase("default", "mysql", "root:root@tcp(localhost:8889)/mahad?charset=utf8", 30)

}

func initDBData() {
	// open the item based backforward index
	//index := openIndex("adengine.item")
	//db, err := sql.Open("mysql", "root:root@tcp(localhost:8889)/mahad")
	//if err != nil {
	//	panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	//}
	//defer db.Close()

	//// Open doesn't open a connection. Validate DSN data:
	//err = db.Ping()
	//if err != nil {
	//	panic(err.Error()) // proper error handling instead of panic in your app
	//}

	o := orm.NewOrm()
	var targetingGroup []TargetGroup
	ids := []int{1}
	num, err := o.Raw("select campaignId, targetType, count(targetCode) from md_targeting where campaignId in (?) group by campaignId,targetType ", ids).QueryRows(&targetingGroup)

	if err == nil {
		log.Println("targeting nums: ", num)
	} else {
		panic(err.Error())
	}

}

func (this *UniqueArray) Add(item string) {
	if this.ListSet[item] != true {
		this.List = append(this.List, item)
		this.ListSet[item] = true
	}
}
