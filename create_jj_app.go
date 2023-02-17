/*
Create: 2023/2/17
Project: Create-JJAPP
Github: https://github.com/landers1037
Copyright Renj
*/

package main

import (
	"fmt"
	"os"

	"github.com/AlecAivazis/survey/v2"
	"github.com/AlecAivazis/survey/v2/terminal"
	"github.com/JJApplication/CJA/internal"
)

func exit(err error) {
	if err == terminal.InterruptErr {
		log.Println("exit")
		os.Exit(0)
	}
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

func main() {
	fmt.Printf("👻 Thanks for using %s to create application of projectJJ\n", internal.APP_NAME)
	fmt.Printf("🐼 Copyright: %s\n🔗 Go to site: http://%s\n📧 Email: %s\n\n",
		internal.COPYRIGHT,
		internal.SITE,
		internal.MAIL,
	)
	// 创建项目名称
	projectName := ""
	promptProjectName := &survey.Input{
		Message: "your project name: ",
	}
	exit(survey.AskOne(promptProjectName, &projectName))

	// 选择项目类型 【web app】 【noengine app】 【empty app】
	chooseType := ""
	promptChooseType := &survey.Select{
		Message: "Choose application type:",
		Options: APPType,
		Default: APPTypeDefault,
	}
	exit(survey.AskOne(promptChooseType, &chooseType))

	// 选择.octopus数据类型 【json】 【yaml】 【pig】
	chooseDataType := ""
	promptChooseDataType := &survey.Select{
		Message: "Choose metadata type:",
		Options: APPMetaType,
		Default: APPMetaTypeDefault,
	}
	exit(survey.AskOne(promptChooseDataType, &chooseDataType))

	// 选择服务通信协议 【http】 【http-based rpc】 【uds】
	var protocol []string
	promptProto := &survey.MultiSelect{
		Message: "What protocols do you prefer:",
		Options: APPProto,
		Default: APPProtoDefault,
	}
	exit(survey.AskOne(promptProto, &protocol))

	// 开始生成
	startTodo := true
	promptStart := &survey.Confirm{
		Message: "Do you like to create the project right now?",
		Default: true,
	}
	exit(survey.AskOne(promptStart, &startTodo))
	if startTodo {
		CreateDir(projectName, chooseType, chooseDataType, protocol)
	}
	log.Println("🥤 Enjoy yourself!!!")
}
