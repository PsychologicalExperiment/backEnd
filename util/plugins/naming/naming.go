package naming

import (
	"context"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/naming/endpoints"
	"io"
	"net"
	"net/http"

	"github.com/PsychologicalExperiment/backEnd/util/plugins/config"
	"github.com/PsychologicalExperiment/backEnd/util/plugins/log"
)

const (
	namespace = "psychology"
)

func init() {
	EtcdRegisterServer()
}

func EtcdRegisterServer() error {
	cli, err := clientv3.NewFromURL(fmt.Sprintf("http://%s:%d", config.Config.NamingServer.IP, config.Config.NamingServer.Port))
	if err != nil {
		return err
	}
	lease, err := cli.Grant(context.Background(), 10)
	if err != nil {
		return err
	}
	ip, err := GetExternalIP()
	if err != nil {
		return err
	}
	key := fmt.Sprintf("%s/%s/%s", namespace, config.Config.Server.ServerName, addr)
	addr := fmt.Sprintf("%s:%d", ip, config.Config.Server.Port)
	em, err := endpoints.NewManager(cli, namspace)
	if err != nil {
		return err
	}
	if err := em.AddEndpoint(ctx, key, endpoints.Endpoint{Addr: addr}, clientv3.WithLease(lease.ID)); err != nil {
		return err
	}
	keepAlive, err := cli.KeepAlive(ctx, lease.ID)
	//ctx, cancel := context.WithCancel(ctx)
	if err != nil || keepAlive == nil {
		return err
	}
	donec := make(chan struct{})
	go func() {
		defer close(donec)
		for range keepAlive {

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

// GetExternalIP 获取外网ip
func GetExternalIP() (string, error) {
	resp, err := http.Get("http://myexternalip.com/raw")
	if err != nil {
		return "", err
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			fmt.Printf("close err: %v", err)
		}
	}()
	b, err := io.ReadAll(resp.Body)
	return string(b), nil
}
