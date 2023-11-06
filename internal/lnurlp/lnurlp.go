package lnurlp

import "fmt"

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
type LAResponse struct {
	Callback       string `json:"callback"`
	MaxSendable    int64  `json:"maxSendable"`
	MinSendable    int64  `json:"minSendable"`
	Metadata       string `json:"metadata"`
	CommentAllowed int64  `json:"commentAllowed"`
	WithdrawLink   string `json:"withdrawLink"`
	Tag            string `json:"tag"`
}

func (l *LAResponse) String() string {
	return fmt.Sprintf(`{
		"callback": "%s",
		"maxSendable": %d,
		"minSendable": %d,
		"metadata": "%s",
		"commentAllowed": %d,
		"withdrawLink": "%s",
		"tag": "%s"
	}`, l.Callback, l.MaxSendable, l.MinSendable, l.Metadata, l.CommentAllowed, l.WithdrawLink, l.Tag)
}

func (l *LAResponse) BuildRespponse(user string) (map[string]interface{}, error) {

	resp := make(map[string]interface{})
	// [ ]: verify if the user is valid, existing, etc.
	// if the user is not valid, return an error response
	// return nil, fmt.Errorf("user %s is not valid", user)

	al := LndServices.NodeAlias
	fmt.Printf("alias: %s\n", al)

	// [ ]: verify if the user has a valid LNURLp endpoint
	resp["callback"] = fmt.Sprintf("https://ivmanto.com/lnurlp/%s", user)

	// [ ]: get the LNURLp endpoint, min, max values etc.
	resp["maxSendable"] = 1000000
	resp["minSendable"] = 1000

	// [ ]: compose the metadata according to the LNURLp spec (see docs/design-notes)
	resp["metadata"] = "[[\"text/plain\", \"lorem ipsum blah blah\"]]"

	// [ ]: implement the commentAllowed feature
	resp["commentAllowed"] = 0

	// [ ]: implement the withdrawLink feature
	resp["withdrawLink"] = ""

	resp["tag"] = "payRequest"

	return resp, nil
}

// --------================   ERROR RESPONSE   ================--------

// {"status": "ERROR", "reason": "error details..."}
type LAErrorResponse struct {
	Status string `json:"status"`
	Reason string `json:"reason"`
}

func (le *LAErrorResponse) String() string {
	return fmt.Sprintf(`{
		"status": "%s",
		"reason": "%s"
	}`, le.Status, le.Reason)
}

func (le *LAErrorResponse) BuildErrorResponse(usr, reason string) (map[string]interface{}, error) {

	ersp := make(map[string]interface{})
	ersp["status"] = "ERROR"
	if reason == "" {
		reason = "unknown error"
	}
	ersp["reason"] = fmt.Sprintf("error while processing request for user: %s, reason: %s", usr, reason)
	return ersp, nil
}
