# Contributing to equation-solver

If you would like to contribute to this project here is a list of steps I would like you to follow.

## Want to add a new feature or make a fix?
### Your branch name
- When you create a new branch, its name should be `equation-solver-<author>`.
- In case of a patch, its name should be `equation-solver-patch-<author>`.

### How commits should be?
- Preferrebly frequents. Or at least avoid making huge modifications without some commits in the middle.
- Each commit in general should refer to a modification in a single file, unless that modification is really small (changing variable name, adding a comment, ...).

### Where your functions should be located?
- Use `equation.go` if the functions have to be accessible by the user.
- Use `internal_functions.go` if not.

### How functions should be made?
- Functions should have a short description of what they do, use the comments on top their names:
```go
// A simple description of what this function should do
//
// Eventually you could write an example of input and output
func exampleFunctions() {
    return
} 
```
- Functions should be tested.

### How to create a pull request?
- When create a pull request write a small description of what changes, or features, you have made.

### Extra
- If you speak another language, like spanish, french, german or other, consider making a `readme.es.md` or `readme.fr.md` file.

Thank you!! :)