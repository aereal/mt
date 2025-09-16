package mt

const (
	ConvertBreaksNone                    ConvertBreaks = "0"
	ConvertBreaksConvert                 ConvertBreaks = "1"
	ConvertBreaksMarkdown                ConvertBreaks = "markdown"
	ConvertBreaksMarkdownWithSmartyPants ConvertBreaks = "markdown_with_smarty_pants"
	ConvertBreaksRichtext                ConvertBreaks = "richtext"
	ConvertBreaksTextile2                ConvertBreaks = "textile_2"
)

type ConvertBreaks string

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
