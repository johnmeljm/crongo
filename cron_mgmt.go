package crongo

import "sync"

// CronMgmt cron management
type CronMgmt struct {
	List   []CronItem
	Status int
}

// CronItem list
type CronItem struct {
	T    Timer
	F    func([]string)
	Args []string
}

// New return a new instance of CronMgmt
func New() CronMgmt {
	return CronMgmt{}
}

// Add add new cron
func (c *CronMgmt) Add(timeString string, f func([]string), args []string) {
	t, _ := timerParser(timeString)
	c.List = append(c.List, CronItem{T: t, F: f, Args: args})
}

// Run start Cron
func (c *CronMgmt) Run() {
	if len(c.List) <= 0 {
		return
	}
	wg := sync.WaitGroup{}
	wg.Add(len(c.List))
	for _, v := range c.List {
		go func(i CronItem) {
			runTask(i.T, i.F, i.Args)
			wg.Done()
		}(v)
	}
	wg.Wait()
}
