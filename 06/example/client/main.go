package main

import (
	"context"
	"example/kitex_gen/api"
	"example/kitex_gen/api/echo"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
	"log"
	"time"
)

// 此时client类型其实就是可以实现idl 服务（更复杂的echo）的接口类型
func main() {
	c, err := echo.NewClient("example", client.WithHostPorts("0.0.0.0:8888"))
	if err != nil {
		log.Fatal(err)
	}
	req := &api.Request{Message: "my request"}
	resp, err := c.Echo(context.Background(), req, callopt.WithRPCTimeout(3*time.Second))
	if err != nil {
		log.Fatal(err)
	}
	log.Println(resp)
}
