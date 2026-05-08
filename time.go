package time

import (
	"database/sql/driver"
	"fmt"
	"time"
)

type Time time.Time

func (t Time) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(time.DateTime)+2)
	b = append(b, '"')
	b = time.Time(t).AppendFormat(b, time.DateTime)
	b = append(b, '"')
	return b, nil
}

func (t *Time) UnmarshalJSON(data []byte) (err error) {
	now, err := time.ParseInLocation(`"`+time.DateTime+`"`, string(data), time.Local)
	if err != nil {
		return err
	}
	*t = Time(now)
	return
}

func (t Time) String() string {
	return time.Time(t).Format(time.DateTime)
}

func (t Time) Value() (driver.Value, error) {
	var _time = time.Time(t)
	if _time.IsZero() {
		return nil, nil
	}
	return _time, nil
}

func (t *Time) Scan(v any) error {
	value, ok := v.(time.Time)
	if ok {
		*t = Time(value)
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}
