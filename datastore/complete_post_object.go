package datastore

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// CompletePostObject sets the CompletePostObject handler function
func (protocol *DataStoreProtocol) CompletePostObject(handler func(err error, client *nex.Client, callID uint32, dataStoreCompletePostParam *DataStoreCompletePostParam)) {
	protocol.CompletePostObjectHandler = handler
}

func (protocol *DataStoreProtocol) HandleCompletePostObject(packet nex.PacketInterface) {
	if protocol.CompletePostObjectHandler == nil {
		globals.Logger.Warning("DataStore::CompletePostObject not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	dataStoreCompletePostParam, err := parametersStream.ReadStructure(NewDataStoreCompletePostParam())
	if err != nil {
		go protocol.CompletePostObjectHandler(err, client, callID, nil)
		return
	}

	go protocol.CompletePostObjectHandler(nil, client, callID, dataStoreCompletePostParam.(*DataStoreCompletePostParam))
}
