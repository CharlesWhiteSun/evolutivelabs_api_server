package service

import (
	"testing"
)

func TestDailySentence_GetSentence(t *testing.T) {

	if testing.Short() {
		t.Skip("skipping test short mode")
	}

	tests := []struct {
		name     string
		sentence *DailySentence
		wantErr  bool
	}{
		{
			name:     "網址為 nil",
			sentence: &DailySentence{},
			wantErr:  true,
		},
		{
			name: "網址為空字串",
			sentence: &DailySentence{
				ReqURL: "",
			},
			wantErr: true,
		},
		{
			name: "網址亂碼",
			sentence: &DailySentence{
				ReqURL: "！＠＃＄％︿＆",
			},
			wantErr: true,
		},
		{
			name: "網址 Google 非法前綴",
			sentence: &DailySentence{
				ReqURL: "www.google.com",
			},
			wantErr: true,
		},
		{
			name: "網址 Google 合法前綴",
			sentence: &DailySentence{
				ReqURL: "https://www.google.com/",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.sentence.GetSentence()
			if (err != nil) != tt.wantErr {
				t.Errorf("DailySentence.GetSentence() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
