package crongo

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// runTask run by time and func
func runTask(t Timer, f func([]string), args []string) {
	var once = make(chan int)

	go func(args []string) {
		for {
			select {
			case <-once:
				go f(args)
			}
		}
	}(args)

	for {
		for _, vMonth := range t.TMonth {
			if int(time.Now().Month()) == int(vMonth) {
				for _, vDay := range t.TDay {
					if time.Now().Day() == int(vDay) {
						for _, vHour := range t.THour {
							if time.Now().Hour() == int(vHour) {
								for _, vMinute := range t.TMinute {
									if time.Now().Minute() == int(vMinute) {
										if time.Now().Second() == 0 {
											once <- 1
											time.Sleep(time.Minute*1 - time.Nanosecond*time.Duration(time.Now().Nanosecond()))
										}
									}
								}
							}
						}
					}
				}
				for _, vWeek := range t.TWeek {
					if int(time.Now().Weekday()) == int(vWeek) {
						for _, vHour := range t.THour {
							if time.Now().Hour() == int(vHour) {
								for _, vMinute := range t.TMinute {
									if time.Now().Minute() == int(vMinute) {
										if time.Now().Second() == 0 {
											once <- 1
											time.Sleep(time.Minute*1 - time.Nanosecond*time.Duration(time.Now().Nanosecond()))
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
}

// Timer return task running time format
type Timer struct {
	// TSecond []uint8 //0-59
	TMinute []uint8 //0-59
	THour   []uint8 //0-23
	TDay    []uint8 //1-31
	TMonth  []uint8 //1-12
	TWeek   []uint8 //0-6
}

// timerParser parse string to Timer format
// The param format is the same as the Linux Crontab
func timerParser(s string) (Timer, error) {
	// reg, _ := regexp.Compile(`[0-9*/- ]+`)
	matched, _ := regexp.MatchString(`[0-9*/\-] [0-9*/\-] [0-9*/\-] [0-9*/\-] [0-9*/\-]`, s)
	if !matched {
		return Timer{}, errors.New("format illegal")
	}
	timeArr := strings.Split(s, " ")
	if len(timeArr) != 5 {
		return Timer{}, errors.New("format illegal")
	}
	var timer Timer
	timer.TMinute = fmtItem(timeArr[0], 0, 59)
	timer.THour = fmtItem(timeArr[1], 0, 23)
	timer.TDay = fmtItem(timeArr[2], 1, 31)
	timer.TMonth = fmtItem(timeArr[3], 1, 12)
	timer.TWeek = fmtItem(timeArr[4], 0, 6)
	return timer, nil
}

func fmtItem(item string, min, max uint8) (result []uint8) {
	// *
	if item == "*" {
		var i uint8
		for i = min; i <= max; i++ {
			result = append(result, i)
		}
	}
	// 1,2,3,4
	if strings.Contains(item, ",") {
		rangeValue := strings.Split(item, ",")
		for _, v := range rangeValue {
			value, _ := strconv.Atoi(v)
			if uint8(value) >= min && uint8(value) <= max {
				result = append(result, uint8(value))
			}
		}
		// 1-10
	} else if strings.Contains(item, "-") {
		rangeValue := strings.Split(item, "-")
		start, _ := strconv.Atoi(rangeValue[0])
		if uint8(start) < min {
			start = int(min)
		}
		end, _ := strconv.Atoi(rangeValue[1])
		if uint8(end) > max {
			end = int(max)
		}
		var i uint8
		for i = uint8(start); i <= uint8(end); i++ {
			result = append(result, i)
		}
		// */5
	} else if strings.Contains(item, "/") {
		arithValue := strings.Split(item, "/")
		denominator, _ := strconv.Atoi(arithValue[1])
		// */5
		if arithValue[0] == "*" {
			var i uint8
			for i = min; i <= max; i++ {
				if (i-min)%uint8(denominator) == 0 {
					result = append(result, i)
				}
			}
		}
		// 1-10/5
		if strings.Contains(arithValue[0], "-") {
			rangeValue := strings.Split(arithValue[0], "-")
			start, _ := strconv.Atoi(rangeValue[0])
			if uint8(start) < min {
				start = int(min)
			}
			end, _ := strconv.Atoi(rangeValue[1])
			if uint8(end) > max {
				end = int(max)
			}
			var i uint8
			for i = uint8(start); i <= uint8(end); i++ {
				if (i-uint8(start))%uint8(denominator) == 0 {
					result = append(result, i)
				}
			}
		}
	}

	return
}
