package polling

import (
	"github.com/Pairin/go-engine.io/transport"
)

var Creater = transport.Creater{
	Name:      "polling",
	Upgrading: false,
	Server:    NewServer,
	Client:    NewClient,
}
