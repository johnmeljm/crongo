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

go func() {
    // call function at the first time, optional.
    demo.DoSomeThing(map[string]interface{}{"a": "Foo", "b": "bar"})

    c := crongo.New()
    // call function with parameter, parameter type is map[string]interface{}
    c.Add("* * * * *", demo.DoSomeThing, map[string]interface{}{"a": "Foo", "b": "bar"})
    // call function without parameter
    c.AddNoParam("* * * * *", demo.DoAnotherThing)
    c.Run()
}()
```