package main

import (
	utils "github.com/TH04e22/go-chatroom/internel/utils"
)

func main() {
	utils.Info_log.Println("User xxx login")
	utils.Debug_log.Println("variable a = 1")
	utils.Error_log.Println("Undefined Object")
}
