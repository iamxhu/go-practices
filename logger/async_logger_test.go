package logger

import "testing"

func TestLog(t *testing.T) {
	ai := []int32{1, 4, 5}
	for i := 0; i < 1000; i++ {
		if i%2 == 0 {
			Logger.Error(ai)
		}

		if i%3 == 0 {
			Logger.Info(ai)
		}

		if i%5 == 0 {
			Logger.Warn(ai)
		}

		if i%7 == 0 {
			Logger.Info(ai)
		}
	}
}
