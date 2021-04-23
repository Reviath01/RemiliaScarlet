package interaction_commands

import (
	"fmt"
	"strconv"

	embedutil "git.randomchars.net/Reviath/RemiliaScarlet/EmbedUtil"
	multiplexer "git.randomchars.net/Reviath/RemiliaScarlet/Multiplexer"
	"git.randomchars.net/Reviath/RemiliaScarlet/interactions"
	"git.randomchars.net/Reviath/RemiliaScarlet/sql"
	"github.com/bwmarrin/discordgo"
)

func GuildInfoCommand(session *discordgo.Session, interaction interactions.Interaction) interactions.InteractionResponse {
	Guild, _ := session.Guild(string(interaction.GuildID))
	if sql.CheckLanguage(interaction.GuildID) == "tr" {
		if sql.IsBlocked(interaction.GuildID, "guild_info") == "true" {
			return multiplexer.CreateResponse("Bu komut bu sunucuda engellenmiş.")
		}

		embed := embedutil.NewEmbed().
			SetColor(0xefff00).
			AddField("Sunucu İsmi", Guild.Name).
			AddField("Kişi Sayısı", strconv.Itoa(Guild.MemberCount)).
			AddField("Bölge", Guild.Region).
			AddField("Sunucu Sahibi", fmt.Sprintf("<@%s>", Guild.OwnerID)).
			AddField("Sunucu Sahibinin ID'si", Guild.OwnerID).
			AddField("Afk Süresi", strconv.Itoa(Guild.AfkTimeout)).
			AddField("ID:", interaction.GuildID).
			AddField("Yer", Guild.PreferredLocale).MessageEmbed
		return multiplexer.CreateEmbedResponse(embed)
	}

	if sql.IsBlocked(interaction.GuildID, "guild_info") == "true" {
		return multiplexer.CreateResponse("This command is blocked on this guild.")
	}

	embed := embedutil.NewEmbed().
		SetColor(0xefff00).
		AddField("Guild Name", Guild.Name).
		AddField("Member Count", strconv.Itoa(Guild.MemberCount)).
		AddField("Region", Guild.Region).
		AddField("Guild Owner", fmt.Sprintf("<@%s>", Guild.OwnerID)).
		AddField("Guild Owner ID", Guild.OwnerID).
		AddField("Afk Timeout", strconv.Itoa(Guild.AfkTimeout)).
		AddField("ID:", interaction.GuildID).
		AddField("Locale", Guild.PreferredLocale).MessageEmbed
	return multiplexer.CreateEmbedResponse(embed)
}
