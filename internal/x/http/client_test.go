package http_test

import (
	"context"
	"testing"

	xhttp "github.com/omissis/goturin-todod/internal/x/http"
)

func TestGet(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		desc       string
		url        string
		wantErr    bool
		wantErrMsg string
	}{
		{
			desc:       "wrong url",
			url:        "-+_@#$%^&*()",
			wantErr:    true,
			wantErrMsg: `cannot create request: parse "-+_@#$%^&*()": invalid URL escape "%^&"`,
		},
		{
			desc:       "server down",
			url:        "http://serverdown:8999",
			wantErr:    true,
			wantErrMsg: `cannot do request: Get "http://serverdown:8999": dial tcp: lookup serverdown: no such host`,
		},
		{
			desc:    "server up",
			url:     "http://1.1.1.1",
			wantErr: false,
		},
	}
	for _, tC := range testCases {
		tC := tC

		t.Run(tC.desc, func(t *testing.T) {
			t.Parallel()

			_, err := xhttp.Get(tC.url, context.Background())

			if (err != nil) != tC.wantErr {
				t.Errorf("Get() error = '%v', wantErr '%v'", err, tC.wantErr)

				return
			}

			if (err != nil) && (err.Error() != tC.wantErrMsg) {
				t.Errorf("Get() error = '%v', wantErrMsg '%v'", err, tC.wantErrMsg)

				return
			}
		})
	}
}
