package main

import (
	"flag"
	"fmt"
	"net/url"
	"strconv"
	"strings"

	v1 "github.com/ivanlsz/go-ssh/v1"
)

func main() {
	var (
		// 请求链接地址
		url string
		// ssh的私有key
		key string
		// ssh的私有证书的密码
		password string
	)
	flag.StringVar(&url, "h", "root@127.0.0.1:22", "远程服务器的地址信息")
	flag.StringVar(&key, "k", "", "私钥")
	flag.StringVar(&password, "p", "", "私钥的密码")
	flag.Parse()

	err := connect(url, key, password)
	if err != nil {
		fmt.Printf("Failed to connect - %s\n", err)
	}
}

func connect(s, key, password string) error {
	var (
		host string
		port int
	)
	if !strings.HasPrefix(s, "tcp://") {
		s = "tcp://" + s
	}
	u, err := url.Parse(s)
	if err != nil {
		return err
	}
	h := strings.Split(u.Host, ":")
	if len(h) > 1 {
		host = h[0]
		port, err = strconv.Atoi(h[1])
		if err != nil {
			return err
		}
	} else {
		host = u.Host
		port = 22
	}
	auth := &v1.Auth{}
	if key != "" {
		auth.Keys = append(auth.Keys, key)
		auth.KeyPasswords = append(auth.KeyPasswords, password)
	} else {
		password, _ = u.User.Password()
		auth.Passwords = append(auth.Passwords, password)
	}
	client, err := v1.NewNativeClient(u.User.Username(), host, "SSH-2.0-IvanLamClient-1.0", port, auth, nil)
	if err != nil {
		return err
	}
	err = client.Shell()
	if err != nil && err.Error() != "exit status 255" {
		return fmt.Errorf("Failed to request shell - %s", err)
	}

	return nil
}
