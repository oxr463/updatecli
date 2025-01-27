package branch

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/updatecli/updatecli/pkg/plugins/utils/version"
)

func TestSource(t *testing.T) {

	tests := []struct {
		name     string
		manifest struct {
			URL           string
			Token         string
			Owner         string
			Repository    string
			VersionFilter version.Filter
		}
		wantResult string
		wantErr    bool
	}{
		{
			name: "repository olblak/updatecli-donotexist should not exist",
			manifest: struct {
				URL           string
				Token         string
				Owner         string
				Repository    string
				VersionFilter version.Filter
			}{
				URL:        "gitlab.com",
				Token:      "",
				Owner:      "updatecli",
				Repository: "updatecli-donotexist",
			},
			wantResult: "",
			wantErr:    true,
		},
		{
			name: "repository should exist with a branch main",
			manifest: struct {
				URL           string
				Token         string
				Owner         string
				Repository    string
				VersionFilter version.Filter
			}{
				Token:      "",
				Owner:      "olblak",
				Repository: "updatecli",
				VersionFilter: version.Filter{
					Kind:    "regex",
					Pattern: "main",
				},
			},
			wantResult: "main",
			wantErr:    false,
		},
		{
			name: "repository should not have branch donotexist",
			manifest: struct {
				URL           string
				Token         string
				Owner         string
				Repository    string
				VersionFilter version.Filter
			}{
				URL:        "gitlab.com",
				Token:      "",
				Owner:      "olblak",
				Repository: "updatecli",
				VersionFilter: version.Filter{
					Kind:    "regex",
					Pattern: "donotexist",
				},
			},
			wantResult: "",
			wantErr:    true,
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {

			// Init gitea object
			g, gotErr := New(tt.manifest)
			require.NoError(t, gotErr)

			gotResult, gotErr := g.Source("")

			if tt.wantErr {
				require.Error(t, gotErr)
			} else {
				require.NoError(t, gotErr)
			}

			assert.Equal(t, tt.wantResult, gotResult)

		})

	}
}
