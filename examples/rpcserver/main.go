package main

import (
	"github.com/micro-plat/hydra/examples/rpcserver/services/order"
	"github.com/micro-plat/hydra/examples/rpcserver/services/user"
	"github.com/micro-plat/hydra/hydra"
)

func main() {
	app := hydra.NewApp(
		hydra.WithPlatName("hydra-20"),
		hydra.WithSystemName("collector"),
		hydra.WithServerTypes("rpc-api"),
		hydra.WithAutoCreateConf(),
		hydra.WithDebug())

	app.Micro("/user/login", user.NewLoginHandler)
	app.Micro("/order/query", order.NewQueryHandler)
	app.Micro("/order/bind", order.NewBindHandler)
	app.Start()
}