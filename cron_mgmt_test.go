package crongo

import (
	"log"
	"testing"
	"time"
)

func TestMgmt(t *testing.T) {
	a := New()
	a.Add("* * * * *", demo1)
	a.AddWithParams("* * * * *", demo2, map[string]interface{}{"a": "a", "b": "b"})
	a.Run()
}

func demo1() {
	log.Printf("%+v\n", "func no params")
}

func demo2(sArr map[string]interface{}) {
	log.Printf("%+v\n", sArr)
	time.Sleep(time.Minute * 2)
}
