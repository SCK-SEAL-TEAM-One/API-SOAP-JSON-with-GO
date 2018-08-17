package timer

import "testing"

func Test_CheckTime_Input_Time_and_TimeOut_Should_Be_True(t *testing.T) {
	time := 20
	timeOut := 30
	expected := true

	actualResult := CheckTime(time, timeOut)

	if expected != actualResult {
		t.Errorf("Expect %v but got it %v", expected, actualResult)
	}
}

func Test_CheckTime_Input_Time_and_TimeOut_Should_Be_False(t *testing.T) {
	time := 31
	timeOut := 30
	expected := false

	actualResult := CheckTime(time, timeOut)

	if expected != actualResult {
		t.Errorf("Expect %v but got it %v", expected, actualResult)
	}
}
