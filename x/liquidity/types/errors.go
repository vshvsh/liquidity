package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// liquidity module sentinel errors
var (
	ErrPoolNotExists                = sdkerrors.Register(ModuleName, 1, "pool not exists")
	ErrPoolTypeNotExists            = sdkerrors.Register(ModuleName, 2, "pool type not exists")
	ErrEqualDenom                   = sdkerrors.Register(ModuleName, 3, "reserve coin denomination are equal")
	ErrInvalidDenom                 = sdkerrors.Register(ModuleName, 4, "invalid denom")
	ErrNumOfReserveCoin             = sdkerrors.Register(ModuleName, 5, "invalid number of reserve coin")
	ErrNumOfPoolCoin                = sdkerrors.Register(ModuleName, 6, "invalid number of pool coin")
	ErrInsufficientPool             = sdkerrors.Register(ModuleName, 7, "insufficient pool")
	ErrInsufficientBalance          = sdkerrors.Register(ModuleName, 8, "insufficient coin balance to escrow")
	ErrLessThanMinInitDeposit       = sdkerrors.Register(ModuleName, 9, "deposit coin less than MinInitDepositToPool")
	ErrNotImplementedYet            = sdkerrors.Register(ModuleName, 10, "not implemented yet")
	ErrPoolAlreadyExists            = sdkerrors.Register(ModuleName, 11, "the pool already exists")
	ErrPoolBatchNotExists           = sdkerrors.Register(ModuleName, 12, "pool batch not exists")
	ErrOrderBookInvalidity          = sdkerrors.Register(ModuleName, 13, "orderbook is not validity")
	ErrBatchNotExecuted             = sdkerrors.Register(ModuleName, 14, "the liquidity pool batch is not executed")
	ErrEmptyPoolCreatorAddr         = sdkerrors.Register(ModuleName, 15, "empty pool creator address")
	ErrEmptyDepositorAddr           = sdkerrors.Register(ModuleName, 16, "empty pool depositor address")
	ErrEmptyWithdrawerAddr          = sdkerrors.Register(ModuleName, 17, "empty pool withdrawer address")
	ErrEmptySwapRequesterAddr       = sdkerrors.Register(ModuleName, 18, "empty pool swap requester address")
	ErrBadPoolCoinAmount            = sdkerrors.Register(ModuleName, 19, "invalid pool coin amount")
	ErrBadDepositCoinsAmount        = sdkerrors.Register(ModuleName, 20, "invalid pool coin amount")
	ErrBadOfferCoinAmount           = sdkerrors.Register(ModuleName, 21, "invalid offer coin amount")
	ErrBadOrderingReserveCoin       = sdkerrors.Register(ModuleName, 22, "reserve coin denoms not ordered alphabetical")
	ErrBadOderPrice                 = sdkerrors.Register(ModuleName, 23, "invalid order price ")
	ErrNumOfReserveCoinDenoms       = sdkerrors.Register(ModuleName, 24, "invalid reserve coin denoms")
	ErrEmptyReserveAccountAddress   = sdkerrors.Register(ModuleName, 25, "empty reserve account address")
	ErrEmptyPoolCoinDenom           = sdkerrors.Register(ModuleName, 26, "empty pool coin denom")
	ErrBadOrderingReserveCoinDenoms = sdkerrors.Register(ModuleName, 27, "bad ordering reserve coin denoms")
	ErrBadReserveAccountAddress     = sdkerrors.Register(ModuleName, 28, "bad reserve account address")
	ErrBadPoolCoinDenom             = sdkerrors.Register(ModuleName, 29, "bad pool coin denom")
)
