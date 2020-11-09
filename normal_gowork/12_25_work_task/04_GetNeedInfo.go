//根据MDB的地址，查询符合要求的udisk信息， 返回的是一个切片
func GetNeedInfo(mDbInfo limax.MDbInfo, aggregateOption bson.M) (needInfo limax.NeedInfo, ok bool) {
	tLCInfo := db.Search(mDbInfo.Url, mDbInfo.DbInfo.Dbname, "t_lc_info", aggregateOption)

	ok := true
	needInfo := new(limax.NeedInfo)

	if len(tLCInfo) != 1 { //？？？？？？？？？是否能用len(tLCInfo)来判断查询到了数据？？？？？
		ok = false
	} else {
		needInfo.InnerId = fmt.Sprint(tLCInfo.Id)
		needInfo.UDiskId = tLCInfo.ExternId
		needInfo.Name = tLCInfo.Name
		needInfo.CompanyId = tLCInfo.TopOid
		needInfo.OrganizationId = tLCInfo.Oid
		needInfo.Size = tLCInfo.Size
		needInfo.RecycleTime = tLCInfo.RecycleTime
		needInfo.RecycledTime = tLCInfo.RecycledTime
		needInfo.CreateTime = tLCInfo.CreateTime
		needInfo.DeleteTime = tLCInfo.DeleteTime
		needInfo.DiskType = tLCInfo.DiskType
		needInfo.Status = tLCInfo.Status
		needInfo.MountStatus = tLCInfo.MountStatus
		needInfo.UHostId = tLCInfo.MountVmId
		needInfo.DeviceName = tLCInfo.MountDeviceName
		needInfo.UtmMode = tLCInfo.UtmMode
		needInfo.GateIp = tLCInfo.GateIp
		needInfo.Archiecture = "new"
		needInfo.Region = limax.GetRegion(mDbInfo.ZoneId)
		needInfo.ZoneShort = mDbInfo.ZoneShort
		needInfo.CmkId = tLCInfo.CmkId
		needInfo.Set = mDbInfo.Set
		needInfo.MDbInfo = &mDbInfo

		logs.Info("dbInfo: ", fmt.Sprintf("%+v", mDbInfo), "GetNeedInfo needInfo:", needInfo)
	}

	return needInfo, ok

}