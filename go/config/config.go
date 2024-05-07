/*
 * @Author: Callay 2415993100@qq.com
 * @Date: 2024-05-04 13:57:52
 * @LastEditors: Callay 2415993100@qq.com
 * @LastEditTime: 2024-05-07 13:19:07
 * @FilePath: \checkIn\go\config\config.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Server   Server   `yaml:"server"`
	Database Database `yaml:"database"`
	Jwt      Jwt      `yaml:"jwt"`
}

type Server struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

type Database struct {
	Mysql Mysql
}

type Mysql struct {
	DbName   string `yaml:"dbname"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

type Jwt struct {
	Key string `yaml:"key"`
}

var Application Config

func init() {
	// 打开 YAML 文件
	file, err := os.Open("./config/config.yml")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// 创建解析器
	decoder := yaml.NewDecoder(file)

	// 解析 YAML 数据
	err = decoder.Decode(&Application)
	if err != nil {
		fmt.Println("Error decoding YAML:", err)
		return
	}
}
