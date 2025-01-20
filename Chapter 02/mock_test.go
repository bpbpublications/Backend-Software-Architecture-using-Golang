package main

import (
    "testing"
    "github.com/golang/mock/gomock"
)

func TestMyFunction(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockService := NewMockMyService(ctrl)
    expected := 44
    mockService.EXPECT().DoSomething().Return(expected, nil)

    result, err := mockService.DoSomething()
    if err != nil {
        t.Fatalf("Unexpected error: %v", err)
    }
    if result != expected {
        t.Errorf("Expected %d, but got %d", expected, result)
    }
}
