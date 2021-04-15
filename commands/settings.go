package commands

import (
	embedutil "git.randomchars.net/Reviath/RemiliaScarlet/EmbedUtil"
	CommandHandler "git.randomchars.net/Reviath/RemiliaScarlet/handler"
	"git.randomchars.net/Reviath/RemiliaScarlet/sql"
)

func SettingsCommand(ctx CommandHandler.Context, _ []string) error {
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
	var welcomechannel string
	var leavechannel string
	var autorole string
	var leavemsg string
	var welcomemsg string
	var logchannel string

	if sql.CheckLanguage(ctx.Guild.ID) == "tr" {
		if sql.IsBlocked(ctx.Guild.ID, "settings") == "true" {
			ctx.Reply("Bu komut bu sunucuda engellenmiş.")
			return nil
		}

		err := db.QueryRow("SELECT channelid FROM welcomechannel WHERE guildid ='" + ctx.Guild.ID + "'").Scan(&tag.welcomechannelid)
		if err == nil {
			welcomechannel = "<#" + tag.welcomechannelid + "> (" + tag.welcomechannelid + ")"
		} else {
			welcomechannel = "Ayarlanmamış."
		}

		err = db.QueryRow("SELECT channelid FROM leavechannel WHERE guildid ='" + ctx.Guild.ID + "'").Scan(&tag.leavechannelid)
		if err == nil {
			leavechannel = "<#" + tag.leavechannelid + "> (" + tag.leavechannelid + ")"
		} else {
			leavechannel = "Ayarlanmamış."
		}

		err = db.QueryRow("SELECT roleid FROM autorole WHERE guildid ='" + ctx.Guild.ID + "'").Scan(&tag.roleid)
		if err == nil {
			autorole = "<@&" + tag.roleid + "> (" + tag.roleid + ")"
		} else {
			autorole = "Ayarlanmamış."
		}

		err = db.QueryRow("SELECT message FROM welcomemessage WHERE guildid ='" + ctx.Guild.ID + "'").Scan(&tag.welcomemessage)
		if err == nil {
			welcomemsg = tag.welcomemessage
		} else {
			welcomemsg = "Ayarlanmamış."
		}

		err = db.QueryRow("SELECT message FROM leavemessage WHERE guildid ='" + ctx.Guild.ID + "'").Scan(&tag.leavemessage)
		if err == nil {
			leavemsg = tag.leavemessage
		} else {
			leavemsg = "Ayarlanmamış."
		}

		err = db.QueryRow("SELECT channelid FROM log WHERE guildid ='" + ctx.Guild.ID + "'").Scan(&tag.logid)
		if err == nil {
			logchannel = "<#" + tag.logid + "> (" + tag.logid + ")"
		} else {
			logchannel = "Ayarlanmamış."
		}

		embed := embedutil.NewEmbed().
			SetTitle(ctx.Guild.Name+" Settings").
			AddField("Hoş Geldin Kanalı:", welcomechannel).
			AddField("Çıkış Kanalı:", leavechannel).
			AddField("Otorol:", autorole).
			AddField("Çıkış Mesajı:", leavemsg).
			AddField("Hoş Geldin Mesajı:", welcomemsg).
			AddField("Log Kanalı:", logchannel).
			SetColor(0x00f0ff).MessageEmbed

		ctx.ReplyEmbed(embed)

		return nil
	}

	if sql.IsBlocked(ctx.Guild.ID, "settings") == "true" {
		ctx.Reply("This command is blocked on this guild.")
		return nil
	}

	err := db.QueryRow("SELECT channelid FROM welcomechannel WHERE guildid ='" + ctx.Guild.ID + "'").Scan(&tag.welcomechannelid)
	if err == nil {
		welcomechannel = "<#" + tag.welcomechannelid + "> (" + tag.welcomechannelid + ")"
	} else {
		welcomechannel = "Not existing."
	}

	err = db.QueryRow("SELECT channelid FROM leavechannel WHERE guildid ='" + ctx.Guild.ID + "'").Scan(&tag.leavechannelid)
	if err == nil {
		leavechannel = "<#" + tag.leavechannelid + "> (" + tag.leavechannelid + ")"
	} else {
		leavechannel = "Not existing."
	}

	err = db.QueryRow("SELECT roleid FROM autorole WHERE guildid ='" + ctx.Guild.ID + "'").Scan(&tag.roleid)
	if err == nil {
		autorole = "<@&" + tag.roleid + "> (" + tag.roleid + ")"
	} else {
		autorole = "Not existing."
	}

	err = db.QueryRow("SELECT message FROM welcomemessage WHERE guildid ='" + ctx.Guild.ID + "'").Scan(&tag.welcomemessage)
	if err == nil {
		welcomemsg = tag.welcomemessage
	} else {
		welcomemsg = "Not existing."
	}

	err = db.QueryRow("SELECT message FROM leavemessage WHERE guildid ='" + ctx.Guild.ID + "'").Scan(&tag.leavemessage)
	if err == nil {
		leavemsg = tag.leavemessage
	} else {
		leavemsg = "Not existing."
	}

	err = db.QueryRow("SELECT channelid FROM log WHERE guildid ='" + ctx.Guild.ID + "'").Scan(&tag.logid)
	if err == nil {
		logchannel = "<#" + tag.logid + "> (" + tag.logid + ")"
	} else {
		logchannel = "Not existing."
	}

	embed := embedutil.NewEmbed().
		SetTitle(ctx.Guild.Name+" Settings").
		AddField("Welcome Channel:", welcomechannel).
		AddField("Leave Channel:", leavechannel).
		AddField("Autorole:", autorole).
		AddField("Leave Message:", leavemsg).
		AddField("Welcome Message:", welcomemsg).
		AddField("Log Channel:", logchannel).
		SetColor(0x00f0ff).MessageEmbed

	ctx.ReplyEmbed(embed)

	return nil
}
