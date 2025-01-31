// Code generated by https://github.com/defany/govk. DO NOT EDIT.

package models

import (
	"encoding/json"
)

// suppress unused package warning
var _ *json.RawMessage

type StoriesClickableArea struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type StoriesClickableStickerStyle string

const (
	StoriesClickableStickerStyleTransparent StoriesClickableStickerStyle = "transparent"
	StoriesClickableStickerStyleBlueGradient StoriesClickableStickerStyle = "blue_gradient"
	StoriesClickableStickerStyleRedGradient StoriesClickableStickerStyle = "red_gradient"
	StoriesClickableStickerStyleUnderline StoriesClickableStickerStyle = "underline"
	StoriesClickableStickerStyleBlue StoriesClickableStickerStyle = "blue"
	StoriesClickableStickerStyleGreen StoriesClickableStickerStyle = "green"
	StoriesClickableStickerStyleWhite StoriesClickableStickerStyle = "white"
	StoriesClickableStickerStyleQuestionReply StoriesClickableStickerStyle = "question_reply"
	StoriesClickableStickerStyleLight StoriesClickableStickerStyle = "light"
	StoriesClickableStickerStyleImpressive StoriesClickableStickerStyle = "impressive"
)

type StoriesClickableStickerSubtype string

const (
	StoriesClickableStickerSubtypeMarketItem StoriesClickableStickerSubtype = "market_item"
	StoriesClickableStickerSubtypeAliexpressProduct StoriesClickableStickerSubtype = "aliexpress_product"
)

type StoriesClickableStickerType string

const (
	StoriesClickableStickerTypeHashtag StoriesClickableStickerType = "hashtag"
	StoriesClickableStickerTypeMention StoriesClickableStickerType = "mention"
	StoriesClickableStickerTypeLink StoriesClickableStickerType = "link"
	StoriesClickableStickerTypeQuestion StoriesClickableStickerType = "question"
	StoriesClickableStickerTypePlace StoriesClickableStickerType = "place"
	StoriesClickableStickerTypeMarketItem StoriesClickableStickerType = "market_item"
	StoriesClickableStickerTypeMusic StoriesClickableStickerType = "music"
	StoriesClickableStickerTypeStoryReply StoriesClickableStickerType = "story_reply"
	StoriesClickableStickerTypeOwner StoriesClickableStickerType = "owner"
	StoriesClickableStickerTypePost StoriesClickableStickerType = "post"
	StoriesClickableStickerTypePoll StoriesClickableStickerType = "poll"
	StoriesClickableStickerTypeSticker StoriesClickableStickerType = "sticker"
	StoriesClickableStickerTypeApp StoriesClickableStickerType = "app"
	StoriesClickableStickerTypeSituationalTheme StoriesClickableStickerType = "situational_theme"
)

type StoriesClickableSticker struct {
	App *AppsAppMin `json:"app,omitempty"`
	// Additional context for app sticker
	AppContext *string `json:"app_context,omitempty"`
	Audio *AudioAudio `json:"audio,omitempty"`
	AudioStartTime *int `json:"audio_start_time,omitempty"`
	ClickableArea []StoriesClickableArea `json:"clickable_area"`
	// Color, hex format
	Color *string `json:"color,omitempty"`
	// Whether current user has unread interaction with this app
	HasNewInteractions *bool `json:"has_new_interactions,omitempty"`
	Hashtag *string `json:"hashtag,omitempty"`
	// Clickable sticker ID
	Id int `json:"id"`
	// Whether current user allowed broadcast notify from this app
	IsBroadcastNotifyAllowed *bool `json:"is_broadcast_notify_allowed,omitempty"`
	LinkObject *BaseLink `json:"link_object,omitempty"`
	MarketItem *MarketMarketItem `json:"market_item,omitempty"`
	Mention *string `json:"mention,omitempty"`
	//  Format: int64
	OwnerId *int `json:"owner_id,omitempty"`
	PlaceId *int `json:"place_id,omitempty"`
	Poll *PollsPoll `json:"poll,omitempty"`
	PostId *int `json:"post_id,omitempty"`
	//  Format: int64
	PostOwnerId *int `json:"post_owner_id,omitempty"`
	Question *string `json:"question,omitempty"`
	QuestionButton *string `json:"question_button,omitempty"`
	SituationalAppUrl *string `json:"situational_app_url,omitempty"`
	SituationalThemeId *int `json:"situational_theme_id,omitempty"`
	// Sticker ID
	StickerId *int `json:"sticker_id,omitempty"`
	// Sticker pack ID
	StickerPackId *int `json:"sticker_pack_id,omitempty"`
	StoryId *int `json:"story_id,omitempty"`
	Style *StoriesClickableStickerStyle`json:"style,omitempty"`
	Subtype *StoriesClickableStickerSubtype`json:"subtype,omitempty"`
	TooltipText *string `json:"tooltip_text,omitempty"`
	Type StoriesClickableStickerType`json:"type"`
}

