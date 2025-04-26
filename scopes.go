package go_twitchAuth

import (
	"bytes"
	"encoding/json"
)

type ScopeType int

const (
	AnalyticsReadExtensions ScopeType = iota + 1
	AnalyticsReadGames
	BitsRead
	ChannelBot
	ChannelManageAds
	ChannelReadAds
	ChannelManageBroadcast
	ChannelReadCharity
	ChannelEditCommercial
	ChannelReadEditors
	ChannelManageExtensions
	ChannelReadGoals
	ChannelReadGuestStar
	ChannelManageGuestStar
	ChannelReadHypeTrain
	ChannelManageModerators
	ChannelReadPolls
	ChannelManagePolls
	ChannelReadPredictions
	ChannelManagePredictions
	ChannelManageRaids
	ChannelReadRedemptions
	ChannelManageRedemptions
	ChannelManageSchedule
	ChannelReadStreamKey
	ChannelReadSubscriptions
	ChannelManageVideos
	ChannelReadVips
	ChannelManageVips
	ChannelModerate
	ClipsEdit
	ModerationRead
	ModeratorManageAnnouncements
	ModeratorManageAutomod
	ModeratorReadAutomodSettings
	ModeratorManageAutomodSettings
	ModeratorReadBannedUsers
	ModeratorManageBannedUsers
	ModeratorReadBlockedTerms
	ModeratorReadChatMessages
	ModeratorManageBlockedTerms
	ModeratorManageChatMessages
	ModeratorReadChatSettings
	ModeratorManageChatSettings
	ModeratorReadChatters
	ModeratorReadFollowers
	ModeratorReadGuestStar
	ModeratorManageGuestStar
	ModeratorReadModerators
	ModeratorReadShieldMode
	ModeratorManageShieldMode
	ModeratorReadShoutouts
	ModeratorManageShoutouts
	ModeratorReadSuspiciousUsers
	ModeratorReadUnbanRequests
	ModeratorManageUnbanRequests
	ModeratorReadVips
	ModeratorReadWarnings
	ModeratorManageWarnings
	UserBot
	UserEdit
	UserEditBroadcast
	UserReadBlockedUsers
	UserManageBlockedUsers
	UserReadBroadcast
	UserReadChat
	UserManageChatColor
	UserReadEmail
	UserReadEmotes
	UserReadFollows
	UserReadModeratedChannels
	UserReadSubscriptions
	UserReadWhispers
	UserManageWhispers
	UserWriteChat
)

