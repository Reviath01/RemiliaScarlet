package commands

import (
	ctx "git.randomchars.net/Reviath/RemiliaScarlet/CommandHandler/Context"
	embedutil "git.randomchars.net/Reviath/RemiliaScarlet/EmbedUtil"
	"git.randomchars.net/Reviath/RemiliaScarlet/sql"
	"github.com/bwmarrin/discordgo"
)

type Settings struct {
}

func (s Settings) Execute(ctx ctx.Ctx, session *discordgo.Session) error {
	db := sql.Connect()

	type Tag struct {
		welcomechannelid string
		leavechannelid   string
		roleid           string
		welcomemessage   string
		leavemessage     string
		logid            string
	}

	var tag Tag

	if sql.CheckLanguage(ctx.Guild().ID) == "tr" {
		if sql.IsBlocked(ctx.Guild().ID, "settings") == "true" {
			_, _ = session.ChannelMessageSend(ctx.Channel().ID, "Bu komut bu sunucuda engellenmiş.")
			return nil
		}

		err := db.QueryRow("SELECT channelid FROM welcomechannel WHERE guildid ='" + ctx.Guild().ID + "'").Scan(&tag.welcomechannelid)
		if err != nil {
			tag.welcomechannelid = "Ayarlanmamış."
		}

		err = db.QueryRow("SELECT channelid FROM leavechannel WHERE guildid ='" + ctx.Guild().ID + "'").Scan(&tag.leavechannelid)
		if err != nil {
			tag.leavechannelid = "Ayarlanmamış."
		}

		err = db.QueryRow("SELECT roleid FROM autorole WHERE guildid ='" + ctx.Guild().ID + "'").Scan(&tag.roleid)
		if err != nil {
			tag.roleid = "Ayarlanmamış."
		}

		err = db.QueryRow("SELECT message FROM welcomemessage WHERE guildid ='" + ctx.Guild().ID + "'").Scan(&tag.welcomemessage)
		if err != nil {
			tag.welcomemessage = "Ayarlanmamış."
		}

		err = db.QueryRow("SELECT message FROM leavemessage WHERE guildid ='" + ctx.Guild().ID + "'").Scan(&tag.leavemessage)
		if err != nil {
			tag.leavemessage = "Ayarlanmamış."
		}

		err = db.QueryRow("SELECT channelid FROM log WHERE guildid ='" + ctx.Guild().ID + "'").Scan(&tag.logid)
		if err != nil {
			tag.logid = "Ayarlanmamış."
		}

		embed := embedutil.NewEmbed().
			SetTitle(ctx.Guild().Name+" Ayarları").
			AddField("Hoş Geldin Kanalı:", tag.welcomechannelid).
			AddField("Çıkış Kanalı:", tag.leavechannelid).
			AddField("Otorol:", tag.roleid).
			AddField("Çıkış Mesajı:", tag.leavemessage).
			AddField("Hoş Geldin Mesajı:", tag.welcomemessage).
			AddField("Log Kanalı:", tag.logid).
			SetColor(0x00f0ff).MessageEmbed

		_, _ = session.ChannelMessageSendEmbed(ctx.Channel().ID, embed)

		return nil
	}

	if sql.IsBlocked(ctx.Guild().ID, "settings") == "true" {
		_, _ = session.ChannelMessageSend(ctx.Channel().ID, "This command is blocked on this guild.")
		return nil
	}
	err := db.QueryRow("SELECT channelid FROM welcomechannel WHERE guildid ='" + ctx.Guild().ID + "'").Scan(&tag.welcomechannelid)
	if err != nil {
		tag.welcomechannelid = "Not existing."
	}

	err = db.QueryRow("SELECT channelid FROM leavechannel WHERE guildid ='" + ctx.Guild().ID + "'").Scan(&tag.leavechannelid)
	if err != nil {
		tag.leavechannelid = "Not existing."
	}

	err = db.QueryRow("SELECT roleid FROM autorole WHERE guildid ='" + ctx.Guild().ID + "'").Scan(&tag.roleid)
	if err != nil {
		tag.roleid = "Not existing."
	}

	err = db.QueryRow("SELECT message FROM welcomemessage WHERE guildid ='" + ctx.Guild().ID + "'").Scan(&tag.welcomemessage)
	if err != nil {
		tag.welcomemessage = "Not existing."
	}

	err = db.QueryRow("SELECT message FROM leavemessage WHERE guildid ='" + ctx.Guild().ID + "'").Scan(&tag.leavemessage)
	if err != nil {
		tag.leavemessage = "Not existing."
	}

	err = db.QueryRow("SELECT channelid FROM log WHERE guildid ='" + ctx.Guild().ID + "'").Scan(&tag.logid)
	if err != nil {
		tag.logid = "Not existing."
	}

	embed := embedutil.NewEmbed().
		SetTitle(ctx.Guild().Name+" Settings").
		AddField("Welcome Channel:", tag.welcomechannelid).
		AddField("Leave Channel:", tag.leavechannelid).
		AddField("Autorole:", tag.roleid).
		AddField("Leave Message:", tag.leavemessage).
		AddField("Welcome Message:", tag.welcomemessage).
		AddField("Log Channel:", tag.logid).
		SetColor(0x00f0ff).MessageEmbed

	_, _ = session.ChannelMessageSendEmbed(ctx.Channel().ID, embed)

	return nil
}
