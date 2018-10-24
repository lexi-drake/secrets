# Secrets #
[![CircleCI](https://circleci.com/gh/lexi-drake/secrets.svg?style=svg)](https://circleci.com/gh/lexi-drake/secrets)

A tool for handling secrets in go.

## Usage ##

Get the library

```bash
go get github.com/lexi-drake/secrets
```

Load secrets

```go
import (
       secrets "github.com/lexi-drake/secrets"
)

type YourSecrets struct {
     Foo string `json:"foo"`
     Bar string `json:"bar"`
}

func (yourSecrets YourSecrets) Validate() bool {
     return (yourSecrets.Foo != "" &&
     	     yourSecrets.Bar != "")
}

...

var s YourSecrets
err := secrets.LoadFromJson(&s, "path/to/secrets.json")
if err != nil {
   // handle error
}

secrets.LoadFromEnvironment(&s)
if ok := s.Validate(); !ok {
   // handle invalid secrets
}
```