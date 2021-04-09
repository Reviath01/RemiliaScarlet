package commands

import (
	"database/sql"
	"strings"

	ctx "git.randomchars.net/Reviath/RemiliaScarlet/Context"
	"github.com/bwmarrin/discordgo"
	_ "github.com/go-sql-driver/mysql"
)

type WelcomeChannel struct {
}

func (w WelcomeChannel) Execute(ctx ctx.Ctx, session *discordgo.Session) error {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/remilia")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	type Tag struct {
		channelid string `json:"channelid"`
		lang      string `json:"language"`
	}

	var tag Tag

	err = db.QueryRow("SELECT language FROM languages WHERE guildid ='" + ctx.Guild().ID + "'").Scan(&tag.lang)
	if err == nil {
		if tag.lang == "tr" {
			perms, err := session.State.UserChannelPermissions(ctx.Author().ID, ctx.Channel().ID)
			if err == nil && (int(perms)&discordgo.PermissionAdministrator == discordgo.PermissionAdministrator) == false {
				_, err := session.ChannelMessageSend(ctx.Channel().ID, "Bu komutu kullanmak için yönetici yetkisine sahip olmalısın.")

				if err != nil {
					return nil
				}

				return err
			}

			var args string
			if len(strings.Join(ctx.Args(), " ")) < 1 {
				_, err := session.ChannelMessageSend(ctx.Channel().ID, "Bir kanal belirtmelisin.")

				if err != nil {
					return nil
				}

				return err
			}
			args = ctx.Args()[0]

			if len(args) < 19 {
				c, err := session.Channel(args)
				if err == nil {
					err = db.QueryRow("SELECT channelid FROM welcomechannel WHERE guildid ='" + ctx.Guild().ID + "'").Scan(&tag.channelid)
					if err == nil {
						_, err = session.ChannelMessageSend(ctx.Channel().ID, "Hoş geldin kanalı zaten ayarlanmış, sıfırlamak için reset_welcome_channel komutunu kullan.")

						if err != nil {
							return nil
						}

						return err
					} else {
						insert, err := db.Query("INSERT INTO welcomechannel (channelid, guildid) VALUES ('" + c.ID + "', '" + ctx.Guild().ID + "')")
						if err != nil {
							_, err = session.ChannelMessageSend(ctx.Channel().ID, "Bir hata oluştu.")

							if err != nil {
								return nil
							}

							return err
						}
						defer insert.Close()

						_, err = session.ChannelMessageSend(ctx.Channel().ID, "Hoş geldin kanalı başarıyla ayarlandı.")

						if err != nil {
							return nil
						}

						return err
					}
				} else {
					_, err = session.ChannelMessageSend(ctx.Channel().ID, "Bir kanal belirtmelisin.")

					if err != nil {
						return nil
					}

					return err
				}
			} else {
				if len(args) > 20 {
					c, err := session.Channel(args[2:][:18])
					if err == nil {
						err = db.QueryRow("SELECT channelid FROM welcomechannel WHERE guildid ='" + ctx.Guild().ID + "'").Scan(&tag.channelid)
						if err == nil {
							_, err = session.ChannelMessageSend(ctx.Channel().ID, "Hoş geldin kanalı zaten ayarlanmış, sıfırlamak için reset_welcome_channel komutunu kullan.")

							if err != nil {
								return nil
							}

							return err
						} else {
							insert, err := db.Query("INSERT INTO welcomechannel (channelid, guildid) VALUES ('" + c.ID + "', '" + ctx.Guild().ID + "')")
							if err != nil {
								_, err = session.ChannelMessageSend(ctx.Channel().ID, "Bir hata oluştu.")

								if err != nil {
									return nil
								}

								return err
							}
							defer insert.Close()

							_, err = session.ChannelMessageSend(ctx.Channel().ID, "Hoş geldin kanalı başarıyla ayarlandı.")

							if err != nil {
								return nil
							}

							return err
						}
					} else {
						_, err = session.ChannelMessageSend(ctx.Channel().ID, "Bir kanal belirtmelisin.")

						if err != nil {
							return nil
						}

						return err
					}
				} else {
					_, err := session.ChannelMessageSend(ctx.Channel().ID, "Bir kanal belirtmelisin.")

					if err != nil {
						return nil
					}

					return err
				}
			}

		}
	}

	perms, err := session.State.UserChannelPermissions(ctx.Author().ID, ctx.Channel().ID)
	if err == nil && (int(perms)&discordgo.PermissionAdministrator == discordgo.PermissionAdministrator) == false {
		_, err := session.ChannelMessageSend(ctx.Channel().ID, "You need administrator permission to run this command.")

		if err != nil {
			return nil
		}

		return err
	}

	var args string
	if len(strings.Join(ctx.Args(), " ")) < 1 {
		_, err := session.ChannelMessageSend(ctx.Channel().ID, "You need to specify the channel.")

		if err != nil {
			return nil
		}

		return err
	}
	args = ctx.Args()[0]

	if len(args) < 19 {
		c, err := session.Channel(args)
		if err == nil {
			err = db.QueryRow("SELECT channelid FROM welcomechannel WHERE guildid ='" + ctx.Guild().ID + "'").Scan(&tag.channelid)
			if err == nil {
				_, err = session.ChannelMessageSend(ctx.Channel().ID, "Welcome channel is already existing (to reset, use reset_welcome_channel command).")

				if err != nil {
					return nil
				}

				return err
			} else {
				insert, err := db.Query("INSERT INTO welcomechannel (channelid, guildid) VALUES ('" + c.ID + "', '" + ctx.Guild().ID + "')")
				if err != nil {
					_, err = session.ChannelMessageSend(ctx.Channel().ID, "An error occurred, please try again.")

					if err != nil {
						return nil
					}

					return err
				}
				defer insert.Close()

				_, err = session.ChannelMessageSend(ctx.Channel().ID, "Welcome channel set successfully.")

				if err != nil {
					return nil
				}

				return err
			}
		} else {
			_, err = session.ChannelMessageSend(ctx.Channel().ID, "You need to specify the channel.")

			if err != nil {
				return nil
			}

			return err
		}
	} else {
		if len(args) > 20 {
			c, err := session.Channel(args[2:][:18])
			if err == nil {
				err = db.QueryRow("SELECT channelid FROM welcomechannel WHERE guildid ='" + ctx.Guild().ID + "'").Scan(&tag.channelid)
				if err == nil {
					_, err = session.ChannelMessageSend(ctx.Channel().ID, "Welcome channel is already existing (to reset, use reset_welcome_channel command).")

					if err != nil {
						return nil
					}

					return err
				} else {
					insert, err := db.Query("INSERT INTO welcomechannel (channelid, guildid) VALUES ('" + c.ID + "', '" + ctx.Guild().ID + "')")
					if err != nil {
						_, err = session.ChannelMessageSend(ctx.Channel().ID, "An error occurred, please try again.")

						if err != nil {
							return nil
						}

						return err
					}
					defer insert.Close()

					_, err = session.ChannelMessageSend(ctx.Channel().ID, "Welcome channel set successfully.")

					if err != nil {
						return nil
					}

					return err
				}
			} else {
				_, err = session.ChannelMessageSend(ctx.Channel().ID, "You need to specify the channel.")

				if err != nil {
					return nil
				}

				return err
			}
		} else {
			_, err := session.ChannelMessageSend(ctx.Channel().ID, "You need to specify the channel.")

			if err != nil {
				return nil
			}

			return err
		}
	}
}
