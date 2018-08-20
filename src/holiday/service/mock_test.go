package service_test

type mockLogger struct {
	ErrorCount int
	Success    int
}

func (mocklog *mockLogger) Error(s string) bool {
	mocklog.ErrorCount++
	return true
}

func (mocklog *mockLogger) Info(s string) bool {
	mocklog.Success++
	return true
}
