package app

import (
	"log"
	"net/http"

	_ "github.com/joho/godotenv/autoload"
	"github.com/ynoacamino/alert-go/config"
	"github.com/ynoacamino/alert-go/internal/adapters/db"
	"github.com/ynoacamino/alert-go/internal/adapters/server"
	"github.com/ynoacamino/alert-go/internal/core/services"
	goahttp "goa.design/goa/v3/http"

	genMailAddress "github.com/ynoacamino/alert-go/gen/mail_addresses"
	genResult "github.com/ynoacamino/alert-go/gen/result"

	genhttpMailAddress "github.com/ynoacamino/alert-go/gen/http/mail_addresses/server"
	genhttpResult "github.com/ynoacamino/alert-go/gen/http/result/server"
)

func Run(cfg *config.Config) {
	sqlDB, err := db.ConnectDB(cfg.DatabaseURL)
	if err != nil {
		log.Fatal(err)
	}
	defer sqlDB.Close()

	resultRepository := db.NewResultRepository(sqlDB)
	mailAddressRepository := db.NewMailAddressRepository(sqlDB)

	resultService := services.NewResultService(resultRepository)
	mailAddressService := services.NewMailAddressService(mailAddressRepository)

	resultEndPointService := server.NewResultEndPointService(resultService)
	mailAddressEndPointService := server.NewMailAddressEndPointService(mailAddressService)

	resultEndPoints := genResult.NewEndpoints(resultEndPointService)
	mailAddressEndPoints := genMailAddress.NewEndpoints(mailAddressEndPointService)

	mux := goahttp.NewMuxer()
	requestDecoder := goahttp.RequestDecoder
	responseEncoder := goahttp.ResponseEncoder

	requestHandler := genhttpResult.New(resultEndPoints, mux, requestDecoder, responseEncoder, nil, nil)
	mailAddressHandler := genhttpMailAddress.New(mailAddressEndPoints, mux, requestDecoder, responseEncoder, nil, nil)

	genhttpResult.Mount(mux, requestHandler)
	genhttpMailAddress.Mount(mux, mailAddressHandler)

	server := &http.Server{Addr: ":" + cfg.Port, Handler: mux}

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
