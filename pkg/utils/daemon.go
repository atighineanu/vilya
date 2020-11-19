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

func (config *VilyaCfg) RunDaemon(duration time.Duration) error {
	return nil
}

func (config *VilyaCfg) ParsePeriod(period string) (map[string]int, error) {
	temp := strings.Split(period, "")
	var number, power int
	timemap := make(map[string]int)
	for i := len(temp) -1; i >= 0; i-- {
		if unicode.IsDigit([]rune(temp[i])[0]) {
			bit, err := strconv.Atoi(temp[i]) 
			if err != nil {
				return nil, fmt.Errorf("Couldn't convert string to int.")
			}
			number += int(float64(bit) * math.Pow(10, float64(power)) )
			power++
		} else {
			if i != len(temp) - 1 {
				return nil, fmt.Errorf("Bad format for input --period")
			} 
		}
		if i == 0 {
			switch temp[len(temp)-1] {
			case "s":
				timemap["seconds"] = number
				return timemap, nil
			case "m":
				timemap["minutes"] = number
				return timemap, nil
			case "h":
				timemap["hours"] = number
				return timemap, nil
			return nil, fmt.Errorf("Bad formatting. Cannot recognize time unit.")
			}
		}
	}
	//fmt.Println(temp)
	return nil, nil
}

func NewTimez(period string) (time.Duration, error) {
	r, v := utf8.DecodeLastRuneInString(period)
	if r == unicode.ReplacementChar {
		return 0, fmt.Errorf("%v invalid input", period)
	}
	n, err := strconv.Atoi(period[:len(period)-v])
	if err != nil {
		return 0, err
	}
	return convTimez(time.Duration(n), r), nil
}

func convTimez(n time.Duration, r rune) time.Duration {
	var res time.Duration
	switch r {
	case 's':
		res =  n * time.Second
	case 'm':
		res= n * time.Minute
	case 'h':
		res= n * time.Hour
	}
	return res	
}


