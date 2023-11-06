package lnurlp

import (
	"log"

	"github.com/lightninglabs/lndclient"
)

var LndServices *lndclient.GrpcLndServices
var LndServicesConfig lndclient.LndServicesConfig = lndclient.LndServicesConfig{
	LndAddress:  "192.168.178.10:10009",
	Network:     "mainnet",
	TLSPath:     "assets/tls.cert",
	MacaroonDir: "assets/macs/",
}

func init() {
	grpcLndSrvcs, err := lndclient.NewLndServices(&LndServicesConfig)
	if err != nil {
		log.Printf("error initializing lnd services: %v", err)
		panic(err)
	}
	LndServices = grpcLndSrvcs
}
