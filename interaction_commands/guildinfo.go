package interaction_commands

import (
	"fmt"
	"strconv"

	embedutil "git.randomchars.net/FreeNitori/EmbedUtil"
	multiplexer "git.randomchars.net/Reviath/RemiliaScarlet/Multiplexer"
	"git.randomchars.net/Reviath/RemiliaScarlet/interactions"
	"git.randomchars.net/Reviath/RemiliaScarlet/sql"
	"github.com/bwmarrin/discordgo"
)

func GuildInfoCommand(session *discordgo.Session, interaction interactions.Interaction) interactions.InteractionResponse {
	Guild, _ := session.Guild(string(interaction.GuildID))
	switch sql.CheckLanguage(interaction.GuildID) {
	case "tr":
		if sql.IsBlocked(interaction.GuildID, "guild_info") == "true" {
			return multiplexer.CreateResponse("Bu komut bu sunucuda engellenmiş.")
		}

		embed := embedutil.New("", "")
		embed.Color = 0xefff00
		embed.AddField("Sunucu İsmi", Guild.Name, true)
		embed.AddField("Kişi Sayısı", strconv.Itoa(Guild.MemberCount), true)
		embed.AddField("Bölge", Guild.Region, true)
		embed.AddField("Sunucu Sahibi", fmt.Sprintf("<@%s>", Guild.OwnerID), true)
		embed.AddField("Sunucu Sahibinin ID'si", Guild.OwnerID, true)
		embed.AddField("Afk Süresi", strconv.Itoa(Guild.AfkTimeout), true)
		embed.AddField("ID:", interaction.GuildID, true)
		embed.AddField("Yer", Guild.PreferredLocale, true)
		return multiplexer.CreateEmbedResponse(embed)
	default:

		if sql.IsBlocked(interaction.GuildID, "guild_info") == "true" {
			return multiplexer.CreateResponse("This command is blocked on this guild.")
		}

		embed := embedutil.New("", "")
		embed.Color = 0xefff00
		embed.AddField("Guild Name", Guild.Name, true)
		embed.AddField("Member Size", strconv.Itoa(Guild.MemberCount), true)
		embed.AddField("Region", Guild.Region, true)
		embed.AddField("Server Owner", fmt.Sprintf("<@%s>", Guild.OwnerID), true)
		embed.AddField("Server Owner ID", Guild.OwnerID, true)
		embed.AddField("AFK Timeout", strconv.Itoa(Guild.AfkTimeout), true)
		embed.AddField("ID", interaction.GuildID, true)
		embed.AddField("Locale", Guild.PreferredLocale, true)
		return multiplexer.CreateEmbedResponse(embed)
	}
}
