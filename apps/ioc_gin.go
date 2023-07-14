package apps

import (
	"github.com/gin-gonic/gin"
	"github.com/solodba/mcube/logger"
)

// 内部实例映射
var (
	httpAppStore = map[string]HttpIocObject{}
)

// 内部实例Ioc接口
type HttpIocObject interface {
	IocObject
	RegistryHandler(gin.IRouter)
}

// Ioc内部实例注册函数
func RegistryHttpApp(obj HttpIocObject) {
	if _, ok := httpAppStore[obj.Name()]; ok {
		logger.L().Panic().Msgf("http %s has registried ioc center, please don't registry again", obj.Name())
	}
	httpAppStore[obj.Name()] = obj
	logger.L().Info().Msgf("http %s registry ioc center success", obj.Name())
}

// 从Ioc获取内部实例
func GetHttpApp(objName string) any {
	if v, ok := httpAppStore[objName]; ok {
		return v
	}
	logger.L().Panic().Msgf("http %s is not registry ioc center, please registry ioc center first", objName)
	return nil
}
