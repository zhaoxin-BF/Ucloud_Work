package db

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"new_mikasa_api/models/uaccount"
	"time"
)

// mongodb 't_cl_info 表结构'
type TLCInfo struct {
	Id              int    `id`
	ExternId        string `extern_id` //
	LcRandomId      string `lc_random_id`
	Owner           string `owner`
	Name            string `name`
	Size            int    `size`
	Oid             int    `oid`          //
	TopOid          int    `top_oid`      // 查新账单需要
	Status          int    `status`       //
	MountStatus     int    `mount_status` //
	MountVmId       string `mount_vm_id`  //
	CreateTime      int    `create_time`
	ThrowTime       int    `throw_time`
	DeleteTime      int    `delete_time`
	RecycleTime     int    `recycle_time`
	RecycledTime    int    `recycled_time`
	LastAttachTime  int    `last_attach_time`
	LastAetachTime  int    `last_detach_time`
	LastResizeTime  int    `last_resize_time`
	ImageUseTime    int    `image_use_time`
	DiskType        int    `disk_type` //
	UtmMode         int    `utm_mode`
	UtmStatus       int    `utm_status`
	UtmModifyTime   int    `utm_modify_time`
	GateIp          string `gate_ip`
	GatePort        int    `gate_port`
	MigrateTime     int    `migrate_time` //迁移时间
	LastActiveTime  int    `last_active_time`
	MountDeviceName string `mount_device_name` //挂载主机设备名
	SnapshotCount   int    `snapshot_count`    //快照个数
	SnapshotLimit   int    `snapshot_limit`    //快照个数限制
	IsInternal      int    `is_internal`
	CmkId           string `bson:"cmk_id"`
}

type SetTLCInfo struct {
	CompanyName       string
	ExternId          string `extern_id` //磁盘extern_id
	Name              string `name`      //磁盘名字
	Size              int    `size`
	OrganizationId    int    `oid`
	TopOrganizationId int    `top_oid`
	Status            int    `status`
	MountStatus       int    `mount_status` //挂载状态
	MountVmId         string `mount_vm_id`  //挂载的主机id
	CreateTime        int    `create_time`
	ThrowTime         int    `throw_time` //进回收站的时间
	DeleteTime        int    `delete_time`
	RecycleTime       int    `recycle_time`     //回收时间
	RecycledTime      int    `recycled_time`    //可回收时间
	LastAttachTime    int    `last_attach_time` //最近的一次挂载时间
	LastDetachTime    int    `last_detach_time` //最近的一次卸载时间
	LastResizeTime    int    `last_resize_time` //最近的一次的扩容时间
	DiskType          int    `disk_type`        //磁盘类型
	UtmMode           int    `utm_mode`         //是否开启方舟
	UtmStatus         int    `utm_status`       //方舟状态
	UtmModifyTime     int    `utm_modify_time`  //方舟修改状态的时间
	GateIp            string `gate_ip`          //宿主机ip
	GatePort          int    `gate_port`        //宿主机端口
	MigrateTime       int    `migrate_time`     //迁移时间
	LastActiveTime    int    `last_active_time`
	MountDeviceName   string `mount_device_name` //挂载主机设备名
	SnapshotCount     int    `snapshot_count`    //快照个数
	SnapshotLimit     int    `snapshot_limit`    //快照个数限制
	SnapCountLimit    string //快照个数（限制）
}

func Search(database *mgo.Database, col string, aggregateOptions []bson.M) []TLCInfo {
	if database == nil {
		return nil
	}
	var results []TLCInfo
	err := database.C(col).Pipe(aggregateOptions).All(&results)
	if err != nil {
		logs.Error("match err:", err)
	}
	return results
}

func GetTLCInfoCountGroupByTopOid(database *mgo.Database) (map[int][]int, error) {
	if database == nil {
		return nil, connectError
	}
	var tlcInfo []struct {
		TopOid int `bson:"top_oid"`
		Size   int `bson:"size"`
	}
	database.C("t_lc_info").Find(bson.M{"status": 0}).
		Select(bson.M{"top_oid": 1, "size": 1}).All(&tlcInfo)

	t := make(map[int][]int)

	for _, tlc := range tlcInfo {
		if t[tlc.TopOid] == nil {
			t[tlc.TopOid] = append(t[tlc.TopOid], []int{0, 0}...)
		}
		t[tlc.TopOid][0] += 1
		t[tlc.TopOid][1] += tlc.Size

	}

	return t, nil
}

