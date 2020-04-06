package organization

import (
	"testing"
)

func TestTwitterHandler_RedirectUrl(t *testing.T) {
	tests := []struct {
		name string
		th   TwitterHandler
		want string
	}{
		{name: "@darth", th: TwitterHandler("@darth"), want: "https://www.twitter.com/darth"},
		{name: "@luke_skywalker", th: TwitterHandler("@luke_skywalker"), want: "https://www.twitter.com/luke_skywalker"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.th.RedirectUrl(); got != tt.want {
				t.Errorf("RedirectUrl() = %v, want %v", got, tt.want)
			}
		})
	}
}
