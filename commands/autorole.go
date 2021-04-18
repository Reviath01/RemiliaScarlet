package commands

import (
	"fmt"
	"strings"

	multiplexer "git.randomchars.net/Reviath/RemiliaScarlet/Multiplexer"
	CommandHandler "git.randomchars.net/Reviath/RemiliaScarlet/handler"
	"git.randomchars.net/Reviath/RemiliaScarlet/interactions"
	"git.randomchars.net/Reviath/RemiliaScarlet/sql"
	"github.com/bwmarrin/discordgo"
)

func AutoRoleCommand(ctx CommandHandler.Context, _ []string) error {
	db := sql.Connect()

	type Tag struct {
		roleid string
	}

	var tag Tag

	if sql.CheckLanguage(ctx.Guild.ID) == "tr" {
		if strings.Join(multiplexer.GetArgs(ctx.Message.Content, multiplexer.GetPrefix()), " ") == "" {
			ctx.Reply("Rolü belirtmelisin.")
			return nil
		} else {
			args := multiplexer.GetRole(multiplexer.GetArgs(ctx.Message.Content, multiplexer.GetPrefix())[0])
			if args == "Error" {
				ctx.Reply("Rolü belirtmelisin.")
				return nil
			}
			err := db.QueryRow(fmt.Sprintf("SELECT roleid FROM autorole WHERE guildid ='%s'", ctx.Guild.ID)).Scan(&tag.roleid)
			if err == nil {
				ctx.Reply("Otorol zaten ayarlanmış, tekrar ayarlamak için reset_autorole komutunu kullan!")
				return nil
			} else {
				insert, err := db.Query(fmt.Sprintf("INSERT INTO autorole (roleid, guildid) VALUES ('%s', '%s')", args, ctx.Guild.ID))
				if err != nil {
					ctx.Reply("Bir hata oluştu.")

					return nil
				}
				defer insert.Close()

				ctx.Reply("Otorol başarıyla ayarlandı.")

				return nil
			}
		}
	}
	if strings.Join(multiplexer.GetArgs(ctx.Message.Content, multiplexer.GetPrefix()), " ") == "" {
		ctx.Reply("You need to specify the role.")
		return nil
	} else {
		args := multiplexer.GetRole(multiplexer.GetArgs(ctx.Message.Content, multiplexer.GetPrefix())[0])
		if args == "Error" {
			ctx.Reply("You need to specify the role.")
			return nil
		}
		err := db.QueryRow(fmt.Sprintf("SELECT roleid FROM autorole WHERE guildid ='%s'", ctx.Guild.ID)).Scan(&tag.roleid)
		if err == nil {
			ctx.Reply("Autorole is already existing, use reset_auto_role command to reset!")
			return nil
		} else {
			insert, err := db.Query(fmt.Sprintf("INSERT INTO autorole (roleid, guildid) VALUES ('%s', '%s')", args, ctx.Guild.ID))
			if err != nil {
				ctx.Reply("An error occurred.")

				return nil
			}
			defer insert.Close()

			ctx.Reply("Successfully set autorole.")

			return nil
		}
	}
}

func InteractiveAutoRoleCommand(session *discordgo.Session, interaction interactions.Interaction) interactions.InteractionResponse {
	db := sql.Connect()

	type Tag struct {
		roleid string
	}

	var tag Tag

	err := db.QueryRow(fmt.Sprintf("SELECT roleid FROM autorole WHERE guildid ='%s'", interaction.GuildID)).Scan(&tag.roleid)
	if err == nil {
		response := interactions.InteractionResponse{
			Type: interactions.InteractionResponseTypeChannelMessageWithSource,
			Data: interactions.InteractionResponseData{
				TTS:     false,
				Content: "Autorole is already existing.",
			},
		}
		return response
	}

	if multiplexer.GetRole(interaction.Data.Options[0].Value.(string)) == "Error" {
		response := interactions.InteractionResponse{
			Type: interactions.InteractionResponseTypeChannelMessageWithSource,
			Data: interactions.InteractionResponseData{
				TTS:     false,
				Content: "You need to specify a role.",
			},
		}
		return response
	}

	insert, _ := db.Query(fmt.Sprintf("INSERT INTO autorole (roleid, guildid) VALUES ('%s', '%s')", multiplexer.GetRole(interaction.Data.Options[0].Value.(string)), interaction.GuildID))
	defer insert.Close()
	response := interactions.InteractionResponse{
		Type: interactions.InteractionResponseTypeChannelMessageWithSource,
		Data: interactions.InteractionResponseData{
			TTS:     false,
			Content: "Successfully set autorole.",
		},
	}
	return response
}
