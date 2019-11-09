package events

import (
	"encoding/json"
	"fmt"

	"github.com/go-pg/pg/v9"
	"github.com/justcompile/midgard/common/dal"
)

type bus struct {
	db *pg.DB
}

var eventsBus *bus

func init() {
	eventsBus = &bus{
		db: dal.Database(),
	}
}

func send(channelName string, data interface{}) error {
	out, err := json.Marshal(data)
	if err != nil {
		return err
	}

	_, err = eventsBus.db.Exec(fmt.Sprintf("NOTIFY %s, ?", channelName), string(out))
	return err
}

// Subscribe does a thing
func Subscribe(channels ...string) *pg.Listener {
	return eventsBus.db.Listen(channels...)
}

// func (b *bus) Shutdown() error {
// 	return nil
// }
