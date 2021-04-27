package commands

import (
	"fmt"
	"strconv"

	embedutil "git.randomchars.net/Reviath/RemiliaScarlet/EmbedUtil"
	CommandHandler "git.randomchars.net/Reviath/RemiliaScarlet/handler"
	"git.randomchars.net/Reviath/RemiliaScarlet/sql"
)

func GuildInfoCommand(ctx CommandHandler.Context, _ []string) error {
	if sql.CheckLanguage(ctx.Guild.ID) == "tr" {
		if sql.IsBlocked(ctx.Guild.ID, "guild_info") == "true" {
			ctx.Reply("Bu komut bu sunucuda engellenmiş.")
			return nil
		}

		embed := embedutil.NewEmbed().
			SetColor(0xefff00).
			AddField("Sunucu İsmi", ctx.Guild.Name).
			AddField("Kişi Sayısı", strconv.Itoa(ctx.Guild.MemberCount)).
			AddField("Bölge", ctx.Guild.Region).
			AddField("Sunucu Sahibi", fmt.Sprintf("<@%s>", ctx.Guild.OwnerID)).
			AddField("Sunucu Sahibinin ID'si", ctx.Guild.OwnerID).
			AddField("Afk Süresi", strconv.Itoa(ctx.Guild.AfkTimeout)).
			AddField("ID:", ctx.Guild.ID).
			AddField("Yer", ctx.Guild.PreferredLocale).MessageEmbed
		ctx.ReplyEmbed(embed)

		return nil
	}

	if sql.IsBlocked(ctx.Guild.ID, "guild_info") == "true" {
		ctx.Reply("This command is blocked on this guild.")
		return nil
	}

	embed := embedutil.NewEmbed().
		SetColor(0xefff00).
		AddField("Guild Name", ctx.Guild.Name).
		AddField("Member Count", strconv.Itoa(ctx.Guild.MemberCount)).
		AddField("Region", ctx.Guild.Region).
		AddField("Guild Owner", fmt.Sprintf("<@%s>", ctx.Guild.OwnerID)).
		AddField("Guild Owner ID", ctx.Guild.OwnerID).
		AddField("Afk Timeout", strconv.Itoa(ctx.Guild.AfkTimeout)).
		AddField("ID:", ctx.Guild.ID).
		AddField("Locale", ctx.Guild.PreferredLocale).MessageEmbed
	ctx.ReplyEmbed(embed)

	return nil
}
