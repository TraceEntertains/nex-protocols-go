// Package protocol implements the DataStoreSuperMarioMaker protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
	datastore_types "github.com/PretendoNetwork/nex-protocols-go/datastore/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleRecommendedCourseSearchObject(packet nex.PacketInterface) {
	if protocol.RecommendedCourseSearchObject == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "DataStoreSuperMarioMaker::RecommendedCourseSearchObject not implemented")

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
	extraData := types.NewList[*types.String]()
	extraData.Type = types.NewString("")

	var err error

	err = param.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.RecommendedCourseSearchObject(fmt.Errorf("Failed to read param from parameters. %s", err.Error()), packet, callID, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	err = extraData.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.RecommendedCourseSearchObject(fmt.Errorf("Failed to read extraData from parameters. %s", err.Error()), packet, callID, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.RecommendedCourseSearchObject(nil, packet, callID, param, extraData)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
