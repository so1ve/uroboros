package identifier

var ReservedWords = map[string]string{
	"3":          "abstract boolean byte char class double enum export extends final float goto implements import int interface long native package private protected public short static super synchronized throws transient volatile",
	"5":          "class enum extends super const export import",
	"6":          "enum",
	"strict":     "implements interface let package private protected public static yield",
	"strictBind": "eval arguments",
}

const ecma5AndLessKeywords = "break case catch continue debugger default do else finally for function if return switch throw try var while with null true false instanceof typeof void delete new in this"

var Keywords = map[string]string{
	"5":       ecma5AndLessKeywords,
	"5module": ecma5AndLessKeywords + " export import",
	"6":       ecma5AndLessKeywords + " const class extends export import super",
}
