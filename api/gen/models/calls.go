// Code generated by https://github.com/defany/govk. DO NOT EDIT.

package models

import (
	"encoding/json"
)

// suppress unused package warning
var _ *json.RawMessage

type CallsCall struct {
	// Call duration
	//  Minimum: 0
	Duration *int `json:"duration,omitempty"`
	// Caller initiator
	//  Minimum: 0
	InitiatorId int `json:"initiator_id"`
	Participants *CallsParticipants `json:"participants,omitempty"`
	// Caller receiver
	//  Minimum: 0
	ReceiverId int `json:"receiver_id"`
	State CallsEndState `json:"state"`
	// Timestamp for call
	Time int `json:"time"`
	// Was this call initiated as video call
	Video *bool `json:"video,omitempty"`
}

// CallsEndState State in which call ended up
type CallsEndState string

const (
	CallsEndStateCanceledByInitiator CallsEndState = "canceled_by_initiator"
	CallsEndStateCanceledByReceiver CallsEndState = "canceled_by_receiver"
	CallsEndStateReached CallsEndState = "reached"
)

type CallsParticipants struct {
	// Participants count
	//  Minimum: 0
	Count *int `json:"count,omitempty"`
	//  Format: int64
	List *[]int `json:"list,omitempty"`
}
