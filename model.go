package rings

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/fox-one/mixin-sdk-go"
	"github.com/fox-one/pkg/uuid"
	"github.com/jmoiron/sqlx/types"
	"github.com/shopspring/decimal"
)

var (
	BlocksPerYear = decimal.NewFromInt(2102400)
)

const (
	// TransactionKeyService key service type :string
	TransactionKeyService = "service"
	// TransactionKeyBlock block index :int64
	TransactionKeyBlock = "block"
	// TransactionKeySymbol symbol key :string
	TransactionKeySymbol = "symbol"
	// TransactionKeyPrice price :decimal
	TransactionKeyPrice = "price"
	// TransactionKeyBorrowRate borrow rate :decimal
	TransactionKeyBorrowRate = "borrow_rate"
	// TransactionKeySupplyRate supply rate : decimal
	TransactionKeySupplyRate = "supply_rate"
	// TransactionKeyAmount amount
	TransactionKeyAmount = "amount"
	// TransactionKeyCToken ctokens
	TransactionKeyCToken = "ctoken"
	// TransactionKeyInterest interest
	TransactionKeyInterest = "interest"
	// TransactionKeyStatus status
	TransactionKeyStatus = "status"
	// TransactionKeyUser user
	TransactionKeyUser = "user"
	// TransactionKeyErrorCode error code
	TransactionKeyErrorCode = "error_code"
	// TransactionKeyReferTrace refer trace
	TransactionKeyReferTrace = "refer_trace"
	// TransactionKeyAssetID asset id
	TransactionKeyAssetID = "asset_id"
	// TransactionKeyTotalCash total cash
	TransactionKeyTotalCash = "total_cash"
	// TransactionKeyTotalBorrows total borrows
	TransactionKeyTotalBorrows = "total_borrows"
	// TransactionKeyReserves reserves
	TransactionKeyReserves = "reserves"
	// TransactionKeyCTokens ctokens
	TransactionKeyCTokens = "ctokens"
	// TransactionKeyCTokenAssetID ctoken asset id
	TransactionKeyCTokenAssetID = "ctoken_asset_id"
	// TransactionKeyOrigin origin
	TransactionKeyOrigin = "origin"
	// TransactionKeySupply supply
	TransactionKeySupply = "supply"
	// TransactionKeyBorrow borrow
	TransactionKeyBorrow = "borrow"
	// TransactionKeyMarket market
	TransactionKeyMarket = "market"
)

type Market struct {
	AssetID       string          `json:"asset_id"`
	Symbol        string          `json:"symbol"`
	CTokenAssetID string          `json:"ctoken_asset_id"`
	TotalCash     decimal.Decimal `json:"total_cash"`
	TotalBorrows  decimal.Decimal `json:"total_borrows"`
	// 保留金
	Reserves decimal.Decimal `json:"reserves"`
	// CToken 累计铸造出来的币的数量
	CTokens decimal.Decimal `json:"ctokens"`
	// 初始兑换率
	InitExchangeRate decimal.Decimal `json:"init_exchange_rate"`
	// 平台保留金率 (0, 1), 默认为 0.10
	ReserveFactor decimal.Decimal `json:"reserve_factor"`
	// 清算激励因子 (0, 1), 一般为0.1
	LiquidationIncentive decimal.Decimal `json:"liquidation_incentive"`
	// 资金池的最小资金量
	BorrowCap decimal.Decimal `json:"borrow_cap"`
	//抵押因子 = 可借贷价值 / 抵押资产价值，目前compound设置为0.75. 稳定币(USDT)的抵押率是0,即不可抵押
	CollateralFactor decimal.Decimal `json:"collateral_factor"`
	//触发清算因子 [0.05, 0.9] 清算人最大可清算的资产比例
	CloseFactor decimal.Decimal `json:"close_factor"`
	//基础利率 per year, 0.025
	BaseRate decimal.Decimal `json:"base_rate"`
	// The multiplier of utilization rate that gives the slope of the interest rate. per year
	Multiplier decimal.Decimal `json:"multiplier"`
	// The multiplierPerBlock after hitting a specified utilization point. per year
	JumpMultiplier decimal.Decimal `json:"jump_multiplier"`
	// Kink
	Kink decimal.Decimal `json:"kink"`
	//当前区块高度
	BlockNumber        int64           `json:"block_number"`
	UtilizationRate    decimal.Decimal `json:"utilization_rate"`
	ExchangeRate       decimal.Decimal `json:"exchange_rate"`
	SupplyRatePerBlock decimal.Decimal `json:"supply_rate_per_block"`
	BorrowRatePerBlock decimal.Decimal `json:"borrow_rate_per_block"`
	Price              decimal.Decimal `json:"price"`
	PriceUpdatedAt     time.Time       `json:"price_updated_at"`
	BorrowIndex        decimal.Decimal `json:"borrow_index"`
	Status             MarketStatus    `json:"status"`
	SupplyApy          decimal.Decimal `json:"supply_apy,omitempty"`
	BorrowApy          decimal.Decimal `json:"borrow_apy,omitempty"`
	Suppliers          int64           `json:"suppliers,omitempty"`
	Borrowers          int64           `json:"borrowers,omitempty"`
}

