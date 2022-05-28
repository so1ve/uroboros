package options

type ESType uint8

const (
	ES3 ESType = iota
	ES5
	ES6
	ES7
	ES8
	ES9
	ES10
	ES11
	ES12
	ES13
	ES2015   = ES6
	ES2016   = ES7
	ES2017   = ES8
	ES2018   = ES9
	ES2019   = ES10
	ES2020   = ES11
	ES2021   = ES12
	ES2022   = ES13
	ESLatest = ES2022
)

type SourceTypeType uint8

const (
	SourceTypeScript SourceTypeType = iota
	SourceTypeModule
)

type Options struct {
	// `ES` indicates the ECMAScript version to parse. Must be
	// either 3, 5, 6 (or 2015), 7 (2016), 8 (2017), 9 (2018), 10
	// (2019), 11 (2020), 12 (2021), 13 (2022), or `"latest"` (the
	// latest version the library supports). This influences support
	// for strict mode, the set of reserved words, and support for
	// new syntax features.
	EcmaScript ESType
	// `sourceType` indicates the mode the code should be parsed in.
	// Can be either `"script"` or `"module"`. This influences global
	// strict mode and parsing of `import` and `export` declarations.
	SourceType SourceTypeType
}

var DefaultOptions = &Options{
	EcmaVersion: ESLatest,
	SourceType:  SourceTypeScript,
}
