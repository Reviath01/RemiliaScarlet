package commands

import (
	"fmt"
	"strconv"

	embedutil "git.randomchars.net/FreeNitori/EmbedUtil"
	CommandHandler "git.randomchars.net/Reviath/RemiliaScarlet/handler"
	"git.randomchars.net/Reviath/RemiliaScarlet/sql"
)

// GuildInfoCommand is a handler for guild info command
func GuildInfoCommand(ctx CommandHandler.Context, _ []string) error {
	switch sql.CheckLanguage(ctx.Guild.ID) {
	case "tr":
		if sql.IsBlocked(ctx.Guild.ID, "guild_info") == "true" {
			ctx.Reply("Bu komut bu sunucuda engellenmiş.")
			return nil
		}

		embed := embedutil.New("", "")
		embed.Color = 0xefff00
		embed.AddField("Sunucu İsmi", ctx.Guild.Name, true)
		embed.AddField("Kişi Sayısı", strconv.Itoa(ctx.Guild.MemberCount), true)
		embed.AddField("Bölge", ctx.Guild.Region, true)
		embed.AddField("Sunucu Sahibi", fmt.Sprintf("<@%s>", ctx.Guild.OwnerID), true)
		embed.AddField("Sunucu Sahibinin ID'si", ctx.Guild.OwnerID, true)
		embed.AddField("Afk Süresi", strconv.Itoa(ctx.Guild.AfkTimeout), true)
		embed.AddField("ID:", ctx.Guild.ID, true)
		embed.AddField("Yer", ctx.Guild.PreferredLocale, true)
		ctx.ReplyEmbed(embed)

		return nil

	default:

		if sql.IsBlocked(ctx.Guild.ID, "guild_info") == "true" {
			ctx.Reply("This command is blocked on this guild.")
			return nil
		}

		embed := embedutil.New("", "")
		embed.Color = 0xefff00
		embed.AddField("Guild Name", ctx.Guild.Name, true)
		embed.AddField("Member Count", strconv.Itoa(ctx.Guild.MemberCount), true)
		embed.AddField("Region", ctx.Guild.Region, true)
		embed.AddField("Guild Owner", fmt.Sprintf("<@%s>", ctx.Guild.OwnerID), true)
		embed.AddField("Guild Owner ID", ctx.Guild.OwnerID, true)
		embed.AddField("Afk Timeout", strconv.Itoa(ctx.Guild.AfkTimeout), true)
		embed.AddField("ID:", ctx.Guild.ID, true)
		embed.AddField("Locale", ctx.Guild.PreferredLocale, true)
		ctx.ReplyEmbed(embed)
		return nil
	}
}
