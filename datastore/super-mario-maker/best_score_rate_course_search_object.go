// Package protocol implements the DataStoreSuperMarioMaker protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	datastore_types "github.com/PretendoNetwork/nex-protocols-go/v2/datastore/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

func (protocol *Protocol) handleBestScoreRateCourseSearchObject(packet nex.PacketInterface) {
	if protocol.BestScoreRateCourseSearchObject == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "DataStoreSuperMarioMaker::BestScoreRateCourseSearchObject not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	param := datastore_types.NewDataStoreSearchParam()
	var extraData types.List[types.String]

	var err error

	err = param.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.BestScoreRateCourseSearchObject(fmt.Errorf("Failed to read param from parameters. %s", err.Error()), packet, callID, param, extraData)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	err = extraData.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.BestScoreRateCourseSearchObject(fmt.Errorf("Failed to read extraData from parameters. %s", err.Error()), packet, callID, param, extraData)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.BestScoreRateCourseSearchObject(nil, packet, callID, param, extraData)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
