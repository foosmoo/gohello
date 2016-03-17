package util

import "testing"

func TestAsUpper(t *testing.T) {
    cases := []struct {
        in, want string
    }{
        {"Hello, world", "HELLO, WORLD"},
        {"HELLO, WORLD", "HELLO, WORLD"},
        {"", ""},
    }
    for _, c := range cases {
        got := AsUpper(c.in)
        if got != c.want {
            t.Errorf("AsUpper(%q) == %q, want %q", c.in, got, c.want)
        }
    }
}
