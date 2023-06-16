package scanpck

import "fmt"

// 定义三个构造器 general user admin boss
type generalUser struct {
	Name     string
	Age      int
	Password string
	Account  string
}

// 构造器 admin
type admin struct {
	Name     string
	Age      int
	Password string
	Account  string
	Level    int
}

// admin可以修改 普通用户的信息
func (a *admin) SetName(u *generalUser, name string) {
	u.Name = name
}

// 构造器 boss
type boss struct {
	Name     string
	Age      int
	Password string
	Account  string
	Level    int
}

// boss可以修改 admin的信息
func (b *boss) SetName(a *admin, name string) {
	a.Name = name
}

func Factory(t string, info map[string]interface{}) (ins interface{}) {
	switch t {
	case "general":
		ins := generalUser{}
		ins.Name = info["name"].(string)
		ins.Age = info["age"].(int)
		ins.Password = info["password"].(string)
		ins.Account = info["account"].(string)
		return ins
	case "admin":
		ins := admin{}
		ins.Name = info["name"].(string)
		ins.Age = info["age"].(int)
		ins.Password = info["password"].(string)
		ins.Account = info["account"].(string)
		ins.Level = info["level"].(int)
		return ins
	case "boss":
		ins := boss{}
		ins.Name = info["name"].(string)
		ins.Age = info["age"].(int)
		ins.Password = info["password"].(string)
		ins.Account = info["account"].(string)
		ins.Level = info["level"].(int)
		return ins
	default:
		fmt.Println("没有这个类型")
		return nil
	}

}
