package utils

import (

	//"math"
	//"unicode/utf8"
	"time"

	"github.com/atighineanu/shoutrrr"
)

func (config *VilyaCfg) RunAndNotify(duration time.Duration, url string, message string) error {
	for {
		shoutrrr.Send(url, message)
		time.Sleep(duration)
	}
	return nil
}
