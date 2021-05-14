package interaction_commands

import (
	"fmt"

	embedutil "git.randomchars.net/FreeNitori/EmbedUtil"
	multiplexer "git.randomchars.net/Reviath/RemiliaScarlet/Multiplexer"
	"git.randomchars.net/Reviath/RemiliaScarlet/interactions"
	"git.randomchars.net/Reviath/RemiliaScarlet/sql"
	"github.com/bwmarrin/discordgo"
)

//Settings slash command.
func SettingsCommand(session *discordgo.Session, interaction interactions.Interaction) interactions.InteractionResponse {
	db := sql.Connect()
	defer db.Close()

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

	switch sql.CheckLanguage(interaction.GuildID) {
	case "tr":
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

		embed := embedutil.New(Guild.Name+" Settings", "")
		embed.AddField("Hoş Geldin Kanalı:", welcomechannel, true)
		embed.AddField("Çıkış Kanalı:", leavechannel, true)
		embed.AddField("Otorol:", autorole, true)
		embed.AddField("Çıkış Mesajı:", leavemsg, true)
		embed.AddField("Hoş Geldin Mesajı:", welcomemsg, true)
		embed.AddField("Log Kanalı:", logchannel, true)
		embed.Color = 0x00f0ff

		return multiplexer.CreateEmbedResponse(embed)
	default:

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

		embed := embedutil.New(Guild.Name+" Settings", "")
		embed.AddField("Welcome Channel:", welcomechannel, true)
		embed.AddField("Leave Channel:", leavechannel, true)
		embed.AddField("Autorole:", autorole, true)
		embed.AddField("Leave Message:", leavemsg, true)
		embed.AddField("Welcome Message:", welcomemsg, true)
		embed.AddField("Log Channel:", logchannel, true)
		embed.Color = 0x00f0ff
		return multiplexer.CreateEmbedResponse(embed)
	}
}
