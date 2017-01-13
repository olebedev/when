# when [![godoc](http://img.shields.io/badge/godoc-reference-blue.svg?style=flat)](https://godoc.org/github.com/olebedev/when) [![wercker status](https://app.wercker.com/status/a04ca8246bf35621b2665a73c1ed765f/s/master "wercker status")](https://app.wercker.com/project/byKey/a04ca8246bf35621b2665a73c1ed765f)

> `when` is a natural language date/time parser with pluggable rules and merge strategies

### Examples

* **tonight at 11:10 pm**
* at **Friday afternoon**
* the deadline is **next tuesday 14:00**
* drop me a line in **next wednesday at 2:25 p.m**
* it could be done at **11 am past tuesday**

Check [EN](https://github.com/olebedev/when/blob/master/rules/en) and [RU](https://github.com/olebedev/when/blob/master/rules/ru) rules and tests for them, for more examples.

**Needed rule not found?**
Open [an issue](https://github.com/olebedev/when/issues/new) with the case and it will be added asap.

### How it works

Usually, there are several rules added to the parser's instance for checking. Each rule has own borders - length and offset in provided string. Meanwhile, each rule yields only the first match over the string. So, the library is checking all the rules and extracting a cluster of matched rules which have distance between each other less or equal [`options.Distance`](https://github.com/olebedev/when/blob/master/when.go#L141-L144), 5 by default. For example:

```
in next wednesday at 2:25 p.m.
   └──────┬─────┘    └───┬───┘
   	   weekday      hour + minute
```

So, we have a cluster of matched rules - `"next wednesday at 2:25 p.m."` in the string representation. 

After that, each rule is being applied to the context. In order of definition or in match order, if [`options.MatchByOrder`](https://github.com/olebedev/when/blob/master/when.go#L141-L144) set in `true`(by default). Each rule could be applied with given merge strategy. By default, it's an [Override](https://github.com/olebedev/when/blob/master/rules/rules.go#L13) strategy. The other strategies are not implemented yet in the rules. **Pull requests are welcome.**

### Usage

```go
w := when.New(nil)
w.Add(en.All...)
w.Add(common.All...)

text := "drop me a line in next wednesday at 2:25 p.m"
r, err := w.Parse(text, time.Now())
if err != nil {
	// an error has occurred
}
if  r == nil {
 	// no matches found
}

fmt.Println(
	"the time",
	r.Time.String(),
	"mentioned in",
	text[r.Index:len(r.Text)],
)
```

### State of the project

The project is in a more-or-less complete state. It used for one project already. Bugs will be fixed as soon as they will be found.

### TODO

- [ ] readme: describe all the existing rules
- [ ] implement rules for all [these examples](https://github.com/mojombo/chronic#examples)
- [ ] add cli and simple rest api server([#2](https://github.com/olebedev/when/issues/2))

### LICENSE

http://www.apache.org/licenses/LICENSE-2.0

