package main

import (
	"macedonia"
	"net/http"
)

func init() {
	http.Handle("/", macedonia.Build())
}
