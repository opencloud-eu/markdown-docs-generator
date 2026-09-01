package dochelpers

import "testing"

func TestRewriteRelativeLinks(t *testing.T) {
	tests := []struct {
		name    string
		content string
		service string
		branch  string
		want    string
	}{
		{
			name:    "sibling file",
			content: "see [MIGRATION.md](MIGRATION.md) for details",
			service: "search",
			branch:  "main",
			want:    "see [MIGRATION.md](https://github.com/opencloud-eu/opencloud/blob/main/services/search/MIGRATION.md) for details",
		},
		{
			name:    "parent dir",
			content: "the [frontend](../frontend) service",
			service: "webdav",
			branch:  "main",
			want:    "the [frontend](https://github.com/opencloud-eu/opencloud/blob/main/services/frontend) service",
		},
		{
			name:    "dot slash prefix",
			content: "[doc](./MIGRATION.md)",
			service: "search",
			branch:  "main",
			want:    "[doc](https://github.com/opencloud-eu/opencloud/blob/main/services/search/MIGRATION.md)",
		},
		{
			name:    "repo root",
			content: "[license](/LICENSE)",
			service: "search",
			branch:  "main",
			want:    "[license](https://github.com/opencloud-eu/opencloud/blob/main/LICENSE)",
		},
		{
			name:    "fragment kept",
			content: "[setup](MIGRATION.md#setup)",
			service: "search",
			branch:  "main",
			want:    "[setup](https://github.com/opencloud-eu/opencloud/blob/main/services/search/MIGRATION.md#setup)",
		},
		{
			name:    "title kept",
			content: `[doc](MIGRATION.md "the migration guide")`,
			service: "search",
			branch:  "main",
			want:    `[doc](https://github.com/opencloud-eu/opencloud/blob/main/services/search/MIGRATION.md "the migration guide")`,
		},
		{
			name:    "image",
			content: "![diagram](flow.png)",
			service: "search",
			branch:  "main",
			want:    "![diagram](https://github.com/opencloud-eu/opencloud/blob/main/services/search/flow.png)",
		},
		{
			name:    "branch respected",
			content: "[doc](MIGRATION.md)",
			service: "search",
			branch:  "stable-1.0",
			want:    "[doc](https://github.com/opencloud-eu/opencloud/blob/stable-1.0/services/search/MIGRATION.md)",
		},
		{
			name:    "multiple links",
			content: "[a](A.md) and [b](https://example.com) and [c](B.md)",
			service: "search",
			branch:  "main",
			want:    "[a](https://github.com/opencloud-eu/opencloud/blob/main/services/search/A.md) and [b](https://example.com) and [c](https://github.com/opencloud-eu/opencloud/blob/main/services/search/B.md)",
		},
		{
			name:    "https untouched",
			content: "[docs](https://docs.opencloud.eu/)",
			service: "search",
			branch:  "main",
			want:    "[docs](https://docs.opencloud.eu/)",
		},
		{
			name:    "anchor untouched",
			content: "[above](#configuration)",
			service: "search",
			branch:  "main",
			want:    "[above](#configuration)",
		},
		{
			name:    "mailto untouched",
			content: "[mail](mailto:alice@example.com)",
			service: "search",
			branch:  "main",
			want:    "[mail](mailto:alice@example.com)",
		},
		{
			name:    "protocol relative untouched",
			content: "[cdn](//example.com/asset.js)",
			service: "search",
			branch:  "main",
			want:    "[cdn](//example.com/asset.js)",
		},
		{
			name:    "no links",
			content: "plain text without links",
			service: "search",
			branch:  "main",
			want:    "plain text without links",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := string(rewriteRelativeLinks([]byte(tt.content), tt.service, tt.branch))
			if got != tt.want {
				t.Errorf("got  %s\nwant %s", got, tt.want)
			}
		})
	}
}
