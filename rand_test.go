package gorand

import (
	"testing"
	"time"
)

func TestRandStrWithCharset(t *testing.T) {
	first := RandStrWithCharset(6, "0123456789")
	time.Sleep(1 * time.Second)
	second := RandStrWithCharset(6, "0123456789")
	if first == second {
		t.Fatal("No randomness even after 1 sec delay ")
	}
}

func TestRandStr(t *testing.T) {
	first := RandStr(6)
	time.Sleep(1 * time.Second)
	second := RandStr(6)
	if first == second {
		t.Fatal("No randomness even after 1 sec delay ")
	}
}
