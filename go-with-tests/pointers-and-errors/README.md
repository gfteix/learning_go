Pointers:

- Go copies values when you pass them to functions/methods, so if you're writing a function that needs to mutate state you'll need it to take a pointer to the thing you want to change.

- The fact that Go takes a copy of values is useful a lot of the time but sometimes you won't want your system to make a copy of something, in which case you need to pass a reference. Examples include referencing very large data structures or things where only one instance is necessary (like database connection pools).


Nil:

- Pointers can be nil

- When a function returns a pointer to something, you need to make sure you check if it's nil or you might raise a runtime exception - the compiler won't help you here.

- Useful for when you want to describe a value that could be missing


Errors:

- Errors are the way to signify failure when calling a function/method.

- By listening to our tests we concluded that checking for a string in an error would result in a flaky test. So we refactored our implementation to use a meaningful value instead and this resulted in easier to test code and concluded this would be easier for users of our API too.

- This is not the end of the story with error handling, you can do more sophisticated things but this is just an intro. Later sections will cover more strategies.

Create new types from existing ones

- Useful for adding more domain specific meaning to values

- Can let you implement interfaces



----

Next steps: Read https://dave.cheney.net/2016/04/27/dont-just-check-errors-handle-them-gracefully


Don't check errors handle them gracefully

- Never inspect the output of error.Error (this advice doesn’t apply to writing tests)

The contents of that string belong in a log file, or displayed on screen. You shouldn’t try to change the behaviour of your program by inspecting it. Comparing the string form of an error is, in my opinion, a code smell, and you should try to avoid it.

-> avoid sentinel errors