package crongo

import (
	"log"
	"testing"
	"time"
)

func TestMgmt(t *testing.T) {
	a := New()
	a.Add("* * * * *", demo, []string{"a", "b"})
	a.Run()
}

func demo(sArr []string) {
	log.Printf("%+v\n", sArr)
	time.Sleep(time.Minute * 2)
}
