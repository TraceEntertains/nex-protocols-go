// Package nexproto provides all the main NEX protocols.
//
// Each folder contains a different package for that specific protocol,
// with all their types and methods needed to parse and build packets with RMC payloads
package nexproto

import (
	"github.com/PretendoNetwork/nex-go"
	account_management_types "github.com/PretendoNetwork/nex-protocols-go/account-management/types"
	authentication_types "github.com/PretendoNetwork/nex-protocols-go/authentication/types"
	match_making_types "github.com/PretendoNetwork/nex-protocols-go/match-making/types"
)

func init() {
	nex.RegisterDataHolderType(account_management_types.NewNintendoCreateAccountData())
	nex.RegisterDataHolderType(account_management_types.NewAccountExtraInfo())
	nex.RegisterDataHolderType(authentication_types.NewNintendoLoginData())
	nex.RegisterDataHolderType(authentication_types.NewAuthenticationInfo())
	nex.RegisterDataHolderType(match_making_types.NewGathering())
	nex.RegisterDataHolderType(match_making_types.NewMatchmakeSession())
}
