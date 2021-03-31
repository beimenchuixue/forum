package main

import "forum/app/web"

func main() {
	webApp := web.NewApp()
	webApp.Run()
}
