// Package uuid
// Date: 2022/9/7 09:49
// Author: Amu
// Description:
package uuid

import "github.com/google/uuid"

// UUID Define alias
type UUID = uuid.UUID

// NewUUID Create uuid
func NewUUID() (UUID, error) {
	return uuid.NewRandom()
}

// MustUUID Create uuid(Throw panic if something goes wrong)
func MustUUID() UUID {
	v, err := NewUUID()
	if err != nil {
		panic(err)
	}
	return v
}

// MustString Create uuid
func MustString() string {
	return MustUUID().String()
}

// MustParseUUIString convert uuid str to uuid
func MustParseUUIString(uuidStr string) UUID {
	return uuid.MustParse(uuidStr)
}
