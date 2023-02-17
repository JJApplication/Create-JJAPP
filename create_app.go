/*
Create: 2022/7/7
Project: FuShin
Github: https://github.com/landers1037
Copyright Renj
*/

// Package private
package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/JJApplication/fushin/utils/stream"
	"github.com/JJApplication/octopus_meta"
)

// 创建app 当前支持web服务 client服务 noengine代理服务 空模板服务
// 空模板只存在.octupus元数据

func createOctopus(appName string, format string) error {
	log.Println("start to create project meta")
	app := octopus_meta.App{
		Name:          appName,
		ID:            "app_" + strings.ToLower(appName),
		Type:          octopus_meta.SourceApp,
		ReleaseStatus: octopus_meta.Unreleased,
		EngDes:        "",
		CHSDes:        "",
		Link:          "",
		ManageCMD: octopus_meta.CMD{
			Start:     "start.sh",
			Stop:      "stop.sh",
			Restart:   "restart.sh",
			ForceKill: "kill.sh",
			Check:     "check.sh",
		},
		Meta: octopus_meta.Meta{
			Version: "0.0.1",
		},
		RunData:       octopus_meta.RunData{},
		Runtime:       octopus_meta.Runtime{},
		ResourceLimit: octopus_meta.ResourceLimit{},
	}

	var data []byte
	var err error
	switch format {
	case "json", "pig", "default":
		data, err = stream.JSON.MarshalIndent(app, "", "  ")
		if err != nil {
			return err
		}
	case "yaml":
		data, err = stream.YAML.Marshal(app)
		if err != nil {
			return err
		}
	default:
		data, err = stream.JSON.MarshalIndent(app, "", "  ")
		if err != nil {
			return err
		}
	}

	log.Printf("%s created", fmt.Sprintf("%s.%s", appName, format))
	return ioutil.WriteFile(fmt.Sprintf("%s.%s", appName, format), data, 0644)
}
