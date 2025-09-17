package mt

const (
	// ConvertBreaksNone means raw html.
	ConvertBreaksNone ConvertBreaks = "0"
	// ConvertBreaksConvert means that convert line breaks to <br />.
	ConvertBreaksConvert ConvertBreaks = "1"
	// ConvertBreaksMarkdown means Markdown.
	ConvertBreaksMarkdown ConvertBreaks = "markdown"
	// ConvertBreaksMarkdownWithSmartyPants means Markdown with Smartypants.
	ConvertBreaksMarkdownWithSmartyPants ConvertBreaks = "markdown_with_smarty_pants"
	// ConvertBreaksRichtext means Rich Text via HTML text editor.
	ConvertBreaksRichtext ConvertBreaks = "richtext"
	// ConvertBreaksTextile2 means Textile 2.
	ConvertBreaksTextile2 ConvertBreaks = "textile_2"
)

// ConvertBreaks is the value for the “convert breaks” flag for the entry.
type ConvertBreaks string

// ParseConvertBreaks returns a [ConvertBreaks] corresponding to the argument.
func ParseConvertBreaks(v string) (ConvertBreaks, error) {
	switch cb := ConvertBreaks(v); cb {
	case ConvertBreaksNone,
		ConvertBreaksConvert,
		ConvertBreaksMarkdown,
		ConvertBreaksMarkdownWithSmartyPants,
		ConvertBreaksRichtext,
		ConvertBreaksTextile2:
		return cb, nil
	default:
		return "", &InvalidConvertBreaksError{Value: v}
	}
}