type StoriesClickableStickers struct {
	ClickableStickers []StoriesClickableSticker `json:"clickable_stickers"`
	//  Minimum: 0
	OriginalHeight int `json:"original_height"`
	//  Minimum: 0
	OriginalWidth int `json:"original_width"`
}

type StoriesFeedItemType string

const (
	StoriesFeedItemTypePromoStories StoriesFeedItemType = "promo_stories"
	StoriesFeedItemTypeStories StoriesFeedItemType = "stories"
	StoriesFeedItemTypeLiveActive StoriesFeedItemType = "live_active"
	StoriesFeedItemTypeLiveFinished StoriesFeedItemType = "live_finished"
	StoriesFeedItemTypeCommunityGroupedStories StoriesFeedItemType = "community_grouped_stories"
	StoriesFeedItemTypeAppGroupedStories StoriesFeedItemType = "app_grouped_stories"
	StoriesFeedItemTypeBirthday StoriesFeedItemType = "birthday"
	StoriesFeedItemTypeDiscover StoriesFeedItemType = "discover"
	StoriesFeedItemTypeAdvices StoriesFeedItemType = "advices"
)

type StoriesFeedItem struct {
	// App, which stories has been grouped (for type app_grouped_stories)
	App *AppsAppMin `json:"app,omitempty"`
	BirthdayUserId *int `json:"birthday_user_id,omitempty"`
	// Grouped stories of various authors (for types community_grouped_stories/app_grouped_stories type)
	Grouped *[]StoriesFeedItem `json:"grouped,omitempty"`
	HasUnseen *bool `json:"has_unseen,omitempty"`
	Id *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	// Additional data for promo stories (for type promo_stories)
	PromoData *StoriesPromoBlock `json:"promo_data,omitempty"`
	// Author stories
	Stories *[]StoriesStory `json:"stories,omitempty"`
	TrackCode *string `json:"track_code,omitempty"`
	// Type of Feed Item
	Type StoriesFeedItemType`json:"type"`
}

// StoriesPromoBlock Additional data for promo stories
type StoriesPromoBlock struct {
	// Promo story title
	Name string `json:"name"`
	// Hide animation for promo story
	NotAnimated bool `json:"not_animated"`
	// RL of square photo of the story with 100 pixels in width
	Photo100 string `json:"photo_100"`
	// RL of square photo of the story with 50 pixels in width
	Photo50 string `json:"photo_50"`
}

type StoriesReplies struct {
	// Replies number.
	//  Minimum: 0
	Count int `json:"count"`
	// New replies number.
	New *int `json:"new,omitempty"`
}

type StoriesStatLine struct {
	//  Minimum: 0
	Counter *int `json:"counter,omitempty"`
	IsUnavailable *bool `json:"is_unavailable,omitempty"`
	Name string `json:"name"`
}

