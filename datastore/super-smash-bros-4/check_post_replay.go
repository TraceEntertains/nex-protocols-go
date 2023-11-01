// Package protocol implements the Super Smash Bros. 4 DataStore protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	datastore_super_smash_bros_4_types "github.com/PretendoNetwork/nex-protocols-go/datastore/super-smash-bros-4/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// CheckPostReplay sets the CheckPostReplay handler function
func (protocol *Protocol) CheckPostReplay(handler func(err error, packet nex.PacketInterface, callID uint32, param *datastore_super_smash_bros_4_types.DataStorePreparePostReplayParam) uint32) {
	protocol.checkPostReplayHandler = handler
}

func (protocol *Protocol) handleCheckPostReplay(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.checkPostReplayHandler == nil {
		globals.Logger.Warning("DataStoreSuperSmashBros4::CheckPostReplay not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	param, err := parametersStream.ReadStructure(datastore_super_smash_bros_4_types.NewDataStorePreparePostReplayParam())
	if err != nil {
		errorCode = protocol.checkPostReplayHandler(fmt.Errorf("Failed to read param from parameters. %s", err.Error()), packet, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.checkPostReplayHandler(nil, packet, callID, param.(*datastore_super_smash_bros_4_types.DataStorePreparePostReplayParam))
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
