package drng

import (
	"net/http"

	drngheader "github.com/iotaledger/goshimmer/packages/binary/drng/payload/header"
	drngpayload "github.com/iotaledger/goshimmer/packages/binary/drng/payload"
	cbpayload "github.com/iotaledger/goshimmer/packages/binary/drng/subtypes/collectiveBeacon/payload"
	"github.com/iotaledger/goshimmer/plugins/issuer"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
)

func SendDrngMsg(c echo.Context) error {
	var request Request
	if err := c.Bind(&request); err != nil {
		log.Info(err.Error())
		return c.JSON(http.StatusBadRequest, Response{Error: err.Error()})
	}

	switch request.Cmd {
	case "data":
        header := drngheader.New(0, 0)
        data := []byte("test")
        p := drngpayload.New(header, data)
        msg, err := issuer.IssuePayload(p)

        if err != nil {
            return c.JSON(http.StatusBadRequest, Response{Error: err.Error()})
        }
        return c.JSON(http.StatusOK, Response{ID: msg.Id().String()})

	case "cb":
        header := drngheader.New(drngheader.TypeCollectiveBeacon, 1)
        p := cbpayload.New(header.InstanceID,
            0,
            []byte("AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA"), // prevSignature
            []byte("BBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBB"), // signature
            []byte("CCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCC"))                                                 // distributed PK
        msg, err := issuer.IssuePayload(p)

        if err != nil {
            return c.JSON(http.StatusBadRequest, Response{Error: err.Error()})
        }
        return c.JSON(http.StatusOK, Response{ID: msg.Id().String()})
    }
    return nil
}

// Response contains the ID of the message sent.
type Response struct {
	ID    string `json:"id,omitempty"`
	Error string `json:"error,omitempty"`
}

// Request contains the parameters of a spammer request.
type Request struct {
	Cmd string `json:"cmd"`
}
