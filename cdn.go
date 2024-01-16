package gorms

import (
	"net/url"
	"strings"
)

// CDN 完整的拼接成完整的路径地址
func CDN(host, path string, replaces ...bool) string {
	if strings.HasPrefix(path, "https://") ||
		strings.HasPrefix(path, "http://") {
		parse, _ := url.Parse(path)
		if len(replaces) > 0 && replaces[0] == true {
			hu, _ := url.Parse(host)
			parse.Host = hu.Host
		}
		return parse.String()
	}
	return host + path
}

// CDNRelativePath
// 将完整的路径地址做成相对路径
func CDNRelativePath(cdnPath string) string {
	if strings.HasPrefix(cdnPath, "https://") ||
		strings.HasPrefix(cdnPath, "http://") {
		parse, _ := url.Parse(cdnPath)
		q := parse.Query()
		newPath := parse.Path
		if len(q) > 0 {
			newPath += "?"
			newPath += q.Encode()
		}
		return newPath
	}
	return cdnPath
}
