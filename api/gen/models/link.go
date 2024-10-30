// Code generated by https://github.com/defany/govk. DO NOT EDIT.

package models

import (
	"encoding/json"
)

// suppress unused package warning
var _ *json.RawMessage

type LinkTargetObject struct {
	// Item ID
	ItemId *int `json:"item_id,omitempty"`
	// Owner ID
	//  Format: int64
	OwnerId *int `json:"owner_id,omitempty"`
	// Object type
	Type *string `json:"type,omitempty"`
}

