package feed

import "testing"

func Test_formatTitle(t *testing.T) {
	type args struct {
		title string
	}
	tests := []struct {
		name            string
		args            args
		wantFormatTitle string
	}{
		{
			name: "format title to string title",
			args: args{
				title: "<span style='color: red;'>播</span><span style='color: red;'>客</span>志",
			},
			wantFormatTitle: "播客志",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotFormatTitle := formatTitle(tt.args.title); gotFormatTitle != tt.wantFormatTitle {
				t.Errorf("formatTitle() = %v, want %v", gotFormatTitle, tt.wantFormatTitle)
			}
		})
	}
}

func Test_formatItemTitle(t *testing.T) {
	type args struct {
		title string
	}
	tests := []struct {
		name            string
		args            args
		wantFormatTitle string
	}{
		{
			name: "format item title",
			args: args{
				title: `#11 和设计师聊AIGC 2：UX设计师"喜迎"AI，是趁手工具，还是竞争对手?`,
			},
			wantFormatTitle: "#11 和设计师聊AIGC 2：UX设计师`喜迎`AI，是趁手工具，还是竞争对手?",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotFormatTitle := formatItemTitle(tt.args.title); gotFormatTitle != tt.wantFormatTitle {
				t.Errorf("formatItemTitle() = %v, want %v", gotFormatTitle, tt.wantFormatTitle)
			}
		})
	}
}
