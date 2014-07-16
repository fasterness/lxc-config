package lxcconfig

import "testing"

func TestConfig(t *testing.T) {
	c := New()
	str := c.String()
	if str == "" {
		t.Errorf("Test failed. string is empty")
	}
}
