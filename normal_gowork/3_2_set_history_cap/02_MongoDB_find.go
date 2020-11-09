package main

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type SetCapHistroyInfo struct {
	Region string
	Set  string
	SetUsage string
	Total int
}

//一、插入数据库
func main11() {
	fmt.Print("开始插入")
	session, err := mgo.Dial("117.50.104.98:27017")
	if err != nil {
		panic(err)
	}
	defer session.Clone()

	session.SetMode(mgo.Monotonic, true)
	c := session.DB("test").C("t_set_cap_history")
	for i := 0; i < 10; i++ {
		err = c.Insert(&SetCapHistroyInfo{"hn02", "1111","ssd_data", 1024})
		err = c.Insert(&SetCapHistroyInfo{"hn02", "1111","rssd_data", 1024})
		err = c.Insert(&SetCapHistroyInfo{"hn02", "1111","ssd_system", 1024})
		err = c.Insert(&SetCapHistroyInfo{"hn02", "1111","sate_system", 1024})
		err = c.Insert(&SetCapHistroyInfo{"hn02", "1111","sate_data", 1024})
		err = c.Insert(&SetCapHistroyInfo{"hn03", "1111","ssd_data", 1024})
		err = c.Insert(&SetCapHistroyInfo{"hn03", "1111","rssd_data", 1024})
		err = c.Insert(&SetCapHistroyInfo{"hn03", "1111","ssd_system", 1024})
		err = c.Insert(&SetCapHistroyInfo{"hn03", "1111","sate_system", 1024})
		err = c.Insert(&SetCapHistroyInfo{"hn03", "1111","sate_data", 1024})
		err = c.Insert(&SetCapHistroyInfo{"hn04", "1111","ssd_data", 1024})
		err = c.Insert(&SetCapHistroyInfo{"hn04", "1111","rssd_data", 1024})
		err = c.Insert(&SetCapHistroyInfo{"hn04", "1111","ssd_system", 1024})
		err = c.Insert(&SetCapHistroyInfo{"hn04", "1111","sate_system", 1024})
		err = c.Insert(&SetCapHistroyInfo{"hn04", "1111","sate_data", 1024})
		//err = c.Remove(&SetCapHistroyInfo{"hn02", "1000", 1024})
		if err != nil {
			panic(err)
		}
	}

	fmt.Println("插入完毕！")
}

//二、查找数据库
//func main() {
//	str := "172.18.183.132:27020"
//	session, err := mgo.Dial(str)
//	if err != nil {
//		panic(err)
//	}
//	defer session.Clone()
//
//	session.SetMode(mgo.Monotonic, true)
//	c := session.DB("test").C("student")
//
//	//stus := make([]Person, 200)
//	var stus []Person
//	err = c.Find(nil).All(&stus)
//
//	fmt.Println(stus[1].Name + ":" + stus[1].Age)
//
//	fmt.Println(len(stus), "条数据查找完毕！")
//}

//三、bson查找数据库
func main() {
	str := "117.50.104.98:27017"
	session, err := mgo.Dial(str)
	if err != nil {
		panic(err)
	}
	defer session.Clone()

	session.SetMode(mgo.Monotonic, true)
	c := session.DB("test").C("t_set_cap_history")

	//所有符合类型条件的集群信息
	var setinfos []SetCapHistroyInfo
	err = c.Find(bson.M{"setusage": "ssd_data"}).All(&setinfos)
	fmt.Println(len(setinfos), "条数据查找完毕！")

	//已机房为key,整合独立机房数据
	regionsetinfos := make(map[string][]SetCapHistroyInfo)
	for _, setinfo := range setinfos {
		regionsetinfos[setinfo.Region] = append(regionsetinfos[setinfo.Region], setinfo)
	}

	//已集群为key,整合独立集群数据
	type mmp map[string][]SetCapHistroyInfo
	allregionsetinfo := make(map[string]mmp)

	for key, region := range regionsetinfos{
		setfos := make(mmp)
		for _, setfo := range region{
			setfos[setfo.Set] = append(setfos[setfo.Set], setfo)
		}
		allregionsetinfo[key] = setfos
	}
	for _, oneregionset := range allregionsetinfo{
		fmt.Println(oneregionset)
	}
}