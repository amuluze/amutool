// Package cert
// Date: 2024/08/12 23:18:43
// Author: Amu
// Description:
// Description:
package cert

import (
	"fmt"
	
	"golang.org/x/crypto/acme/autocert"
)

const (
	contactEmail = "amuluze@163.com"
	domain       = "amprobe.amuluze.com"
)

func GenerateCert() {
	// 生成证书
	certManager := autocert.Manager{
		Prompt:     autocert.AcceptTOS,
		HostPolicy: autocert.HostWhitelist(domain), //your domain here
		Cache:      autocert.DirCache("certs"),     //folder for storing certificates
		Email:      contactEmail,
	}
	fmt.Printf("dir: %#v\n", certManager.Cache)
	certManager.GetCertificate()
}
