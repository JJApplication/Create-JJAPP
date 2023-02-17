/*
Create: 2022/8/7
Project: FuShin
Github: https://github.com/landers1037
Copyright Renj
*/

// Package private
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/JJApplication/CJA/internal"
	"github.com/JJApplication/fushin/utils/json"
)

// 创建服务的目录

// CreateTag 生成服务的标签
func createTag(projectName, Type, DataType string, protocol []string) {
	log.Println("start to create project tag")
	var err error
	data, err := json.Json.MarshalIndent(map[string]interface{}{
		"create":    time.Now(),
		"generator": internal.APP_NAME,
		"project":   projectName,
		"type":      Type,
		"metaType":  DataType,
		"protocol":  protocol,
	}, "", "  ")
	if err != nil {
		log.Printf("create tag error: %s", err.Error())
		return
	}
	// 保证在当前目录下的projectName/projectName.meta.json
	if err = ioutil.WriteFile(fmt.Sprintf("%s.meta.json", projectName), data, 0644); err != nil {
		log.Printf("create tag error: %s", err.Error())
		return
	}
}

func CreateDir(projectName, Type, DataType string, protocol []string) {
	log.Println("start to create project")
	if projectName == "" {
		return
	}
	var err error
	if err = os.Mkdir(projectName, 0644); err != nil {
		log.Printf("create project dir error: %s", err.Error())
		return
	}

	if err = os.Chdir(projectName); err != nil {
		log.Printf("change to projectDir [%s] error: %s", projectName, err.Error())
	}
	createTag(projectName, Type, DataType, protocol)
	if err = createOctopus(projectName, DataType); err != nil {
		log.Printf("create app meta error: %s", err.Error())
	}
	if err = createShell(); err != nil {
		log.Printf("create app shell templates error: %s", err.Error())
	}
}
