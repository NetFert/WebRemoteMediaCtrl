package api

import (
	"WebRemoteMediaCtrl/api/handler"
	"WebRemoteMediaCtrl/config"
	"embed"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Server(embedF embed.FS) {
	printNetworkInterfaces()

	gin.SetMode(config.Mode)
	r := gin.Default()

	r.GET("/", func(c *gin.Context) { c.FileFromFS("web/", http.FS(embedF)) })
	r.GET("/qrcode.min.js", func(c *gin.Context) { c.FileFromFS("web/qrcode.min.js", http.FS(embedF)) })
	r.GET("/media", handler.Media)

	if err := r.Run(":" + config.Port); err != nil {
		log.Panicln("Fail to run:", err)
	}
}

func printNetworkInterfaces() {
	ifaces, err := net.Interfaces()
	if err != nil {
		log.Fatal(err)
	}

	for _, iface := range ifaces {
		if iface.Flags&net.FlagUp == 0 || iface.Flags&net.FlagLoopback != 0 {
			continue
		}

		addrs, err := iface.Addrs()
		if err != nil {
			continue
		}

		for _, addr := range addrs {
			ipnet, ok := addr.(*net.IPNet)
			if !ok || ipnet == nil || ipnet.IP == nil {
				continue
			}

			ip := ipnet.IP
			if ip.IsLinkLocalUnicast() {
				continue
			}

			if ip.To4() != nil {
				fmt.Printf("IPv4：http://%s:%s\n", ip.String(), config.Port)
			} else {
				fmt.Printf("IPv6：http://[%s]:%s\n", ip.String(), config.Port)
			}
		}
	}

	fmt.Printf("\nLoop：http://127.0.0.1:%s\n", config.Port)
	fmt.Printf("Loop (IPv6)：http://[::1]:%s\n", config.Port)
}
