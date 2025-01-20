package main

import (
    "testing"
)

func TestAdd(t *testing.T) {
    tests := []struct {
        a, b, want int
    }{
        {2, 3, 5},
        {4, 5, 9},
        {10, 20, 30},
    }

    for _, tt := range tests {
        got := Add(tt.a, tt.b)
        if got != tt.want {
            t.Errorf("Add(%d, %d) = %d; want %d", tt.a, tt.b, got, tt.want)
        }
    }
}
