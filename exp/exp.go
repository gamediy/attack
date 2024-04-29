package exp

import (
	"attack/exp/source_ip/host_ip"
	"context"
)

type Exp interface {
	Exp(ctx context.Context, payload string) error
}

var (
	ExpPayLoad map[string]Exp
)

const (
	SourceIPHost = "SourceIPHost"
)

func init() {
	ExpPayLoad = make(map[string]Exp, 10)
	ExpPayLoad[SourceIPHost] = host_ip.HostIP{}
}
