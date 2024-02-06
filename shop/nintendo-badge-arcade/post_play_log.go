// Package protocol implements the Nintendo Badge Arcade Shop protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	shop_nintendo_badge_arcade_types "github.com/PretendoNetwork/nex-protocols-go/shop/nintendo-badge-arcade/types"
)

func (protocol *Protocol) handlePostPlayLog(packet nex.PacketInterface) {
	var err error

	if protocol.PostPlayLog == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "ShopNintendoBadgeArcade::PostPlayLog not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	param := shop_nintendo_badge_arcade_types.NewShopPostPlayLogParam()
	err = param.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.PostPlayLog(fmt.Errorf("Failed to read param from parameters. %s", err.Error()), packet, callID, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.PostPlayLog(nil, packet, callID, param)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
