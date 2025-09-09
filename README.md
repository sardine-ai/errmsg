# errmsg
[![CircleCI](https://circleci.com/gh/sardine-ai/errmsg.svg?style=svg)](https://circleci.com/gh/sardine-ai/errmsg)
[![codecov](https://codecov.io/gh/sardine-ai/errmsg/branch/master/graph/badge.svg?token=NSE90I0AVQ)](https://codecov.io/gh/sardine-ai/errmsg)

This library wraps errors from golang's [json package](https://golang.org/pkg/encoding/json)
and famous [go-playground/validator](https://github.com/go-playground/validator) library. 
The idea is to provide error message that is more friendly for consumer of your json APIs.

For instance, instead of `json: cannot unmarshal string into Go struct field birthday.birthday.month of type int8` (error message from json pkg)
 this library says `'birthday.month' should be integer but received string`.

Note for go-playground/validator one can just use translation: eg https://github.com/go-playground/validator/blob/master/_examples/translations/main.go

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

