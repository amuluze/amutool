// Package query
// Date:   2025/4/14 09:47
// Description:
package query

import (
	"fmt"
	"testing"
)

func TestTimeRangeQueryMarshal(t *testing.T) {
	trq := TimeRangeQuery{
		Operator: "=",
		Value: TimeRange{
			From: 1743782400,
			To:   1743868800,
		},
	}

	jsonData, err := trq.Marshal()
	if err != nil {
		t.Errorf("Marshal failed: %v", err)
		return
	}

	fmt.Printf("Marshaled JSON: %s\n", string(jsonData))
}

func TestTimeRangeQueryUnmarshal(t *testing.T) {
	trq := TimeRangeQuery{
		Operator: "=",
		Value: TimeRange{
			From: 1743782400,
			To:   1743868800,
		},
	}

	jsonData, err := trq.Marshal()
	if err != nil {
		t.Errorf("Marshal failed: %v", err)
		return
	}

	unmarshaled := TimeRangeQuery{}
	err = unmarshaled.Unmarshal(jsonData)
	if err != nil {
		t.Errorf("Unmarshal failed: %v", err)
		return
	}

	if unmarshaled.Operator != trq.Operator {
		t.Errorf("Operator mismatch: got %s, want %s", unmarshaled.Operator, trq.Operator)
	}

	if unmarshaled.Value.From != trq.Value.From {
		t.Errorf("From time mismatch: got %v, want %v", unmarshaled.Value.From, trq.Value.From)
	}

	if unmarshaled.Value.To != trq.Value.To {
		t.Errorf("To time mismatch: got %v, want %v", unmarshaled.Value.To, trq.Value.To)
	}

	t.Logf("Unmarshaled: %#v", unmarshaled)
}
