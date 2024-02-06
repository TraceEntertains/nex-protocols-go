// Package protocol implements the DataStoreSuperMarioMaker protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
	datastore_types "github.com/PretendoNetwork/nex-protocols-go/datastore/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handlePreparePostObjectWithOwnerIDAndDataID(packet nex.PacketInterface) {
	var err error

	if protocol.PreparePostObjectWithOwnerIDAndDataID == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "DataStoreSuperMarioMaker::PreparePostObjectWithOwnerIDAndDataID not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	ownerID := types.NewPrimitiveU32(0)
	err = ownerID.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.PreparePostObjectWithOwnerIDAndDataID(fmt.Errorf("Failed to read ownerID from parameters. %s", err.Error()), packet, callID, nil, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	dataID := types.NewPrimitiveU64(0)
	err = dataID.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.PreparePostObjectWithOwnerIDAndDataID(fmt.Errorf("Failed to read dataID from parameters. %s", err.Error()), packet, callID, nil, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	param := datastore_types.NewDataStorePreparePostParam()
	err = param.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.PreparePostObjectWithOwnerIDAndDataID(fmt.Errorf("Failed to read param from parameters. %s", err.Error()), packet, callID, nil, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.PreparePostObjectWithOwnerIDAndDataID(nil, packet, callID, ownerID, dataID, param)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
