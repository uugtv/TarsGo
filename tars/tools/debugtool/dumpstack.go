package main

import (
	"fmt"
	"github.com/uugtv/TarsGo/tars"
	"github.com/uugtv/TarsGo/tars/protocol/res/adminf"
)

func main() {
	comm := tars.NewCommunicator()
	obj := fmt.Sprintf("_APP_._SERVER_._SERVANT_Obj@tcp -h 127.0.0.1 -p 10014 -t 60000")
	app := new(adminf.AdminF)
	comm.StringToProxy(obj, app)
	ret, err := app.Notify("tars.dumpstack")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(ret)
}
