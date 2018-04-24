package rematch

import (
	"testing"

	"github.com/dannypsnl/redux/action"
)

func TestRematchReducer(t *testing.T) {
	counter := Reducer{
		State: 0,
		Reducers: Reducers{
			"INC": func(state interface{}, act action.Action) interface{} {
				return state.(int) + act.Args["payload"].(int)
			},
			"DEC": func(state interface{}, act action.Action) interface{} {
				return state.(int) - act.Args["payload"].(int)
			},
		},
	}
	testInit(t, counter)
	testAction(t, counter)
	testNotExistAction(t, counter)
}

func testInit(t *testing.T, counter Reducer) {
	if counter.State != 0 {
		t.Error("initialize bug")
	}
}
func testAction(t *testing.T, counter Reducer) {
	act := counter.Action("INC").Arg("payload", 10)
	if act.Type != "INC" || act.Args["payload"] != 10 {
		t.Error("rematch::Reducer can't generate correct action")
	}
}
func testNotExistAction(t *testing.T, counter Reducer) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("generate the action not existing in rematch::Reducer should cause panic")
		}
	}()
	counter.Action("PLUS")
}
