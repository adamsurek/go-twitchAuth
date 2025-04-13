package go_twitchAuth

import (
	"bytes"
	"encoding/json"
)

type ScopeType int

const (
	UserManageBlockedUsers ScopeType = iota + 1
	UserReadChat
	UserReadEmotes
	UserReadFollows
	UserReadModeratedChannels
	UserReadSubscriptions
	UserReadWhispers
	UserManageWhispers
	UserWriteChat
)

var ScopeTypeId = map[string]ScopeType{
	"user:manage:blocked_users":    UserManageBlockedUsers,
	"user:read:chat":               UserReadChat,
	"user:read:emotes":             UserReadEmotes,
	"user:read:follows":            UserReadFollows,
	"user:read:moderated_channels": UserReadModeratedChannels,
	"user:read:subscriptions":      UserReadSubscriptions,
	"user:read:whispers":           UserReadWhispers,
	"user:manage:whispers":         UserManageWhispers,
	"user:write:chat":              UserWriteChat,
}

var ScopeTypeName = map[ScopeType]string{
	UserManageBlockedUsers:    "user:manage:blocked_users",
	UserReadChat:              "user:read:chat",
	UserReadEmotes:            "user:read:emotes",
	UserReadFollows:           "user:read:follows",
	UserReadModeratedChannels: "user:read:moderated_channels",
	UserReadSubscriptions:     "user:read:subscriptions",
	UserReadWhispers:          "user:read:whispers",
	UserManageWhispers:        "user:manage:whispers",
	UserWriteChat:             "user:write:chat",
}

func (t *ScopeType) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(ScopeTypeName[*t])
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

func (t *ScopeType) UnmarshalJSON(b []byte) error {
	var s string
	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*t = ScopeTypeId[s]
	return nil
}
