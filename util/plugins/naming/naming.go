package naming

import (
	"context"
	"fmt"
	"io"
	"net"
	"net/http"

	clientv3 "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/naming/endpoints"

	"github.com/PsychologicalExperiment/backEnd/util/plugins/config"
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
	addr := fmt.Sprintf("%s:%d", ip, config.Config.Server.Port)
	key := fmt.Sprintf("%s/%s/%s", namespace, config.Config.Server.ServerName, addr)
	em, err := endpoints.NewManager(cli, namespace)
	if err != nil {
		return err
	}
	if err := em.AddEndpoint(context.Background(), key, endpoints.Endpoint{Addr: addr}, clientv3.WithLease(lease.ID)); err != nil {
		return err
	}
	keepAlive, err := cli.KeepAlive(context.Background(), lease.ID)
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
	if err != nil {
		return "", err
	}
	return string(b), nil
}
