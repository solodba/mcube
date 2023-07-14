package format

import (
	"encoding/json"

	"github.com/solodba/mcube/logger"
)

// json格式化输出
func PrettifyJson(v any) string {
	dj, err := json.MarshalIndent(v, "", "	")
	if err != nil {
		logger.L().Panic().Msgf("json format error, err: %s", err.Error())
	}
	return string(dj)
}
