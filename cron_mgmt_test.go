package crongo

import (
	"log"
	"testing"
	"time"
)

func TestMgmt(t *testing.T) {
	a := New()
	a.Add("* * * * *", demo, map[string]interface{}{"a": "a", "b": "b"})
	a.AddNoParam("* * * * *", demo2)
	a.Run()
}

func demo(sArr map[string]interface{}) {
	log.Printf("%+v\n", sArr)
	time.Sleep(time.Minute * 2)
}

func demo2() {
	log.Printf("%+v\n", "func no params")
}
