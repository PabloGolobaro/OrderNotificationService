package utils

import (
	"WhatsappOrderServer/config"
	tele "gopkg.in/telebot.v3"
	"log"
	"testing"
	"time"
)

func TestSendToAdmins(t *testing.T) {
	type args struct {
		bot    *tele.Bot
		admins []string
	}
	config, _ := config.LoadConfig("C:\\Users\\Professional\\GolandProjects\\WhatsappOrderServer")
	pref := tele.Settings{
		Token:  config.Token,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
		return
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "First", args: args{
			bot:    b,
			admins: []string{"1048928261"},
		}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := SendToAdmins(tt.args.bot, tt.args.admins, nil); (err != nil) != tt.wantErr {
				t.Errorf("SendToAdmins() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
