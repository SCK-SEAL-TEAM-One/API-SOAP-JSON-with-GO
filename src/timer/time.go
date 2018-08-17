package timer

func CheckTime(time, timeOut int) bool {
	if time > timeOut {
		return false
	}
	return true
}
