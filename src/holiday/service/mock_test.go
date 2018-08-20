package service_test

type mockLogger struct {
	Error   int
	Success int
}

func (mocklog *mockLogger) LogError(s string) bool {
	mocklog.Error++
	return true
}

func (mocklog *mockLogger) LogInfo(s string) bool {
	mocklog.Success++
	return true
}
