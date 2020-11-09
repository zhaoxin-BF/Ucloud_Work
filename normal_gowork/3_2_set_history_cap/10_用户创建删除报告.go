package main

//package model
//
//import (
//"github.com/astaxie/beego/logs"
//
//"io/ioutil"
//)
//
//type AllCreateDeleteInfo struct {
//	TopID       int
//	Company     string
//	Level       string
//	InnerMark   int                    //是否内部用户
//	CreateCount int
//	CreateSize  int
//	DeleteCount int
//	DeleteSize  int
//	Region      string
//	GrowthCount int
//	GrowthSize  int
//}
//
//type CreDelInfo struct {
//	TopID       int                    //公司ID
//	Company     string                 //公司名称
//	Level       string                 //公司等级
//	CreateCount int                    //创建数量
//	CreateSize  int                    //创建size
//	DeleteCount int                    //删除数量
//	DeleteSize  int                    //删除size
//	Region      string                 //机房
//	SetName     int                    //集群名
//	//SetType     string                 //集群类型
//	InnerMark   int                    //是否内部用户
//	Usage       string                 //此条数据的所在集群的磁盘类型
//}
//
//type Create struct {
//	CreateSize  int `size`
//	CreateTime  int `create_time`
//	TopID       int `top_oid`
//	CreateCount int
//	Region      string
//	SetName     string
//}
//
//type Delete struct {
//	TopID       int `top_oid`
//	DeleteSize  int `size`
//	DeleteTime  int `delete_time`
//	DeleteCount int
//	Region      string
//	SetName     string
//}
//
//type CreDelBody struct {
//	Zone      string
//	BeginTime int
//	EndTime   int
//	InnerMark int
//}
//
//var Usage = []string{"sata_data", "sata_system", "ssd_data", "ssd_system", "rssd_data"} //0,1,2,3,4

