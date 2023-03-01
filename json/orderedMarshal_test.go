package json

import (
	"testing"
)

func Test_MarshalJSON(t *testing.T) {
	ms := &MapSlice{}
	result := string(ms.Add("name", "jack").Add("age", "23").Marshal())
	expect := `{"name:"jack","age:"23"}`
	if result != expect {
		t.Errorf("expect %s but got: %s", expect, result)
	}
}
