package crongo

import (
	"sync"
	"time"
)

// CronMgmt cron management
type CronMgmt struct {
	List   []CronItem
	Status int
}

// CronItem list
type CronItem struct {
	T    Timer
	F    func(map[string]interface{})
	FNP  func()
	Args map[string]interface{}
}

// New return a new instance of CronMgmt
func New() CronMgmt {
	return CronMgmt{}
}

// Add add new cron
func (c *CronMgmt) Add(timeString string, f func(map[string]interface{}), args map[string]interface{}) {
	t, _ := timerParser(timeString)
	c.List = append(c.List, CronItem{T: t, F: f, FNP: nil, Args: args})
}

// AddNoParam add new cron
func (c *CronMgmt) AddNoParam(timeString string, f func()) {
	t, _ := timerParser(timeString)
	c.List = append(c.List, CronItem{T: t, F: nil, FNP: f, Args: map[string]interface{}{}})
}

// Run start Cron
func (c *CronMgmt) Run_bak() {
	if len(c.List) <= 0 {
		return
	}
	wg := sync.WaitGroup{}
	wg.Add(len(c.List))
	for _, v := range c.List {
		go func(i CronItem) {
			runTask(i.T, i.F, i.FNP, i.Args)
			wg.Done()
		}(v)
	}
	wg.Wait()
}

// Run start Cron
func (c *CronMgmt) Run() {
	if len(c.List) <= 0 {
		return
	}
	t := time.NewTicker(time.Minute)
	for {
		select {
		case <-t.C:
			for _, v := range c.List {
				go func(i CronItem) {
					runTask(i.T, i.F, i.FNP, i.Args)
				}(v)
			}
		}
	}
}
