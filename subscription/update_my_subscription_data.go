// Package protocol implements the Subscription protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleUpdateMySubscriptionData(packet nex.PacketInterface) {
	var err error

	if protocol.UpdateMySubscriptionData == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "SubscriptionProtocol::UpdateMySubscriptionData not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	unk := types.NewPrimitiveU32(0)
	err = unk.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.UpdateMySubscriptionData(fmt.Errorf("Failed to read unk from parameters. %s", err.Error()), packet, callID, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	// * This is done since the server doesn't need to care about the data here (it's game-specific)
	// * so we just pass it along to store however the handler wants
	rmcMessage, rmcError := protocol.UpdateMySubscriptionData(nil, packet, callID, unk, parametersStream.ReadRemaining())
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
