// Package protocol implements the Subscriber protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleDeleteContent(packet nex.PacketInterface) {
	var err error

	if protocol.DeleteContent == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "Subscriber::DeleteContent not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	unknown1 := types.NewList[*types.String]()
	unknown1.Type = types.NewString("")
	err = unknown1.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.DeleteContent(fmt.Errorf("Failed to read unknown1 from parameters. %s", err.Error()), packet, callID, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	unknown2 := types.NewPrimitiveU64(0)
	err = unknown2.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.DeleteContent(fmt.Errorf("Failed to read unknown2 from parameters. %s", err.Error()), packet, callID, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.DeleteContent(nil, packet, callID, unknown1, unknown2)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
