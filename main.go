package main

import (
	middleware2 "github.com/deeean/klaytn-reverseproxy/middleware"
	"github.com/deeean/klaytn-reverseproxy/security"
	"github.com/deeean/klaytn-reverseproxy/utils"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

func main() {
	e := echo.New()
	e.HideBanner = true
	e.Use(
		middleware.CORS(),
		middleware.Recover(),
		middleware.LoggerWithConfig(middleware.LoggerConfig{
			Format: "[${time_rfc3339}] ${remote_ip} ${status} ${method} ${host}${path} ${latency_human}\n",
		}),
		middleware2.HeaderOverwriteMiddleware,
	)

	username := utils.GetEnvOrDefault("USERNAME", "root")
	if username == "" {
		e.Logger.Fatal("'username' must not be empty string")
	}

	password := utils.GetEnvOrDefault("PASSWORD", "root")
	if password == "" {
		e.Logger.Fatal("'password' must not be empty string")
	}

	cypressRpcUrl, err := utils.GetEnvURLOrDefault("CYPRESS_RPC_URL", "http://localhost:8551")
	if err != nil {
		e.Logger.Fatal(err)
	}

	cypressWsUrl, err := utils.GetEnvURLOrDefault("CYPRESS_WS_URL", "http://localhost:8552")
	if err != nil {
		e.Logger.Fatal(err)
	}

	baobabRpcUrl, err := utils.GetEnvURLOrDefault("BAOBAB_RPC_URL", "http://localhost:8551")
	if err != nil {
		e.Logger.Fatal(err)
	}

	baobabWsUrl, err := utils.GetEnvURLOrDefault("BAOBAB_WS_URL", "http://localhost:8552")
	if err != nil {
		e.Logger.Fatal(err)
	}

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "ok")
	})

	v1 := e.Group("/v1")
	{
		v1.Use(middleware.BasicAuth(security.BasicAuth(username, password)))

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
