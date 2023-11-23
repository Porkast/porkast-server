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
				tokenString: "eyJhbGciOiJIUzI1NiIsImtpZCI6IlQwbHZ3L3pQbjNNL2gzSEMiLCJ0eXAiOiJKV1QifQ.eyJhdWQiOiJhdXRoZW50aWNhdGVkIiwiZXhwIjoxNzAwNzMzODA1LCJpYXQiOjE3MDA3MzAyMDUsImlzcyI6Imh0dHBzOi8vcW9qcHlnam54YWN1eGRydHpkdHouc3VwYWJhc2UuY28vYXV0aC92MSIsInN1YiI6IjQ5NjJiYjNjLTFiN2ItNDE1NC1iY2Y5LWE3YWQwZjIzNWExMyIsImVtYWlsIjoiY2hlbmp1bnFpYW4wODEwQGdtYWlsLmNvbSIsInBob25lIjoiIiwiYXBwX21ldGFkYXRhIjp7InByb3ZpZGVyIjoiZ29vZ2xlIiwicHJvdmlkZXJzIjpbImdvb2dsZSJdfSwidXNlcl9tZXRhZGF0YSI6eyJhdmF0YXJfdXJsIjoiaHR0cHM6Ly9saDMuZ29vZ2xldXNlcmNvbnRlbnQuY29tL2EvQUNnOG9jS203UEpDMHpQSUh3VFUtTzVYRVVZWU96WHlHX0FiUFdPMVZSUmhFZ3RKZzIwPXM5Ni1jIiwiZW1haWwiOiJjaGVuanVucWlhbjA4MTBAZ21haWwuY29tIiwiZW1haWxfdmVyaWZpZWQiOnRydWUsImZ1bGxfbmFtZSI6IumZiOS_iuS5viIsImlzcyI6Imh0dHBzOi8vYWNjb3VudHMuZ29vZ2xlLmNvbSIsIm5hbWUiOiLpmYjkv4rkub4iLCJwaG9uZV92ZXJpZmllZCI6ZmFsc2UsInBpY3R1cmUiOiJodHRwczovL2xoMy5nb29nbGV1c2VyY29udGVudC5jb20vYS9BQ2c4b2NLbTdQSkMwelBJSHdUVS1PNVhFVVlZT3pYeUdfQWJQV08xVlJSaEVndEpnMjA9czk2LWMiLCJwcm92aWRlcl9pZCI6IjExMzYxMzM2MzMwNzA1ODM2NTU4OSIsInN1YiI6IjExMzYxMzM2MzMwNzA1ODM2NTU4OSJ9LCJyb2xlIjoiYXV0aGVudGljYXRlZCIsImFhbCI6ImFhbDEiLCJhbXIiOlt7Im1ldGhvZCI6Im9hdXRoIiwidGltZXN0YW1wIjoxNzAwNzMwMjA1fV0sInNlc3Npb25faWQiOiJhOTM3ZGZiYi1kMjRiLTQ5NGUtYjIyYy05YmM3ZGU1ZDc5MWIifQ.OuZBgLChAQMs9YTNE9m0SeHD_ac_k3gpC6Rfju2a0Rk",
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
