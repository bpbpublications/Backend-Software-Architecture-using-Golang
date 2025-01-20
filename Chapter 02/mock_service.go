// Run this to generate mock_test_gen.go (sample included)
// `mockgen -source=mock_service.go -destination=mock_test_gen.go  -package=main . MyService`
// The output mock is used in mock_test.go for actual testing.

package main

type MyService interface {
    DoSomething() (int, error)
}

type MockService struct{}

func (ms *MockService) DoSomething() (int, error) {
    return 45, nil
}

func main() {
}

