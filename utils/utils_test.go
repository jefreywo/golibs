package utils

import (
	"fmt"
	"testing"
)

func TestFilterEmoji(t *testing.T) {
	fmt.Println(FilterEmoji("Thats a nice joke 😆😆😆 😛"))
}

func TestRound(t *testing.T) {
	if r := Round(2.4515345326, 0); r != 2 {
		t.Fatalf("不保留小数error:%v\n", r)
	}
	if r := Round(2.4515345326, 3); r != 2.452 {
		t.Fatalf("保留三位小数error:%v\n", r)
	}
}

func TestIsEqual(t *testing.T) {
	if IsEqual(2.33333334, 2.2, 0.000001) != false {
		t.Fatal("error")
	}
	if IsEqual(2.33333334, 2.333333, 0.000001) != true {
		t.Fatal("error")
	}
}