var ScopeTypeId = map[string]ScopeType{
	"analytics:read:extensions":         AnalyticsReadExtensions,
	"analytics:read:games":              AnalyticsReadGames,
	"bits:read":                         BitsRead,
	"channel:bot":                       ChannelBot,
	"channel:manage:ads":                ChannelManageAds,
	"channel:read:ads":                  ChannelReadAds,
	"channel:manage:broadcast":          ChannelManageBroadcast,
	"channel:read:charity":              ChannelReadCharity,
	"channel:edit:commercial":           ChannelEditCommercial,
	"channel:read:editors":              ChannelReadEditors,
	"channel:manage:extensions":         ChannelManageExtensions,
	"channel:read:goals":                ChannelReadGoals,
	"channel:read:guest_star":           ChannelReadGuestStar,
	"channel:manage:guest_star":         ChannelManageGuestStar,
	"channel:read:hype_train":           ChannelReadHypeTrain,
	"channel:manage:moderators":         ChannelManageModerators,
	"channel:read:polls":                ChannelReadPolls,
	"channel:manage:polls":              ChannelManagePolls,
	"channel:read:predictions":          ChannelReadPredictions,
	"channel:manage:predictions":        ChannelManagePredictions,
	"channel:manage:raids":              ChannelManageRaids,
	"channel:read:redemptions":          ChannelReadRedemptions,
	"channel:manage:redemptions":        ChannelManageRedemptions,
	"channel:manage:schedule":           ChannelManageSchedule,
	"channel:read:stream_key":           ChannelReadStreamKey,
	"channel:read:subscriptions":        ChannelReadSubscriptions,
	"channel:manage:videos":             ChannelManageVideos,
	"channel:read:vips":                 ChannelReadVips,
	"channel:manage:vips":               ChannelManageVips,
	"channel:moderate":                  ChannelModerate,
	"clips:edit":                        ClipsEdit,
	"moderation:read":                   ModerationRead,
	"moderator:manage:announcements":    ModeratorManageAnnouncements,
	"moderator:manage:automod":          ModeratorManageAutomod,
	"moderator:read:automod_settings":   ModeratorReadAutomodSettings,
	"moderator:manage:automod_settings": ModeratorManageAutomodSettings,
	"moderator:read:banned_users":       ModeratorReadBannedUsers,
	"moderator:manage:banned_users":     ModeratorManageBannedUsers,
	"moderator:read:blocked_terms":      ModeratorReadBlockedTerms,
	"moderator:read:chat_messages":      ModeratorReadChatMessages,
	"moderator:manage:blocked_terms":    ModeratorManageBlockedTerms,
	"moderator:manage:chat_messages":    ModeratorManageChatMessages,
	"moderator:read:chat_settings":      ModeratorReadChatSettings,
	"moderator:manage:chat_settings":    ModeratorManageChatSettings,
	"moderator:read:chatters":           ModeratorReadChatters,
	"moderator:read:followers":          ModeratorReadFollowers,
	"moderator:read:guest_star":         ModeratorReadGuestStar,
	"moderator:manage:guest_star":       ModeratorManageGuestStar,
	"moderator:read:moderators":         ModeratorReadModerators,
	"moderator:read:shield_mode":        ModeratorReadShieldMode,
	"moderator:manage:shield_mode":      ModeratorManageShieldMode,
	"moderator:read:shoutouts":          ModeratorReadShoutouts,
	"moderator:manage:shoutouts":        ModeratorManageShoutouts,
	"moderator:read:suspicious_users":   ModeratorReadSuspiciousUsers,
	"moderator:read:unban_requests":     ModeratorReadUnbanRequests,
	"moderator:manage:unban_requests":   ModeratorManageUnbanRequests,
	"moderator:read:vips":               ModeratorReadVips,
	"moderator:read:warnings":           ModeratorReadWarnings,
	"moderator:manage:warnings":         ModeratorManageWarnings,
	"user:bot":                          UserBot,
	"user:edit":                         UserEdit,
	"user:edit:broadcast":               UserEditBroadcast,
	"user:read:blocked_users":           UserReadBlockedUsers,
	"user:manage:blocked_users":         UserManageBlockedUsers,
	"user:read:broadcast":               UserReadBroadcast,
	"user:read:chat":                    UserReadChat,
	"user:manage:chat_color":            UserManageChatColor,
	"user:read:email":                   UserReadEmail,
	"user:read:emotes":                  UserReadEmotes,
	"user:read:follows":                 UserReadFollows,
	"user:read:moderated_channels":      UserReadModeratedChannels,
	"user:read:subscriptions":           UserReadSubscriptions,
	"user:read:whispers":                UserReadWhispers,
	"user:manage:whispers":              UserManageWhispers,
	"user:write:chat":                   UserWriteChat,
}

