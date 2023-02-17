/*
Create: 2023/2/18
Project: Create-JJAPP
Github: https://github.com/landers1037
Copyright Renj
*/

package main

import (
	"fmt"

	"github.com/JJApplication/CJA/internal"
)

type logger struct {
	name string
}

var log = logger{internal.APP_NAME}

func (l *logger) Printf(f string, v ...interface{}) {
	fmt.Println(l.fmt(fmt.Sprintf(f, v...)))
}

func (l *logger) Println(v ...interface{}) {
	fmt.Println(l.fmt(fmt.Sprint(v...)))
}

func (l *logger) fmt(s string) string {
	return fmt.Sprintf("[%s] %s", l.name, s)
}
