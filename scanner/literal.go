package scanner

import (
	"regexp"
	"github.com/modern-go/parse"
)

func parseLiteral(src *parse.Source, literal string) interface{} {
	ctx := src.Attachment.(*buildingContext)
	ctx.fullRe = append(ctx.fullRe, regexp.QuoteMeta(literal)...)
	return nil
}
