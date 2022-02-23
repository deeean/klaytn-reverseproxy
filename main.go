package main

import (
	"github.com/deeean/klaytn-reverseproxy/config"
	"github.com/deeean/klaytn-reverseproxy/handler"
	"github.com/deeean/klaytn-reverseproxy/util"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

func main() {
	e := echo.New()
	e.HideBanner = true
	e.Use(
		middleware.Recover(),
		middleware.LoggerWithConfig(config.LoggerConfig),
	)

	username := util.GetEnvOrDefault("USERNAME", "root")
	if username == "" {
		e.Logger.Fatal("'username' must not be empty string")
	}

	password := util.GetEnvOrDefault("PASSWORD", "root")
	if password == "" {
		e.Logger.Fatal("'password' must not be empty string")
	}

	cypressRpcUrl, err := util.GetEnvURLOrDefault("CYPRESS_RPC_URL", "http://localhost:8551")
	if err != nil {
		e.Logger.Fatal(err)
	}

	cypressWsUrl, err := util.GetEnvURLOrDefault("CYPRESS_WS_URL", "http://localhost:8552")
	if err != nil {
		e.Logger.Fatal(err)
	}

	baobabRpcUrl, err := util.GetEnvURLOrDefault("BAOBAB_RPC_URL", "http://localhost:8551")
	if err != nil {
		e.Logger.Fatal(err)
	}

	baobabWsUrl, err := util.GetEnvURLOrDefault("BAOBAB_WS_URL", "http://localhost:8552")
	if err != nil {
		e.Logger.Fatal(err)
	}

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "ok")
	})

	v1 := e.Group("/v1")
	{
		v1.Use(middleware.BasicAuth(handler.BasicAuth(username, password)))

		rpc := v1.Group("/rpc")
		{

			cypress := rpc.Group("/cypress")
			cypress.Use(middleware.Proxy(middleware.NewRoundRobinBalancer([]*middleware.ProxyTarget{
				{
					URL: cypressRpcUrl,
				},
			})))

			baobab := rpc.Group("/baobab")
			baobab.Use(middleware.Proxy(middleware.NewRoundRobinBalancer([]*middleware.ProxyTarget{
				{
					URL: baobabRpcUrl,
				},
			})))
		}

		ws := v1.Group("/ws")
		{
			cypress := ws.Group("/cypress")
			cypress.Use(middleware.Proxy(middleware.NewRoundRobinBalancer([]*middleware.ProxyTarget{
				{
					URL: cypressWsUrl,
				},
			})))

			baobab := ws.Group("/baobab")
			baobab.Use(middleware.Proxy(middleware.NewRoundRobinBalancer([]*middleware.ProxyTarget{
				{
					URL: baobabWsUrl,
				},
			})))
		}
	}

	e.Logger.Fatal(e.Start(":3000"))
}
