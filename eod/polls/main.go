package polls

import (
	"time"

	"github.com/Nv7-Github/Nv7Haven/eod/base"
	"github.com/Nv7-Github/Nv7Haven/eod/eodb"
	"github.com/bwmarrin/discordgo"
)

type Polls struct {
	*eodb.Data

	dg   *discordgo.Session
	base *base.Base
}

func NewPolls(data *eodb.Data, dg *discordgo.Session, base *base.Base) *Polls {
	p := &Polls{
		Data: data,

		dg:   dg,
		base: base,
	}
	go func() {
		for {
			time.Sleep(time.Minute * 10)
			p.CheckPollTime()
		}
	}()
	return p
}