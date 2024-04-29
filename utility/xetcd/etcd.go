package xetcd

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	clientv3 "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/concurrency"
	"time"
)

var (
	config  clientv3.Config
	Client  *clientv3.Client
	err     error
	Address []string
)

func InitEtcd(ctx context.Context) {
	Address = g.Cfg().MustGet(ctx, "etcd.address").Strings()
	fmt.Println(Address)
	config = clientv3.Config{
		Endpoints:   Address,
		DialTimeout: 5 * time.Second,
	}
	if Client, err = clientv3.New(config); err != nil {
		panic(err)
	}
}
func Watch(ctx context.Context, taskKey string, prefix bool, fc func(response clientv3.WatchResponse)) {
	var watch clientv3.WatchChan
	if prefix {
		watch = Client.Watch(ctx, taskKey, clientv3.WithPrefix())
	} else {
		watch = Client.Watch(ctx, taskKey)
	}

	for {
		select {
		case event := <-watch:
			if fc != nil {
				go fc(event)
			}
		}
	}
}

func Lock(ctx context.Context, lockKey string, fc func()) error {
	session, err := concurrency.NewSession(Client)
	// 创建一个互斥锁
	mutex := concurrency.NewMutex(session, lockKey)
	// 加锁
	err = mutex.TryLock(ctx)
	defer mutex.Unlock(ctx)
	if err != nil {
		fmt.Println(err, "locked")
		return err
	}
	if fc != nil {
		fc()
	}

	return nil

}
