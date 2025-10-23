package main

import (
	"WebRemoteMediaCtrl/api"
	"embed"
)

//go:embed web
var embedF embed.FS

func main() {
	api.Server(embedF)
}
