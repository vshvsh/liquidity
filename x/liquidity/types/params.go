package types

import (
	"fmt"
	"gopkg.in/yaml.v2"

	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

// Const value of liquidity module
const (
	CancelOrderLifeSpan int64 = 0

	// min number of reserveCoins for LiquidityPoolType only 2 is allowed on this spec
	MinReserveCoinNum uint32 = 2

	// max number of reserveCoins for LiquidityPoolType only 2 is allowed on this spec
	MaxReserveCoinNum uint32 = 2

	// Number of blocks in one batch
	DefaultUnitBatchSize uint32 = 1

	// index of target pool type, only 1 is allowed on this version.
	DefaultPoolTypeIndex uint32 = 1

	// swap type of available swap request, only 1 is allowed on this version.
	DefaultSwapType uint32 = 1
)

// Parameter store keys
var (
	KeyLiquidityPoolTypes       = []byte("LiquidityPoolTypes")
	KeyMinInitDepositToPool     = []byte("MinInitDepositToPool")
	KeyInitPoolCoinMintAmount   = []byte("InitPoolCoinMintAmount")
	KeyReserveCoinLimitAmount   = []byte("ReserveCoinLimitAmount")
	KeySwapFeeRate              = []byte("SwapFeeRate")
	KeyLiquidityPoolCreationFee = []byte("LiquidityPoolCreationFee")
	KeyUnitBatchSize            = []byte("UnitBatchSize")
	KeyWithdrawFeeRate          = []byte("WithdrawFeeRate")
	KeyMaxOrderAmountRatio      = []byte("MaxOrderAmountRatio")

	DefaultMinInitDepositToPool     = sdk.NewInt(1000000)
	DefaultInitPoolCoinMintAmount   = sdk.NewInt(1000000)
	DefaultReserveCoinLimitAmount   = sdk.ZeroInt()
	DefaultSwapFeeRate              = sdk.NewDecWithPrec(3, 3) // "0.003000000000000000"
	DefaultWithdrawFeeRate          = sdk.NewDecWithPrec(3, 3) // "0.003000000000000000"
	DefaultMaxOrderAmountRatio      = sdk.NewDecWithPrec(1, 1) // "0.100000000000000000"
	DefaultLiquidityPoolCreationFee = sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(100000000)))
	MinOfferCoinAmount              = sdk.NewInt(100)

	halfRatio, _ = sdk.NewDecFromStr("0.5")
	HalfRatio    = halfRatio

	DecimalErrThreshold3 = sdk.NewDecFromIntWithPrec(sdk.OneInt(), 3)
	DecimalErrThreshold10 = sdk.NewDecFromIntWithPrec(sdk.OneInt(), 10)

	DefaultLiquidityPoolType = LiquidityPoolType{
		PoolTypeIndex:     1,
		Name:              "DefaultPoolType",
		MinReserveCoinNum: MinReserveCoinNum,
		MaxReserveCoinNum: MaxReserveCoinNum,
	}
)

// NewParams liquidity paramtypes constructor
func NewParams(liquidityPoolTypes []LiquidityPoolType, minInitDeposit, initPoolCoinMint, reserveCoinLimit sdk.Int, creationFee sdk.Coins,
	swapFeeRate, withdrawFeeRate, maxOrderAmtRatio sdk.Dec, unitBatchSize uint32) Params {
	return Params{
		LiquidityPoolTypes:       liquidityPoolTypes,
		MinInitDepositToPool:     minInitDeposit,
		InitPoolCoinMintAmount:   initPoolCoinMint,
		ReserveCoinLimitAmount:   reserveCoinLimit,
		LiquidityPoolCreationFee: creationFee,
		SwapFeeRate:              swapFeeRate,
		WithdrawFeeRate:          withdrawFeeRate,
		MaxOrderAmountRatio:      maxOrderAmtRatio,
		UnitBatchSize:            unitBatchSize,
	}
}

// ParamTypeTable returns the TypeTable for liquidity module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// KeyValuePairs implements paramtypes.KeyValuePairs
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {

	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyLiquidityPoolTypes, &p.LiquidityPoolTypes, validateLiquidityPoolTypes),
		paramtypes.NewParamSetPair(KeyMinInitDepositToPool, &p.MinInitDepositToPool, validateMinInitDepositToPool),
		paramtypes.NewParamSetPair(KeyInitPoolCoinMintAmount, &p.InitPoolCoinMintAmount, validateInitPoolCoinMintAmount),
		paramtypes.NewParamSetPair(KeyReserveCoinLimitAmount, &p.ReserveCoinLimitAmount, validateReserveCoinLimitAmount),
		paramtypes.NewParamSetPair(KeyLiquidityPoolCreationFee, &p.LiquidityPoolCreationFee, validateLiquidityPoolCreationFee),
		paramtypes.NewParamSetPair(KeySwapFeeRate, &p.SwapFeeRate, validateSwapFeeRate),
		paramtypes.NewParamSetPair(KeyWithdrawFeeRate, &p.WithdrawFeeRate, validateWithdrawFeeRate),
		paramtypes.NewParamSetPair(KeyMaxOrderAmountRatio, &p.MaxOrderAmountRatio, validateMaxOrderAmountRatio),
		paramtypes.NewParamSetPair(KeyUnitBatchSize, &p.UnitBatchSize, validateUnitBatchSize),
	}
}

