# crongo
golang cron package

# usage
Func New return a instance of cron management, 
Func Add has 3 params, the first param format is the same as the Linux Crontab,
the second param is the calling function, the third param is args for the 
function.
Func Run starting to run the task
```
import github.com/zmwater/crongo

c := crongo.New()
c.Add("* * * * *", demo, []string{"a", "b"})
c.Run()
```