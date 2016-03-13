package main

import (
	"github.com/oxfeeefeee/appgo"
)

// This is an example of how to expose REST API:

// 1. Define a struct as this with any name:
type Hello struct {
	// 2. Define a attribute named "META" of type "struct{}", any type should do, but
	// this way we know it stores nothing,
	// 3. Add Tag for META specifying the path of this API
	META struct{} `path:"/hello"`
}

// 4. Define the input type, if this API needs authorization, define an attribute
//    named "UserId__" of type appgo.Id
//    You can control how the HTTP perameters are paresed by added tags.
type helloInput struct {
	UserId__ appgo.Id
	Echo     string `schema:"echo"`
}

// 5. Define Methods
// 	  a. GET POST PUT DELETE are supported
//    b. It has to have one argument, use appgo.DummyInput if needed
//    c. It can have one or two return values, and the last one should always be
//       of *appgo.ApiError. And if there are two, the first one is the returned
//       value to the API caller.

func (h Hello) GET(in *helloInput) (string, *appgo.ApiError) {
	return "Hello " + in.UserId__.String() + ", " + in.Echo, nil //appgo.NewApiErrWithCode(1000)
}
