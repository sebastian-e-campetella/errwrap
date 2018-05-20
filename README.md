# errwrap
Error wrapper package for basic errors

# import package

```
import "errwrap"
```

# Use Catch methods

Catch method  receive two parameters. First parameter is a result and second is an error.
For example the method http.Get() return one result and error. You can use the output inside to "Catch" call.

```
var ew errwrap.ErrorWrapper

var result = ew.Catch(http.Get("http://www.example.com"))
```

The Catch method save params into ew.Error and ew.Any
To use "ew.Any" you must parse property by the correct format.

E.g:
```
  if ew.Any != nil {
    SomeMethod(ew.Any.(*http.Response)).StatusCode)
  }
```


