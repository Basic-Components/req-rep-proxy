package loadconfig

import (
	"encoding/json"
	"io/ioutil"
	"strings"

	error "github.com/Basic-Components/req-rep-proxy/error"
)

const defaultConfig string = `
{
"frontend_url":"tcp://*:5559",
"backend_url":"tcp://*:5560",
"debug":false,
"log_format":"json",
"log_output":"req-rep-proxy.log"
}
`

// Config json解析出来的配置结构
type Config struct {
	FrontendURL string `json:"frontend_url"`
	BackendURL  string `json:"backend_url"`
	LogFormat   string `json:"log_format"`
	LogOutput   string `json:"log_output"`
	Debug       bool   `json:"debug"`
}

func loadConfig(content string) Config {
	var jsonData Config
	if err := json.Unmarshal([]byte(content), &jsonData); err != nil {
		panic(err.Error())
	} else {
		return jsonData
	}
}

// LoadConfig 从json格式的配置文件中解析出配置对象
func LoadConfig(configPath string) Config {
	if configPath == "" {
		return loadConfig(defaultConfig)
	}
	if strings.HasSuffix(configPath, ".json") {
		b, err := ioutil.ReadFile(configPath)
		if err != nil {
			panic(err.Error())
		} else {
			return loadConfig(string(b))
		}
	} else {
		panic(error.ERR_CONFIG_TYPE)
	}

}