var ScopeTypeName = map[ScopeType]string{
	AnalyticsReadExtensions:        "analytics:read:extensions",
	AnalyticsReadGames:             "analytics:read:games",
	BitsRead:                       "bits:read",
	ChannelBot:                     "channel:bot",
	ChannelManageAds:               "channel:manage:ads",
	ChannelReadAds:                 "channel:read:ads",
	ChannelManageBroadcast:         "channel:manage:broadcast",
	ChannelReadCharity:             "channel:read:charity",
	ChannelEditCommercial:          "channel:edit:commercial",
	ChannelReadEditors:             "channel:read:editors",
	ChannelManageExtensions:        "channel:manage:extensions",
	ChannelReadGoals:               "channel:read:goals",
	ChannelReadGuestStar:           "channel:read:guest_star",
	ChannelManageGuestStar:         "channel:manage:guest_star",
	ChannelReadHypeTrain:           "channel:read:hype_train",
	ChannelManageModerators:        "channel:manage:moderators",
	ChannelReadPolls:               "channel:read:polls",
	ChannelManagePolls:             "channel:manage:polls",
	ChannelReadPredictions:         "channel:read:predictions",
	ChannelManagePredictions:       "channel:manage:predictions",
	ChannelManageRaids:             "channel:manage:raids",
	ChannelReadRedemptions:         "channel:read:redemptions",
	ChannelManageRedemptions:       "channel:manage:redemptions",
	ChannelManageSchedule:          "channel:manage:schedule",
	ChannelReadStreamKey:           "channel:read:stream_key",
	ChannelReadSubscriptions:       "channel:read:subscriptions",
	ChannelManageVideos:            "channel:manage:videos",
	ChannelReadVips:                "channel:read:vips",
	ChannelManageVips:              "channel:manage:vips",
	ChannelModerate:                "channel:moderate",
	ClipsEdit:                      "clips:edit",
	ModerationRead:                 "moderation:read",
	ModeratorManageAnnouncements:   "moderator:manage:announcements",
	ModeratorManageAutomod:         "moderator:manage:automod",
	ModeratorReadAutomodSettings:   "moderator:read:automod_settings",
	ModeratorManageAutomodSettings: "moderator:manage:automod_settings",
	ModeratorReadBannedUsers:       "moderator:read:banned_users",
	ModeratorManageBannedUsers:     "moderator:manage:banned_users",
	ModeratorReadBlockedTerms:      "moderator:read:blocked_terms",
	ModeratorReadChatMessages:      "moderator:read:chat_messages",
	ModeratorManageBlockedTerms:    "moderator:manage:blocked_terms",
	ModeratorManageChatMessages:    "moderator:manage:chat_messages",
	ModeratorReadChatSettings:      "moderator:read:chat_settings",
	ModeratorManageChatSettings:    "moderator:manage:chat_settings",
	ModeratorReadChatters:          "moderator:read:chatters",
	ModeratorReadFollowers:         "moderator:read:followers",
	ModeratorReadGuestStar:         "moderator:read:guest_star",
	ModeratorManageGuestStar:       "moderator:manage:guest_star",
	ModeratorReadModerators:        "moderator:read:moderators",
	ModeratorReadShieldMode:        "moderator:read:shield_mode",
	ModeratorManageShieldMode:      "moderator:manage:shield_mode",
	ModeratorReadShoutouts:         "moderator:read:shoutouts",
	ModeratorManageShoutouts:       "moderator:manage:shoutouts",
	ModeratorReadSuspiciousUsers:   "moderator:read:suspicious_users",
	ModeratorReadUnbanRequests:     "moderator:read:unban_requests",
	ModeratorManageUnbanRequests:   "moderator:manage:unban_requests",
	ModeratorReadVips:              "moderator:read:vips",
	ModeratorReadWarnings:          "moderator:read:warnings",
	ModeratorManageWarnings:        "moderator:manage:warnings",
	UserBot:                        "user:bot",
	UserEdit:                       "user:edit",
	UserEditBroadcast:              "user:edit:broadcast",
	UserReadBlockedUsers:           "user:read:blocked_users",
	UserManageBlockedUsers:         "user:manage:blocked_users",
	UserReadBroadcast:              "user:read:broadcast",
	UserReadChat:                   "user:read:chat",
	UserManageChatColor:            "user:manage:chat_color",
	UserReadEmail:                  "user:read:email",
	UserReadEmotes:                 "user:read:emotes",
	UserReadFollows:                "user:read:follows",
	UserReadModeratedChannels:      "user:read:moderated_channels",
	UserReadSubscriptions:          "user:read:subscriptions",
	UserReadWhispers:               "user:read:whispers",
	UserManageWhispers:             "user:manage:whispers",
	UserWriteChat:                  "user:write:chat",
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
