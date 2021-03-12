package models

import (
	"fmt"
	"strconv"
	"sync"

	"class_notice/helper"

	"github.com/beego/beego/v2/client/orm"
	_ "github.com/go-sql-driver/mysql"
)

var singleOrmInstance orm.Ormer
var once sync.Once

func GetSharedOrmer() orm.Ormer {
	once.Do(func() {
		singleOrmInstance = orm.NewOrm()
	})
	return singleOrmInstance
}

func addUser(name string, password string) {
	id, err := AddUsers(&Users{Name: name, Password: helper.StringToMd5(password)})
	if err == nil {
		fmt.Println("用户创建成功: " + strconv.FormatInt(id, 10))
	}
}

func init() {
	tag := "【Model.go】"

	driver, _ := helper.Env("dbdriver")
	username, _ := helper.Env("dbusername")
	password, _ := helper.Env("dbpassword")
	database, _ := helper.Env("dbdatabase")
	host, _ := helper.Env("dbhost")

	orm.RegisterDriver("mysql", orm.DRMySQL)
	connection_url := helper.GetConnectionURL(username, password, host, database)

	fmt.Println(tag + "连接URL是: " + connection_url)

	orm.RegisterDataBase("default", driver, connection_url)

	fmt.Println(tag + "注册数据模型")
	orm.RegisterModel(
		new(Users),    // 用户
		new(Configs),  // 配置
		new(Accounts), // 账号
		new(Times),    // 时间
		new(Courses),  // 课表
	)

	// 第二个参数为 true 则强制重新建表
	orm.RunSyncdb("default", false, true)

	// addUser("bin", "123456") // 添加默认用户

}