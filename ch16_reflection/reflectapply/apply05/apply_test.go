package test

import (
	"reflect"
	"testing"
)

type user struct {
	UserID string
	Name   string
}

func TestReflectStructPtr(t *testing.T) {
	var (
		model *user
		st    reflect.Type
		elem  reflect.Value
	)

	st = reflect.TypeOf(model)
	t.Log("reflect.TypeOf", st.Kind().String())
	st = st.Elem()
	t.Log("reflect.TypeOf.Elem", st.Kind().String())
	elem = reflect.New(st)

	t.Log("reflect.New", elem.Kind().String())
	t.Log("reflect.New.Elem", elem.Elem().Kind().String())

	model = elem.Interface().(*user)
	elem = elem.Elem()
	elem.FieldByName("UserID").SetString("369")
	elem.FieldByName("Name").SetString("999")
	t.Log("model model.Name", model, model.Name)
}
