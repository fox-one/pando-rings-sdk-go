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

// RequestPayURL request pay url with mixin schema, e.g: mixin://pay?xxxxxxxxxx
func RequestPayURL(ctx context.Context, payRequest *PayRequest) (string, error) {
	if payRequest == nil {
		return "", errors.New("nil payRequest")
	}

	if err := payRequest.Validate(); err != nil {
		return "", err
	}

	url := fmt.Sprintf("%s/api/v1/pay-requests", endPoint())
	resp, err := request(ctx).SetBody(payRequest).Post(url)
	if err != nil {
		return "", err
	}

	var respBody struct {
		Url string `json:"url"`
	}

	if err := parseResponse(resp, &respBody); err != nil {
		return "", err
	}

	return respBody.Url, nil
}

func mtgEncodingToBase64(args []interface{}) (string, error) {
	bytes, err := mtg.Encode(args...)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(bytes), nil
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
func RequestSupply(ctx context.Context, followID string, assetID string, amount decimal.Decimal) (string, string, error) {
	fID, memoValues, err := NewBasicMemoValues(ActionTypeSupply, followID)
	if err != nil {
		return "", "", err
	}

	memoBase64, err := mtgEncodingToBase64(memoValues)
	if err != nil {
		return "", "", err
	}

	payload := PayRequest{
		MemoBase64: memoBase64,
		AssetID:    assetID,
		Amount:     amount,
		TraceID:    uuid.New(),
		FollowID:   fID,
	}

	url, err := RequestPayURL(ctx, &payload)
	if err != nil {
		return "", "", err
	}

	return fID, url, nil
}

// RequestPledge request pledge action url
func RequestPledge(ctx context.Context, followID string, ctokenAssetID string, amount decimal.Decimal) (string, string, error) {
	fID, memoValues, err := NewBasicMemoValues(ActionTypePledge, followID)
	if err != nil {
		return "", "", err
	}

	memoBase64, err := mtgEncodingToBase64(memoValues)
	if err != nil {
		return "", "", err
	}

	payload := PayRequest{
		MemoBase64: memoBase64,
		AssetID:    ctokenAssetID,
		Amount:     amount,
		TraceID:    uuid.New(),
		FollowID:   fID,
	}
	url, err := RequestPayURL(ctx, &payload)
	if err != nil {
		return "", "", err
	}

	return fID, url, nil
}

// RequestUnpledge request unpledge action url
func RequestUnpledge(ctx context.Context, followID string, ctokenAssetID string, ctokenAmount decimal.Decimal) (string, string, error) {
	fID, memoValues, err := NewBasicMemoValues(ActionTypeUnpledge, followID)
	if err != nil {
		return "", "", err
	}

	ctokenAsset, err := uuid.FromString(ctokenAssetID)
	if err != nil {
		return "", "", err
	}

	memoValues = append(memoValues, ctokenAsset)
	memoValues = append(memoValues, ctokenAmount)

	memoBase64, err := mtgEncodingToBase64(memoValues)
	if err != nil {
		return "", "", err
	}

	payload := PayRequest{
		MemoBase64: memoBase64,
		WithGas:    true,
		TraceID:    uuid.New(),
		FollowID:   fID,
	}
	url, err := RequestPayURL(ctx, &payload)
	if err != nil {
		return "", "", err
	}

	return fID, url, nil
}

// RequestQuickPledge request quick_pledge action url
func RequestQuickPledge(ctx context.Context, followID string, assetID string, amount decimal.Decimal) (string, string, error) {
	fID, memoValues, err := NewBasicMemoValues(ActionTypeQuickPledge, followID)
	if err != nil {
		return "", "", err
	}

	memoBase64, err := mtgEncodingToBase64(memoValues)
	if err != nil {
		return "", "", err
	}

	payload := PayRequest{
		MemoBase64: memoBase64,
		AssetID:    assetID,
		Amount:     amount,
		TraceID:    uuid.New(),
		FollowID:   fID,
	}
	url, err := RequestPayURL(ctx, &payload)
	if err != nil {
		return "", "", err
	}

	return fID, url, nil
}

// RequestRedeem request redeem action url
func RequestRedeem(ctx context.Context, followID string, ctokenAssetID string, redeemAmount decimal.Decimal) (string, string, error) {
	fID, memoValues, err := NewBasicMemoValues(ActionTypeRedeem, followID)
	if err != nil {
		return "", "", err
	}

	memoBase64, err := mtgEncodingToBase64(memoValues)
	if err != nil {
		return "", "", err
	}

	payload := PayRequest{
		MemoBase64: memoBase64,
		AssetID:    ctokenAssetID,
		Amount:     redeemAmount,
		TraceID:    uuid.New(),
		FollowID:   fID,
	}
	url, err := RequestPayURL(ctx, &payload)
	if err != nil {
		return "", "", err
	}

	return fID, url, nil
}

// RequestQuickRedeem request quick_redeem action url
func RequestQuickRedeem(ctx context.Context, followID string, ctokenAssetID string, redeemAmount decimal.Decimal) (string, string, error) {
	fID, memoValues, err := NewBasicMemoValues(ActionTypeQuickRedeem, followID)
	if err != nil {
		return "", "", err
	}

	ctokenAsset, err := uuid.FromString(ctokenAssetID)
	if err != nil {
		return "", "", err
	}

	memoValues = append(memoValues, ctokenAsset)
	memoValues = append(memoValues, redeemAmount)

	memoBase64, err := mtgEncodingToBase64(memoValues)
	if err != nil {
		return "", "", err
	}

	payload := PayRequest{
		MemoBase64: memoBase64,
		WithGas:    true,
		TraceID:    uuid.New(),
		FollowID:   fID,
	}
	url, err := RequestPayURL(ctx, &payload)
	if err != nil {
		return "", "", err
	}

	return fID, url, nil
}

// RequestBorrow request borrow action url
func RequestBorrow(ctx context.Context, followID string, assetID string, borrowAmount decimal.Decimal) (string, string, error) {
	fID, memoValues, err := NewBasicMemoValues(ActionTypeBorrow, followID)
	if err != nil {
		return "", "", err
	}

	asset, err := uuid.FromString(assetID)
	if err != nil {
		return "", "", err
	}

	memoValues = append(memoValues, asset)
	memoValues = append(memoValues, borrowAmount)

	memoBase64, err := mtgEncodingToBase64(memoValues)
	if err != nil {
		return "", "", err
	}

	payload := PayRequest{
		MemoBase64: memoBase64,
		WithGas:    true,
		TraceID:    uuid.New(),
		FollowID:   fID,
	}
	url, err := RequestPayURL(ctx, &payload)
	if err != nil {
		return "", "", err
	}

	return fID, url, nil
}

// RequestQuickBorrow request quick_borrow action url
func RequestQuickBorrow(ctx context.Context, followID string, supplyAssetID string, supplyAmount decimal.Decimal, borrowAssetID string, borrowAmount decimal.Decimal) (string, string, error) {
	fID, memoValues, err := NewBasicMemoValues(ActionTypeQuickBorrow, followID)
	if err != nil {
		return "", "", err
	}

	borrowAsset, err := uuid.FromString(borrowAssetID)
	if err != nil {
		return "", "", err
	}

	memoValues = append(memoValues, borrowAsset)
	memoValues = append(memoValues, borrowAmount)

	memoBase64, err := mtgEncodingToBase64(memoValues)
	if err != nil {
		return "", "", err
	}

	payload := PayRequest{
		MemoBase64: memoBase64,
		AssetID:    supplyAssetID,
		Amount:     supplyAmount,
		TraceID:    uuid.New(),
		FollowID:   fID,
	}
	url, err := RequestPayURL(ctx, &payload)
	if err != nil {
		return "", "", err
	}

	return fID, url, nil
}

// RequestRepay request repay action url
func RequestRepay(ctx context.Context, followID string, assetID string, amount decimal.Decimal) (string, string, error) {
	fID, memoValues, err := NewBasicMemoValues(ActionTypeRepay, followID)
	if err != nil {
		return "", "", err
	}

	memoBase64, err := mtgEncodingToBase64(memoValues)
	if err != nil {
		return "", "", err
	}

	payload := PayRequest{
		MemoBase64: memoBase64,
		AssetID:    assetID,
		Amount:     amount,
		TraceID:    uuid.New(),
		FollowID:   fID,
	}
	url, err := RequestPayURL(ctx, &payload)
	if err != nil {
		return "", "", err
	}

	return fID, url, nil
}

// RequestLiquidate request liquidate action url
func RequestLiquidate(ctx context.Context, followID string, supplyUserID string, supplyCTokenAssetID string, borrowAssetID string, repayAmount decimal.Decimal) (string, string, error) {
	fID, memoValues, err := NewBasicMemoValues(ActionTypeLiquidate, followID)
	if err != nil {
		return "", "", err
	}

	seizedAddress := NewUserAddress(supplyUserID)
	address, err := uuid.FromString(seizedAddress)
	if err != nil {
		return "", "", err
	}
	memoValues = append(memoValues, address)

	supplyAsset, err := uuid.FromString(supplyCTokenAssetID)
	if err != nil {
		return "", "", err
	}
	memoValues = append(memoValues, supplyAsset)

	memoBase64, err := mtgEncodingToBase64(memoValues)
	if err != nil {
		return "", "", err
	}

	payload := PayRequest{
		MemoBase64: memoBase64,
		AssetID:    borrowAssetID,
		Amount:     repayAmount,
		TraceID:    uuid.New(),
		FollowID:   fID,
	}
	url, err := RequestPayURL(ctx, &payload)
	if err != nil {
		return "", "", err
	}

	return fID, url, nil
}
