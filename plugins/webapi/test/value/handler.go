package value

import (
	"net/http"

	"github.com/iotaledger/goshimmer/plugins/issuer"
	"github.com/labstack/echo"

	"github.com/iotaledger/hive.go/crypto/ed25519"
	"github.com/iotaledger/goshimmer/dapps/valuetransfers/packages/address"
	"github.com/iotaledger/goshimmer/dapps/valuetransfers/packages/address/signaturescheme"
	"github.com/iotaledger/goshimmer/dapps/valuetransfers/packages/balance"
	"github.com/iotaledger/goshimmer/dapps/valuetransfers/packages/transaction"
	valuepayload "github.com/iotaledger/goshimmer/dapps/valuetransfers/packages/payload"
)

func SendValueMsg(c echo.Context) error {
	addressKeyPair1 := ed25519.GenerateKeyPair()
	addressKeyPair2 := ed25519.GenerateKeyPair()

	p := valuepayload.New(
		valuepayload.GenesisID,
		valuepayload.GenesisID,
		transaction.New(
			transaction.NewInputs(
				transaction.NewOutputID(address.FromED25519PubKey(addressKeyPair1.PublicKey), transaction.RandomID()),
				transaction.NewOutputID(address.FromED25519PubKey(addressKeyPair2.PublicKey), transaction.RandomID()),
			),

			transaction.NewOutputs(map[address.Address][]*balance.Balance{
				address.Random(): {
					balance.New(balance.ColorIOTA, 1337),
				},
			}),
		).Sign(
			signaturescheme.ED25519(addressKeyPair1),
		),
	)

    msg, err := issuer.IssuePayload(p)
    if err != nil {
        return c.JSON(http.StatusBadRequest, Response{Error: err.Error()})
    }

    return c.JSON(http.StatusOK, Response{ID: msg.Id().String()})
}

// Response contains the ID of the message sent.
type Response struct {
	ID    string `json:"id,omitempty"`
	Error string `json:"error,omitempty"`
}
