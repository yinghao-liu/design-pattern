package singleton

// 可以被包外的使用者通过new创建
type Singleton2 struct {
	instance int
}

// 包外的使用者不能直接使用
type singleton struct {
	instance int
}

// 单例-懒汉式
var single *singleton

// 单例-饿汉式
//var single *singleton = new(singleton)

func (s singleton) ShowInstance() int {
	return s.instance
}

// 未加锁  多线程不安全的
func GetInstance() *singleton {
	if nil == single {
		single = new(singleton)
	}
	return single
}
