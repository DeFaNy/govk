// Code generated by https://github.com/defany/govk. DO NOT EDIT.

package models

import (
	"encoding/json"
)

// suppress unused package warning
var _ *json.RawMessage

// StatsActivity Activity stats
type StatsActivity struct {
	// Comments number
	//  Minimum: 0
	Comments *int `json:"comments,omitempty"`
	// Reposts number
	//  Minimum: 0
	Copies *int `json:"copies,omitempty"`
	// Hidden from news count
	//  Minimum: 0
	Hidden *int `json:"hidden,omitempty"`
	// Likes number
	//  Minimum: 0
	Likes *int `json:"likes,omitempty"`
	// New subscribers count
	//  Minimum: 0
	Subscribed *int `json:"subscribed,omitempty"`
	// Unsubscribed count
	//  Minimum: 0
	Unsubscribed *int `json:"unsubscribed,omitempty"`
}

type StatsCity struct {
	// Visitors number
	//  Minimum: 0
	Count *int `json:"count,omitempty"`
	// City name
	Name *string `json:"name,omitempty"`
	// City ID
	Value *int `json:"value,omitempty"`
}

type StatsCountry struct {
	// Country code
	Code *string `json:"code,omitempty"`
	// Visitors number
	//  Minimum: 0
	Count *int `json:"count,omitempty"`
	// Country name
	Name *string `json:"name,omitempty"`
	// Country ID
	Value *int `json:"value,omitempty"`
}

type StatsPeriod struct {
	Activity *StatsActivity `json:"activity,omitempty"`
	// Unix timestamp
	PeriodFrom *int `json:"period_from,omitempty"`
	// Unix timestamp
	PeriodTo *int `json:"period_to,omitempty"`
	Reach *StatsReach `json:"reach,omitempty"`
	Visitors *StatsViews `json:"visitors,omitempty"`
}

// StatsReach Reach stats
type StatsReach struct {
	Age *[]StatsSexAge `json:"age,omitempty"`
	Cities *[]StatsCity `json:"cities,omitempty"`
	Countries *[]StatsCountry `json:"countries,omitempty"`
	// Reach count from mobile devices
	//  Minimum: 0
	MobileReach *int `json:"mobile_reach,omitempty"`
	// Reach count
	//  Minimum: 0
	Reach *int `json:"reach,omitempty"`
	// Subscribers reach count
	//  Minimum: 0
	ReachSubscribers *int `json:"reach_subscribers,omitempty"`
	Sex *[]StatsSexAge `json:"sex,omitempty"`
	SexAge *[]StatsSexAge `json:"sex_age,omitempty"`
}

type StatsSexAge struct {
	// Visitors number
	//  Minimum: 0
	Count *int `json:"count,omitempty"`
	CountSubscribers *int `json:"count_subscribers,omitempty"`
	Reach *int `json:"reach,omitempty"`
	ReachSubscribers *int `json:"reach_subscribers,omitempty"`
	// Sex/age value
	Value string `json:"value"`
}

// StatsViews Views stats
type StatsViews struct {
	Age *[]StatsSexAge `json:"age,omitempty"`
	Cities *[]StatsCity `json:"cities,omitempty"`
	Countries *[]StatsCountry `json:"countries,omitempty"`
	// Number of views from mobile devices
	//  Minimum: 0
	MobileViews *int `json:"mobile_views,omitempty"`
	Sex *[]StatsSexAge `json:"sex,omitempty"`
	SexAge *[]StatsSexAge `json:"sex_age,omitempty"`
	// Views number
	//  Minimum: 0
	Views *int `json:"views,omitempty"`
	// Visitors number
	//  Minimum: 0
	Visitors *int `json:"visitors,omitempty"`
}

type StatsWallpostStat struct {
	// Hidings number
	Hide *int `json:"hide,omitempty"`
	// People have joined the group
	JoinGroup *int `json:"join_group,omitempty"`
	// Link clickthrough
	Links *int `json:"links,omitempty"`
	PostId *int `json:"post_id,omitempty"`
	ReachAds *int `json:"reach_ads,omitempty"`
	// Subscribers reach
	ReachSubscribers *int `json:"reach_subscribers,omitempty"`
	ReachSubscribersCount *int `json:"reach_subscribers_count,omitempty"`
	// Total reach
	ReachTotal *int `json:"reach_total,omitempty"`
	ReachTotalCount *int `json:"reach_total_count,omitempty"`
	ReachViral *int `json:"reach_viral,omitempty"`
	// Reports number
	Report *int `json:"report,omitempty"`
	SexAge *[]StatsSexAge `json:"sex_age,omitempty"`
	// Clickthrough to community
	ToGroup *int `json:"to_group,omitempty"`
	// Unsubscribed members
	Unsubscribe *int `json:"unsubscribe,omitempty"`
}

type StatsGetPostReachResponse struct {
	Response []StatsWallpostStat `json:"response"`
}

type StatsGetResponse struct {
	Response []StatsPeriod `json:"response"`
}

