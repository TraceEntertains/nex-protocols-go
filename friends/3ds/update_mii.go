package friends_3ds

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// UpdateMii sets the UpdateMii handler function
func (protocol *Friends3DSProtocol) UpdateMii(handler func(err error, client *nex.Client, callID uint32, mii *Mii)) {
	protocol.UpdateMiiHandler = handler
}

func (protocol *Friends3DSProtocol) HandleUpdateMii(packet nex.PacketInterface) {
	if protocol.UpdateMiiHandler == nil {
		globals.Logger.Warning("Friends3DS::UpdateMii not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	mii, err := parametersStream.ReadStructure(NewMii())
	if err != nil {
		go protocol.UpdateMiiHandler(err, client, callID, nil)
		return
	}

	go protocol.UpdateMiiHandler(nil, client, callID, mii.(*Mii))
}
