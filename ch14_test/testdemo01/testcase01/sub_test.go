package main

import (
	"testing"
)

func TestGetSub(t *testing.T) {
	res := getSub(10, 5)
	if res != 5 {
		t.Fatalf("getSub錯誤,返回值:%v,預期值:%v\n", res, 5)
	}
	t.Logf("agetSub正確,返回值:%v,預期值:%v\n", res, 5)
}
