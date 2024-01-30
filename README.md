# passwd

![version](https://img.shields.io/github/v/tag/dtreshy/passwd?include_prereleases&label=Version)
![licence](https://img.shields.io/github/license/DTreshy/passwd)
[![codecov coverage](https://codecov.io/gh/DTreshy/passwd/graph/badge.svg?token=d4LlXcn5Kx)](https://codecov.io/gh/DTreshy/passwd)
![Tests passing](https://github.com/DTreshy/passwd/actions/workflows/test.yml/badge.svg)
[![Go Reference](https://pkg.go.dev/badge/github.com/DTreshy/passwd.svg)](https://pkg.go.dev/github.com/DTreshy/passwd)
[![Go Report Card](https://goreportcard.com/badge/github.com/DTreshy/passwd)](https://goreportcard.com/report/github.com/DTreshy/passwd)
[![Maintainability](https://api.codeclimate.com/v1/badges/d1513ec368bd20fe49d3/maintainability)](https://codeclimate.com/github/DTreshy/passwd/maintainability)

## A very simple library for your passwords

### How to use?

```go
hash, err := passwd.Hash("myverystrongpassword123!")
if err != nil {
    return err
}

result := passwd.Check("myverystrongpassword123!", hash)
```
