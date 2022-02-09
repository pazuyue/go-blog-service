package model

import (
	"blog-service/global"
	"errors"
	"github.com/casbin/casbin"
	"github.com/casbin/casbin/util"
	gormadapter "github.com/casbin/gorm-adapter"
	"github.com/jinzhu/gorm"
	"strings"
)

type CasbinModel struct {
	PType  string `json:"p_type" gorm:"column:p_type" description:"策略类型"`
	RoleId string `json:"role_id" gorm:"column:v0" description:"角色ID"`
	Path   string `json:"path" gorm:"column:v1" description:"api路径"`
	Method string `json:"method" gorm:"column:v2" description:"访问方法"`
}

func (c *CasbinModel) TableName() string {
	return "casbin_rule"
}

func (c *CasbinModel) Create(db *gorm.DB) error {
	e := Casbin()
	if success := e.AddPolicy(c.RoleId, c.Path, c.Method); success == false {
		return errors.New("存在相同的API，添加失败")
	}
	return nil
}

func (c *CasbinModel) Update(db *gorm.DB, values interface{}) error {
	if err := db.Model(c).Where("v1 = ? AND v2 = ?", c.Path, c.Method).Update(values).Error; err != nil {
		return err
	}
	return nil
}

func (c *CasbinModel) List(db *gorm.DB) [][]string {
	e := Casbin()
	policy := e.GetFilteredPolicy(0, c.RoleId)
	return policy
}

//@function: Casbin
//@description: 持久化到数据库  引入自定义规则
//@return: *casbin.Enforcer
func Casbin() *casbin.Enforcer {
	adapter := gormadapter.NewAdapterByDB(global.DBEngine)
	enforcer := casbin.NewEnforcer(global.CasbinSetting.ModelPath, adapter)
	enforcer.AddFunction("ParamsMatch", ParamsMatchFunc)
	_ = enforcer.LoadPolicy()
	return enforcer
}

//@function: ParamsMatch
//@description: 自定义规则函数
//@param: fullNameKey1 string, key2 string
//@return: bool
func ParamsMatch(fullNameKey1 string, key2 string) bool {
	key1 := strings.Split(fullNameKey1, "?")[0]
	// 剥离路径后再使用casbin的keyMatch2
	return util.KeyMatch2(key1, key2)
}

//@function: ParamsMatchFunc
//@description: 自定义规则函数
//@param: args ...interface{}
//@return: interface{}, error
func ParamsMatchFunc(args ...interface{}) (interface{}, error) {
	name1 := args[0].(string)
	name2 := args[1].(string)

	return ParamsMatch(name1, name2), nil
}
