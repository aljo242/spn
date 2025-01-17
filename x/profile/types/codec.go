package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	// this line is used by starport scaffolding # 1
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgUpdateValidatorDescription{}, "profile/UpdateValidatorDescription", nil)
	cdc.RegisterConcrete(&MsgDeleteValidator{}, "profile/DeleteValidator", nil)
	cdc.RegisterConcrete(&MsgSetValidatorConsAddress{}, "profile/SetValidatorConsAddress", nil)
	cdc.RegisterConcrete(&MsgCreateCoordinator{}, "profile/CreateCoordinator", nil)
	cdc.RegisterConcrete(&MsgUpdateCoordinatorDescription{}, "profile/UpdateCoordinatorDescription", nil)
	cdc.RegisterConcrete(&MsgUpdateCoordinatorAddress{}, "profile/UpdateCoordinatorAddress", nil)
	cdc.RegisterConcrete(&MsgDeleteCoordinator{}, "profile/DeleteCoordinator", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgUpdateValidatorDescription{},
		&MsgDeleteValidator{},
		&MsgSetValidatorConsAddress{},
		&MsgCreateCoordinator{},
		&MsgUpdateCoordinatorDescription{},
		&MsgUpdateCoordinatorAddress{},
		&MsgDeleteCoordinator{},
	)
	// this line is used by starport scaffolding # 3
	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
