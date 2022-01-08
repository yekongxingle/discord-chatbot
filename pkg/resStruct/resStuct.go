package resStruct

import "time"

type AutoGenerated []struct {
	ID                string            `json:"id"`
	Type              int               `json:"type"`
	Content           string            `json:"content"`
	ChannelID         string            `json:"channel_id"`
	Author            Author            `json:"author"`
	Attachments       []interface{}     `json:"attachments"`
	Embeds            []interface{}     `json:"embeds"`
	Mentions          []Mentions        `json:"mentions"`
	MentionRoles      []interface{}     `json:"mention_roles"`
	Pinned            bool              `json:"pinned"`
	MentionEveryone   bool              `json:"mention_everyone"`
	Tts               bool              `json:"tts"`
	Timestamp         time.Time         `json:"timestamp"`
	EditedTimestamp   interface{}       `json:"edited_timestamp"`
	Flags             int               `json:"flags"`
	Components        []interface{}     `json:"components"`
	MessageReference  MessageReference  `json:"message_reference,omitempty"`
	ReferencedMessage ReferencedMessage `json:"referenced_message,omitempty"`
}
type Author struct {
	ID            string `json:"id"`
	Username      string `json:"username"`
	Avatar        string `json:"avatar"`
	Discriminator string `json:"discriminator"`
	PublicFlags   int    `json:"public_flags"`
}
type Mentions struct {
	ID            string `json:"id"`
	Username      string `json:"username"`
	Avatar        string `json:"avatar"`
	Discriminator string `json:"discriminator"`
	PublicFlags   int    `json:"public_flags"`
}
type MessageReference struct {
	ChannelID string `json:"channel_id"`
	GuildID   string `json:"guild_id"`
	MessageID string `json:"message_id"`
}
type ReferencedMessage struct {
	ID              string        `json:"id"`
	Type            int           `json:"type"`
	Content         string        `json:"content"`
	ChannelID       string        `json:"channel_id"`
	Author          Author        `json:"author"`
	Attachments     []interface{} `json:"attachments"`
	Embeds          []interface{} `json:"embeds"`
	Mentions        []interface{} `json:"mentions"`
	MentionRoles    []interface{} `json:"mention_roles"`
	Pinned          bool          `json:"pinned"`
	MentionEveryone bool          `json:"mention_everyone"`
	Tts             bool          `json:"tts"`
	Timestamp       time.Time     `json:"timestamp"`
	EditedTimestamp interface{}   `json:"edited_timestamp"`
	Flags           int           `json:"flags"`
	Components      []interface{} `json:"components"`
}
