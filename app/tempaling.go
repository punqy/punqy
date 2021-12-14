package app

import (
	"encoding/json"
	logger "github.com/sirupsen/logrus"
	"html/template"
	"strconv"
	"strings"
	"time"
)

func TemplatingConfig() template.FuncMap {
	return template.FuncMap{
		"to_upper": strings.ToUpper,
		"inc": func(i int) int {
			return i + 1
		},
		"format": func(t time.Time, format string) string {
			return t.Format(format)
		},
		"minus": func(a, b int) int {
			return a - b
		},
		"mul": func(a, b int) int {
			return a * b
		},
		"plus": func(a, b int) int {
			return a + b
		},
		"to_string": func(value interface{}) string {
			switch v := value.(type) {
			case string:
				return v
			case int:
				return strconv.Itoa(v)
			case uint32:
				return strconv.FormatUint(uint64(v), 10)
			case uint64:
				return strconv.FormatUint(v, 10)
			// Add whatever other types you need
			default:
				return ""
			}
		},
		"to_json": func(arg interface{}) string {
			marshal, err := json.Marshal(arg)
			if err != nil {
				logger.Error(err)
			}
			return string(marshal)
		},
	}
}
