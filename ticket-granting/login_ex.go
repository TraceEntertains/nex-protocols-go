// Package protocol implements the Ticket Granting protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleLoginEx(packet nex.PacketInterface) {
	var err error

	if protocol.LoginEx == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "TicketGranting::LoginEx not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	strUserName := types.NewString("")
	err = strUserName.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.LoginEx(fmt.Errorf("Failed to read strUserName from parameters. %s", err.Error()), packet, callID, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	oExtraData := types.NewAnyDataHolder()
	err = oExtraData.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.LoginEx(fmt.Errorf("Failed to read oExtraData from parameters. %s", err.Error()), packet, callID, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.LoginEx(nil, packet, callID, strUserName, oExtraData)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
