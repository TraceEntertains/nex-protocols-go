// Package protocol implements the DataStore protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	datastore_types "github.com/PretendoNetwork/nex-protocols-go/datastore/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// PreparePostObjectV1 sets the PreparePostObjectV1 handler function
func (protocol *Protocol) PreparePostObjectV1(handler func(err error, client *nex.Client, callID uint32, dataStorePreparePostParamV1 *datastore_types.DataStorePreparePostParamV1)) {
	protocol.preparePostObjectV1Handler = handler
}

func (protocol *Protocol) handlePreparePostObjectV1(packet nex.PacketInterface) {
	if protocol.preparePostObjectV1Handler == nil {
		globals.Logger.Warning("DataStore::PreparePostObjectV1 not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	dataStorePreparePostParamV1, err := parametersStream.ReadStructure(datastore_types.NewDataStorePreparePostParamV1())
	if err != nil {
		go protocol.preparePostObjectV1Handler(fmt.Errorf("Failed to read dataStorePreparePostParamV1 from parameters. %s", err.Error()), client, callID, nil)
		return
	}

	go protocol.preparePostObjectV1Handler(nil, client, callID, dataStorePreparePostParamV1.(*datastore_types.DataStorePreparePostParamV1))
}
