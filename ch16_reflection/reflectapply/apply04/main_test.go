package main

import (
	"reflect"
	"testing"
)

type user struct {
	UserID string
	Name   string
}

func TestReflectStruct(t *testing.T) {
	var (
		model *user
		sv    reflect.Value
	)
	model = &user{}
	sv = reflect.ValueOf(model)
	t.Log("reflect.ValueOf", sv.Kind().String())
	sv = sv.Elem()
	t.Log("reflect.ValueOf.Elem", sv.Kind().String())
	sv.FieldByName("UserID").SetString("369")
	sv.FieldByName("Name").SetString("999")
	t.Log("model", model)
}
