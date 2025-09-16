package mt

const (
	FieldNone            Field = ""
	FieldAllowComments   Field = "ALLOW COMMENTS"
	FieldAllowPings      Field = "ALLOW PINGS"
	FieldAuthor          Field = "AUTHOR"
	FieldBasename        Field = "BASENAME"
	FieldBody            Field = "BODY"
	FieldCategory        Field = "CATEGORY"
	FieldComment         Field = "COMMENT"
	FieldConvertBreaks   Field = "CONVERT BREAKS"
	FieldDate            Field = "DATE"
	FieldEmail           Field = "EMAIL"
	FieldExcerpt         Field = "EXCERPT"
	FieldExtendedBody    Field = "EXTENDED BODY"
	FieldIP              Field = "IP"
	FieldPrimaryCategory Field = "PRIMARY CATEGORY"
	FieldStatus          Field = "STATUS"
	FieldTags            Field = "TAGS"
	FieldTitle           Field = "TITLE"
	FieldURL             Field = "URL"
)

type Field string
