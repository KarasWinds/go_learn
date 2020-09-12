package main

import (
	"fmt"
	"testing"
)

func TestAddUpper(t *testing.T) {
	res := addUpper(10)
	if res != 55 {
		t.Fatalf("addUpper錯誤,返回值:%v,預期值:%v\n", res, 55)
	}
	t.Logf("addUpper正確,返回值:%v,預期值:%v\n", res, 55)
}

func TestHello(t *testing.T) {
	fmt.Println("Test Hello...")
}
