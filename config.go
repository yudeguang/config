package config

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

/*
读取配置文件，简单用=分隔键与值，出现多个相同键时，以最后一个值为准
*/

type ConfigStruct struct {
	mapConfigs map[string]string
}

//配置文件初始化函数，如果配置文件放置在所在程序同目录下并且名称为config.ini则无需执行此函数
func NewConfig(fileName string) (*ConfigStruct, error) {
	var configs ConfigStruct
	configs.mapConfigs = make(map[string]string)
	if fileName == "" {
		fileName = "config.ini"
	}
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		createExampleConfig()
		return nil, fmt.Errorf("配置文件不存在，请参考程序所在目录下,config_example.ini")
	}
	//如果本来就是utf-8就不需要再转换,utf-8文件头前三位为0xef, 0xbb, 0xbf
	if bytes.Equal(data[0:3], []byte{0xef, 0xbb, 0xbf}) {
		configs.parse(string(data))
	} else {
		configs.parse(string(data))
	}
	return &configs, nil
}

//拆分以\r\n换行分隔的数据到对象
func (this *ConfigStruct) parse(sour string) {
	equal, newLine := "=", "\r\n"
	if sour != "" {
		arrLines := strings.Split(sour, newLine)
		for _, line := range arrLines {
			pos := strings.Index(line, equal)
			if pos > 0 {
				key := strings.ToLower(line[:pos])
				val := line[pos+1:]
				this.mapConfigs[key] = val
			}
		}
	}
}

//根据键值获得一个字符串
func (this *ConfigStruct) Get(key string) string {
	smallkey := strings.ToLower(key)
	return this.mapConfigs[smallkey]
}

//根据键值获得一个数字
func (this *ConfigStruct) GetInt(key string) int {

	smallkey := strings.ToLower(key)
	val, ok := this.mapConfigs[smallkey]
	if !ok {
		return 0
	}
	n, err := strconv.ParseInt(val, 10, 64)
	if err != nil {
		return 0
	}
	return int(n)
}

//生成样本配置文件
func createExampleConfig() {
	var configText = "[config]"
	configText = configText + "\r\nKeyExample=ValueExample"
	ioutil.WriteFile("./config_example.ini", []byte(configText), 0x666)
}