type StoriesStory struct {
	// Access key for private object.
	AccessKey *string `json:"access_key,omitempty"`
	BirthdayWishUserId *int `json:"birthday_wish_user_id,omitempty"`
	// Information whether story has question sticker and current user can send question to the author
	CanAsk *BaseBoolInt `json:"can_ask,omitempty"`
	// Information whether story has question sticker and current user can send anonymous question to the author
	CanAskAnonymous *BaseBoolInt `json:"can_ask_anonymous,omitempty"`
	// Information whether current user can comment the story (0 - no, 1 - yes).
	CanComment *BaseBoolInt `json:"can_comment,omitempty"`
	// Information whether current user can hide the story (0 - no, 1 - yes).
	CanHide *BaseBoolInt `json:"can_hide,omitempty"`
	// Information whether current user can like the story.
	CanLike *bool `json:"can_like,omitempty"`
	// Information whether current user can reply to the story (0 - no, 1 - yes).
	CanReply *BaseBoolInt `json:"can_reply,omitempty"`
	// Information whether current user can see the story (0 - no, 1 - yes).
	CanSee *BaseBoolInt `json:"can_see,omitempty"`
	// Information whether current user can share the story (0 - no, 1 - yes).
	CanShare *BaseBoolInt `json:"can_share,omitempty"`
	CanUseInNarrative *bool `json:"can_use_in_narrative,omitempty"`
	ClickableStickers *StoriesClickableStickers `json:"clickable_stickers,omitempty"`
	// Date when story has been added in Unixtime.
	//  Minimum: 0
	Date *int `json:"date,omitempty"`
	// Story expiration time. Unixtime.
	//  Minimum: 0
	ExpiresAt *int `json:"expires_at,omitempty"`
	FirstNarrativeTitle *string `json:"first_narrative_title,omitempty"`
	// Story ID.
	Id int `json:"id"`
	// Information whether the story is deleted (false - no, true - yes).
	IsDeleted *bool `json:"is_deleted,omitempty"`
	// Information whether the story is expired (false - no, true - yes).
	IsExpired *bool `json:"is_expired,omitempty"`
	Link *StoriesStoryLink `json:"link,omitempty"`
	NarrativesCount *int `json:"narratives_count,omitempty"`
	// Story owner's ID.
	//  Format: int64
	OwnerId int `json:"owner_id"`
	ParentStory *StoriesStory `json:"parent_story,omitempty"`
	// Access key for private object.
	ParentStoryAccessKey *string `json:"parent_story_access_key,omitempty"`
	// Parent story ID.
	ParentStoryId *int `json:"parent_story_id,omitempty"`
	// Parent story owner's ID.
	ParentStoryOwnerId *int `json:"parent_story_owner_id,omitempty"`
	Photo *PhotosPhoto `json:"photo,omitempty"`
	// Replies counters to current story.
	Replies *StoriesReplies `json:"replies,omitempty"`
	// Information whether current user has seen the story or not (0 - no, 1 - yes).
	Seen *BaseBoolInt `json:"seen,omitempty"`
	Type *StoriesStoryType `json:"type,omitempty"`
	Video *VideoVideoFull `json:"video,omitempty"`
	// Views number.
	//  Minimum: 0
	Views *int `json:"views,omitempty"`
}

type StoriesStoryLink struct {
	// How to open url
	LinkUrlTarget *string `json:"link_url_target,omitempty"`
	// Link text
	Text string `json:"text"`
	// Link URL
	//  Format: uri
	Url string `json:"url"`
}

type StoriesStoryStats struct {
	Answer StoriesStoryStatsStat `json:"answer"`
	Bans StoriesStoryStatsStat `json:"bans"`
	Likes StoriesStoryStatsStat `json:"likes"`
	OpenLink StoriesStoryStatsStat `json:"open_link"`
	Replies StoriesStoryStatsStat `json:"replies"`
	Shares StoriesStoryStatsStat `json:"shares"`
	Subscribers StoriesStoryStatsStat `json:"subscribers"`
	Views StoriesStoryStatsStat `json:"views"`
}

type StoriesStoryStatsStat struct {
	// Stat value
	//  Minimum: 0
	Count *int `json:"count,omitempty"`
	State StoriesStoryStatsState `json:"state"`
}

// StoriesStoryStatsState Statistic state
type StoriesStoryStatsState string

const (
	StoriesStoryStatsStateOn StoriesStoryStatsState = "on"
	StoriesStoryStatsStateOff StoriesStoryStatsState = "off"
	StoriesStoryStatsStateHidden StoriesStoryStatsState = "hidden"
)

// StoriesStoryType Story type.
type StoriesStoryType string

const (
	StoriesStoryTypePhoto StoriesStoryType = "photo"
	StoriesStoryTypeVideo StoriesStoryType = "video"
	StoriesStoryTypeLiveActive StoriesStoryType = "live_active"
	StoriesStoryTypeLiveFinished StoriesStoryType = "live_finished"
	StoriesStoryTypeBirthdayInvite StoriesStoryType = "birthday_invite"
)

type StoriesUploadLinkText string

const (
	StoriesUploadLinkTextToStore StoriesUploadLinkText = "to_store"
	StoriesUploadLinkTextVote StoriesUploadLinkText = "vote"
	StoriesUploadLinkTextMore StoriesUploadLinkText = "more"
	StoriesUploadLinkTextBook StoriesUploadLinkText = "book"
	StoriesUploadLinkTextOrder StoriesUploadLinkText = "order"
	StoriesUploadLinkTextEnroll StoriesUploadLinkText = "enroll"
	StoriesUploadLinkTextFill StoriesUploadLinkText = "fill"
	StoriesUploadLinkTextSignup StoriesUploadLinkText = "signup"
	StoriesUploadLinkTextBuy StoriesUploadLinkText = "buy"
	StoriesUploadLinkTextTicket StoriesUploadLinkText = "ticket"
	StoriesUploadLinkTextWrite StoriesUploadLinkText = "write"
	StoriesUploadLinkTextOpen StoriesUploadLinkText = "open"
	StoriesUploadLinkTextLearnMore StoriesUploadLinkText = "learn_more"
	StoriesUploadLinkTextView StoriesUploadLinkText = "view"
	StoriesUploadLinkTextGoTo StoriesUploadLinkText = "go_to"
	StoriesUploadLinkTextContact StoriesUploadLinkText = "contact"
	StoriesUploadLinkTextWatch StoriesUploadLinkText = "watch"
	StoriesUploadLinkTextPlay StoriesUploadLinkText = "play"
	StoriesUploadLinkTextInstall StoriesUploadLinkText = "install"
	StoriesUploadLinkTextRead StoriesUploadLinkText = "read"
	StoriesUploadLinkTextCalendar StoriesUploadLinkText = "calendar"
)

