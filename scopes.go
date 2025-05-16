package go_twitchAuth

import (
	"bytes"
	"encoding/json"
)

type ScopeType int

const (
	ScopeAnalyticsReadExtensions ScopeType = iota + 1
	ScopeAnalyticsReadGames
	ScopeBitsRead
	ScopeChannelBot
	ScopeChannelManageAds
	ScopeChannelReadAds
	ScopeChannelManageBroadcast
	ScopeChannelReadCharity
	ScopeChannelEditCommercial
	ScopeChannelReadEditors
	ScopeChannelManageExtensions
	ScopeChannelReadGoals
	ScopeChannelReadGuestStar
	ScopeChannelManageGuestStar
	ScopeChannelReadHypeTrain
	ScopeChannelManageModerators
	ScopeChannelReadPolls
	ScopeChannelManagePolls
	ScopeChannelReadPredictions
	ScopeChannelManagePredictions
	ScopeChannelManageRaids
	ScopeChannelReadRedemptions
	ScopeChannelManageRedemptions
	ScopeChannelManageSchedule
	ScopeChannelReadStreamKey
	ScopeChannelReadSubscriptions
	ScopeChannelManageVideos
	ScopeChannelReadVips
	ScopeChannelManageVips
	ScopeChannelModerate
	ScopeClipsEdit
	ScopeModerationRead
	ScopeModeratorManageAnnouncements
	ScopeModeratorManageAutomod
	ScopeModeratorReadAutomodSettings
	ScopeModeratorManageAutomodSettings
	ScopeModeratorReadBannedUsers
	ScopeModeratorManageBannedUsers
	ScopeModeratorReadBlockedTerms
	ScopeModeratorReadChatMessages
	ScopeModeratorManageBlockedTerms
	ScopeModeratorManageChatMessages
	ScopeModeratorReadChatSettings
	ScopeModeratorManageChatSettings
	ScopeModeratorReadChatters
	ScopeModeratorReadFollowers
	ScopeModeratorReadGuestStar
	ScopeModeratorManageGuestStar
	ScopeModeratorReadModerators
	ScopeModeratorReadShieldMode
	ScopeModeratorManageShieldMode
	ScopeModeratorReadShoutouts
	ScopeModeratorManageShoutouts
	ScopeModeratorReadSuspiciousUsers
	ScopeModeratorReadUnbanRequests
	ScopeModeratorManageUnbanRequests
	ScopeModeratorReadVips
	ScopeModeratorReadWarnings
	ScopeModeratorManageWarnings
	ScopeUserBot
	ScopeUserEdit
	ScopeUserEditBroadcast
	ScopeUserReadBlockedUsers
	ScopeUserManageBlockedUsers
	ScopeUserReadBroadcast
	ScopeUserReadChat
	ScopeUserManageChatColor
	ScopeUserReadEmail
	ScopeUserReadEmotes
	ScopeUserReadFollows
	ScopeUserReadModeratedChannels
	ScopeUserReadSubscriptions
	ScopeUserReadWhispers
	ScopeUserManageWhispers
	ScopeUserWriteChat
)

