package crongo

import (
	"log"
	"strconv"
	"testing"
	"time"
)

func TestTask(t *testing.T) {
	timer, err := timerParser("33-50/3 * 1-6 * 2-3")
	log.Printf("%+v\n", timer)
	log.Printf("%+v\n", err)
	runTask(
		timer,
		testFunc,
		nil,
		map[string]interface{}{
			"second":     strconv.Itoa(time.Now().Second()),
			"nanosecond": strconv.Itoa(time.Now().Nanosecond()),
		},
		true,
	)
}

func testFunc(args map[string]interface{}) {
	log.Printf("%+v\n", args)
	log.Printf("%+v\n", time.Now().Second())
	log.Printf("%+v\n", time.Now().Nanosecond())
}