type StoriesViewersItem struct {
	// user has like for this object
	IsLiked bool `json:"is_liked"`
	User *UsersUserFull `json:"user,omitempty"`
	// user id
	//  Format: int64
	UserId int `json:"user_id"`
}

type StoriesGetBannedExtendedResponse struct {
	Stories struct {
		// Stories count
		Count int `json:"count"`
		Groups []GroupsGroupFull `json:"groups"`
		Items []int `json:"items"`
		Profiles []UsersUserFull `json:"profiles"`
	} `json:"stories"`
}

type StoriesGetBannedResponse struct {
	Stories struct {
		// Stories count
		Count int `json:"count"`
		//  Format: int64
		Items []int `json:"items"`
	} `json:"stories"`
}

type StoriesGetByIdExtendedResponse struct {
	Stories struct {
		// Stories count
		Count int `json:"count"`
		Groups []GroupsGroupFull `json:"groups"`
		Items []StoriesStory `json:"items"`
		Profiles []UsersUserFull `json:"profiles"`
	} `json:"stories"`
}

type StoriesGetPhotoUploadServerResponse struct {
	Stories struct {
		// Upload URL
		UploadUrl string `json:"upload_url"`
		// Users ID who can to see story.
		//  Minimum: 0
		UserIds []int `json:"user_ids"`
	} `json:"stories"`
}

type StoriesGetStatsResponse struct {
	Response StoriesStoryStats `json:"response"`
}

type StoriesGetVideoUploadServerResponse struct {
	Stories struct {
		// Upload URL
		UploadUrl string `json:"upload_url"`
		// Users ID who can to see story.
		//  Minimum: 0
		UserIds []int `json:"user_ids"`
	} `json:"stories"`
}

type StoriesGetViewersExtendedV5115Response struct {
	Stories struct {
		// Viewers count
		Count int `json:"count"`
		HiddenReason *string `json:"hidden_reason,omitempty"`
		Items []StoriesViewersItem `json:"items"`
		NextFrom *string `json:"next_from,omitempty"`
	} `json:"stories"`
}

type StoriesGetViewersExtendedResponse struct {
	Stories struct {
		// Viewers count
		Count int `json:"count"`
		Items []UsersUserFull `json:"items"`
	} `json:"stories"`
}

type StoriesGetV5113Response struct {
	Stories struct {
		Count int `json:"count"`
		Groups *[]GroupsGroup `json:"groups,omitempty"`
		Items []StoriesFeedItem `json:"items"`
		NeedUploadScreen *bool `json:"need_upload_screen,omitempty"`
		NextFrom *string `json:"next_from,omitempty"`
		Profiles *[]UsersUserFull `json:"profiles,omitempty"`
	} `json:"stories"`
}

type StoriesGetResponse struct {
	Stories struct {
		// Stories count
		Count int `json:"count"`
		Groups *[]GroupsGroup `json:"groups,omitempty"`
		Items [][]StoriesStory `json:"items"`
		NeedUploadScreen *bool `json:"need_upload_screen,omitempty"`
		Profiles *[]UsersUserFull `json:"profiles,omitempty"`
		PromoData *StoriesPromoBlock `json:"promo_data,omitempty"`
	} `json:"stories"`
}

type StoriesSaveResponse struct {
	Stories struct {
		Count int `json:"count"`
		Groups *[]GroupsGroup `json:"groups,omitempty"`
		Items []StoriesStory `json:"items"`
		Profiles *[]UsersUser `json:"profiles,omitempty"`
	} `json:"stories"`
}

type StoriesUploadResponse struct {
	Stories struct {
		// A string hash that is used in the stories.save method
		UploadResult *string `json:"upload_result,omitempty"`
	} `json:"stories"`
}

