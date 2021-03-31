package intersystems

import (
	"github.com/caretdev/go-irisnative/src/connection"
)

type InterSystems struct {
	addr      string
	namespace string
	user      string
	password  string
}

func NewInterSystems(addr, namespace, user, password string) InterSystems {
	return InterSystems{addr, namespace, user, password}
}

func (i InterSystems) Connect() (connection.Connection, error) {
	return connection.Connect(i.addr, i.namespace, i.user, i.password)
}

func (i InterSystems) ConnectSYS() (connection.Connection, error) {
	return connection.Connect(i.addr, "%SYS", i.user, i.password)
}