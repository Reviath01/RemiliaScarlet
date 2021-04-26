package interaction_commands

import (
	"fmt"

	embedutil "git.randomchars.net/Reviath/RemiliaScarlet/EmbedUtil"
	multiplexer "git.randomchars.net/Reviath/RemiliaScarlet/Multiplexer"
	"git.randomchars.net/Reviath/RemiliaScarlet/interactions"
	"git.randomchars.net/Reviath/RemiliaScarlet/sql"
	"github.com/bwmarrin/discordgo"
)

func SettingsCommand(session *discordgo.Session, interaction interactions.Interaction) interactions.InteractionResponse {
	db := sql.Connect()

	Guild, _ := session.Guild(interaction.GuildID)

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

	if sql.CheckLanguage(interaction.GuildID) == "tr" {
		if sql.IsBlocked(interaction.GuildID, "settings") == "true" {
			return multiplexer.CreateResponse("Bu komut bu sunucuda engellenmiş.")
		}

		err := db.QueryRow(fmt.Sprintf("SELECT channelid FROM welcomechannel WHERE guildid ='%s'", interaction.GuildID)).Scan(&tag.welcomechannelid)
		if err == nil {
			welcomechannel = fmt.Sprintf("<#%s> (%s)", tag.welcomechannelid, tag.welcomechannelid)
		} else {
			welcomechannel = "Ayarlanmamış."
		}

		err = db.QueryRow(fmt.Sprintf("SELECT channelid FROM leavechannel WHERE guildid ='%s'", interaction.GuildID)).Scan(&tag.leavechannelid)
		if err == nil {
			leavechannel = fmt.Sprintf("<#%s> (%s)", tag.leavechannelid, tag.leavechannelid)
		} else {
			leavechannel = "Ayarlanmamış."
		}

		err = db.QueryRow(fmt.Sprintf("SELECT roleid FROM autorole WHERE guildid ='%s'", interaction.GuildID)).Scan(&tag.roleid)
		if err == nil {
			autorole = fmt.Sprintf("<@&%s> (%s)", tag.roleid, tag.roleid)
		} else {
			autorole = "Ayarlanmamış."
		}

		err = db.QueryRow(fmt.Sprintf("SELECT message FROM welcomemessage WHERE guildid ='%s'", interaction.GuildID)).Scan(&tag.welcomemessage)
		if err == nil {
			welcomemsg = tag.welcomemessage
		} else {
			welcomemsg = "Ayarlanmamış."
		}

		err = db.QueryRow(fmt.Sprintf("SELECT message FROM leavemessage WHERE guildid ='%s'", interaction.GuildID)).Scan(&tag.leavemessage)
		if err == nil {
			leavemsg = tag.leavemessage
		} else {
			leavemsg = "Ayarlanmamış."
		}

		err = db.QueryRow(fmt.Sprintf("SELECT channelid FROM log WHERE guildid ='%s'", interaction.GuildID)).Scan(&tag.logid)
		if err == nil {
			logchannel = fmt.Sprintf("<#%s> (%s)", tag.logid, tag.logid)
		} else {
			logchannel = "Ayarlanmamış."
		}

		embed := embedutil.NewEmbed().
			SetTitle(Guild.Name+" Settings").
			AddField("Hoş Geldin Kanalı:", welcomechannel).
			AddField("Çıkış Kanalı:", leavechannel).
			AddField("Otorol:", autorole).
			AddField("Çıkış Mesajı:", leavemsg).
			AddField("Hoş Geldin Mesajı:", welcomemsg).
			AddField("Log Kanalı:", logchannel).
			SetColor(0x00f0ff).MessageEmbed

		return multiplexer.CreateEmbedResponse(embed)
	}

	if sql.IsBlocked(interaction.GuildID, "settings") == "true" {
		return multiplexer.CreateResponse("This command is blocked on this guild.")
	}

	err := db.QueryRow(fmt.Sprintf("SELECT channelid FROM welcomechannel WHERE guildid ='%s'", interaction.GuildID)).Scan(&tag.welcomechannelid)
	if err == nil {
		welcomechannel = fmt.Sprintf("<#%s> (%s)", tag.welcomechannelid, tag.welcomechannelid)
	} else {
		welcomechannel = "Not existing."
	}

	err = db.QueryRow(fmt.Sprintf("SELECT channelid FROM leavechannel WHERE guildid ='%s'", interaction.GuildID)).Scan(&tag.leavechannelid)
	if err == nil {
		leavechannel = fmt.Sprintf("<#%s> (%s)", tag.leavechannelid, tag.leavechannelid)
	} else {
		leavechannel = "Not existing."
	}

	err = db.QueryRow(fmt.Sprintf("SELECT roleid FROM autorole WHERE guildid ='%s'", interaction.GuildID)).Scan(&tag.roleid)
	if err == nil {
		autorole = fmt.Sprintf("<@&%s> (%s)", tag.roleid, tag.roleid)
	} else {
		autorole = "Not existing."
	}

	err = db.QueryRow(fmt.Sprintf("SELECT message FROM welcomemessage WHERE guildid ='%s'", interaction.GuildID)).Scan(&tag.welcomemessage)
	if err == nil {
		welcomemsg = tag.welcomemessage
	} else {
		welcomemsg = "Not existing."
	}

	err = db.QueryRow(fmt.Sprintf("SELECT message FROM leavemessage WHERE guildid ='%s'", interaction.GuildID)).Scan(&tag.leavemessage)
	if err == nil {
		leavemsg = tag.leavemessage
	} else {
		leavemsg = "Not existing."
	}

	err = db.QueryRow(fmt.Sprintf("SELECT channelid FROM log WHERE guildid ='%s'", interaction.GuildID)).Scan(&tag.logid)
	if err == nil {
		logchannel = fmt.Sprintf("<#%s> (%s)", tag.logid, tag.logid)
	} else {
		logchannel = "Not existing."
	}

	embed := embedutil.NewEmbed().
		SetTitle(Guild.Name+" Settings").
		AddField("Welcome Channel:", welcomechannel).
		AddField("Leave Channel:", leavechannel).
		AddField("Autorole:", autorole).
		AddField("Leave Message:", leavemsg).
		AddField("Welcome Message:", welcomemsg).
		AddField("Log Channel:", logchannel).
		SetColor(0x00f0ff).MessageEmbed
	return multiplexer.CreateEmbedResponse(embed)
}
