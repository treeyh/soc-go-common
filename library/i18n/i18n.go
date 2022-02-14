package i18n

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/treeyh/soc-go-common/core/config"
	"github.com/treeyh/soc-go-common/core/logger"
	"io/ioutil"
	"path"
	"strings"
)

var (
	log = logger.Logger()

	// langMaps 语言map
	langMaps = make(map[string]*viper.Viper)

	defaultLang = ""
)

// InitI18n 初始化i18n
func InitI18n(i18nConf *config.I18nConfig) {
	if i18nConf == nil || !i18nConf.Enable {
		return
	}
	defaultLang = i18nConf.DefaultLang
	files, err := ioutil.ReadDir(i18nConf.Path)
	if err != nil {
		panic(fmt.Sprintf("init i18n fail. err: %+v", err))
	}
	maps := make(map[string]*viper.Viper)
	for _, f := range files {
		ext := strings.ToLower(path.Ext(f.Name()))
		if ext != ".yaml" && ext != ".yml" {
			continue
		}
		index := strings.LastIndex(f.Name(), ".")
		if index < 1 {
			continue
		}
		langKey := f.Name()[0:index]
		filePath := path.Join(i18nConf.Path, f.Name())
		maps[langKey] = loadLangMap(filePath)
	}
	langMaps = maps
}

func GetDefaultLang() string {
	return defaultLang
}

func loadLangMap(filePath string) *viper.Viper {
	conf := viper.New()
	conf.SetConfigFile(filePath)
	conf.SetConfigType("yaml")
	//尝试进行配置读取
	if err := conf.ReadInConfig(); err != nil {
		panic(fmt.Sprintf("init i18n fail. path: %s; err: %+v", filePath, err))
	}
	return conf
}

func Get(lang, key string) string {
	if v, ok := langMaps[lang]; ok {
		return v.GetString(key)
	}
	return ""
}

func GetByDefault(lang, key, defaultValue string) string {
	value := Get(lang, key)
	if value == "" {
		value = defaultValue
	}
	return value
}
