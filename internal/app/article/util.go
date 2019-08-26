package article

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/pthethanh/robusta/internal/pkg/uuid"
)

var (
	defaultArticleID string
)

func init() {
	defaultArticleID = articleID("")
}

// articleID return a unique article ID base on the title
func articleID(title string) string {
	v := strings.ToLower(title)
	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		// should not happen
		return strings.Replace(title, " ", "-", -1)
	}
	// append with UUID to make it unique
	v += fmt.Sprintf("_%s", uuid.New())
	v = reg.ReplaceAllString(v, "-")
	return v
}

// isArticleID check if the id is article ID is an ArticleID or just an ID
// this simply checks the length for now
func isArticleID(id string) bool {
	if len(id) >= len(defaultArticleID) {
		return true
	}
	return false
}
