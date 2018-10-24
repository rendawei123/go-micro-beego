package models

import (
	"github.com/Unknwon/com"
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
	"os"
	"path"
	"strconv"
)

const (
	_DB_NAME        = "data/beeblog.db"
	_SQLITE3_DRIVER = "sqlite3"
)

type Category struct {
	Id         int64
	Title      string
	TopicCount int64 // 统计这个分类有多少篇文章
}

type Topic struct {
	Id              int64
	Title           string
	Content         string `orm:"size(5000)"` // 文章内容
	Views           int64  `orm:"index"`      // 浏览数量
	Author          string
	ReplyCount      int64
	ReplyLastUserId int64
}

func RegisterDB() {
	if !com.IsExist(_DB_NAME) { // 判断此数据库路径是否存在
		os.MkdirAll(path.Dir(_DB_NAME), os.ModePerm) // MkdirAll一次性创建多级目录，path.Dir将名称变成路径
		os.Create(_DB_NAME)
	}
	orm.RegisterModel(new(Category), new(Topic))                   // 注册模型
	orm.RegisterDriver(_SQLITE3_DRIVER, orm.DRSqlite)              // 注册驱动
	orm.RegisterDataBase("default", _SQLITE3_DRIVER, _DB_NAME, 10) // 创建数据库
}

func AddCategory(name string) error {

	// 获取orm对象
	o := orm.NewOrm()

	// 创建对象
	cate := &Category{Title: name}

	// 查询判断name是否已用
	qs := o.QueryTable("category")
	err := qs.Filter("title", name).One(cate)
	if err == nil {
		return err
	}

	_, err1 := o.Insert(cate)

	if err1 != nil {
		return err
	}
	return nil
}

func GetAllCategories() ([]*Category, error) {
	o := orm.NewOrm()
	cates := make([]*Category, 0)
	qs := o.QueryTable("Category")
	_, err := qs.All(&cates)

	return cates, err
}

func DeleteCategory(id string) error {

	// 将string类型的id转化为int类型
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}

	o := orm.NewOrm()
	cate := &Category{Id: cid}
	_, err1 := o.Delete(cate)
	return err1
}
