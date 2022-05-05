package service

import (
	"testing"
)

func TestDailySentence_GetSentence(t *testing.T) {

	if testing.Short() {
		t.Skip("skipping test short mode")
	}

	tests := []struct {
		name    string
		d       *DailySentence
		wantErr bool
	}{
		{
			name:    "網址為 nil",
			d:       &DailySentence{},
			wantErr: true,
		},
		{
			name: "網址為空字串",
			d: &DailySentence{
				ReqURL: "",
			},
			wantErr: true,
		},
		{
			name: "網址亂碼",
			d: &DailySentence{
				ReqURL: "！＠＃＄％︿＆",
			},
			wantErr: true,
		},
		{
			name: "網址 Google 非法前綴",
			d: &DailySentence{
				ReqURL: "www.google.com",
			},
			wantErr: true,
		},
		{
			name: "網址 Google 合法前綴",
			d: &DailySentence{
				ReqURL: "https://www.google.com/",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.d.GetSentence()
			if (err != nil) != tt.wantErr {
				t.Errorf("DailySentence.GetSentence() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