var ScopeTypeId = map[string]ScopeType{
	"analytics:read:extensions":         ScopeAnalyticsReadExtensions,
	"analytics:read:games":              ScopeAnalyticsReadGames,
	"bits:read":                         ScopeBitsRead,
	"channel:bot":                       ScopeChannelBot,
	"channel:manage:ads":                ScopeChannelManageAds,
	"channel:read:ads":                  ScopeChannelReadAds,
	"channel:manage:broadcast":          ScopeChannelManageBroadcast,
	"channel:read:charity":              ScopeChannelReadCharity,
	"channel:edit:commercial":           ScopeChannelEditCommercial,
	"channel:read:editors":              ScopeChannelReadEditors,
	"channel:manage:extensions":         ScopeChannelManageExtensions,
	"channel:read:goals":                ScopeChannelReadGoals,
	"channel:read:guest_star":           ScopeChannelReadGuestStar,
	"channel:manage:guest_star":         ScopeChannelManageGuestStar,
	"channel:read:hype_train":           ScopeChannelReadHypeTrain,
	"channel:manage:moderators":         ScopeChannelManageModerators,
	"channel:read:polls":                ScopeChannelReadPolls,
	"channel:manage:polls":              ScopeChannelManagePolls,
	"channel:read:predictions":          ScopeChannelReadPredictions,
	"channel:manage:predictions":        ScopeChannelManagePredictions,
	"channel:manage:raids":              ScopeChannelManageRaids,
	"channel:read:redemptions":          ScopeChannelReadRedemptions,
	"channel:manage:redemptions":        ScopeChannelManageRedemptions,
	"channel:manage:schedule":           ScopeChannelManageSchedule,
	"channel:read:stream_key":           ScopeChannelReadStreamKey,
	"channel:read:subscriptions":        ScopeChannelReadSubscriptions,
	"channel:manage:videos":             ScopeChannelManageVideos,
	"channel:read:vips":                 ScopeChannelReadVips,
	"channel:manage:vips":               ScopeChannelManageVips,
	"channel:moderate":                  ScopeChannelModerate,
	"clips:edit":                        ScopeClipsEdit,
	"moderation:read":                   ScopeModerationRead,
	"moderator:manage:announcements":    ScopeModeratorManageAnnouncements,
	"moderator:manage:automod":          ScopeModeratorManageAutomod,
	"moderator:read:automod_settings":   ScopeModeratorReadAutomodSettings,
	"moderator:manage:automod_settings": ScopeModeratorManageAutomodSettings,
	"moderator:read:banned_users":       ScopeModeratorReadBannedUsers,
	"moderator:manage:banned_users":     ScopeModeratorManageBannedUsers,
	"moderator:read:blocked_terms":      ScopeModeratorReadBlockedTerms,
	"moderator:read:chat_messages":      ScopeModeratorReadChatMessages,
	"moderator:manage:blocked_terms":    ScopeModeratorManageBlockedTerms,
	"moderator:manage:chat_messages":    ScopeModeratorManageChatMessages,
	"moderator:read:chat_settings":      ScopeModeratorReadChatSettings,
	"moderator:manage:chat_settings":    ScopeModeratorManageChatSettings,
	"moderator:read:chatters":           ScopeModeratorReadChatters,
	"moderator:read:followers":          ScopeModeratorReadFollowers,
	"moderator:read:guest_star":         ScopeModeratorReadGuestStar,
	"moderator:manage:guest_star":       ScopeModeratorManageGuestStar,
	"moderator:read:moderators":         ScopeModeratorReadModerators,
	"moderator:read:shield_mode":        ScopeModeratorReadShieldMode,
	"moderator:manage:shield_mode":      ScopeModeratorManageShieldMode,
	"moderator:read:shoutouts":          ScopeModeratorReadShoutouts,
	"moderator:manage:shoutouts":        ScopeModeratorManageShoutouts,
	"moderator:read:suspicious_users":   ScopeModeratorReadSuspiciousUsers,
	"moderator:read:unban_requests":     ScopeModeratorReadUnbanRequests,
	"moderator:manage:unban_requests":   ScopeModeratorManageUnbanRequests,
	"moderator:read:vips":               ScopeModeratorReadVips,
	"moderator:read:warnings":           ScopeModeratorReadWarnings,
	"moderator:manage:warnings":         ScopeModeratorManageWarnings,
	"user:bot":                          ScopeUserBot,
	"user:edit":                         ScopeUserEdit,
	"user:edit:broadcast":               ScopeUserEditBroadcast,
	"user:read:blocked_users":           ScopeUserReadBlockedUsers,
	"user:manage:blocked_users":         ScopeUserManageBlockedUsers,
	"user:read:broadcast":               ScopeUserReadBroadcast,
	"user:read:chat":                    ScopeUserReadChat,
	"user:manage:chat_color":            ScopeUserManageChatColor,
	"user:read:email":                   ScopeUserReadEmail,
	"user:read:emotes":                  ScopeUserReadEmotes,
	"user:read:follows":                 ScopeUserReadFollows,
	"user:read:moderated_channels":      ScopeUserReadModeratedChannels,
	"user:read:subscriptions":           ScopeUserReadSubscriptions,
	"user:read:whispers":                ScopeUserReadWhispers,
	"user:manage:whispers":              ScopeUserManageWhispers,
	"user:write:chat":                   ScopeUserWriteChat,
}

