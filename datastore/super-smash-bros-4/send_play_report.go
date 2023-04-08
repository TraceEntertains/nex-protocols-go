package datastore_super_smash_bros_4

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// SendPlayReport sets the SendPlayReport handler function
func (protocol *DataStoreSuperSmashBros4Protocol) SendPlayReport(handler func(err error, client *nex.Client, callID uint32, playReport []int32)) {
	protocol.SendPlayReportHandler = handler
}

func (protocol *DataStoreSuperSmashBros4Protocol) HandleSendPlayReport(packet nex.PacketInterface) {
	if protocol.SendPlayReportHandler == nil {
		globals.Logger.Warning("DataStoreSmash4::SendPlayReport not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	playReport := parametersStream.ReadListInt32LE()

	go protocol.SendPlayReportHandler(nil, client, callID, playReport)
}
