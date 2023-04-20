// Package uuidx
// Date: 2022/9/7 09:49
// Author: Amu
// Description:
package uuidx

import (
	"time"

	"github.com/google/uuid"
	"github.com/sony/sonyflake"
)

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

// MustParseUUIToString convertx uuid str to uuid
func MustParseUUIToString(uuidStr string) UUID {
	return uuid.MustParse(uuidStr)
}

// SnowID 雪花id
func SnowID() uint64 {
	sf := sonyflake.NewSonyflake(sonyflake.Settings{
		StartTime: time.Date(2021, 7, 28, 0, 0, 0, 0, time.UTC),
	})
	id, err := sf.NextID()
	if err == nil {
		return id
	}

	sleep := 1
	for {
		time.Sleep(time.Duration(sleep) * time.Millisecond)
		id, err := sf.NextID()
		if err == nil {
			return id
		}
		sleep *= 2
	}
}
