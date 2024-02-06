// Package protocol implements the Friends 3DS protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	friends_3ds_types "github.com/PretendoNetwork/nex-protocols-go/friends-3ds/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleUpdateMiiList(packet nex.PacketInterface) {
	var err error

	if protocol.UpdateMiiList == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "Friends3DS::UpdateMiiList not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	miiList := friends_3ds_types.NewMiiList()
	err = miiList.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.UpdateMiiList(fmt.Errorf("Failed to read miiList from parameters. %s", err.Error()), packet, callID, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.UpdateMiiList(nil, packet, callID, miiList)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
