/*
Create: 2023/2/17
Project: Create-JJAPP
Github: https://github.com/landers1037
Copyright Renj
*/

package main

import (
	"io/ioutil"
)

var shellTemplates = []string{"start.sh", "stop.sh", "check.sh", "restart.sh", "kill.sh"}

func createShell() error {
	log.Println("start to create project scripts")
	for _, s := range shellTemplates {
		data := "# !/bin/bash"
		log.Printf("create script %s", s)
		if err := ioutil.WriteFile(s, []byte(data), 0755); err != nil {
			return err
		}
	}

	return nil
}
