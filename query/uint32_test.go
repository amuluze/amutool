package query

import (
	"fmt"
	"testing"
)

func TestUInt32InArrayQueryMultiToSQL(t *testing.T) {
	saq := UInt32InArrayQuery{
		Operator: "=",
		Value:    []uint32{123, 32},
	}
	fmt.Println(saq.QuerySQL("src_port", ""))
}

func TestUInt32InArrayQueryOneToSQL(t *testing.T) {
	saq := UInt32InArrayQuery{
		Operator: "!=",
		Value:    []uint32{123},
	}
	fmt.Println(saq.QuerySQL("src_port", ""))
}

func TestUInt32InArrayMarshal(t *testing.T) {
	saq := UInt32InArrayQuery{
		Operator: "=",
		Value:    []uint32{23},
	}
	fmt.Println(saq.Marshal())
}

func TestUInt32InArrayUnmarshal(t *testing.T) {
	saq := UInt32InArrayQuery{
		Operator: "=",
		Value:    []uint32{213, 232},
	}
	req, _ := saq.Marshal()
	ss := UInt32InArrayQuery{}
	err := ss.Unmarshal(req)
	if err != nil {
		t.Error(err)
	}
	t.Logf("ss: %#v", ss)
}
