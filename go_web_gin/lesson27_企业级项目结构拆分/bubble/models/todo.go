package models

import (
	"bubble/dao"
	"bubble/logs"
)

// Todo Model
type Todo struct {
	ID       int         `json:"id"`
	Title    string      `json:"title"`
	Status   bool        `json:"status"`
}

/*
	Todo这个Model的增删改查操作都放在这里
 */
// CreateATodo 创建todo
func CreateATodo(todo *Todo) (err error){
	err = dao.DB.Create(&todo).Error
	logs.Info("创建一条信息！")
	return
}

func GetAllTodo() (todoList []*Todo, err error){
	if err = dao.DB.Find(&todoList).Error; err != nil{
		return nil, err
	}
	logs.Info("获得所有信息！")
	return
}

func GetATodo(id string)(todo *Todo, err error){
	todo = new(Todo)
	if err = dao.DB.Debug().Where("id=?", id).First(todo).Error; err!=nil{
		return nil, err
	}
	logs.Info("获取一条信息！")
	return
}

func UpdateATodo(todo *Todo)(err error){
	err = dao.DB.Save(todo).Error
	logs.Info("更新一条信息！")
	return
}

func DeleteATodo(id string)(err error){
	err = dao.DB.Where("id=?", id).Delete(&Todo{}).Error
	logs.Info("删除一条信息！")
	return
}

