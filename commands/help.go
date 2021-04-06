package commands

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"

	embedutil "git.randomchars.net/Reviath/embed-util"
	ctx "git.randomchars.net/Reviath/handlers/Context"
	"github.com/bwmarrin/discordgo"
	_ "github.com/go-sql-driver/mysql"
)

type Help struct {

}

func (h Help) Execute(ctx ctx.Ctx, session *discordgo.Session) error {
    file, err := ioutil.ReadFile("./config.json")

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

    type configStruct struct {
        Token     string `json:"Token"`
        BotPrefix string `json:"BotPrefix"`
        Presence string `json:"Presence"`
        Owner string `json:"Owner"`
    }
    
    var (
        config *configStruct
    ) 
    
    err = json.Unmarshal(file, &config)

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

    db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/remilia")

    if err != nil {
        panic(err.Error())
    }

    defer db.Close()

    type Tag struct {
        lang string `json:"language"`
    }

    var tag Tag

    err = db.QueryRow("SELECT language FROM languages WHERE guildid ='" + ctx.Guild().ID +"'").Scan(&tag.lang)
    if err == nil {
        if tag.lang == "tr" {
    u, err := session.User("@me")
    if err != nil {
        return nil
    }
    embed := embedutil.NewEmbed().
        SetTitle(u.Username + " Yardım Menüsü!").
        SetColor(0x2ecc71).
        AddField("Kullanıcı Komutları", config.BotPrefix + "afk: `AFK moduna geçersiniz.` \n" + config.BotPrefix + "avatar: `Belirttiğiniz kişinin profil fotoğrafını gönderir.` \n" + config.BotPrefix + "ping: `Botun gecikmesini gönderir.` \n" + config.BotPrefix + "invite: `Botun davet linkini gönderir.` \n" + config.BotPrefix + "roles: `Rollerin listesini gösterir.` \n" + config.BotPrefix + "guild_info: `Sunucuyla ilgili bilgi verir.` \n" + config.BotPrefix + "stats: `Botun istatistiklerini gösterir.` \n" + config.BotPrefix + "settings: `Sunucu ayarlarını gösterir.`").
        AddField("Moderasyon Komutları", config.BotPrefix + "ban: `Belirtilen kişiyi sunucudan yasaklar.` \n" + config.BotPrefix + "kick: `Belirtilen kişiyi sunucudan atar.` \n" + config.BotPrefix + "unban: `Belirtilen kişinin yasağını kaldırır.` \n" + config.BotPrefix + "start_vote: `Oylama başlatır.` \n" + config.BotPrefix + "welcome_channel: `Yeni bir üye katıldığında mesaj atılacak kanalı belirler.` \n" + config.BotPrefix + "leave_channel: `Bir üye ayrıldığında mesaj atılacak kanalı belirler.` \n" + config.BotPrefix + "reset_welcome_channel: `Yeni bir üye katıldığında mesaj atılacak kanalı sıfırlar.` \n" + config.BotPrefix + "reset_leave_channel: `Bir üye ayrıldığında mesaj atılacak kanalı sıfırlar.` \n" + config.BotPrefix + "auto_role: `Otorol ayarlar.` \n" + config.BotPrefix + "reset_auto_role: `Otorolü sıfırlar.` \n" + config.BotPrefix + "welcome_message: `Yeni bir üye katıldığında atılacak mesajı belirler.` \n" + config.BotPrefix + "reset_welcome_message: `Yeni bir üye katıldığında atılacak mesajı sıfırlar.` \n" + config.BotPrefix + "leave_message: `Bir üye ayrıldığında atılacak mesajı belirler.` \n" + config.BotPrefix + "reset_leave_message: `Bir üye ayrıldığında atılacak mesajı sıfırlar.` \n" + config.BotPrefix + "log: `Log kanalını belirler.`").
        AddField("Moderasyon Komutları (Devamı)",  config.BotPrefix + "reset_log: `Log kanalını sıfırlamanızı sağlar.` \n" + config.BotPrefix + "disable: `Belirtilen komutu devre dışı bırakır.` \n" + config.BotPrefix + "enable: `Belirtilen komutu etkinleştirir.` \n" + config.BotPrefix + "language: `Botun dilini değiştirmenizi sağlar`").
        AddField("Eğlence Komutları", config.BotPrefix + "embed: `Mesajınızı embed olarak gönderir.` \n" + config.BotPrefix + "hug: `Belirttiğiniz kişiye sarılırsınız.` \n" + config.BotPrefix + "icon: `Sunucu icon'unu gönderir.` \n" + config.BotPrefix + "kiss: `Belirttiğiniz kişiyi öpersiniz.` \n" + config.BotPrefix + "slap: `Belirttiğiniz kişiye vurursunuz.` \n" + config.BotPrefix + "spoiler: `Mesajınızı spoiler olarak gönderir.`").MessageEmbed
    _, err = session.ChannelMessageSendEmbed(ctx.Channel().ID, embed)

    if err != nil {
        return nil
    }

    return err
        }
    }
    
    u, err := session.User("@me")

    if err != nil {
        return nil
    }

    embed := embedutil.NewEmbed().
        SetTitle(u.Username + " Help Menu!").
        SetColor(0x2ecc71).
        AddField("User Commands", config.BotPrefix + "afk: `Sets you as AFK.` \n" + config.BotPrefix + "avatar: `Fetch the profile picture of a user.` \n" + config.BotPrefix + "ping: `Sends latency of the bot.` \n" + config.BotPrefix + "invite: `Sends invite link.` \n" + config.BotPrefix + "roles: `Shows list of roles.` \n" + config.BotPrefix + "guild_info: `Gives information about guild.` \n" + config.BotPrefix + "stats: `Shows stats for bot.` \n" + config.BotPrefix + "settings: `Shows server settings.`").
        AddField("Moderation Commands", config.BotPrefix + "ban: `Bans the user.` \n" + config.BotPrefix + "kick: `Kicks the user.` \n" + config.BotPrefix + "unban: `Unbans the user.` \n" + config.BotPrefix + "start_vote: `Starts a vote.` \n" + config.BotPrefix + "welcome_channel: `When someone joins your server, the bot sends a message to the channel you set.` \n" + config.BotPrefix + "leave_channel: `When someone leaves your server, the bot sends a message to the channel you set.` \n" + config.BotPrefix + "reset_welcome_channel: `Resets welcome channel.` \n" + config.BotPrefix + "reset_leave_channel: `Resets leave channel.` \n" + config.BotPrefix + "auto_role: `Sets autorole.` \n" + config.BotPrefix + "reset_auto_role: `Resets autorole.` \n" + config.BotPrefix + "welcome_message: `Sets welcome message.` \n" + config.BotPrefix + "reset_welcome_message: `Resets welcome message.` \n" + config.BotPrefix + "leave_message: `Sets leave message.` \n"+ config.BotPrefix + "reset_leave_message: `Resets leave message.` \n" + config.BotPrefix + "log: `Sets log channel.` \n" + config.BotPrefix + "reset_log: `Resets log channel.` \n" + config.BotPrefix + "disable: `Disables specified command.` \n" + config.BotPrefix + "enable: `Enables specified command.` \n" + config.BotPrefix + "language: `Sets bot language.`").
        AddField("Fun Commands", config.BotPrefix + "embed: `Sends your message as an embed.` \n" + config.BotPrefix + "hug: `Allows you to hug someone.` \n" + config.BotPrefix + "icon: `Sends guild icon.` \n" + config.BotPrefix + "kiss: `Allows you to kiss someone.` \n" + config.BotPrefix + "slap: `Sends slap gif.` \n" + config.BotPrefix + "spoiler: `Sends your message as a spoiler.`").MessageEmbed
    _, err = session.ChannelMessageSendEmbed(ctx.Channel().ID, embed)

    if err != nil {
        return nil
    }

    return err
}
