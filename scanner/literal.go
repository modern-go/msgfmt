package scanner

import (
	"github.com/modern-go/parse"
	"regexp"
)

func parseLiteral(src *parse.Source, literal string) interface{} {
	ctx := src.Attachment.(*buildingContext)
	ctx.fullRe = append(ctx.fullRe, regexp.QuoteMeta(literal)...)
	return nil
}
