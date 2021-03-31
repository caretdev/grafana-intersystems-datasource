package intersystems

import (
	"strconv"
	"strings"
	"time"
)

func FromHorolog(h string) (time.Time, error) {
	dateParts := strings.Split(h, ",")
	if datePart, err := strconv.ParseInt(dateParts[0], 10, 64); err != nil {
		return time.Time{}, err
	} else if timePart, err := strconv.ParseInt(dateParts[1], 10, 64); err != nil {
		return time.Time{}, err
	} else {
		tm := time.Unix((datePart - 47117) * 86400 + timePart, 0)
		return tm, nil
	}

}