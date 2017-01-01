# goluhn

[![Build Status](https://travis-ci.org/abadojack/goluhn.svg?branch=master)](https://travis-ci.org/abadojack/goluhn) [![GoDoc](https://godoc.org/github.com/abado    jack/goluhn?status.png)](http://godoc.org/github.com/abadojack/goluhn)


Validates credit card numbers using Luhn's algorithm and regular expressions.

### Installation

```sh
go get -u github.com/abadojack/goluhn
```

### Usage
```go
import "github.com/abadojack/goluhn"
```

### Validate credit/debit card number

```go
goluhn.IsValid("5500005555555559")	//returns true
```

### Validate credit/debit card number and return type

```go
goluhn.CardType("5500005555555559")	//returns mastercard
```

