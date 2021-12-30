package rings

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/fox-one/pando-rings-sdk-go/mtg"
	"github.com/fox-one/pkg/uuid"
	"github.com/shopspring/decimal"
)

var Endpoint string

// RequestAllMarkets response all markets' data
func RequestAllMarkets(ctx context.Context) ([]*Market, error) {
	url := fmt.Sprintf("%s/api/v1/markets/all", endPoint())

	resp, err := request(ctx).Get(url)
	if err != nil {
		return nil, err
	}

	var respBody struct {
		Data []*Market `json:"data"`
	}

	if err := parseResponse(resp, &respBody); err != nil {
		return nil, err
	}

	return respBody.Data, nil
}

// RequestTransactions pull compound transactions
//
// #limit: page limit
// #offset: transaction start time
func RequestTransactions(ctx context.Context, limit int, offset time.Time) ([]*Transaction, error) {
	url := fmt.Sprintf("%s/api/v1/transactions", endPoint())
	req := request(ctx).
		SetQueryParam("limit", strconv.Itoa(limit)).
		SetQueryParam("offset", offset.Format(time.RFC3339Nano))

	resp, err := req.Get(url)
	if err != nil {
		return nil, err
	}

	var transactions []*Transaction
	if err := parseResponse(resp, &transactions); err != nil {
		return nil, err
	}

	return transactions, nil
}

// RequestPayURL
func RequestPayURL(ctx context.Context, payRequest *PayRequest) (*PayInput, error) {
	if payRequest == nil {
		return nil, errors.New("nil payRequest")
	}

	if err := payRequest.Validate(); err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s/api/v1/pay-requests", endPoint())
	resp, err := request(ctx).SetBody(payRequest).Post(url)
	if err != nil {
		return nil, err
	}

	var payInput PayInput
	if err := parseResponse(resp, &payInput); err != nil {
		return nil, err
	}

	return &payInput, nil
}

func encodeParams(followID string, args ...interface{}) (string, error) {
	data, err := mtg.Encode(args...)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(data), nil
}

func endPoint() string {
	if Endpoint == "" {
		panic("endpoint is empty")
	}

	if !strings.HasPrefix(Endpoint, "https://") &&
		!strings.HasPrefix(Endpoint, "HTTPS://") &&
		!strings.HasPrefix(Endpoint, "http://") &&
		!strings.HasPrefix(Endpoint, "HTTP://") {
		panic("no protocol")
	}

	return Endpoint
}

// RequestSupply request supply action url
func RequestSupply(ctx context.Context, followID string, assetID string, amount decimal.Decimal) (*PayInput, error) {
	memo, err := encodeParams(followID, ActionTypeSupply)
	if err != nil {
		return nil, err
	}

	payload := PayRequest{
		MemoBase64: memo,
		AssetID:    assetID,
		Amount:     amount,
		TraceID:    uuid.New(),
		FollowID:   followID,
	}

	payInput, err := RequestPayURL(ctx, &payload)
	if err != nil {
		return nil, err
	}

	return payInput, nil
}

// RequestPledge request pledge action url
func RequestPledge(ctx context.Context, followID string, ctokenAssetID string, amount decimal.Decimal) (*PayInput, error) {
	memo, err := encodeParams(followID, ActionTypePledge)
	if err != nil {
		return nil, err
	}

	payload := PayRequest{
		MemoBase64: memo,
		AssetID:    ctokenAssetID,
		Amount:     amount,
		TraceID:    uuid.New(),
		FollowID:   followID,
	}
	payInput, err := RequestPayURL(ctx, &payload)
	if err != nil {
		return nil, err
	}

	return payInput, nil
}

// RequestUnpledge request unpledge action url
func RequestUnpledge(ctx context.Context, followID string, ctokenAssetID string, ctokenAmount decimal.Decimal) (*PayInput, error) {
	ctokenAsset, err := uuid.FromString(ctokenAssetID)
	if err != nil {
		return nil, err
	}

	memo, err := encodeParams(followID, ActionTypeUnpledge, ctokenAsset, ctokenAmount)
	if err != nil {
		return nil, err
	}

	payload := PayRequest{
		MemoBase64: memo,
		WithGas:    true,
		TraceID:    uuid.New(),
		FollowID:   followID,
	}
	payInput, err := RequestPayURL(ctx, &payload)
	if err != nil {
		return nil, err
	}

	return payInput, nil
}

// RequestQuickPledge request quick_pledge action url
func RequestQuickPledge(ctx context.Context, followID string, assetID string, amount decimal.Decimal) (*PayInput, error) {
	memo, err := encodeParams(followID, ActionTypeQuickPledge)
	if err != nil {
		return nil, err
	}

	payload := PayRequest{
		MemoBase64: memo,
		AssetID:    assetID,
		Amount:     amount,
		TraceID:    uuid.New(),
		FollowID:   followID,
	}
	payInput, err := RequestPayURL(ctx, &payload)
	if err != nil {
		return nil, err
	}

	return payInput, nil
}

