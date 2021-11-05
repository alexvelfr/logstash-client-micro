package logstashclientmicro

import (
	"context"
	"errors"
	"testing"
)

func Test_client_LogError(t *testing.T) {
	type args struct {
		ctx context.Context
		msg Message
	}
	tests := []struct {
		name    string
		c       Client
		args    args
		wantErr bool
	}{
		{
			name: "t1",
			c:    NewClient("service1", "http://localhost:3337", false),
			args: args{ctx: context.Background(), msg: Message{
				XReqID: "1",
				Data:   "my data",
				File:   "client.go",
				Action: "some action",
				Error:  errors.New("something went wrong"),
			}},
			wantErr: false,
		},
		{
			name: "t2",
			c:    NewClient("service1", "http://localhost:3337/", false),
			args: args{ctx: context.Background(), msg: Message{
				XReqID: "2",
				Data:   "my data2",
				File:   "client.go",
				Action: "some action2",
				Error:  errors.New("something went wrong2"),
			}},
			wantErr: false,
		},
		{
			name: "t3",
			c:    NewClient("service2", "http://localhost:3337/", false),
			args: args{ctx: context.Background(), msg: Message{
				XReqID: "2",
				Data:   "my data2",
				File:   "client.go",
				Action: "some action2",
				Error:  errors.New("something went wrong2"),
			}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.c.LogError(tt.args.ctx, tt.args.msg); (err != nil) != tt.wantErr {
				t.Errorf("client.LogError() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
