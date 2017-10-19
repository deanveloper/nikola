# nikola
A Golang Testing Library which integrates with standard Golang tests. (Requires Go-1.9)

Go 1.9 added the `t.Helper()` function, which allows functions that are not the `TestXx` function call 
test-specific functions like `t.Fail()`, but still have the error occur in the `TestXx` function, rather in
the helper function you wrote. This is extremely useful, and enables us to make testing libraries such as this one!
