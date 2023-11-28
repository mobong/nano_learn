package src

import (
	"fmt"
	"github.com/lonng/nano"
	"github.com/lonng/nano/component"
	"github.com/lonng/nano/serialize/protobuf"
	"github.com/spf13/viper"
	"nano_learn/db"
	"nano_learn/src/logic"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

// 初始化游戏服务器
func StartUp() {
	// start mysql
	closer := db.DbStartUp()
	defer closer()
	comps := &component.Components{}
	comps.Register(logic.NewManager())

	addr := fmt.Sprintf(":%d", viper.GetInt("logic.port"))
	nano.Listen(addr,
		nano.WithIsWebsocket(true),
		nano.WithSerializer(protobuf.NewSerializer()),
		nano.WithComponents(comps),
		nano.WithCheckOriginFunc(func(_ *http.Request) bool { return true }),
	)
	sg := make(chan os.Signal)
	signal.Notify(sg, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGKILL)
	// stop server
	select {
	case s := <-sg:
		fmt.Printf("stop server!!! :%s", s.String())
	}
}
