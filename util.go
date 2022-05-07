package indigo

import (
	"reflect"
	"time"
)

func convDate(dateIf interface{}) string {
	if dateIf == nil {
		return ""
	}
	var t time.Time
	var err error

	if reflect.TypeOf(dateIf).String() == "string" {
		dateStr := dateIf.(string)
		t, err = time.Parse("2006-01-02 15:04:05", dateStr)
	} else {
		dateMap := dateIf.(map[string]interface{})
		t, err = time.Parse("2006-01-02 15:04:05.000000", dateMap["date"].(string))
	}
	if err != nil {
		return ""
	}
	return t.Format(time.RFC3339)
}
