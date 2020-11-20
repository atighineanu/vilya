package utils

import (
	"fmt"
	"strings"
	"unicode"
	"strconv"
	//"math"
	//"unicode/utf8"
	"time"
)

func (config *VilyaCfg) RunAndNotify(duration time.Duration, url string, message string) error {
	for {
		fmt.Println("tzatzatza")
		time.Sleep(duration)
	}
	return nil
}

func (config *VilyaCfg) ParsePeriod(period string) (time.Duration, error) {
	number, err := strconv.Atoi(strings.TrimRightFunc(period, func(r rune) bool {
		return unicode.IsLetter(r)
	}))
	if err != nil {
		return 0, fmt.Errorf("Couldn't convert string to int.")
	}
	if unicode.IsLetter([]rune(period[len(period)-1:])[0]) {
		switch period[len(period)-1:] {
		case "s":
			return time.Duration(number) * time.Second, nil
		case "m":
			return time.Duration(number) * time.Minute, nil
		case "h":
			return time.Duration(number) * time.Hour, nil
		default:
			return 0, fmt.Errorf("Bad formatting. Cannot recognize time unit.")
		}
	} else {
		return 0, fmt.Errorf("Bad formatting. Cannot recognize time unit.")
	}
}


