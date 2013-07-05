golang-set
==========

A simple set type for the Go language.

Coming from Python one of the things I miss is the superbly wonderful set collection.  This is my attempt to mimic the primary features of the set from Python.
You can of course argue that there is no need for a set in Go, otherwise the creators would have added one to the standard library.  To those I say simply ignore this repository
and carry-on and to the rest that find this useful please contribute in helping me make it better by:

* Helping to make more idiomatic improvements to the code.
* Helping to make it better for more generic use across types.
* Helping to increase the performance of it. (So far, no attempt has been made, but since it uses a map internally, I expect it to be mostly performant.)
* Helping to make the unit-tests more robust and kick-ass.
* Helping to fill in the [documentation.](http://godoc.org/github.com/deckarep/golang-set)
* Simply offering feedback and suggestions, since I am a Go n00b.  (Positive, constructive feedback is appreciated.)

I have to give some credit for helping seed the idea with this post on [stackoverflow.](http://programmers.stackexchange.com/questions/177428/sets-data-structure-in-golang)

Please see the unit test file for additional usage examples.  The Python set documentation will also do a better job than I can of explaining how a set typically [works.](http://docs.python.org/2/library/sets.html)    Please keep in mind 
however that the Python set is a built-in type and supports additional features and syntax that make it awesome.  This set for Go is nowhere near as comprehensive as the Python set
also, this set has not been battle-tested or used in production.  Also, this set is not goroutine safe...you have been warned.

Examples but not exhaustive
```go
requiredClasses := NewSet()
requiredClasses.Add("Cooking")
requiredClasses.Add("English")
requiredClasses.Add("Math")
requiredClasses.Add("Biology")

scienceClasses := NewSet()
scienceClasses.Add("Biology")
scienceClasses.Add("Chemistry")

electiveClasses := NewSet()
electiveClasses.Add("Welding")
electiveClasses.Add("Music")
electiveClasses.Add("Automotive")

bonusClasses := NewSet()
bonusClasses.Add("Go Programming")
bonusClasses.Add("Python Programming")

//Show me all the available classes I can take
allClasses := requiredClasses.Union(scienceClasses).Union(electiveClasses).Union(bonusClasses)
fmt.Println(allClasses) //Set{English, Chemistry, Automotive, Cooking, Math, Biology, Welding, Music, Go Programming}

//Is cooking considered a science class?
fmt.Println(scienceClasses.Contains("Cooking")) //false

//Show me all classes that are not science classes, since I hate science.
fmt.Println(allClasses.Difference(scienceClasses)) //Set{English, Automotive, Cooking, Math, Welding, Music, Go Programming}

//Which science classes are also required classes?
fmt.Println(scienceClasses.Intersect(requiredClasses)) //Set{Biology}

//How many bonus classes do you offer?
fmt.Println(bonusClasses.Size()) //2
```

Thanks!

-Ralph



