package naming

import (
	"context"
	"fmt"
	"net"

	"github.com/PsychologicalExperiment/backEnd/util/plugins/config"
	clientv3 "go.etcd.io/etcd/client/v3"
)

func EtcdRegisterServer(ctx context.Context, server, addr string, ttl int64) error {
	cli, err := clientv3.NewFromURL(fmt.Sprintf("http://%s:%d", config.Config.NamingServer.IP, config.Config.NamingServer.Port))
	if err != nil {
		return err
	}
	lease, err := cli.Grant(ctx, ttl)
	if err != nil {
		return err
	}
	ip, err := GetLocalIP()
	if err != nil {
		return err
	}
	key := fmt.Sprintf("/%s/%s", config.Config.Server.ServerName,
		fmt.Sprintf("%s:%d", ip, config.Config.Server.Port))
	_, err = cli.Put(ctx, key, addr, clientv3.WithLease(lease.ID))
	if err != nil {
		return err
	}
	keepAlive, err := cli.KeepAlive(ctx, lease.ID)
	if err != nil {
		return err
	}
	go func() {
		for {
			<-keepAlive
		}
	}()
	return nil
}

// GetLocalIP 获取本地ip
func GetLocalIP() (string, error) {
	addrs, err := net.InterfaceAddrs()

	if err != nil {
		return "", err
	}

	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String(), nil
			}
		}
	}
	return "", fmt.Errorf("get ip error")
}
