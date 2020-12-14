package telegram

import (
	"fmt"
	tb "gopkg.in/tucnak/telebot.v2"
	"github.com/TraiOi/games"
)

func HandleBountyLocation(b *tb.Bot) {
	bounty := new(games.BOUNTY_LOCATION)
	listBounty := bounty.GetData()

	selector := &tb.ReplyMarkup{}

	btn := getBtn(selector, listBounty)
	row := getRow(selector, btn)
	selector.Inline(row...)

	for _, v := range btn {
		tmp := fmt.Sprintf("%s (%s)", bounty.GetVName(v.Unique),
					      bounty.GetEName(v.Unique))
                b.Handle(&v, func(c *tb.Callback) {
                        b.Respond(c, &tb.CallbackResponse{
                                ShowAlert: false,
                        })
                        b.Send(c.Message.Chat, tmp)
                })
        }


	b.Handle("/bounty", func(m *tb.Message) {
		b.Send(m.Chat, "Danh sách Truy", selector)
	})
}

func getBtn(selector *tb.ReplyMarkup, content []games.BOUNTY_DATA) []tb.Btn {
	var result []tb.Btn

	for _, v := range content {
		btn := selector.Data(v.Hint, v.Shiki)
		result = append(result, btn)
	}
	return result
}

func getRow(selector *tb.ReplyMarkup, btn []tb.Btn) []tb.Row {
	var result []tb.Row
	for _, v := range btn {
		result = append(result, selector.Row(v))
	}
	return result
}
