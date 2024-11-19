# crongo
golang cron package

# usage
- Func New return a instance of cron management.
- Func NewDebug return a instance of cron management with debug info.
- Func AddWithParams has 3 parameters, the first parameter format is the same as the Linux Crontab, the second parameter is the calling function, the third parameter is args for the function.
- Func AddNoParam has 2 parameters, both parameters are the same as Func Add's first two parameters.
- Func Run starting to run the task.

```
import github.com/johnmeljm/crongo

go func() {
    // call function at the first time, optional.
    demo.DoSomeThing(map[string]interface{}{"a": "Foo", "b": "Bar"})

    c := crongo.New()
    // c := crongo.NewDebug()
    // call function with parameter, parameter type is map[string]interface{}
    c.AddWithParams("* * * * *", demo.DoSomeThing, map[string]interface{}{"a": "Foo", "b": "Bar"})
    // call function without parameter
    c.Add("* * * * *", demo.DoAnotherThing)
    c.Run()
}()
```