module onmyoji

go 1.15

require (
	github.com/TraiOi/telegram v0.0.0
	github.com/TraiOi/games v0.0.0
	github.com/TraiOi/tpl v0.0.0
	github.com/spf13/viper v1.7.1
	gopkg.in/tucnak/telebot.v2 v2.3.5 // indirect
)

replace github.com/TraiOi/telegram => ./telegram
replace github.com/TraiOi/games => ./games
replace github.com/TraiOi/tpl => ./tpl
