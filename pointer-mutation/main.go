package main

import (
	"fmt"
)

type queryResource struct {
	resource nullString
}

func main() {
	// Not only is queryResource a pointer value, but so are all its fields.
	qr := &queryResource{}
	fmt.Printf("qr type: %T, value: %#[1]v\n", qr)

	// Expect no change because there is no return value in setValue method.
	// Note that pointers are accepted in non-pointer value method receivers.
	qr.resource.setValue("pods")
	fmt.Printf("qr.resource 'set' without pointer; qr type: %T, qr value: %#[1]v\n", qr)
	fmt.Printf("qr.resource type: %T, value: %#[1]v\n\n", qr)

	// The nullString set method mutates the queryResource.resource field value.
	qr.resource.set("nodes")
	fmt.Printf("qr.resource set using pointer; qr type: %T, qr value: %#[1]v\n", qr)
	fmt.Printf("qr.resource type: %T, value: %#[1]v\n", qr)
}
