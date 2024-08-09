// file: calculator_test.go
package test

import (
    "testing"
)

func Add(a int, b int) int {
    return a + b
}

func TestAdd(t *testing.T) {
    tests := []struct {
        a, b     int
        expected int
    }{
        {1, 1, 2},
        {2, 3, 5},
        {-1, -1, -2},
        {0, 0, 0},
    }

    for _, test := range tests {
        result := Add(test.a, test.b)
        if result != test.expected {
            t.Errorf("Add(%d, %d) = %d; expected %d", test.a, test.b, result, test.expected)
        }
    }
}