// DefaultParams returns the default liquidity module parameters
func DefaultParams() Params {
	var defaultLiquidityPoolTypes []LiquidityPoolType
	defaultLiquidityPoolTypes = append(defaultLiquidityPoolTypes, DefaultLiquidityPoolType)

	return NewParams(
		defaultLiquidityPoolTypes,
		DefaultMinInitDepositToPool,
		DefaultInitPoolCoinMintAmount,
		DefaultReserveCoinLimitAmount,
		DefaultLiquidityPoolCreationFee,
		DefaultSwapFeeRate,
		DefaultWithdrawFeeRate,
		DefaultMaxOrderAmountRatio,
		DefaultUnitBatchSize)
}

// String returns a human readable string representation of the parameters.
func (p Params) String() string {
	out, _ := yaml.Marshal(p)
	return string(out)
}

// Validate returns err if Params is invalid
func (p Params) Validate() error {
	if err := validateLiquidityPoolTypes(p.LiquidityPoolTypes); err != nil {
		return err
	}

	if err := validateMinInitDepositToPool(p.MinInitDepositToPool); err != nil {
		return err
	}

	if err := validateInitPoolCoinMintAmount(p.InitPoolCoinMintAmount); err != nil {
		return err
	}

	if err := validateReserveCoinLimitAmount(p.ReserveCoinLimitAmount); err != nil {
		return err
	}

	if err := validateLiquidityPoolCreationFee(p.LiquidityPoolCreationFee); err != nil {
		return err
	}

	if err := validateSwapFeeRate(p.SwapFeeRate); err != nil {
		return err
	}

	if err := validateWithdrawFeeRate(p.WithdrawFeeRate); err != nil {
		return err
	}

	if err := validateMaxOrderAmountRatio(p.MaxOrderAmountRatio); err != nil {
		return err
	}

	if err := validateUnitBatchSize(p.UnitBatchSize); err != nil {
		return err
	}
	// TODO: add detail validate logic

	return nil
}

// check validity of the list of liquidity pool type
func validateLiquidityPoolTypes(i interface{}) error {
	v, ok := i.([]LiquidityPoolType)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}
	if v == nil {
		return fmt.Errorf("empty parameter: LiquidityPoolTypes #{i}")
	}
	for i, p := range v {
		if i+1 != int(p.PoolTypeIndex) {
			return fmt.Errorf("LiquidityPoolTypes index must be sorted")
		}
	}
	if len(v) > 1 {
		return fmt.Errorf("only default pool type allowed on this version")
	}
	if len(v) < 1 {
		return fmt.Errorf("need to default pool type")
	}
	if !v[0].Equal(DefaultLiquidityPoolType) {
		return fmt.Errorf("only default pool type allowed")
	}
	return nil
}

// Validate that the minimum deposit.
func validateMinInitDepositToPool(i interface{}) error {
	v, ok := i.(sdk.Int)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}
	if !v.IsPositive() {
		return fmt.Errorf("MinInitDepositToPool must be positive: %s", v)
	}

	return nil
}

// Validate that the minimum deposit for initiating pool.
func validateInitPoolCoinMintAmount(i interface{}) error {
	v, ok := i.(sdk.Int)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}
	if !v.IsPositive() {
		return fmt.Errorf("InitPoolCoinMintAmount must be positive: %s", v)
	}
	if v.LT(DefaultInitPoolCoinMintAmount) {
		return fmt.Errorf("InitPoolCoinMintAmount should over default value: %s", v)
	}
	return nil
}

// Validate that the Limit the size of each liquidity pool.
func validateReserveCoinLimitAmount(i interface{}) error {
	v, ok := i.(sdk.Int)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}
	if v.IsNegative() {
		return fmt.Errorf("InitPoolCoinMintAmount must be positive or zero: %s", v)
	}
	return nil
}

// Check if the swap fee rate is between 0 and 1
func validateSwapFeeRate(i interface{}) error {
	v, ok := i.(sdk.Dec)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}
	if v.IsNegative() {
		return fmt.Errorf("SwapFeeRate cannot be negative: %s", v)
	}
	if v.GT(sdk.OneDec()) {
		return fmt.Errorf("SwapFeeRate too large: %s", v)
	}
	return nil
}

// Check if the withdraw fee rate is between 0 and 1
func validateWithdrawFeeRate(i interface{}) error {
	v, ok := i.(sdk.Dec)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}
	if v.IsNegative() {
		return fmt.Errorf("WithdrawFeeRate cannot be negative: %s", v)
	}
	if v.GT(sdk.OneDec()) {
		return fmt.Errorf("WithdrawFeeRate too large: %s", v)
	}
	return nil
}

// Check if the max order amount ratio is between 0 and 1
func validateMaxOrderAmountRatio(i interface{}) error {
	v, ok := i.(sdk.Dec)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}
	if v.IsNegative() {
		return fmt.Errorf("WithdrawFeeRate cannot be negative: %s", v)
	}
	if v.GT(sdk.OneDec()) {
		return fmt.Errorf("WithdrawFeeRate too large: %s", v)
	}
	return nil
}

// Check if the pool creation fee is valid
func validateLiquidityPoolCreationFee(i interface{}) error {
	coins, ok := i.(sdk.Coins)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}
	if err := coins.Validate(); err != nil {
		return err
	}
	if coins.Empty() {
		return fmt.Errorf("LiquidityPoolCreationFee cannot be Empty: %s", coins)
	}
	return nil
}

// Check if the liquidity Msg fee is valid
func validateUnitBatchSize(i interface{}) error {
	int, ok := i.(uint32)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}
	if int == 0 {
		return fmt.Errorf("UnitBatchSize cannot be zero")
	}
	return nil
}
