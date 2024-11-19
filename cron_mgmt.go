package crongo

import (
	"sync"
	"time"
)

// CronMgmt cron management
type CronMgmt struct {
	List  []CronItem
	Debug int
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

// NewDebug return a new instance of CronMgmt with debug info
func NewDebug() CronMgmt {
	return CronMgmt{Debug: 1}
}

// Add add new cron
func (c *CronMgmt) Add(timeString string, f func()) {
	t, _ := timerParser(timeString)
	c.List = append(c.List, CronItem{T: t, F: nil, FNP: f, Args: map[string]interface{}{}})
}

// AddWithParam add new cron with params
func (c *CronMgmt) AddWithParams(timeString string, f func(map[string]interface{}), args map[string]interface{}) {
	t, _ := timerParser(timeString)
	c.List = append(c.List, CronItem{T: t, F: f, FNP: nil, Args: args})
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
			runTask(i.T, i.F, i.FNP, i.Args, c.Debug != 0)
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
	// 固定每分钟2秒时执行, 避免趋近于整分钟时, 会跳过的情况
	time.Sleep(time.Duration((62 - time.Now().Second()) % 60))
	t := time.NewTicker(time.Minute)
	defer t.Stop()
	for range t.C {
		for _, v := range c.List {
			go func(i CronItem) {
				runTask(i.T, i.F, i.FNP, i.Args, c.Debug != 0)
			}(v)
		}
	}
}