// TODO 03:boreas.zhao 2.24编写,因代码重构待优化
func GetOneSetTlcInfo(database *mgo.Database) ([]SetTLCInfo, error) {
	if database == nil {
		return nil, connectError
	}
	var tlcInfo []TLCInfo
	nowtime := (int(time.Now().Unix()) - 7*86400)

	database.C("t_lc_info").Find(bson.M{"$or": []bson.M{{"status": bson.M{"$in": []int{0,3, 4}}}, {"$and": []bson.M{{"status":1}, {"delete_time":bson.M{"$gte": nowtime}}}}, {"$and": []bson.M{{"status":2}, {"recycled_time":bson.M{"$gte": nowtime}}}}}}).All(&tlcInfo)

	//查询所有正常、回收站、正在回收的资源
	//database.C("t_lc_info").Find(bson.M{"status": bson.M{"$in": []int{0,3, 4}}}).All(&tlcInfo)

	//7天内，标记删除
	//database.C("t_lc_info").Find(bson.M{"$and": []bson.M{bson.M{"status":1}, bson.M{"delete_time":bson.M{"$gte": nowtime}}}}).All(&tlcInfo)

	//7天内，确定回收
	//database.C("t_lc_info").Find(bson.M{"$and": []bson.M{bson.M{"status":2}, bson.M{"recycled_time":bson.M{"$gte": nowtime}}}}).All(&tlcInfo)

	//$and的使用
	//database.C("t_lc_info").Find(bson.M{"$and": []bson.M{bson.M{"status": bson.M{"$in": []int{1, 2}}}, bson.M{"delete_time":bson.M{"$gte": nowtime}}}}).All(&tlcInfo)

	//7天内，标记删除
	//database.C("t_lc_info").Find(bson.M{"status":1}).All(&tlcInfo)
	//database.C("t_lc_info").Find(bson.M{"delete_time":bson.M{"$gte": nowtime}}).All(&tlcInfo)

	//7天内，确定回收
	//database.C("t_lc_info").Find(bson.M{"status":2}).All(&tlcInfo)
	//database.C("t_lc_info").Find(bson.M{"recycled_time":bson.M{"$gte": nowtime}}).All(&tlcInfo)

	var CompanyName string
	var setTlcInfo []SetTLCInfo
	for _, lc := range tlcInfo {
		ICompanyInfo, err := uaccount.IGetCompanyInfo([]int{lc.TopOid})
		if err == nil {
			CompanyName = ICompanyInfo[0].CompanyName
		} else {
			CompanyName = "undefined"
		}

		var oneTlcInfo SetTLCInfo

		oneTlcInfo.CompanyName = CompanyName
		oneTlcInfo.ExternId = lc.ExternId
		oneTlcInfo.Name = lc.Name
		oneTlcInfo.Size = lc.Size
		oneTlcInfo.OrganizationId = lc.Oid
		oneTlcInfo.TopOrganizationId = lc.TopOid
		oneTlcInfo.Status = lc.Status
		oneTlcInfo.MountStatus = lc.MountStatus
		oneTlcInfo.MountVmId = lc.MountVmId
		oneTlcInfo.CreateTime = lc.CreateTime
		oneTlcInfo.ThrowTime = lc.ThrowTime
		oneTlcInfo.DeleteTime = lc.DeleteTime
		oneTlcInfo.RecycleTime = lc.RecycleTime
		oneTlcInfo.RecycledTime = lc.RecycledTime
		oneTlcInfo.LastAttachTime = lc.LastAttachTime
		oneTlcInfo.LastDetachTime = lc.LastAetachTime
		oneTlcInfo.LastResizeTime = lc.LastResizeTime
		oneTlcInfo.DiskType = lc.DiskType
		oneTlcInfo.UtmMode = lc.UtmMode
		oneTlcInfo.UtmStatus = lc.UtmStatus
		oneTlcInfo.UtmModifyTime = lc.UtmModifyTime
		oneTlcInfo.GateIp = lc.GateIp
		oneTlcInfo.GatePort = lc.GatePort
		oneTlcInfo.MigrateTime = lc.MigrateTime
		oneTlcInfo.LastActiveTime = lc.LastActiveTime
		oneTlcInfo.MountDeviceName = lc.MountDeviceName
		oneTlcInfo.SnapshotCount = lc.SnapshotCount
		oneTlcInfo.SnapshotLimit = lc.SnapshotLimit
		oneTlcInfo.SnapCountLimit = fmt.Sprint(lc.SnapshotCount) + "(" + fmt.Sprint(lc.SnapshotLimit) + ")"

		setTlcInfo = append(setTlcInfo, oneTlcInfo)
	}

	logs.Info("SetTlcInfo query success")

	return setTlcInfo, nil
}

func GetTLCInfoCountByTopOid(database *mgo.Database, topOid int) (map[int]int, error) {
	if database == nil {
		return nil, connectError
	}
	var tlcInfo []struct {
		DiskType int `bson:"disk_type"`
	}
	database.C("t_lc_info").Find(bson.M{"top_oid": topOid, "status": 0}).
		Select(bson.M{"disk_type": 1}).All(&tlcInfo)

	t := make(map[int]int)

	for _, tlc := range tlcInfo {
		t[tlc.DiskType] += 1
	}

	return t, nil
}

/*
udiskId 业务库资源状态恢复
status: 0 正常 1 删除
*/
func Update(database *mgo.Database, col, externId string) error {
	if database == nil {
		return connectError
	}
	//status:0 正常 1 删除
	err := database.C(col).Update(bson.M{"extern_id": externId, "status": 1}, bson.M{"$set": bson.M{"status": 0}})
	if err != nil {
		logs.Error("Update status failed! error: %v!  col: %s ExternId: %s", err, col, externId)
		return err
	}
	return nil
}

func GetMountStatus20And40(database *mgo.Database, col string) (info []*TLCInfo, err error) {
	if database == nil {
		return nil, connectError
	}
	err = database.C(col).Find(bson.M{"$or": []bson.M{{"mount_status": 20, "status": 0}, {"mount_status": 40, "status": 0}}}).All(&info)
	if err != nil {
		logs.Error("Get t_lc_info failed! error: %v! col: %s", err, col)
		return
	}
	return
}
