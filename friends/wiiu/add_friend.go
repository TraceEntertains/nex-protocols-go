package friends_wiiu

import (
	"errors"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// AddFriend sets the AddFriend handler function
func (protocol *FriendsWiiUProtocol) AddFriend(handler func(err error, client *nex.Client, callID uint32, pid uint32)) {
	protocol.AddFriendHandler = handler
}

func (protocol *FriendsWiiUProtocol) HandleAddFriend(packet nex.PacketInterface) {
	if protocol.AddFriendHandler == nil {
		globals.Logger.Warning("FriendsWiiU::AddFriend not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	if len(parametersStream.Bytes()[parametersStream.ByteOffset():]) < 4 {
		err := errors.New("[FriendsWiiU::AddFriend] Data holder not long enough for PID")
		go protocol.AddFriendHandler(err, client, callID, 0)
		return
	}

	pid := parametersStream.ReadUInt32LE()

	go protocol.AddFriendHandler(nil, client, callID, pid)
}
