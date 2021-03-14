package converter

import (
	"reflect"
	"testing"
)

type Pointer struct {
	X int `convert:"field:X-pointer,from:XP"`
	Y int `convert:"field:Y-pointer"`
	Z int
}

func TestTag(t *testing.T) {
	p := Pointer{1, 2, 3}

	typ := reflect.TypeOf(p)
	x, has := typ.FieldByName("X")
	if !has {
		t.Fatal()
	}
	tx := newTag(&x)
	if tx.field != "X-pointer" && tx.from != "XP" {
		t.Fatal()
	}

	y, has := typ.FieldByName("Y")
	if !has {
		t.Fatal()
	}
	ty := newTag(&y)
	if ty.field != "Y-pointer" && ty.field != "Y" {
		t.Fatal()
	}

	z, has := typ.FieldByName("Z")
	if !has {
		t.Fatal()
	}
	tz := newTag(&z)
	if tz.field != "Z" && tz.from != "Z" {
		t.Fatal()
	}
}