////客户创建删除报告代码重构，author:boreas.zhao date:2020-3-12
//
////0、
////客户创建删除报告信息
//func CreDelDataInfo_bor(zone string, beginTime int, endTime int, innerMark int) interface{} {
//return GetAllSetCreateDelete_bor(beginTime, endTime, zone, innerMark)                     //获取所有机房集群删除创建信息
//}
//
////2、获取所有集群、创建删除数据
////获取所有机房的创建删除
//func GetAllSetCreateDelete_bor(beginTime int, endTime int, zone string, innerMark int) map[string][]AllCreateDeleteInfo {
//var UMgoGroup sync.WaitGroup
////var allCreateDeletes []CreDelInfo
//allSetMongoInfos := db.GetUDiskDatabase()                                                   //获得所要查询的机房集群数据库mgo信息
//ch := make(chan []CreDelInfo, 1000)
//for _, oneSetMongoInfo := range allSetMongoInfos {
//if zone == "all" || oneSetMongoInfo.ZoneShort == zone{
//UMgoGroup.Add(1)
//go func(oneSetMongoInfo *db.Database) {
//defer UMgoGroup.Done()
//ch <- GetOneSetCreateDeleteInfo_bor(oneSetMongoInfo, beginTime, endTime)
//}(oneSetMongoInfo)
//}
//}
//UMgoGroup.Wait()
//close(ch)
//var allZoneCompanyInfos []CreDelInfo
//for setinfo := range ch {
//allZoneCompanyInfos = append(allZoneCompanyInfos, setinfo...)
//}
//
//logs.Info("所有机房的创建删除信息：汇总完毕！")
//fiveTypeCreateDeleteInfo := make(map[string][]CreDelInfo)
////按磁盘类型初步分类共5类
//finalyTypeCreareDeleteInfo := make(map[string][]AllCreateDeleteInfo)                //封装好的数据
////根据磁盘类型分类
//for _, val := range allZoneCompanyInfos {
////fmt.Println(val)
//fiveTypeCreateDeleteInfo[val.Usage] = append(fiveTypeCreateDeleteInfo[val.Usage], val)
//}
//
////汇总整合个一个类型的创建删除数据
////companyCre := make(map[int][]CreDelInfo)
////for _, val := range fiveTypeCreateDeleteInfo["ssd_data"]{
////	fmt.Println(val)
////	companyCre[val.TopID] = append(companyCre[val.TopID], val)
////}
////  计算每个五个分类的公司数据
//for _, usage := range Usage {
//cdcompany := make(map[int][]CreDelInfo) //公司键值对切片
//var typeTotal AllCreateDeleteInfo
//
//for _, cd := range fiveTypeCreateDeleteInfo[usage] {
//cdcompany[cd.TopID] = append(cdcompany[cd.TopID], cd)
//}
//
//for key, val := range cdcompany { //	一个公司信息汇总
//var total AllCreateDeleteInfo
//for _, v := range val {
//total.DeleteCount += v.DeleteCount
//total.DeleteSize += v.DeleteSize
//total.CreateCount += v.CreateCount
//total.CreateSize += v.CreateSize
//if !strings.Contains(total.Region, v.Region) { //true //判断字符串是否在前一个串中
//total.Region += v.Region + ","
//}
//}
//total.TopID = key //公司ID
//total.GrowthCount = total.CreateCount - total.DeleteCount
//total.GrowthSize = total.CreateSize - total.DeleteSize
//
//companyInfo := udisk.IGetCompanyInfo(strconv.Itoa(key))
//if companyInfo == nil && len(companyInfo) == 0 {
//total.Company = " "
//} else {
//total.Company = companyInfo["CompanyName"].(string)
//total.InnerMark = int(companyInfo["InnerMark"].(float64))
//}
//
////companyInfos ,err := uaccount.IGetCompanyInfo([]int{key})   //获得公司名称、是否内部用户标签
////if err == nil {
////	companyInfo := companyInfos[0]
////	if companyInfo == nil {
////		total.Company = " "
////	} else {
////		total.Company = companyInfo.CompanyName
////		total.InnerMark = companyInfo.InnerMark
////	}
////}
//
////companyFuzzy, err:= uaccount.GetCompanysFuzzy(key)     //获得公司的等级
////if err == nil{
////	total.Level = companyFuzzy.Level
////}
//companyFuzzy := udisk.GetCompanysFuzzy(strconv.Itoa(key))
//total.Level = companyFuzzy.Level
//if total.InnerMark == innerMark { //去除内部客户 InnerMark
//typeTotal.CreateCount += total.CreateCount
//typeTotal.CreateSize += total.CreateSize
//
//typeTotal.DeleteCount += total.DeleteCount
//typeTotal.DeleteSize += total.DeleteSize
//
//typeTotal.GrowthCount += total.GrowthCount
//typeTotal.GrowthSize += total.GrowthSize
//
//finalyTypeCreareDeleteInfo[usage] = append(finalyTypeCreareDeleteInfo[usage], total)
////fmt.Println("del Data ", delData)
//}
//}
//
//typeTotal.Company = usage                         //磁盘汇总的类型
//typeTotal.TopID = len(finalyTypeCreareDeleteInfo) //公司数量
//
//finalyTypeCreareDeleteInfo[usage] = append(finalyTypeCreareDeleteInfo[usage], typeTotal)
////for _, oneCompany := range finalyTypeCreareDeleteInfo[usage]{
////	fmt.Println(
////		"topID: ", oneCompany.TopID,
////		"compantName: ", oneCompany.Company,
////		"lever: ", oneCompany.Level,
////		"creCount:", oneCompany.CreateCount,
////		"creSize: ", oneCompany.CreateSize,
////		"delCount: ", oneCompany.DeleteCount,
////		"delSize: ", oneCompany.DeleteSize,
////		"GrowthCount: ", oneCompany.GrowthCount,
////		"GrowthSize: ", oneCompany.GrowthSize,
////		"region: ", oneCompany.Region)
////}
//}
//logs.Info("创建删除set分类汇总:完毕！")
//
//return finalyTypeCreareDeleteInfo
//}
////3、获取单个集群、创建删除数据
//// 操作mongo数据库获取一个set创建删除报告
//func GetOneSetCreateDeleteInfo_bor(setdatabase *db.Database, beginTime int, endTime int) []CreDelInfo {
//
//db := setdatabase.Database              //数据库查询链接
//set := setdatabase.Set                  //集群名字
//regionID := setdatabase.ZoneShort       //机房名字
//
//usage := Usage[set/1000 - 1]            //集群类型
//
////fmt.Println("set:", set)               /////////////////////////////////////////////////////////////////////////测试打印数据
////fmt.Println("regionID:", regionID)
////fmt.Println("usage:", usage)
//
//companyCre := make(map[int][]Create)       //一个公司的     多条数据
//companyCreTotals := make(map[int]CreDelInfo)   //一个集群的     多个公司的创建汇总数据
//companyDel := make(map[int][]Delete)
//companyDelTotals := make(map[int]CreDelInfo)
//
//
//
//var cre []Create
//var del []Delete
//
//var MgoGroup sync.WaitGroup
//MgoGroup.Add(2)
//
//go func() {
//defer MgoGroup.Done()
//db.C("t_lc_info").Find(bson.M{"create_time": bson.M{"$gt": beginTime, "$lt": endTime}}).All(&cre)
//}()
//
//go func() {
//defer MgoGroup.Done()
//db.C("t_lc_info").Find(bson.M{"delete_time": bson.M{"$gt": beginTime, "$lt": endTime}}).All(&del)
//}()
//
//MgoGroup.Wait()
//logs.Info("mongdb query over")
////fmt.Println("111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111")
//for _, c := range cre {
////fmt.Println(c)//整合单个公司的创建数据
//companyCre[c.TopID] = append(companyCre[c.TopID], c)
//}
//
//for key, val := range companyCre {                              //计算单个公司的创建总量数据
//var creTotal CreDelInfo
//for _, v := range val {
//creTotal.CreateSize += v.CreateSize
//creTotal.CreateCount += 1
//}
//creTotal.Region     = regionID
//creTotal.SetName    = set
//creTotal.Usage      = usage
//creTotal.TopID      = key
//companyCreTotals[key] = creTotal
//}
//
//for _, d := range del {                                       //整合单个公司的删除数据
//companyDel[d.TopID] = append(companyDel[d.TopID], d)
//}
//for key, val := range companyDel {                            //计算单个公司的创建总量数据
//var delTotal CreDelInfo
//for _, v := range val {
//delTotal.DeleteSize += v.DeleteSize
//delTotal.DeleteCount += 1
//}
//delTotal.Region     = regionID
//delTotal.SetName    = set
//delTotal.Usage      = usage
//delTotal.TopID      = key
//companyDelTotals[key] = delTotal
//}
//
////整合一个公司创建、删除数据
//for key, companyinfo := range companyDelTotals{
//if _, ok := companyCreTotals[key]; ok {
//var info CreDelInfo
//info.TopID       = companyinfo.TopID
//info.DeleteSize  = companyinfo.DeleteSize
//info.DeleteCount = companyinfo.DeleteCount
//info.CreateCount = companyCreTotals[key].CreateCount
//info.CreateSize  = companyCreTotals[key].CreateSize
//info.Region      = companyinfo.Region
//info.SetName     = companyinfo.SetName
//info.Usage       = companyinfo.Usage
//companyCreTotals[key] = info
//}else {
//companyCreTotals[key] = companyinfo
//}
//}
//var oneSetCompanyInfos []CreDelInfo
////fmt.Println("2222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222")
//for _, companyinfo := range companyCreTotals{
//oneSetCompanyInfos = append(oneSetCompanyInfos, companyinfo)
//}
//logs.Info("oneSetCreDelTotal，整合完毕！")
//
////for _, oneCompany := range oneSetCompanyInfos{
////	fmt.Println(
////		"topID: ", oneCompany.TopID,
////		"creCount:", oneCompany.CreateCount,
////		"creSize: ", oneCompany.CreateSize,
////		"delCount: ", oneCompany.DeleteCount,
////		"delSize: ", oneCompany.DeleteSize,
////		"region: ", oneCompany.Region,
////		"SetName: ", oneCompany.SetName,
////		"Usage: ", oneCompany.Usage)
////}
//return oneSetCompanyInfos
//}
