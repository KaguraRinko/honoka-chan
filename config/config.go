// Copyright (C) 2021 YumeMichi
//
// SPDX-License-Identifier: Apache-2.0
package config

import (
	"honoka-chan/utils"
	"os"

	_ "github.com/mattn/go-sqlite3"
	"xorm.io/xorm"
)

var (
	ConfName  = "config.yml"
	Conf      = &AppConfigs{}
	ExampleDb = "assets/data.example.db"
	MainDb    = "assets/main.db"
	UserDb    = "assets/data.db"
	MainEng   *xorm.Engine
	UserEng   *xorm.Engine
)

func init() {
	Conf = Load(ConfName)

	_, err := os.Stat(UserDb)
	if err != nil {
		utils.WriteAllText(UserDb, utils.ReadAllText(ExampleDb))
	}

	eng, err := xorm.NewEngine("sqlite3", MainDb)
	if err != nil {
		panic(err)
	}
	err = eng.Ping()
	if err != nil {
		panic(err)
	}
	MainEng = eng
	MainEng.SetMaxOpenConns(50)
	MainEng.SetMaxIdleConns(10)

	eng, err = xorm.NewEngine("sqlite3", UserDb)
	if err != nil {
		panic(err)
	}
	err = eng.Ping()
	if err != nil {
		panic(err)
	}
	UserEng = eng
	UserEng.SetMaxOpenConns(50)
	MainEng.SetMaxIdleConns(10)
}
