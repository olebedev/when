# when [![godoc](http://img.shields.io/badge/godoc-reference-blue.svg?style=flat)](https://godoc.org/github.com/olebedev/when) [![wercker status](https://app.wercker.com/status/a04ca8246bf35621b2665a73c1ed765f/s/master "wercker status")](https://app.wercker.com/project/byKey/a04ca8246bf35621b2665a73c1ed765f)

> `when` is a natural language date/time parser with pluggable rules and merge strategies

### Examples

* **tonight at 11:10 pm**
* at **Friday afternoon**
* the deadline is **next tuesday 14:00**
* drop me a line **next wednesday at 2:25 p.m**
* it could be done at **11 am past tuesday**

Check [EN](https://github.com/olebedev/when/blob/master/rules/en), [RU](https://github.com/olebedev/when/blob/master/rules/ru) and [BR](https://github.com/olebedev/when/blob/master/rules/br) rules and tests for them, for more examples.

**Needed rule not found?**
Open [an issue](https://github.com/olebedev/when/issues/new) with the case and it will be added asap.

### How it works

Usually, there are several rules added to the parser's instance for checking. Each rule has its own borders - length and offset in provided string. Meanwhile, each rule yields only the first match over the string. So, the library checks all the rules and extracts a cluster of matched rules which have distance between each other less or equal to [`options.Distance`](https://github.com/olebedev/when/blob/master/when.go#L141-L144), which is 5 by default. For example:

```
on next wednesday at 2:25 p.m.
   └──────┬─────┘    └───┬───┘
       weekday      hour + minute
```

So, we have a cluster of matched rules - `"next wednesday at 2:25 p.m."` in the string representation. 

After that, each rule is applied to the context. In order of definition or in match order, if [`options.MatchByOrder`](https://github.com/olebedev/when/blob/master/when.go#L141-L144) is set to `true`(which it is by default). Each rule could be applied with given merge strategy. By default, it's an [Override](https://github.com/olebedev/when/blob/master/rules/rules.go#L13) strategy. The other strategies are not implemented yet in the rules. **Pull requests are welcome.**

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
	text[r.Index:r.Index+len(r.Text)],
)
```

#### Distance Option

```go
w := when.New(nil)
w.Add(en.All...)
w.Add(common.All...)

text := "February 23, 2019 | 1:46pm"

// With default distance (5):
// February 23, 2019 | 1:46pm
//            └───┬───┘
//           distance: 9 (1:46pm will be ignored)

r, _ := w.Parse(text, time.Now())
fmt.Printf(r.Time.String())
// "2019-02-23 09:21:21.835182427 -0300 -03"
// 2019-02-23 (correct)
//   09:21:21 ("wrong")

// With custom distance (10):
w.SetOptions(&rules.Options{
	Distance:     10,
	MatchByOrder: true})

r, _ = w.Parse(text, time.Now())
fmt.Printf(r.Time.String())
// "2019-02-23 13:46:21.559521554 -0300 -03"
// 2019-02-23 (correct)
//   13:46:21 (correct)
```

### State of the project

The project is in a more-or-less complete state. It's used for one project already. Bugs will be fixed as soon as they will be found.

### TODO

- [ ] readme: describe all the existing rules
- [ ] implement missed rules for [these examples](https://github.com/mojombo/chronic#examples)
- [ ] add cli and simple rest api server([#2](https://github.com/olebedev/when/issues/2))

### LICENSE

http://www.apache.org/licenses/LICENSE-2.0

