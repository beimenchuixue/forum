package main

import "bbs/app/web"

func main() {
	webApp := web.NewApp()
	webApp.Run()
}
