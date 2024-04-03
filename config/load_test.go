package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

// func Test_clientSecret(t *testing.T) {
// 	type args struct {
// 		context *Context
// 		opts    Opts
// 	}
// 	tests := []struct {
// 		name   string
// 		args   args
// 		before func()
// 		after  func()
// 		want   *string
// 	}{
// 		{
// 			name:   "from env",
// 			args:   args{},
// 			before: func() { os.Setenv("CF_OIDC_CLIENT_SECRET", "secret") },
// 			after:  func() { os.Unsetenv("CF_OIDC_CLIENT_SECRET") },
// 			want:   grab.Ptr("secret"),
// 		},
// 		{
// 			name: "from opts",
// 			args: args{
// 				opts: Opts{
// 					ClientSecret: "another secret",
// 				},
// 			},
// 			want: grab.Ptr("another secret"),
// 		},
// 		{
// 			name: "from context",
// 			args: args{
// 				context: &Context{
// 					OIDCClientSecret: grab.Ptr("context"),
// 				},
// 			},
// 			want: grab.Ptr("context"),
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if tt.before != nil {
// 				tt.before()
// 			}
// 			if tt.after != nil {
// 				defer tt.after()
// 			}
// 			if got := clientSecret(grab.Value(tt.args.context).OIDCClientSecret, tt.args.opts); got != tt.want {
// 				assert.EqualValues(t, tt.want, got)
// 			}
// 		})
// 	}
// }

func Test_loadFromSources(t *testing.T) {
	type args struct {
		value   string
		key     Key
		envVars map[string]string
		sources []configSource
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "ok",
			want: "test",
			args: args{
				sources: []configSource{StaticSource{Result: "test"}},
			},
		},
		{
			name: "doesnt_override_set_value",
			want: "existing",
			args: args{
				value:   "existing",
				sources: []configSource{StaticSource{Result: "test"}},
			},
		},
		{
			name: "respects_source_priority",
			// in this test case, there are two config sources.
			// we should take the value that we get from the first config source.
			want: "first",
			args: args{
				sources: []configSource{StaticSource{Result: "first"}, StaticSource{Result: "second"}},
			},
		},
		{
			name: "env_var",
			want: "http://example.com",
			args: args{
				envVars: map[string]string{
					"CF_API_URL": "http://example.com",
				},
				key:     APIURLKey,
				sources: []configSource{EnvSource{}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for k, v := range tt.args.envVars {
				os.Setenv(k, v)
			}

			value := tt.args.value

			loadFromSources(&value, tt.args.key, tt.args.sources)

			assert.Equal(t, tt.want, value)
		})
	}
}
