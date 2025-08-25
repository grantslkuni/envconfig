package envconfig

import (
	"regexp"
	"strings"
)

type SplitFunction = func(string) []string

type FormatFunction func(string) string

type JoinFunction func(...string) string

type Formatter interface {
	Split(value string) []string
	Join(values ...string) string
}

func FunctionalFormatter(split SplitFunction, join JoinFunction) Formatter {
	return &functionalFormatter{
		split: split,
		join:  join,
	}
}

type functionalFormatter struct {
	split SplitFunction
	join  JoinFunction
}

func (f functionalFormatter) Split(value string) []string {
	return f.split(value)
}

func (f functionalFormatter) Join(values ...string) string {
	return f.join(values...)
}

func OriginalNames() *FormatBuilder {
	return NewFormatBuilder(func(value string) []string {
		return []string{value}
	})
}

var camelCase = regexp.MustCompile("[A-Z][^A-Z]*")

func SplitCamelCase() *FormatBuilder {
	return NewFormatBuilder(func(value string) []string {
		return camelCase.FindAllString(value, -1)
	})
}

func NewFormatBuilder(split SplitFunction) *FormatBuilder {
	return &FormatBuilder{split: split}
}

type FormatBuilder struct {
	split  SplitFunction
	format []FormatFunction
}

func (b *FormatBuilder) Each(format FormatFunction) *FormatBuilder {
	b.format = append(b.format, format)
	return b
}

func (b *FormatBuilder) UpperCase() *FormatBuilder {
	b.Each(strings.ToUpper)
	return b
}

func (b *FormatBuilder) LowerCase() *FormatBuilder {
	b.Each(strings.ToLower)
	return b
}

func (b *FormatBuilder) TrimSpace() *FormatBuilder {
	b.Each(strings.TrimSpace)
	return b
}

func (b *FormatBuilder) Join(join JoinFunction) Formatter {
	return functionalFormatter{
		split: b.split,
		join: func(values ...string) string {
			if len(b.format) > 0 {
				for index, value := range values {
					for _, format := range b.format {
						value = format(value)
					}
					values[index] = value
				}
			}
			return join(values...)
		},
	}
}

func (b *FormatBuilder) JoinWith(separator string) Formatter {
	return b.Join(func(values ...string) string {
		return strings.Join(values, separator)
	})
}
