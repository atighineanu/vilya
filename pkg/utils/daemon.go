package utils

import (
	"fmt"
	"strings"
	"unicode"
	"strconv"
	"math"
	"unicode/utf8"
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
	temp := strings.Split(period, "")
	var number, power int
	for i := len(temp) -1; i >= 0; i-- {
		if unicode.IsDigit([]rune(temp[i])[0]) {
			bit, err := strconv.Atoi(temp[i]) 
			if err != nil {
				return 0, fmt.Errorf("Couldn't convert string to int.")
			}
			number += int(float64(bit) * math.Pow(10, float64(power)) )
			power++
		} else {
			if i != len(temp) - 1 {
				return 0, fmt.Errorf("Bad format for input --period")
			} 
		}
		if i == 0 {
			switch temp[len(temp)-1] {
			case "s":
				return time.Duration(number) * time.Second, nil
			case "m":
				return time.Duration(number) * time.Minute, nil
			case "h":
				return time.Duration(number) * time.Hour, nil
			return 0, fmt.Errorf("Bad formatting. Cannot recognize time unit.")
			}
		}
	}
	return 0, nil
}


