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

var zones =[]string{"hn02","hn03","hn04"}
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

	var setinfos []SetCapHistroyInfo
	err = c.Find(bson.M{"setusage": "ssd_data"}).All(&setinfos)
	fmt.Println(len(setinfos), "条数据查找完毕！")

	regionmaps := make(map[string][]SetCapHistroyInfo)
	for _, setinfo := range setinfos {
		for _, zone := range zones {
			if zone == setinfo.Region {
				regionmaps[zone] = append(regionmaps[zone], setinfo)
				break
			}
		}
	}

	fmt.Println(regionmaps,"dsfsf")

	type mmp map[string][]SetCapHistroyInfo
	regionsetmap := make(map[string]mmp)
	//for key, region := range regionmaps{
	//
	//	setfos := make(mmp)
	//	for _, setfo := range region{
	//		if len(setfos) == 0{
	//			setfos[setfo.Set] = append(setfos[setfo.Set], setfo)
	//		}else{
	//			for key, _ := range setfos{
	//				if key == setfo.Set{
	//					setfos[setfo.Set] = append(setfos[setfo.Set],setfo)
	//					break
	//				}
	//			}
	//		}
	//	}
	//	regionsetmap[key] = setfos
	//}

	for key, region := range regionmaps{

		setfos := make(mmp)
		for _, setfo := range region{
			setfos[setfo.Set] = append(setfos[setfo.Set], setfo)
		}
		regionsetmap[key] = setfos
	}
	for _, regionsetma := range regionsetmap{
		fmt.Println(regionsetma)
	}

	fmt.Println(regionsetmap["hn02"])
}