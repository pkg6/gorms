package gorms

import "testing"

func TestCDN(t *testing.T) {
	tests := []struct {
		name     string
		host     string
		path     string
		replaces bool
		want     string
	}{
		{
			name:     "path not  with url",
			host:     "https://github.com/",
			path:     "test.png",
			replaces: false,
			want:     "https://github.com/test.png",
		},
		{
			name:     "path  with url",
			host:     "https://github.com/",
			replaces: false,
			path:     "https://github.com/test.png",
			want:     "https://github.com/test.png",
		},
		{
			name:     "path replace  with url p",
			host:     "https://github.io/",
			replaces: false,
			path:     "https://github.com/test.png?id=1",
			want:     "https://github.com/test.png?id=1",
		},
		{
			name:     "path replace  with url",
			host:     "https://github.io/",
			replaces: true,
			path:     "https://github.com/test.png",
			want:     "https://github.io/test.png",
		},
		{
			name:     "path replace  with url p",
			host:     "https://github.io/",
			replaces: true,
			path:     "https://github.com/test.png?id=1",
			want:     "https://github.io/test.png?id=1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CDN(tt.host, tt.path, tt.replaces); got != tt.want {
				t.Errorf("CDN() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCDNRelativePath(t *testing.T) {
	tests := []struct {
		name    string
		cdnPath string
		want    string
	}{
		{
			name:    "t1",
			cdnPath: "https://github.io/test.png?id=1",
			want:    "/test.png?id=1",
		},
		{
			name:    "t2",
			cdnPath: "https://github.io/test.png",
			want:    "/test.png",
		},
		{
			name:    "t3",
			cdnPath: "/test.png",
			want:    "/test.png",
		},
		{
			name:    "t4",
			cdnPath: "/test.png?id=1",
			want:    "/test.png?id=1",
		},
		{
			name:    "t5",
			cdnPath: "test.png?id=1",
			want:    "test.png?id=1",
		},
		{
			name:    "t6",
			cdnPath: "test.png",
			want:    "test.png",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CDNRelativePath(tt.cdnPath); got != tt.want {
				t.Errorf("CDNRelativePath() = %v, want %v", got, tt.want)
			}
		})
	}
}
