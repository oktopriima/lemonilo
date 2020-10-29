/*
 * Name : Okto Prima Jaya
 * Github : https://github.com/oktopriima
 * Email : octoprima93@gmail.com
 * Created At : 29/10/2020, 20:52
 * Copyright (c) 2020
 */

package helper

import (
	"errors"
	"io/ioutil"
	"path/filepath"
	"strings"

	"github.com/BurntSushi/toml"
)

type x struct {
	key    string
	values interface{}
}

type translation struct {
	locales      []string
	translations map[string]map[string]interface{}
}

var trans *translation
var locale string

func InitLocales(trPath string, lang string) error {
	locale = lang
	trans = &translation{translations: make(map[string]map[string]interface{})}
	return loadTranslations(trPath)
}

func Tr(trKey string) string {
	trValue, ok := trans.translations[locale][trKey]
	if ok {
		return trValue.(string)
	}

	return trKey
}

func loadTranslations(trPath string) error {
	files, err := filepath.Glob(trPath + "/*.toml")
	if err != nil {
		return err
	}

	if len(files) == 0 {
		return errors.New("No translations found")
	}

	for _, file := range files {
		err := loadFileToMap(file)
		if err != nil {
			return err
		}
	}
	return nil
}

func loadFileToMap(filename string) error {
	var objmap map[string]interface{}

	localName := strings.Replace(filepath.Base(filename), ".toml", "", 1)

	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	err = toml.Unmarshal(content, &objmap)
	if err != nil {
		return err
	}
	trans.translations[localName] = objmap
	trans.locales = append(trans.locales, localName)

	return nil
}
