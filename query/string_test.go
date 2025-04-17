package query

import (
	"fmt"
	"testing"
)

func TestStringInArrayQueryMultiToSQL(t *testing.T) {
	saq := StringInArrayQuery{
		Operator: "=",
		Value:    []string{"123", "345"},
	}
	fmt.Println(saq.QuerySQL("src_ip", ""))
}

func TestStringInArrayQueryOneToSQL(t *testing.T) {
	saq := StringInArrayQuery{
		Operator: "=",
		Value:    []string{"123"},
	}
	fmt.Println(saq.QuerySQL("src_ip", ""))
}

func TestStringInArrayMarshal(t *testing.T) {
	saq := StringInArrayQuery{
		Operator: "=",
		Value:    []string{"123"},
	}
	fmt.Println(saq.Marshal())
}

func TestStringInArrayUnmarshal(t *testing.T) {
	saq := StringInArrayQuery{
		Operator: "=",
		Value:    []string{"123", "234"},
	}
	req, _ := saq.Marshal()
	ss := StringInArrayQuery{}
	err := ss.Unmarshal(req)
	if err != nil {
		t.Error(err)
	}
	t.Logf("ss: %#v", ss)
}
