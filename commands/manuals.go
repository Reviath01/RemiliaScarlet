package commands

import (
	"fmt"
	embedutil "git.randomchars.net/Reviath/remilia-scarlet-embed-util"
	"../state"
	"git.randomchars.net/Reviath/remilia-scarlet-command-handler"
	"strings"
)

func init() {
	state.Multiplexer.Route(&multiplexer.Route{
		Pattern:       "man",
		AliasPatterns: []string{"manuals", "help"},
		Description:   "An interface to the system reference manuals.",
		Category:      multiplexer.ManualsCategory,
		Handler:       manuals,
	})
}

func manuals(context *multiplexer.Context) {
	guildPrefix := context.Prefix()

	switch {
	case len(context.Fields) == 1:
		{
			embed := embedutil.New("Manuals",
				fmt.Sprintf("Issue `%sman <category>` for category-specific manuals.", guildPrefix))
			embed.Color = 0x3492c4

			var catText string
			for _, category := range state.Multiplexer.Categories {

				if category.Description == "" {
					continue
				}

				catText += fmt.Sprintf("%s - %s \n", category.Title, category.Description)

			}

			embed.AddField("Categories", catText, false)
			_ = context.SendEmbed("", embed)
		}

	case len(context.Fields) == 2:
		{

			var desiredCat *multiplexer.CommandCategory
			for _, cat := range state.Multiplexer.Categories {
				if strings.EqualFold(cat.Title, context.Fields[1]) {
					desiredCat = cat
					break
				}
			}

			if desiredCat == nil {
				context.SendMessage(multiplexer.InvalidArgument)
				break
			}

			embed := embedutil.TitleWithDescriptionAndColorEmbed(desiredCat.Title,
				desiredCat.Description, 0x007bff)
			
			for _, route := range desiredCat.Routes {

				if route.Description == "" {
					continue
				}

				embed.AddField(route.Pattern, route.Description, false)
			}
			context.SendEmbed("", embed)

		}

	case len(context.Fields) > 2:
		{
			context.SendMessage(multiplexer.InvalidArgument)
		}
	}
}
