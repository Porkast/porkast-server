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
