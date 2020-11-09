package main

import (
	"fmt"
	"gopkg.in/mgo.v2"
)

type SetCapHistroyInfo struct {
	Region              string     `bson:"region"`              //机房简称
	RegionName          string     `bson:"region_name"`         //机房名称
	Set                 int        `bson:"set"`                 //集群号
	Name                string     `bson:"name"`                //集群名
	SetUsage            string     `bson:"set_usage"`           //集群类型
	RecordTime          int        `bson:"record_time"`         //记录时间/年/月/日
	ZoneRegion          string     `bson:"zone_region"`         //机房地域/国内/浪潮/金融
	ZoneRegionType      string     `bson:"zone_region_type"`    //机房地域类型
	Total               int        `bson:"total"`               //总量
	Inuse               int        `bson:"inuse"`               //已使用
	HeartCapacity       int        `bson:"heart_capacity"`      //物理剩余（实际剩余）
	Available           int        `bson:"available"`           //剩余
	AvailablePercentage float64    `bson:"available_percentage"`//剩余百分比
	Recyclebin          int        `bson:"recyclebin"`          //回收站容量
	Recyclable          int        `bson:"recyclable"`          //可回收容量
}
var setcaphistroyinfos []SetCapHistroyInfo

//一、插入数据库
func main() {
	fmt.Println("处理数据")
	for i:=0; i<10; i++ {
		setcaphistoryinfo := SetCapHistroyInfo{"hn03", "广州二B", 1102, "hn02_1102", "ssd_data", 1583219330, "国内", "国内ssd_data", 1024, 512, 800, 512, 20.24, 120, 220}
		setcaphistroyinfos = append(setcaphistroyinfos, setcaphistoryinfo)
	}
	var docs []interface{}
	for _,v := range setcaphistroyinfos{
		docs = append(docs,v)
	}


	fmt.Print("开始插入")
	session, err := mgo.Dial("117.50.104.98:27017")
	if err != nil {
		panic(err)
	}
	defer session.Clone()
	session.SetMode(mgo.Eventual, true)
	c := session.DB("test").C("set_cap_history")
	err = c.Insert(docs...)
	//err = c.Insert(&SetCapHistroyInfo{"hn02", "广州二B", 1101, "hn02_1101", "ssd_data", 1583219330, "国内", "国内ssd_data", 1024, 512, 800, 512, 20.24, 120, 220})
	if err != nil {
		panic(err)
	}

	fmt.Println("插入完毕！")
}


//二、bson查找数据库
//func main() {
//	str := "117.50.104.98:27017"
//	session, err := mgo.Dial(str)
//	if err != nil {
//		panic(err)
//	}
//	defer session.Clone()
//
//	session.SetMode(mgo.Eventual, true)
//	c := session.DB("test").C("set_cap_history")
//
//	//所有符合类型条件的集群信息
//	var setinfos []SetCapHistroyInfo
//	err = c.Find(bson.M{"setusage": "ssd_data"}).All(&setinfos)
//	fmt.Println(len(setinfos), "条数据查找完毕！")
//
//	//已机房为key,整合独立机房数据
//	regionsetinfos := make(map[string][]SetCapHistroyInfo)
//	for _, setinfo := range setinfos {
//		regionsetinfos[setinfo.Region] = append(regionsetinfos[setinfo.Region], setinfo)
//	}
//
//	//已集群为key,整合独立集群数据
//	type mmp map[int][]SetCapHistroyInfo
//	allregionsetinfo := make(map[string]mmp)
//
//	for key, region := range regionsetinfos{
//		setfos := make(mmp)
//		for _, setfo := range region{
//			setfos[setfo.Set] = append(setfos[setfo.Set], setfo)
//		}
//		allregionsetinfo[key] = setfos
//	}
//	for _, oneregionset := range allregionsetinfo{
//		fmt.Println(oneregionset)
//	}
//}