package api

import (
	"WebRemoteMediaCtrl/api/handler"
	"WebRemoteMediaCtrl/config"
	"fmt"
	"log"
	"net"

	"github.com/gin-gonic/gin"
)

func Server() {
	printNetworkInterfaces()

	gin.SetMode(config.Mode)
	r := gin.Default()

	r.GET("/", func(c *gin.Context) { c.File("web/index.html") })
	r.GET("/qrcode.min.js", func(c *gin.Context) { c.File("web/qrcode.min.js") })
	r.GET("/media", handler.Media)

	if err := r.Run(":" + config.Port); err != nil {
		log.Panicln("Fail to run:", err)
	}
}

func printNetworkInterfaces() {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		log.Fatal(err)
	}

	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok {
			if ipnet.IP.To4() != nil {
				if ipnet.IP.IsLoopback() {
					fmt.Printf("Loopback：http://%s:%s\n", ipnet.IP.String(), config.Port)
				} else {
					fmt.Printf("Address：http://%s:%s\n", ipnet.IP.String(), config.Port)
				}
			}
		}
	}
}
