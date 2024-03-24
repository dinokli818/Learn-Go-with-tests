package iteration

import (
	"testing"
)

func TestHttpAddress(t *testing.T) {
	t.Run("www.douyin.com", func(t *testing.T) {
		repeated := IsHttpAddress("www.douyin.com")
		expected := true
		if repeated != expected {
			t.Errorf("expected %v but got %v", expected, repeated)
		}
	})
	t.Run("www.douyin", func(t *testing.T) {
		repeated := IsHttpAddress("www.douyin")
		expected := false
		if repeated != expected {
			t.Errorf("expected %v but got %v", expected, repeated)
		}
	})
}
