package dbSync

import "time"

var (
	incrSyncReadeTimeout = time.Duration(10) * time.Minute
	incrSyncWriteTimeout = time.Duration(10) * time.Minute
)

type delayNode struct {
	t  time.Time // timestamp
	id int64     // id
}

type cmdDetail struct {
	Cmd  string
	Args [][]byte
}

func (c *cmdDetail) String() string {
	str := c.Cmd
	for _, s := range c.Args {
		str += " " + string(s)
	}
	return str
}