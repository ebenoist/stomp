package stomp

import (
	"github.com/jjeffery/stomp/message"
	. "launchpad.net/gocheck"
	"math"
	"time"
)

type UtilSuite struct {}

var _ = Suite(&UtilSuite{})

func (s *UtilSuite) TestAsMilliseconds(c *C) {
	d := time.Duration(30) * time.Millisecond
	c.Check(asMilliseconds(d, math.MaxInt32), Equals, 30)
	
	// approximately one year
	d = time.Duration(365) * time.Duration(24) * time.Hour
	c.Check(asMilliseconds(d, math.MaxInt32), Equals, math.MaxInt32)
	
	d = time.Duration(365) * time.Duration(24) * time.Hour
	c.Check(asMilliseconds(d, message.MaxHeartBeat), Equals, message.MaxHeartBeat)
}