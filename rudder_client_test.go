package rudderstack_echo_middeware

import (
	"github.com/rudderlabs/analytics-go/v4"
	"testing"
)

func TestNewClient(t *testing.T) {
	type args struct {
		key          string
		dataPanelUrl string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test New Client",
			args: args{
				key:          "XXXXX",
				dataPanelUrl: "hillfolk.org",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewClient(tt.args.key, tt.args.dataPanelUrl).GetClient() == nil; got != tt.want {
				t.Errorf("NewClient() is nil ")
			}
		})
	}
}

func TestClient_Enqueue(t *testing.T) {
	type fields struct {
		rudderClient analytics.Client
	}
	type args struct {
		msg analytics.Message
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Test Enqueue",
			fields: fields{
				rudderClient: analytics.New("XXXXX", "hillfolk.org"),
			},
			args: args{
				msg: analytics.Track{
					UserId: "1hKOmRA4GRlm",
					Event:  "Test Event",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				rudderClient: tt.fields.rudderClient,
			}
			if err := c.Enqueue(tt.args.msg); (err != nil) != tt.wantErr {
				t.Errorf("Enqueue() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
