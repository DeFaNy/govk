// Code generated by https://github.com/defany/govk. DO NOT EDIT.

package models

import (
	"encoding/json"
)

// suppress unused package warning
var _ *json.RawMessage

type DonutDonatorSubscriptionInfoStatus string

const (
	DonutDonatorSubscriptionInfoStatusActive DonutDonatorSubscriptionInfoStatus = "active"
	DonutDonatorSubscriptionInfoStatusExpiring DonutDonatorSubscriptionInfoStatus = "expiring"
)

// DonutDonatorSubscriptionInfo Info about user VK Donut subscription
type DonutDonatorSubscriptionInfo struct {
	Amount int `json:"amount"`
	NextPaymentDate int `json:"next_payment_date"`
	//  Format: int64
	OwnerId int `json:"owner_id"`
	Status DonutDonatorSubscriptionInfoStatus`json:"status"`
}

type DonutGetSubscriptionResponse struct {
	Response DonutDonatorSubscriptionInfo `json:"response"`
}

type DonutGetSubscriptionsResponse struct {
	Donut struct {
		//  Minimum: 0
		Count *int `json:"count,omitempty"`
		Groups *[]GroupsGroupFull `json:"groups,omitempty"`
		Profiles *[]UsersUserFull `json:"profiles,omitempty"`
		Subscriptions []DonutDonatorSubscriptionInfo `json:"subscriptions"`
	} `json:"donut"`
}

