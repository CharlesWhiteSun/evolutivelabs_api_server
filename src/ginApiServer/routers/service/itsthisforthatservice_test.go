package service

import (
	"testing"
)

func TestItsThisForThatDailySentence_GetItsThisForThatSentence(t *testing.T) {

	if testing.Short() {
		t.Skip("skipping test short mode")
	}

	tests := []struct {
		name     string
		sentence *ItsThisForThatDailySentence
		wantErr  bool
	}{
		{
			name:     "網址為 nil",
			sentence: &ItsThisForThatDailySentence{},
			wantErr:  true,
		},
		{
			name: "網址為空字串",
			sentence: &ItsThisForThatDailySentence{
				ReqURL: "",
			},
			wantErr: true,
		},
		{
			name: "網址亂碼",
			sentence: &ItsThisForThatDailySentence{
				ReqURL: "！＠＃＄％︿＆",
			},
			wantErr: true,
		},
		{
			name: "網址 Google 非法前綴",
			sentence: &ItsThisForThatDailySentence{
				ReqURL: "www.google.com",
			},
			wantErr: true,
		},
		{
			name: "網址 Google 合法前綴",
			sentence: &ItsThisForThatDailySentence{
				ReqURL: "https://www.google.com/",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.sentence.GetItsThisForThatSentence()
			if (err != nil) != tt.wantErr {
				t.Errorf("ItsThisForThatDailySentence.GetItsThisForThatSentence() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
