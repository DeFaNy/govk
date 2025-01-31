// Code generated by https://github.com/defany/govk. DO NOT EDIT.

package models

import (
	"encoding/json"
)

// suppress unused package warning
var _ *json.RawMessage

type OrdersAmount struct {
	Amounts *[]OrdersAmountItem `json:"amounts,omitempty"`
	// Currency name
	Currency *string `json:"currency,omitempty"`
}

type OrdersAmountItem struct {
	// Votes amount in user's currency
	Amount *float64 `json:"amount,omitempty"`
	// Amount description
	Description *string `json:"description,omitempty"`
	// Votes number
	Votes *string `json:"votes,omitempty"`
}

type OrdersOrderStatus string

const (
	OrdersOrderStatusCreated OrdersOrderStatus = "created"
	OrdersOrderStatusCharged OrdersOrderStatus = "charged"
	OrdersOrderStatusRefunded OrdersOrderStatus = "refunded"
	OrdersOrderStatusChargeable OrdersOrderStatus = "chargeable"
	OrdersOrderStatusCancelled OrdersOrderStatus = "cancelled"
	OrdersOrderStatusDeclined OrdersOrderStatus = "declined"
)

type OrdersOrder struct {
	// Amount
	Amount string `json:"amount"`
	// App order ID
	AppOrderId string `json:"app_order_id"`
	// Cancel transaction ID
	CancelTransactionId *string `json:"cancel_transaction_id,omitempty"`
	// Date of creation in Unixtime
	Date string `json:"date"`
	// Order ID
	Id string `json:"id"`
	// Order item
	Item string `json:"item"`
	// Receiver ID
	ReceiverId string `json:"receiver_id"`
	// Order status
	Status OrdersOrderStatus`json:"status"`
	// Transaction ID
	TransactionId *string `json:"transaction_id,omitempty"`
	// User ID
	UserId string `json:"user_id"`
}

type OrdersSubscription struct {
	// Subscription's application id
	AppId *int `json:"app_id,omitempty"`
	// Subscription's application name
	ApplicationName *string `json:"application_name,omitempty"`
	// Cancel reason
	CancelReason *string `json:"cancel_reason,omitempty"`
	// Date of creation in Unixtime
	CreateTime int `json:"create_time"`
	// Subscription expiration time in Unixtime
	ExpireTime *int `json:"expire_time,omitempty"`
	// Subscription ID
	Id int `json:"id"`
	// Subscription order item
	ItemId string `json:"item_id"`
	// Date of next bill in Unixtime
	NextBillTime *int `json:"next_bill_time,omitempty"`
	// Pending cancel state
	PendingCancel *bool `json:"pending_cancel,omitempty"`
	// Subscription period
	Period int `json:"period"`
	// Date of last period start in Unixtime
	PeriodStartTime int `json:"period_start_time"`
	// Item photo image url
	PhotoUrl *string `json:"photo_url,omitempty"`
	// Subscription price
	Price int `json:"price"`
	// Subscription status
	Status string `json:"status"`
	// Is test subscription
	TestMode *bool `json:"test_mode,omitempty"`
	// Subscription name
	Title *string `json:"title,omitempty"`
	// Date of trial expire in Unixtime
	TrialExpireTime *int `json:"trial_expire_time,omitempty"`
	// Date of last change in Unixtime
	UpdateTime int `json:"update_time"`
}

type OrdersCancelSubscriptionResponse struct {
	// Result
	Response BaseBoolInt `json:"response"`
}

type OrdersChangeStateResponse struct {
	// New state
	Response string `json:"response"`
}

type OrdersGetAmountResponse struct {
	Response []OrdersAmount `json:"response"`
}

type OrdersGetByIdResponse struct {
	Response []OrdersOrder `json:"response"`
}

type OrdersGetUserSubscriptionByIdResponse struct {
	Response OrdersSubscription `json:"response"`
}

type OrdersGetUserSubscriptionsResponse struct {
	Orders struct {
		// Total number
		Count *int `json:"count,omitempty"`
		Items *[]OrdersSubscription `json:"items,omitempty"`
	} `json:"orders"`
}

type OrdersGetResponse struct {
	Response []OrdersOrder `json:"response"`
}

type OrdersUpdateSubscriptionResponse struct {
	// Result
	Response BaseBoolInt `json:"response"`
}

