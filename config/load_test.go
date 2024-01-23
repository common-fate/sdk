package config

import (
	"os"
	"testing"

	"github.com/common-fate/grab"
	"github.com/stretchr/testify/assert"
)

func Test_clientSecret(t *testing.T) {
	type args struct {
		context *Context
		opts    Opts
	}
	tests := []struct {
		name   string
		args   args
		before func()
		after  func()
		want   *string
	}{
		{
			name:   "from env",
			args:   args{},
			before: func() { os.Setenv("CF_OIDC_CLIENT_SECRET", "secret") },
			after:  func() { os.Unsetenv("CF_OIDC_CLIENT_SECRET") },
			want:   grab.Ptr("secret"),
		},
		{
			name: "from opts",
			args: args{
				opts: Opts{
					ClientSecret: "another secret",
				},
			},
			want: grab.Ptr("another secret"),
		},
		{
			name: "from context",
			args: args{
				context: &Context{
					OIDCClientSecret: grab.Ptr("context"),
				},
			},
			want: grab.Ptr("context"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.before != nil {
				tt.before()
			}
			if tt.after != nil {
				defer tt.after()
			}
			if got := clientSecret(grab.Value(tt.args.context).OIDCClientSecret, tt.args.opts); got != tt.want {
				assert.EqualValues(t, tt.want, got)
			}
		})
	}
}