var ScopeTypeName = map[ScopeType]string{
	ScopeAnalyticsReadExtensions:        "analytics:read:extensions",
	ScopeAnalyticsReadGames:             "analytics:read:games",
	ScopeBitsRead:                       "bits:read",
	ScopeChannelBot:                     "channel:bot",
	ScopeChannelManageAds:               "channel:manage:ads",
	ScopeChannelReadAds:                 "channel:read:ads",
	ScopeChannelManageBroadcast:         "channel:manage:broadcast",
	ScopeChannelReadCharity:             "channel:read:charity",
	ScopeChannelEditCommercial:          "channel:edit:commercial",
	ScopeChannelReadEditors:             "channel:read:editors",
	ScopeChannelManageExtensions:        "channel:manage:extensions",
	ScopeChannelReadGoals:               "channel:read:goals",
	ScopeChannelReadGuestStar:           "channel:read:guest_star",
	ScopeChannelManageGuestStar:         "channel:manage:guest_star",
	ScopeChannelReadHypeTrain:           "channel:read:hype_train",
	ScopeChannelManageModerators:        "channel:manage:moderators",
	ScopeChannelReadPolls:               "channel:read:polls",
	ScopeChannelManagePolls:             "channel:manage:polls",
	ScopeChannelReadPredictions:         "channel:read:predictions",
	ScopeChannelManagePredictions:       "channel:manage:predictions",
	ScopeChannelManageRaids:             "channel:manage:raids",
	ScopeChannelReadRedemptions:         "channel:read:redemptions",
	ScopeChannelManageRedemptions:       "channel:manage:redemptions",
	ScopeChannelManageSchedule:          "channel:manage:schedule",
	ScopeChannelReadStreamKey:           "channel:read:stream_key",
	ScopeChannelReadSubscriptions:       "channel:read:subscriptions",
	ScopeChannelManageVideos:            "channel:manage:videos",
	ScopeChannelReadVips:                "channel:read:vips",
	ScopeChannelManageVips:              "channel:manage:vips",
	ScopeChannelModerate:                "channel:moderate",
	ScopeClipsEdit:                      "clips:edit",
	ScopeModerationRead:                 "moderation:read",
	ScopeModeratorManageAnnouncements:   "moderator:manage:announcements",
	ScopeModeratorManageAutomod:         "moderator:manage:automod",
	ScopeModeratorReadAutomodSettings:   "moderator:read:automod_settings",
	ScopeModeratorManageAutomodSettings: "moderator:manage:automod_settings",
	ScopeModeratorReadBannedUsers:       "moderator:read:banned_users",
	ScopeModeratorManageBannedUsers:     "moderator:manage:banned_users",
	ScopeModeratorReadBlockedTerms:      "moderator:read:blocked_terms",
	ScopeModeratorReadChatMessages:      "moderator:read:chat_messages",
	ScopeModeratorManageBlockedTerms:    "moderator:manage:blocked_terms",
	ScopeModeratorManageChatMessages:    "moderator:manage:chat_messages",
	ScopeModeratorReadChatSettings:      "moderator:read:chat_settings",
	ScopeModeratorManageChatSettings:    "moderator:manage:chat_settings",
	ScopeModeratorReadChatters:          "moderator:read:chatters",
	ScopeModeratorReadFollowers:         "moderator:read:followers",
	ScopeModeratorReadGuestStar:         "moderator:read:guest_star",
	ScopeModeratorManageGuestStar:       "moderator:manage:guest_star",
	ScopeModeratorReadModerators:        "moderator:read:moderators",
	ScopeModeratorReadShieldMode:        "moderator:read:shield_mode",
	ScopeModeratorManageShieldMode:      "moderator:manage:shield_mode",
	ScopeModeratorReadShoutouts:         "moderator:read:shoutouts",
	ScopeModeratorManageShoutouts:       "moderator:manage:shoutouts",
	ScopeModeratorReadSuspiciousUsers:   "moderator:read:suspicious_users",
	ScopeModeratorReadUnbanRequests:     "moderator:read:unban_requests",
	ScopeModeratorManageUnbanRequests:   "moderator:manage:unban_requests",
	ScopeModeratorReadVips:              "moderator:read:vips",
	ScopeModeratorReadWarnings:          "moderator:read:warnings",
	ScopeModeratorManageWarnings:        "moderator:manage:warnings",
	ScopeUserBot:                        "user:bot",
	ScopeUserEdit:                       "user:edit",
	ScopeUserEditBroadcast:              "user:edit:broadcast",
	ScopeUserReadBlockedUsers:           "user:read:blocked_users",
	ScopeUserManageBlockedUsers:         "user:manage:blocked_users",
	ScopeUserReadBroadcast:              "user:read:broadcast",
	ScopeUserReadChat:                   "user:read:chat",
	ScopeUserManageChatColor:            "user:manage:chat_color",
	ScopeUserReadEmail:                  "user:read:email",
	ScopeUserReadEmotes:                 "user:read:emotes",
	ScopeUserReadFollows:                "user:read:follows",
	ScopeUserReadModeratedChannels:      "user:read:moderated_channels",
	ScopeUserReadSubscriptions:          "user:read:subscriptions",
	ScopeUserReadWhispers:               "user:read:whispers",
	ScopeUserManageWhispers:             "user:manage:whispers",
	ScopeUserWriteChat:                  "user:write:chat",
}

func (t *ScopeType) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(ScopeTypeName[*t])
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

func (t *ScopeType) UnmarshalJSON(b []byte) error {
	var s string
	err := json.Unmarshal(b, &s)
	if err != nil {
		return err
	}

	*t = ScopeTypeId[s]
	return nil
}
