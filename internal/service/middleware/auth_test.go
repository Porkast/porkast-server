package middleware

import "testing"

func TestVerifyJWTToken(t *testing.T) {
	type args struct {
		tokenString string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "verify JWT token",
			args: args{
				tokenString: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJhdXRoZW50aWNhdGVkIiwiZXhwIjoxNjk3OTkzMjQyLCJpYXQiOjE2OTc5ODk2NDIsInN1YiI6ImZkMDg1NGVhLTUzZGUtNGQwYy05ZTA4LTZhMGIxYjBhOTk0NSIsImVtYWlsIjoiY2hlbmp1bnFpYW4wODEwQGZveG1haWwuY29tIiwicGhvbmUiOiIiLCJhcHBfbWV0YWRhdGEiOnsicHJvdmlkZXIiOiJlbWFpbCIsInByb3ZpZGVycyI6WyJlbWFpbCJdfSwidXNlcl9tZXRhZGF0YSI6e30sInJvbGUiOiJhdXRoZW50aWNhdGVkIiwiYWFsIjoiYWFsMSIsImFtciI6W3sibWV0aG9kIjoibWFnaWNsaW5rIiwidGltZXN0YW1wIjoxNjk3NzI4MDM1fV0sInNlc3Npb25faWQiOiI5ZjFmODk2Zi1hNjI2LTRkOTktYjRlNS1lNDE3NzA1ZDdkY2IifQ.QskpZDtIRzZmGKLRxtGkowjqd4xmtd9w3PkvdFiFaLY",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := VerifyJWTToken(tt.args.tokenString); (err != nil) != tt.wantErr {
				t.Errorf("VerifyJWTToken() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
