// Package protocol implements the Secure Connection protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// RegisterEx sets the RegisterEx handler function
func (protocol *Protocol) RegisterEx(handler func(err error, client *nex.Client, callID uint32, vecMyURLs []*nex.StationURL, hCustomData *nex.DataHolder)) {
	protocol.registerExHandler = handler
}

func (protocol *Protocol) handleRegisterEx(packet nex.PacketInterface) {
	if protocol.registerExHandler == nil {
		globals.Logger.Warning("SecureConnection::RegisterEx not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	vecMyURLs, err := parametersStream.ReadListStationURL()
	if err != nil {
		go protocol.registerExHandler(fmt.Errorf("Failed to read vecMyURLs from parameters. %s", err.Error()), client, callID, nil, nil)
		return
	}

	hCustomData, err := parametersStream.ReadDataHolder()
	if err != nil {
		go protocol.registerExHandler(fmt.Errorf("Failed to read hCustomData from parameters. %s", err.Error()), client, callID, nil, nil)
		return
	}

	go protocol.registerExHandler(nil, client, callID, vecMyURLs, hCustomData)
}
