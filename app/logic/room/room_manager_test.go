// Author: sheppard(ysf1026@gmail.com) 2014-01-08

package room

import (
	"testing"
)

func TestParseConfig(t *testing.T) {
	if config.Games[0].Name != "Five Chess" {
		t.Fail()
	}
}
