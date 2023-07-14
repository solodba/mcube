package apps

import (
	"fmt"
	"strings"

	"github.com/emicklei/go-restful/v3"
	"github.com/gin-gonic/gin"
	"github.com/infraboard/mcube/http/restful/accessor/form"
	"github.com/infraboard/mcube/http/restful/accessor/yaml"
	"github.com/infraboard/mcube/http/restful/accessor/yamlk8s"
	"github.com/solodba/mcube/common/logger"
	"google.golang.org/grpc"
)

// 初始化内部实例配置
func InitInternalApps() error {
	for _, v := range internalAppStore {
		if err := v.Conf(); err != nil {
			return err
		}
		logger.L().Info().Msgf("%s object config initial success", v.Name())
	}
	return nil
}

// 初始化GRPC实例配置
func InitGrpcApps(s *grpc.Server) error {
	for _, v := range grpcAppStore {
		if err := v.Conf(); err != nil {
			return err
		}
		v.RegistryHandler(s)
		logger.L().Info().Msgf("grpc %s object config initial success", v.Name())
	}
	return nil
}

// 初始化restful实例配置
func InitRestfulApps(pathPrefix string, root *restful.Container) error {
	for _, v := range restfulAppStore {
		if err := v.Conf(); err != nil {
			return err
		}
		pathPrefix = strings.TrimSuffix(pathPrefix, "/")
		ws := new(restful.WebService)
		ws.Path(fmt.Sprintf("%s/%s/%s", pathPrefix, v.Version(), v.Name())).
			Consumes(restful.MIME_JSON, restful.MIME_XML, form.MIME_POST_FORM, form.MIME_MULTIPART_FORM, yaml.MIME_YAML, yamlk8s.MIME_YAML).
			Produces(restful.MIME_JSON, restful.MIME_XML, form.MIME_POST_FORM, form.MIME_MULTIPART_FORM, yaml.MIME_YAML, yamlk8s.MIME_YAML)
		v.RegistryHandler(ws)
		root.Add(ws)
		logger.L().Info().Msgf("restful %s object config initial success", v.Name())
	}
	return nil
}

// 初始化gin实例配置
func InitHttpApps(r gin.IRouter) error {
	for _, v := range httpAppStore {
		if err := v.Conf(); err != nil {
			return err
		}
		v.RegistryHandler(r.Group(v.Name()))
		logger.L().Info().Msgf("http %s object config initial success", v.Name())
	}
	return nil
}