// RequestRedeem request redeem action url
func RequestRedeem(ctx context.Context, followID string, ctokenAssetID string, redeemAmount decimal.Decimal) (*PayInput, error) {
	memo, err := encodeParams(followID, ActionTypeRedeem)
	if err != nil {
		return nil, err
	}

	payload := PayRequest{
		MemoBase64: memo,
		AssetID:    ctokenAssetID,
		Amount:     redeemAmount,
		TraceID:    uuid.New(),
		FollowID:   followID,
	}
	payInput, err := RequestPayURL(ctx, &payload)
	if err != nil {
		return nil, err
	}

	return payInput, nil
}

// RequestQuickRedeem request quick_redeem action url
func RequestQuickRedeem(ctx context.Context, followID string, ctokenAssetID string, redeemAmount decimal.Decimal) (*PayInput, error) {
	ctokenAsset, err := uuid.FromString(ctokenAssetID)
	if err != nil {
		return nil, err
	}

	memo, err := encodeParams(followID, ActionTypeQuickRedeem, ctokenAsset, redeemAmount)
	if err != nil {
		return nil, err
	}

	payload := PayRequest{
		MemoBase64: memo,
		WithGas:    true,
		TraceID:    uuid.New(),
		FollowID:   followID,
	}
	payInput, err := RequestPayURL(ctx, &payload)
	if err != nil {
		return nil, err
	}

	return payInput, nil
}

// RequestBorrow request borrow action url
func RequestBorrow(ctx context.Context, followID string, assetID string, borrowAmount decimal.Decimal) (*PayInput, error) {
	asset, err := uuid.FromString(assetID)
	if err != nil {
		return nil, err
	}

	memo, err := encodeParams(followID, ActionTypeBorrow, asset, borrowAmount)
	if err != nil {
		return nil, err
	}

	payload := PayRequest{
		MemoBase64: memo,
		WithGas:    true,
		TraceID:    uuid.New(),
		FollowID:   followID,
	}
	payInput, err := RequestPayURL(ctx, &payload)
	if err != nil {
		return nil, err
	}

	return payInput, nil
}

// RequestQuickBorrow request quick_borrow action url
func RequestQuickBorrow(ctx context.Context, followID string, supplyAssetID string, supplyAmount decimal.Decimal, borrowAssetID string, borrowAmount decimal.Decimal) (*PayInput, error) {
	borrowAsset, err := uuid.FromString(borrowAssetID)
	if err != nil {
		return nil, err
	}

	memo, err := encodeParams(followID, ActionTypeQuickBorrow, borrowAsset, borrowAmount)
	if err != nil {
		return nil, err
	}

	payload := PayRequest{
		MemoBase64: memo,
		AssetID:    supplyAssetID,
		Amount:     supplyAmount,
		TraceID:    uuid.New(),
		FollowID:   followID,
	}
	payInput, err := RequestPayURL(ctx, &payload)
	if err != nil {
		return nil, err
	}

	return payInput, nil
}

// RequestRepay request repay action url
func RequestRepay(ctx context.Context, followID string, assetID string, amount decimal.Decimal) (*PayInput, error) {
	memo, err := encodeParams(followID, ActionTypeRepay)
	if err != nil {
		return nil, err
	}

	payload := PayRequest{
		MemoBase64: memo,
		AssetID:    assetID,
		Amount:     amount,
		TraceID:    uuid.New(),
		FollowID:   followID,
	}
	payInput, err := RequestPayURL(ctx, &payload)
	if err != nil {
		return nil, err
	}

	return payInput, nil
}

// RequestLiquidate request liquidate action url
func RequestLiquidate(ctx context.Context, followID string, userAddress string, supplyCTokenAssetID string, borrowAssetID string, repayAmount decimal.Decimal) (*PayInput, error) {
	address, err := uuid.FromString(userAddress)
	if err != nil {
		return nil, err
	}

	supplyAsset, err := uuid.FromString(supplyCTokenAssetID)
	if err != nil {
		return nil, err
	}

	memo, err := encodeParams(followID, ActionTypeLiquidate, address, supplyAsset)
	if err != nil {
		return nil, err
	}

	payload := PayRequest{
		MemoBase64: memo,
		AssetID:    borrowAssetID,
		Amount:     repayAmount,
		TraceID:    uuid.New(),
		FollowID:   followID,
	}
	payInput, err := RequestPayURL(ctx, &payload)
	if err != nil {
		return nil, err
	}

	return payInput, nil
}
