package main

import (
	"log"
	"reflect"
	"strconv"
	"strings"
)

// User используется для тестирования.
type User struct {
	Nick string
	Age  int `limit:"18"`
	Rate int `limit:"0,100"`
}

// Str2Int конвертирует строку в int.
func Str2Int(s string) int {
	v, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return v
}

// Validate проверяет min и max для int c тегом limit.
func Validate(obj interface{}) bool {
	vobj := reflect.ValueOf(obj)
	objType := vobj.Type() // получаем описание типа

	// перебираем все поля структуры
	for i := 0; i < objType.NumField(); i++ {
		// берём значение текущего поля и проверяем, что это int
		if v, ok := vobj.Field(i).Interface().(int); ok {
			// подсказка: тег limit надо искать в поле objType.Field(i)
			// objType.Field(i).Tag.Lookup или objType.Field(i).Tag.Get

			value := objType.Field(i).Tag.Get("limit")
			log.Printf("value=%s", value)
			if value != "" {
				nums_string := strings.Split(value, ",")
				log.Printf("len=%d", len(nums_string))
				if len(nums_string) == 1 {
					log.Printf("%d > %d", v, Str2Int(nums_string[0]))
					if v < Str2Int(nums_string[0]) {
						return false
					} 
				} else {
					log.Printf("%d < %d < %d", Str2Int(nums_string[0]), v, Str2Int(nums_string[1]))
					if v < Str2Int(nums_string[0]) || v > Str2Int(nums_string[1]) {
						return false
					}
				}
			}
		}
	}
	return true
}
