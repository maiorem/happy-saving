package common

import (
	"encoding/json"
	"time"
)

type MyTime time.Time

const MyTimeFormat = "2000-05-10"
const MyLocation = "Asia/Seoul"

var _ json.Unmarshaler = &MyTime{}

func (mt *MyTime) UnmarshalJSON(bs []byte) error {
	var s string

	err := json.Unmarshal(bs, &s)
	if err != nil {
		return err
	}

	location, err := time.LoadLocation(MyLocation)
	if err != nil {
		return err
	}

	t, err := time.ParseInLocation(MyTimeFormat, s, location)
	if err != nil {
		return err
	}

	*mt = MyTime(t)

	return nil
}
