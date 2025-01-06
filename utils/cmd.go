package utils

import "reflect"

type Cmd interface {
	Output() ([]byte, error)
	Run() error
}

func Marshal(any interface{}) []string {
	var args []string

	t := reflect.TypeOf(any)
	v := reflect.ValueOf(any)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		tag := field.Tag.Get("ArgsPre")
		value := v.Field(i).String()
		if value == "" {
			continue
		}
		args = append(args, tag, value)
	}
	return args
}

//func Unmarshal(data []byte, any interface{}) error {
//
//}
