package main

import "testing"

func TestFoo(t *testing.T) {
	r := Foo(10, 20)
	if r != 30 {
		t.Errorf("bad answer")
	}
}

func TestFoo2(t *testing.T) {
	r := Foo(1, 20)
	if r != 21 {
		t.Errorf("bad answer")
	}
}
