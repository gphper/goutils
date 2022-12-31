package timeutil

import (
	"testing"
	"time"
)

func TestYMD(t *testing.T) {
	if time.Now().Format(YMD) != time.Now().Format("2006-01-02") {
		t.Fail()
	}

	if time.Now().Format(YMDHIS) != time.Now().Format("2006-01-02 15:04:05") {
		t.Fail()
	}

	if time.Now().Format(HIS) != time.Now().Format("15:04:05") {
		t.Fail()
	}
}
