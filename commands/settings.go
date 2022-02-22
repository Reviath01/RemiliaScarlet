package commands

import (
	"fmt"

	CommandHandler "git.randomchars.net/Reviath/RemiliaScarlet/handler"
	"git.randomchars.net/Reviath/RemiliaScarlet/sql"
	embedutil "git.randomchars.net/freenitori/embedutil"
)

// SettingsCommand is a handler for settings command
func SettingsCommand(ctx CommandHandler.Context, _ []string) error {
	db := sql.Connect()
	defer db.Close()
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

	switch sql.CheckLanguage(ctx.Guild.ID) {
	case "tr":
		if sql.IsBlocked(ctx.Guild.ID, "settings") == "true" {
			ctx.Reply("Bu komut bu sunucuda engellenmiş.")
			return nil
		}

		err := db.QueryRow(fmt.Sprintf("SELECT channelid FROM welcomechannel WHERE guildid ='%s'", ctx.Guild.ID)).Scan(&tag.welcomechannelid)
		if err == nil {
			welcomechannel = fmt.Sprintf("<#%s> (%s)", tag.welcomechannelid, tag.welcomechannelid)
		} else {
			welcomechannel = "Ayarlanmamış."
		}

		err = db.QueryRow(fmt.Sprintf("SELECT channelid FROM leavechannel WHERE guildid ='%s'", ctx.Guild.ID)).Scan(&tag.leavechannelid)
		if err == nil {
			leavechannel = fmt.Sprintf("<#%s> (%s)", tag.leavechannelid, tag.leavechannelid)
		} else {
			leavechannel = "Ayarlanmamış."
		}

		err = db.QueryRow(fmt.Sprintf("SELECT roleid FROM autorole WHERE guildid ='%s'", ctx.Guild.ID)).Scan(&tag.roleid)
		if err == nil {
			autorole = fmt.Sprintf("<@&%s> (%s)", tag.roleid, tag.roleid)
		} else {
			autorole = "Ayarlanmamış."
		}

		err = db.QueryRow(fmt.Sprintf("SELECT message FROM welcomemessage WHERE guildid ='%s'", ctx.Guild.ID)).Scan(&tag.welcomemessage)
		if err == nil {
			welcomemsg = tag.welcomemessage
		} else {
			welcomemsg = "Ayarlanmamış."
		}

		err = db.QueryRow(fmt.Sprintf("SELECT message FROM leavemessage WHERE guildid ='%s'", ctx.Guild.ID)).Scan(&tag.leavemessage)
		if err == nil {
			leavemsg = tag.leavemessage
		} else {
			leavemsg = "Ayarlanmamış."
		}

		err = db.QueryRow(fmt.Sprintf("SELECT channelid FROM log WHERE guildid ='%s'", ctx.Guild.ID)).Scan(&tag.logid)
		if err == nil {
			logchannel = fmt.Sprintf("<#%s> (%s)", tag.logid, tag.logid)
		} else {
			logchannel = "Ayarlanmamış."
		}

		embed := embedutil.New(fmt.Sprintf("%s Settings", ctx.Guild.Name), "")
		embed.AddField("Hoş Geldin Kanalı:", welcomechannel, true)
		embed.AddField("Çıkış Kanalı:", leavechannel, true)
		embed.AddField("Otorol:", autorole, true)
		embed.AddField("Çıkış Mesajı:", leavemsg, true)
		embed.AddField("Hoş Geldin Mesajı:", welcomemsg, true)
		embed.AddField("Log Kanalı:", logchannel, true)
		embed.AddField("Dil:", "Türkçe.", true)
		embed.Color = 0x00f0ff

		ctx.ReplyEmbed(embed)

		return nil
	default:

		if sql.IsBlocked(ctx.Guild.ID, "settings") == "true" {
			ctx.Reply("This command is blocked on this guild.")
			return nil
		}

		err := db.QueryRow(fmt.Sprintf("SELECT channelid FROM welcomechannel WHERE guildid ='%s'", ctx.Guild.ID)).Scan(&tag.welcomechannelid)
		if err == nil {
			welcomechannel = fmt.Sprintf("<#%s> (%s)", tag.welcomechannelid, tag.welcomechannelid)
		} else {
			welcomechannel = "Not existing."
		}

		err = db.QueryRow(fmt.Sprintf("SELECT channelid FROM leavechannel WHERE guildid ='%s'", ctx.Guild.ID)).Scan(&tag.leavechannelid)
		if err == nil {
			leavechannel = fmt.Sprintf("<#%s> (%s)", tag.leavechannelid, tag.leavechannelid)
		} else {
			leavechannel = "Not existing."
		}

		err = db.QueryRow(fmt.Sprintf("SELECT roleid FROM autorole WHERE guildid ='%s'", ctx.Guild.ID)).Scan(&tag.roleid)
		if err == nil {
			autorole = fmt.Sprintf("<@&%s> (%s)", tag.roleid, tag.roleid)
		} else {
			autorole = "Not existing."
		}

		err = db.QueryRow(fmt.Sprintf("SELECT message FROM welcomemessage WHERE guildid ='%s'", ctx.Guild.ID)).Scan(&tag.welcomemessage)
		if err == nil {
			welcomemsg = tag.welcomemessage
		} else {
			welcomemsg = "Not existing."
		}

		err = db.QueryRow(fmt.Sprintf("SELECT message FROM leavemessage WHERE guildid ='%s'", ctx.Guild.ID)).Scan(&tag.leavemessage)
		if err == nil {
			leavemsg = tag.leavemessage
		} else {
			leavemsg = "Not existing."
		}

		err = db.QueryRow(fmt.Sprintf("SELECT channelid FROM log WHERE guildid ='%s'", ctx.Guild.ID)).Scan(&tag.logid)
		if err == nil {
			logchannel = fmt.Sprintf("<#%s> (%s)", tag.logid, tag.logid)
		} else {
			logchannel = "Not existing."
		}

		embed := embedutil.New(fmt.Sprintf("%s Settings", ctx.Guild.Name), "")
		embed.AddField("Welcome Channel:", welcomechannel, true)
		embed.AddField("Leave Channel:", leavechannel, true)
		embed.AddField("Autorole:", autorole, true)
		embed.AddField("Leave Message:", leavemsg, true)
		embed.AddField("Welcome Message:", welcomemsg, true)
		embed.AddField("Log Channel:", logchannel, true)
		embed.AddField("Language:", "English.", true)
		embed.Color = 0x00f0ff
		ctx.ReplyEmbed(embed)
		return nil
	}
}
