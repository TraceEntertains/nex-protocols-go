package account_management

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

const (
	// ProtocolID is the protocol ID for the Account Management protocol
	ProtocolID = 0x19

	// MethodNintendoCreateAccount is the method ID for the method NintendoCreateAccount
	MethodNintendoCreateAccount = 0x1B
)

// AccountManagementProtocol handles the Account Management nex protocol
type AccountManagementProtocol struct {
	Server                       *nex.Server
	NintendoCreateAccountHandler func(err error, client *nex.Client, callID uint32, strPrincipalName string, strKey string, uiGroups uint32, strEmail string, oAuthData *nex.DataHolder)
}

// Setup initializes the protocol
func (protocol *AccountManagementProtocol) Setup() {
	nex.RegisterDataHolderType(NewNintendoCreateAccountData())
	nex.RegisterDataHolderType(NewAccountExtraInfo())

	protocol.Server.On("Data", func(packet nex.PacketInterface) {
		request := packet.RMCRequest()

		if request.ProtocolID() == ProtocolID {
			protocol.HandlePacket(packet)
		}
	})
}

// HandlePacket handles in incoming PRUDP packet
func (protocol *AccountManagementProtocol) HandlePacket(packet nex.PacketInterface) {
	request := packet.RMCRequest()

	switch request.MethodID() {
	case MethodNintendoCreateAccount:
		go protocol.HandleNintendoCreateAccount(packet)
	default:
		go globals.RespondNotImplemented(packet, ProtocolID)
		fmt.Printf("Unsupported AccountManagement method ID: %#v\n", request.MethodID())
	}
}

// NewAccountManagementProtocol returns a new AccountManagementProtocol
func NewAccountManagementProtocol(server *nex.Server) *AccountManagementProtocol {
	protocol := &AccountManagementProtocol{Server: server}

	protocol.Setup()

	return protocol
}
