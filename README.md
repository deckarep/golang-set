golang-set
==========

A simple set type for the Go language.

Coming from Python one of the things I miss is the superbly wonderful set collection.  This is my attempt to mimic the primary features of the set from Python.
You can of course argue that there is no need for a set in Go, otherwise the creators would have added one to the standard library.  To those I say simply ignore this repository
and to those that find this useful please contribute in helping me make it better by:

* Helping to make Go idiomatic improvements
* Helping to make it better for more generic use across types.
* Helping to increase the performance of it. (So far, no attempt has been made, but since it uses a map internally, I expect it to be mostly performant.)
* Helping to make the unit-tests more robust and kick-ass.
* Helping to add documentation.
* Simply offering feedback and suggestions, since I am a Go n00b.  (Positive, constructive feedback is appreciated.)

I have to give some credit for helping seed the idea with this post from stackoverflow: [stackoverflow.com post](http://programmers.stackexchange.com/questions/177428/sets-data-structure-in-golang)

Please see the unit test file for usage examples.  The Python set documentation will also do a better job than I can of explaining how a set typically works. [Python Set](http://docs.python.org/2/library/sets.html)    Please keep in mind 
however that the Python set is a built-in type and supports additional features and syntax that make it awesome.  This set for Go is nowhere near as comprehensive as the Python set
also, this set is not battle-tested, used in production or 

