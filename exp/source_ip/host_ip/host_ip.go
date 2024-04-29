package host_ip

import (
	"context"
	"fmt"
)

type HostIP struct {
}

func (ip HostIP) Exp(ctx context.Context, playload string) error {
	fmt.Println("playlod host ip", playload)
	return nil
}
