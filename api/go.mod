module phl-skate-sharpening-api/api

go 1.19

replace phl-skate-sharpening-api/utils => ../utils

replace phl-skate-sharpening-api/config => ../config

replace phl-skate-sharpening-api/core => ../core

replace phl-skate-sharpening-api/apibuilder => ../apibuilder

replace phl-skate-sharpening-api/controllers => ../controllers

require (
	github.com/Tomasen/realip v0.0.0-20180522021738-f0c99a92ddce
	github.com/go-sql-driver/mysql v1.7.0
	github.com/gorilla/context v1.1.1
	github.com/julienschmidt/httprouter v1.3.0
	github.com/justinas/alice v1.2.0
	github.com/rs/cors v1.8.3
	github.com/sirupsen/logrus v1.9.0
	gopkg.in/guregu/null.v3 v3.5.0
	phl-skate-sharpening-api/apibuilder v0.0.0-00010101000000-000000000000
	phl-skate-sharpening-api/config v0.0.0-00010101000000-000000000000
	phl-skate-sharpening-api/controllers v0.0.0-00010101000000-000000000000
	phl-skate-sharpening-api/core v0.0.0-00010101000000-000000000000
	phl-skate-sharpening-api/utils v0.0.0-00010101000000-000000000000
)
