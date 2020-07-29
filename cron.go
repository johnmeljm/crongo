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
		matchWeekday := intInList(uint8(time.Now().Weekday()), t.TWeekday)
		matchMonth := intInList(uint8(time.Now().Month()), t.TMonth)
		matchDay := intInList(uint8(time.Now().Day()), t.TDay)
		matchHour := intInList(uint8(time.Now().Hour()), t.THour)
		matchMinute := intInList(uint8(time.Now().Minute()), t.TMinute)
		if matchMonth && (matchDay || matchWeekday) && matchHour && matchMinute {
			if time.Now().Second() == 0 {
				once <- 1
				time.Sleep(time.Minute*1 - time.Nanosecond*time.Duration(time.Now().Nanosecond()))
			}
		}
	}
}

func intInList(item uint8, list []uint8) bool {
	if len(list) <= 0 {
		return false
	}
	for _, v := range list {
		if item == v {
			return true
		}
	}
	return false
}

// Timer return task running time format
type Timer struct {
	// TSecond  []uint8 //0-59
	TMinute  []uint8 //0-59
	THour    []uint8 //0-23
	TDay     []uint8 //1-31
	TMonth   []uint8 //1-12
	TWeekday []uint8 //0-6
}

// timerParser parse string to Timer format
// The param format is the same as the Linux Crontab
func timerParser(s string) (Timer, error) {
	matched, _ := regexp.MatchString(`[0-9*/\-]* [0-9*/\-]* [0-9*/\-]* [0-9*/\-]* [0-9*/\-]*`, s)
	if !matched {
		return Timer{}, errors.New("format illegal")
	}
	timeArr := strings.Split(s, " ")
	if len(timeArr) != 5 {
		return Timer{}, errors.New("length illegal")
	}
	var timer Timer
	// timer.TSecond = fmtItem(timeArr[0], 0, 59)
	timer.TMinute = fmtItem(timeArr[0], 0, 59)
	timer.THour = fmtItem(timeArr[1], 0, 23)
	timer.TDay = fmtItem(timeArr[2], 1, 31)
	timer.TMonth = fmtItem(timeArr[3], 1, 12)
	timer.TWeekday = fmtItem(timeArr[4], 0, 6)
	return timer, nil
}

func fmtItem(item string, min, max uint8) (result []uint8) {
	if item == "*" { // *
		var i uint8
		for i = min; i <= max; i++ {
			result = append(result, i)
		}
	}
	if strings.Contains(item, "/") { // []/5
		arithValue := strings.Split(item, "/")
		denominator, _ := strconv.Atoi(arithValue[1])
		if arithValue[0] == "*" { // */5
			var i uint8
			for i = min; i <= max; i++ {
				if (i-min)%uint8(denominator) == 0 {
					result = append(result, i)
				}
			}
		}
		if strings.Contains(arithValue[0], "-") { // 1-10/5
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
	} else if strings.Contains(item, ",") { // 1,2,3,4
		rangeValue := strings.Split(item, ",")
		for _, v := range rangeValue {
			value, _ := strconv.Atoi(v)
			if uint8(value) >= min && uint8(value) <= max {
				result = append(result, uint8(value))
			}
		}
	} else if strings.Contains(item, "-") { // 1-10
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
	}

	return
}
