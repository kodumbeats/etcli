package main

import (
	"context"
	// "fmt"
	// "log"
	"os"
	// "strconv"
	"time"

	"github.com/coreos/etcd/clientv3"
	cli "github.com/urfave/cli/v2"
)

var (
	dialTimeout    = 2 * time.Second
	requestTimeout = 10 * time.Second
)

func main() {
	ctx, _ := context.WithTimeout(context.Background(), requestTimeout)
	client, _ := clientv3.New(clientv3.Config{
		DialTimeout: dialTimeout,
		Endpoints:   []string{"127.0.0.1:2379"},
	})
	defer client.Close()
	kv := clientv3.NewKV(client)
	if ctx == nil {
	}
	if kv == nil {
	}

	(&cli.App{}).Run(os.Args)
}
