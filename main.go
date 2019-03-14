package main

import (
	"encoding/json"
	"fmt"
	"go_jwt/controllers"
	"go_jwt/libs"
	_ "go_jwt/routers"
	"runtime"
	"time"

	"net"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/ulule/limiter/v3"
	"github.com/ulule/limiter/v3/drivers/store/memory"
)

type rateLimiter struct {
	generalLimiter *limiter.Limiter
	loginLimiter   *limiter.Limiter
}

func init() {
	libs.Init()
}

func main() {
	//指定使用多核，核心数为CPU的实际核心数量
	runtime.GOMAXPROCS(runtime.NumCPU())
	if beego.BConfig.RunMode == "dev" {
		orm.Debug = true
	}
	// 记录启动时间
	beego.AppConfig.Set("up_time", fmt.Sprintf("%d", time.Now().Unix()))
	beego.ErrorController(&controllers.ErrorController{})
	orm.DefaultTimeLoc = time.UTC

	r := &rateLimiter{}

	rate, err := limiter.NewRateFromFormatted("5-S")
	PanicOnError(err)
	r.generalLimiter = limiter.New(memory.NewStore(), rate)

	loginRate, err := limiter.NewRateFromFormatted("5-M")
	PanicOnError(err)
	r.loginLimiter = limiter.New(memory.NewStore(), loginRate)

	//More on Beego filters here https://beego.me/docs/mvc/controller/filter.md
	beego.InsertFilter("/*", beego.BeforeRouter, func(c *context.Context) {
		rateLimit(r, c)
	}, true)

	beego.Run()
}

func rateLimit(r *rateLimiter, ctx *context.Context) {
	var (
		limiterCtx limiter.Context
		ip         net.IP
		err        error
		req        = ctx.Request
	)

	if strings.HasPrefix(ctx.Input.URL(), "/v1/user/login") {
		ip = r.loginLimiter.GetIP(req)
		limiterCtx, err = r.loginLimiter.Get(req.Context(), ip.String())
	} else {
		ip = r.generalLimiter.GetIP(req)
		limiterCtx, err = r.generalLimiter.Get(req.Context(), ip.String())
	}
	if err != nil {
		PanicOnError(err)
	}
	h := ctx.ResponseWriter.Header()
	h.Add("X-RateLimit-Limit", strconv.FormatInt(limiterCtx.Limit, 10))
	h.Add("X-RateLimit-Remaining", strconv.FormatInt(limiterCtx.Remaining, 10))
	h.Add("X-RateLimit-Reset", strconv.FormatInt(limiterCtx.Reset, 10))

	if limiterCtx.Reached {
		fmt.Println("Too Many Requests from %s on %s", ip, ctx.Input.URL())
		ctx.Output.SetStatus(429)
		resBytes, err := json.Marshal(controllers.OutResponse(429, nil, ""))
		if err != nil {
			fmt.Println(err)
		}
		ctx.Output.Body(resBytes)
		return
	}

}

func PanicOnError(e error) {
	if e != nil {
		panic(e)
	}
}