type Transaction struct {
	ID              int64           `json:"id,omitempty"`
	Action          ActionType      `json:"action,omitempty"`
	TraceID         string          `json:"trace_id,omitempty"`
	UserID          string          `json:"user_id,omitempty"`
	FollowID        string          `json:"follow_id,omitempty"`
	SnapshotTraceID string          `json:"snapshot_trace_id,omitempty"`
	AssetID         string          `json:"asset_id,omitempty"`
	Amount          decimal.Decimal `json:"amount,omitempty"`
	Data            types.JSONText  `json:"data,omitempty"`
	CreatedAt       time.Time       `json:"created_at,omitempty"`
	UpdatedAt       time.Time       `json:"updated_at,omitempty"`
}

// MarketStatus market status
type MarketStatus int

const (
	_ MarketStatus = iota
	// MarketStatusOpen open
	MarketStatusOpen
	// MarketStatusClose close
	MarketStatusClose
)

// IsValid is valid status
func (s MarketStatus) IsValid() bool {
	return s == MarketStatusClose ||
		s == MarketStatusOpen
}

type PayRequest struct {
	MemoBase64 string          `json:"memo_base64,omitempty"`
	AssetID    string          `json:"asset_id,omitempty"`
	Amount     decimal.Decimal `json:"amount,omitempty"`
	TraceID    string          `json:"trace_id,omitempty"`
	FollowID   string          `json:"follow_id,omitempty"`
	WithGas    bool            `json:"with_gas,omitempty"`
}

func (p *PayRequest) Validate() error {
	if p.MemoBase64 == "" {
		return errors.New("invalid memo_base64")
	}

	if !p.WithGas && (p.AssetID == "" || !p.Amount.IsPositive()) {
		return errors.New("invalid asset or amount")
	}

	if p.TraceID == "" {
		return errors.New("nil trace_id")
	}

	if p.FollowID == "" {
		return errors.New("nil follow_id")
	}

	if _, err := uuid.FromString(p.TraceID); err != nil {
		return errors.New("trace_id is not uuid")
	}

	if _, err := uuid.FromString(p.FollowID); err != nil {
		return errors.New("follow_id is not uuid")
	}

	return nil
}

type PayInput struct {
	URL           string               `json:"url"`
	TransferInput *mixin.TransferInput `json:"transfer_input"`
}

// TransactionExtraData extra data
type TransactionExtraData map[string]interface{}

// NewTransactionExtra new transaction extra instance
func NewTransactionExtra() TransactionExtraData {
	d := make(TransactionExtraData)
	return d
}

// Put put data
func (t TransactionExtraData) Put(key string, value string) {
	t[key] = value
}

// Format format as []byte by default
func (t TransactionExtraData) Format() []byte {
	bs, e := json.Marshal(t)
	if e != nil {
		return []byte("{}")
	}

	return bs
}
