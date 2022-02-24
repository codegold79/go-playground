package nulls

import "fmt"

// nullString is not exported so nothing can be access from this struct.
type nullString struct {
	String string
	IsSet  bool
}

// NullString is public/exported and all but one field is public/exported.
type NullString struct {
	String string
	IsSet  bool
	hidden string
}

// Set will use pointers to mutate the method receiver without having to return
// the NullString struct.
func (ns *NullString) Set(value, msg string) {
	if ns != nil {
		ns.String = value
		ns.IsSet = true
		ns.hidden = msg
	} else {
		fmt.Println("Pointer to NullString is nil, nothing done.")
	}
}
