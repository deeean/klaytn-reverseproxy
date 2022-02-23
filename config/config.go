package config

import "github.com/labstack/echo/v4/middleware"

var LoggerConfig = middleware.LoggerConfig{
	Format: "[${time_rfc3339}] ${remote_ip} ${status} ${method} ${host}${path} ${latency_human}\n",
}
