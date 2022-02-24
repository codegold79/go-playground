package main

type nullString struct {
	str   string
	isSet bool
}

func (ns *nullString) set(value string) {
	if ns != nil {
		ns.str = value
		ns.isSet = true
	}

	// Note that if ns is nil, it has no pointer address. It will do nothing to set a pointer address value.
	// As an example, the following will not do anything to the method receiver, ns:
	// ns = &NullString{
	// 	String: value,
	// 	IsSet:  true,
	// 	hidden: msg,
	// }
}

func (ns nullString) setValue(value string) {
	ns.str = value
	ns.isSet = true
}
