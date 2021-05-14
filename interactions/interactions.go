package interactions

import (
	"encoding/json"

	"github.com/bwmarrin/discordgo"
)

var (
	EndpointApplicationNoOauth       = func(aID string) string { return discordgo.EndpointAPI + "/applications/" + aID }
	EndpointApplicationCommands      = func(aID string) string { return EndpointApplicationNoOauth(aID) + "/commands" }
	EndpointApplicationGuildCommands = func(aID string, gID string) string {
		return EndpointApplicationNoOauth(aID) + "/guilds/" + gID + "/commands"
	}
	EndpointInteractionResponse = func(iID string, iToken string) string {
		return discordgo.EndpointAPI + "/interactions/" + iID + "/" + iToken + "/callback"
	}
)

type Command struct {
	Name        string          `json:"name"`
	Description string          `json:"description"`
	Options     []CommandOption `json:"options,omitempty"`
}

type CommandOptionType int

const (
	OptionTypeSubCommand CommandOptionType = iota + 1
	OptionTypeSubCommandGroup
	OptionTypeString
	OptionTypeInteger
	OptionTypeBoolean
	OptionTypeUser
	OptionTypeChannel
	OptionTypeRole
)

type CommandOption struct {
	Type        CommandOptionType     `json:"type"`
	Name        string                `json:"name"`
	Description string                `json:"description"`
	Default     bool                  `json:"default"`
	Required    bool                  `json:"required"`
	Choices     []CommandOptionChoice `json:"choices,omitempty"`
	Options     []CommandOption       `json:"options,omitempty"`
}

type CommandOptionChoice struct {
	Name  string      `json:"name"`
	Value interface{} `json:"value"`
}

// GlobalCommandCreate returns an error, it posts your command to all servers your bot is in.
func GlobalCommandCreate(s *discordgo.Session, aID string, command Command) (err error) {
	_, err = s.RequestWithBucketID("POST", EndpointApplicationCommands(aID), command, EndpointApplicationCommands(""))
	return
}

// GuildCommandCreate returns an error, it posts your command to specified guild.
func GuildCommandCreate(s *discordgo.Session, aID string, gID string, command Command) (err error) {
	_, err = s.RequestWithBucketID("POST", EndpointApplicationGuildCommands(aID, gID), command, EndpointApplicationGuildCommands("", ""))
	return
}

type InteractionType int

const (
	InteractionTypePing InteractionType = iota + 1
	InteractionTypeApplicationCommand
)

type InteractionData struct {
	ID      string                  `json:"id"`
	Name    string                  `json:"name"`
	Options []InteractionDataOption `json:"options"`
}

type InteractionDataOption struct {
	Name    string                  `json:"name"`
	Value   interface{}             `json:"value"`
	Options []InteractionDataOption `json:"options"`
}

type Interaction struct {
	ID        string           `json:"id"`
	Type      InteractionType  `json:"type"`
	Data      InteractionData  `json:"data"`
	GuildID   string           `json:"guild_id"`
	ChannelID string           `json:"channel_id"`
	Member    discordgo.Member `json:"member"`
	Token     string           `json:"token"`
	Version   int              `json:"version"`
}

func InteractionFromRaw(raw json.RawMessage) (inter Interaction, err error) {
	err = json.Unmarshal(raw, &inter)
	return
}

type InteractionResponseType int

const (
	InteractionResponseTypePong InteractionResponseType = iota + 1
	InteractionResponseTypeAcknowledge
	InteractionResponseTypeChannelMessage
	InteractionResponseTypeChannelMessageWithSource
	InteractionResponseTypeACKWithSource
)

type InteractionResponseData struct {
	TTS             bool                         `json:"tts"`
	Content         string                       `json:"content"`
	Embeds          []discordgo.MessageEmbed     `json:"embeds,omitempty"`
	AllowedMentions discordgo.AllowedMentionType `json:"allowed_mentions,omitempty"`
}

type InteractionResponse struct {
	Type InteractionResponseType `json:"type"`
	Data InteractionResponseData `json:"data,omitempty"`
}

// Respond send a response to interaction, returns to error
func (interaction *Interaction) Respond(s *discordgo.Session, resp InteractionResponse) (err error) {
	_, err = s.RequestWithBucketID("POST", EndpointInteractionResponse(interaction.ID, interaction.Token), resp, EndpointInteractionResponse("", ""))
	return
}
