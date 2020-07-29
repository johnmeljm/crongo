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
	runTask(timer, test, []string{strconv.Itoa(time.Now().Second()), strconv.Itoa(time.Now().Nanosecond())})
}

func test(args []string) {
	log.Printf("%+v\n", args)
	log.Printf("%+v\n", time.Now().Second())
	log.Printf("%+v\n", time.Now().Nanosecond())
}
