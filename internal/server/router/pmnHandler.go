package router

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/dasiyes/ivmtla/internal/lnurlp"
	"github.com/go-chi/chi"
)

type pmnHandler struct {
	l *log.Logger
	// place any dependencies ...
}

func (ph *pmnHandler) router() chi.Router {
	rtr := chi.NewRouter()

	rtr.Route("/", func(r chi.Router) {
		// lnurl-pay
		r.Get("/lnurlp/{user}", ph.convert)
		// r.Post("/send", rh.home)
		// r.Get("/nostr", rh.nostr)
	})

	return rtr
}

// [ ]: return the expected LNURLp response
// `{
//
//		// The URL from LN SERVICE which will accept the pay request parameters
//		callback: String,
//
//	 // Max amount LN SERVICE is willing to receive
//		maxSendable: MilliSatoshi,
//
//	 // Min amount LN SERVICE is willing to receive, can not be less than 1 or more than `maxSendable`
//		minSendable: MilliSatoshi,
//
//	 // Metadata json which must be presented as raw string here, this is required to pass signature verification at a later step
//		metadata: String,
//
//	 // Optional number of characters accepted for the `comment` query parameter on subsequent callback, defaults to 0 if not provided. (no comment allowed)
//		commentAllowed: Number,
//
//	 // Optional lnurl-withdraw link (for explanation see justification.md)
//		withdrawLink: String,
//
//	 // Type of LNURL
//		tag: "payRequest"
//
// }`
func (ph *pmnHandler) convert(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	usr := chi.URLParam(r, "user")
	ph.l.Printf("start processing lightning payment request to user %s ...", usr)

	var rsp lnurlp.LAResponse = lnurlp.LAResponse{}

	resp, err := rsp.BuildRespponse(usr)
	if err != nil {
		respondError(w, usr, err, http.StatusInternalServerError)
	}

	// Convert the map object to a JSON object
	jsonBytes, err := json.Marshal(resp)
	if err != nil {
		respondError(w, usr, err, http.StatusInternalServerError)
	}

	// Send the JSON object back to the client
	_, _ = w.Write(jsonBytes)

}

// Send http response with error
func respondError(w http.ResponseWriter, usr string, err error, code int) {

	var ersp lnurlp.LAErrorResponse
	er, _ := ersp.BuildErrorResponse(usr, err.Error())
	jeBytes, errm := json.Marshal(er)
	if errm != nil {
		log.Printf("error marshalling error response: %v", errm)
		http.Error(w, fmt.Sprintf("error marshalling error response: %v", errm), http.StatusInternalServerError)
	}
	http.Error(w, string(jeBytes), code)
}
