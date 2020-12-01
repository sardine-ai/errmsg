# errmsg
[![CircleCI](https://circleci.com/gh/daisy1754/errmsg.svg?style=svg)](https://circleci.com/gh/daisy1754/errmsg)
[![codecov](https://codecov.io/gh/daisy1754/errmsg/branch/master/graph/badge.svg?token=NSE90I0AVQ)](https://codecov.io/gh/daisy1754/errmsg)

This library wraps errors from golang's [json package](https://golang.org/pkg/encoding/json)
and famous [go-playground/validator](https://github.com/go-playground/validator) library. 
The idea is to provide error message that is more friendly for consumer of your json APIs.

For instance, instead of 
`Key: 'request.UserID' Error:Field validation for 'UserID' failed on the 'required' tag` (error message from go-playground/validator),
this library would say `'UserID' is a required field`. 

Or instead of `json: cannot unmarshal string into Go struct field birthday.birthday.month of type int8` (error message from json pkg)
 this library says `'birthday.month' should be integer but received string`.

### Usage

Simply wrap error from json package or validator library

```
err := json.Unmarshal([]byte(input), &struct)
errmsg.Message(err)
``` 

```
err := validate.Struct(r)
errmsg.Message(err)
```

