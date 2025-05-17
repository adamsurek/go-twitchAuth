package go_twitchAuth

import (
	"bytes"
	"encoding/json"
)

/*
ScopeType represents the level of access an app has on the Twitch API.

Full list of scopes: https://dev.twitch.tv/docs/authentication/scopes/
*/
type ScopeType int

const (
	// ScopeAnalyticsReadExtensions allows app to view analytics data for the Twitch extensions owned by the
	// authenticated account.
	ScopeAnalyticsReadExtensions ScopeType = iota + 1

	// ScopeAnalyticsReadGames allows app to view analytics data for the games owned by the authenticated account.
	ScopeAnalyticsReadGames

	// ScopeBitsRead allows app to view bits information for a channel.
	ScopeBitsRead

	// ScopeChannelBot allows app to join the user's channel as a bot user and perform chat-related actions.
	ScopeChannelBot

	// ScopeChannelManageAds allows app to manage the ads schedule on a channel.
	ScopeChannelManageAds

	// ScopeChannelReadAds allows app to read the ads schedule and details on the user's channel.
	ScopeChannelReadAds

	// ScopeChannelManageBroadcast allows app to manage a channel's broadcast config, including updating channel config
	// and managing stream markers and tags.
	ScopeChannelManageBroadcast

	// ScopeChannelReadCharity allows app to read charity campaign details and user donations on the user's channel.
	ScopeChannelReadCharity

	// ScopeChannelEditCommercial allows app to run commercials on a channel.
	ScopeChannelEditCommercial

	// ScopeChannelReadEditors allows app to view a list of editors in a channel.
	ScopeChannelReadEditors

	// ScopeChannelManageExtensions allows app to manage a channel's Extension config, including activation Extensions.
	ScopeChannelManageExtensions

	// ScopeChannelReadGoals allows app to view Creator Goals for a channel.
	ScopeChannelReadGoals

	// ScopeChannelReadGuestStar allows app to read Guest Star details for the user's channel.
	ScopeChannelReadGuestStar

	// ScopeChannelManageGuestStar allows app to manage Guest Star for the user's channel.
	ScopeChannelManageGuestStar

	// ScopeChannelReadHypeTrain allows app to view Hype Train information for a channel.
	ScopeChannelReadHypeTrain

	// ScopeChannelManageModerators allows app to add and remove moderators on the user's channel.
	ScopeChannelManageModerators

	// ScopeChannelReadPolls allows app to view a channel's polls.
	ScopeChannelReadPolls

	// ScopeChannelManagePolls allows app to manage a channel's polls.
	ScopeChannelManagePolls

	// ScopeChannelReadPredictions allows app to view a channel's Channel Point Predictions.
	ScopeChannelReadPredictions

	// ScopeChannelManagePredictions allows app to manage a channel's Channel Point Predictions.
	ScopeChannelManagePredictions

	// ScopeChannelManageRaids allows app to manage a channel raiding another channel.
	ScopeChannelManageRaids

	// ScopeChannelReadRedemptions allows app to view Channel Points custom rewards and their redemptions on a channel.
	ScopeChannelReadRedemptions

	// ScopeChannelManageRedemptions allows app to manage Channel Points custom rewards and their redemptions on a channel.
	ScopeChannelManageRedemptions

	// ScopeChannelManageSchedule allows app to manage a channel's stream schedule.
	ScopeChannelManageSchedule

	// ScopeChannelReadStreamKey allows app to view an authorized user's stream key.
	ScopeChannelReadStreamKey

	// ScopeChannelReadSubscriptions allows app to view a list of all subscribers to a channel, and check if the user is
	// subscribed to a channel.
	ScopeChannelReadSubscriptions

	// ScopeChannelManageVideos allows app to manage a channel's videos, including deleting videos.
	ScopeChannelManageVideos

	// ScopeChannelReadVips allows app to view a list of VIPs in the user's channel.
	ScopeChannelReadVips

	// ScopeChannelManageVips allows app to add and remove VIPs in the user's channel.
	ScopeChannelManageVips

	// ScopeChannelModerate allows app to perform moderation actions in a channel.
	ScopeChannelModerate

	// ScopeClipsEdit allows app to manage Clips for a channel.
	ScopeClipsEdit

	// ScopeModerationRead allows app to view moderation data including Moderators, Bans, Timeouts, and
	// Automod settings for channels where the authenticated user is a moderator.
	ScopeModerationRead

	// ScopeModeratorManageAnnouncements allows app to send announcements in channels where the authenticated user
	// is a moderator.
	ScopeModeratorManageAnnouncements

	// ScopeModeratorManageAutomod allows app to manage messages held for review by AutoMod in channels where
	// the user is a moderator.
	ScopeModeratorManageAutomod

	// ScopeModeratorReadAutomodSettings allows app to view a broadcaster's AutoMod settings for channels where the
	// user is a moderator.
	ScopeModeratorReadAutomodSettings

	// ScopeModeratorManageAutomodSettings allows app to manage a broadcaster's AutoMod settings for channels where
	// the user is a moderator.
	ScopeModeratorManageAutomodSettings

	// ScopeModeratorReadBannedUsers allows app to view a list of bans and unbans for channels where the authenticated
	//user is a moderator.
	ScopeModeratorReadBannedUsers

	// ScopeModeratorManageBannedUsers allows app to ban and unban users in channels where the authenticated user is
	// a moderator.
	ScopeModeratorManageBannedUsers

	// ScopeModeratorReadBlockedTerms allows app to view a broadcaster's list of blocked terms.
	ScopeModeratorReadBlockedTerms

	// ScopeModeratorReadChatMessages allows app to read deleted chat messages in a channel.
	ScopeModeratorReadChatMessages

	// ScopeModeratorManageBlockedTerms allows app to manage a broadcaster's list of blocked terms.
	ScopeModeratorManageBlockedTerms

	// ScopeModeratorManageChatMessages allows app to delete chat messages in channels where the authenticated user
	// is a moderator.
	ScopeModeratorManageChatMessages

	// ScopeModeratorReadChatSettings allows app to view a broadcaster's chat room settings.
	ScopeModeratorReadChatSettings

	// ScopeModeratorManageChatSettings allows app to manage a broadcaster's chat room settings.
	ScopeModeratorManageChatSettings

	// ScopeModeratorReadChatters allows app to view the chatters in a broadcaster's chatroom.
	ScopeModeratorReadChatters

	// ScopeModeratorReadFollowers allows app to view the followers of a broadcaster.
	ScopeModeratorReadFollowers

	// ScopeModeratorReadGuestStar allows app to view Guest Star details for channels where the authenticated user
	// is a Guest Star moderator.
	ScopeModeratorReadGuestStar

	// ScopeModeratorManageGuestStar allows app to manage Guest Star details for channels where the authenticated
	// user is a Guest Star moderator.
	ScopeModeratorManageGuestStar

	// ScopeModeratorReadModerators allows app to view a list of moderators in channels where the authenticated
	// user is a moderator.
	ScopeModeratorReadModerators

	// ScopeModeratorReadShieldMode allows app to view a broadcaster's Shield Mode status.
	ScopeModeratorReadShieldMode

	// ScopeModeratorManageShieldMode allows app to manage a broadcaster's Shield Mode status.
	ScopeModeratorManageShieldMode

	// ScopeModeratorReadShoutouts allows app to view a broadcaster's shoutouts.
	ScopeModeratorReadShoutouts

	// ScopeModeratorManageShoutouts allows app to manage a broadcaster's shoutouts.
	ScopeModeratorManageShoutouts

	// ScopeModeratorReadSuspiciousUsers allows app to view chat messages from suspicious users and see users flagged
	// as suspicious in channels where the authenticated user is a moderator.
	ScopeModeratorReadSuspiciousUsers

	// ScopeModeratorReadUnbanRequests allows app to view a broadcaster's unban requests.
	ScopeModeratorReadUnbanRequests

	// ScopeModeratorManageUnbanRequests allows app to manage a broadcaster's unban requests.
	ScopeModeratorManageUnbanRequests

	// ScopeModeratorReadVips allows app to view the list of VIPs for channels where the authenticated user is
	// a moderator.
	ScopeModeratorReadVips

	// ScopeModeratorReadWarnings allows app to view warnings in channels where the authenticated user is a moderator.
	ScopeModeratorReadWarnings

	// ScopeModeratorManageWarnings allows app to warn users in channels where the authenticated user is a moderator.
	ScopeModeratorManageWarnings

	// ScopeUserBot allows app to join a chat channel as the authenticated user but appearing as a bot and perform
	// actions as the user.
	ScopeUserBot

	// ScopeUserEdit allows app to update the authenticated user's information.
	ScopeUserEdit

	// ScopeUserEditBroadcast allows app to view and edit the authenticated user's broadcasting config, including
	// Extension configs.
	ScopeUserEditBroadcast

	// ScopeUserReadBlockedUsers allows app to view the authenticated user's block list.
	ScopeUserReadBlockedUsers

	// ScopeUserManageBlockedUsers allows app to manage the authenticated user's block list.
	ScopeUserManageBlockedUsers

	// ScopeUserReadBroadcast allows app to view the authenticated user's broadcasting config, including
	// Extension configs.
	ScopeUserReadBroadcast

	// ScopeUserReadChat allows app to receive chatroom messages and informational notifications related to a
	// channel's chatroom.
	ScopeUserReadChat

	// ScopeUserManageChatColor allows app to update the color used for the authenticated user's name in chat.
	ScopeUserManageChatColor

	// ScopeUserReadEmail allows app to view the authenticated user's email address,
	ScopeUserReadEmail

	// ScopeUserReadEmotes allows app to view the emotes available to the authenticated user.
	ScopeUserReadEmotes

	// ScopeUserReadFollows allows app to view the list of channels that the authenticated user follows.
	ScopeUserReadFollows

	// ScopeUserReadModeratedChannels allows app to view the list of channels where the authenticated user is a
	// moderator.
	ScopeUserReadModeratedChannels

	// ScopeUserReadSubscriptions allows app to view the list of channels that the authenticated user is subscribed to.
	ScopeUserReadSubscriptions

	// ScopeUserReadWhispers allows app to receive whispers sent to the authenticated user.
	ScopeUserReadWhispers

	// ScopeUserManageWhispers allows app to receive whispers sent to the authenticated user, and send whispers on their
	// behalf.
	ScopeUserManageWhispers

	// ScopeUserWriteChat allows app to send chat messages as the authenticated user.
	ScopeUserWriteChat
)

// scopeTypeId translates the string version of an access scope to its enum value.
var scopeTypeId = map[string]ScopeType{
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

// scopeTypeName translates the enum value version of an access scope to its string value.
var scopeTypeName = map[ScopeType]string{
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
	buffer.WriteString(scopeTypeName[*t])
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

func (t *ScopeType) UnmarshalJSON(b []byte) error {
	var s string
	err := json.Unmarshal(b, &s)
	if err != nil {
		return err
	}

	*t = scopeTypeId[s]
	return nil
}
