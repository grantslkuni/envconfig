package envconfig

import (
	"errors"
	"fmt"
	"os"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

var ErrInvalidSpecification = errors.New("specification must be a struct pointer or map")

type setterFunc func(target reflect.Value, tokens ...string) error

var camelCase = regexp.MustCompile("[A-Z][^A-Z]*")

type Variable[PATTERN any] struct {
	pattern PATTERN
	set     setterFunc
}

func (v Variable[PATTERN]) String() string {
	return fmt.Sprint(v.pattern)
}

type Options struct {
	Prefix    string
	MatchCase bool

	Map   MapOptions
	Slice SliceOptions

	Formatters []Formatter
}

type MapOptions struct {
	KeyPattern        string
	EntrySeparator    string
	KeyValueSeparator string
}

type SliceOptions struct {
	IndexPattern     string
	ElementSeparator string
	FirstIndex       int
}

type SplitFunction func(string) []string

type JoinFunction func([]string) string

type Formatter struct {
	Split SplitFunction
	Join  JoinFunction
}

func DefaultOptions() Options {
	return Options{
		Map: MapOptions{
			KeyPattern:        "(.+)",
			EntrySeparator:    ",",
			KeyValueSeparator: ":",
		},
		Slice: SliceOptions{
			IndexPattern:     "([0-9]+)",
			ElementSeparator: ",",
		},
		Formatters: []Formatter{
			{
				Split: func(name string) []string {
					return []string{name}
				},
				Join: func(names []string) string {
					return strings.Join(names, "_")
				},
			},
			{
				Split: func(name string) []string {
					return camelCase.FindAllString(name, -1)
				},
				Join: func(names []string) string {
					return strings.ToUpper(strings.Join(names, "_"))
				},
			},
		},
	}
}

func Init(spec any) error {
	return InitWithOptions(spec, DefaultOptions())
}

func InitWithOptions(spec any, options Options) error {
	target := reflect.ValueOf(spec)

	if target.Kind() != reflect.Pointer {
		return ErrInvalidSpecification
	}

	variables, templates := options.collectVariables(target.Type())

	for _, variable := range variables {
		value := os.Getenv(variable.pattern)
		if value != "" {
			err := variable.set(target, value)
			if err != nil {
				return err
			}
		}
	}

	if len(templates) > 0 {
		for key, value := range getEnvironment() {
			for _, template := range templates {
				tokens := template.pattern.FindStringSubmatch(key)
				if len(tokens) > 0 {
					tokens = append(tokens[1:], value)
					err := template.set(target, tokens...)
					if err != nil {
						return err
					}
				}
			}
		}
	}

	return nil
}

func getEnvironment() map[string]string {
	variables := make(map[string]string)
	for _, variable := range os.Environ() {
		tokens := strings.SplitN(variable, "=", 2)
		variables[tokens[0]] = tokens[1]
	}
	return variables
}

func (o Options) collectVariables(spec reflect.Type) ([]Variable[string], []Variable[*regexp.Regexp]) {
	variables := make([]Variable[string], 0)
	templates := make([]Variable[*regexp.Regexp], 0)

	o.analyze(spec, func(setter setterFunc, fragments ...fragment) {
		if o.Prefix != "" {
			fragments = append(fragments, fragment{o.Prefix, false})
		}

		for _, formatter := range o.Formatters {
			pattern, dynamic := format(formatter, fragments)
			//fmt.Printf("Variable: %q (dynamic: %v)\n", name, dynamic)

			if dynamic {
				pattern = "^" + pattern + "$"
				if !o.MatchCase {
					pattern = "(?i)" + pattern
				}
				templates = append(templates, Variable[*regexp.Regexp]{
					pattern: regexp.MustCompile(pattern),
					set:     setter,
				})
			} else {
				variables = append(variables, Variable[string]{
					pattern: pattern,
					set:     setter,
				})
			}
		}

	})

	return variables, templates
}

func format(formatter Formatter, fragments []fragment) (string, bool) {
	tokens := make([]string, 0)
	dynamic := false

	for i := len(fragments) - 1; i >= 0; i-- {
		f := fragments[i]
		if f.dynamic {
			dynamic = true
			tokens = append(tokens, f.pattern)
		} else {
			tokens = append(tokens, formatter.Split(f.pattern)...)
		}
	}

	return formatter.Join(tokens), dynamic
}

type fragment struct {
	pattern string
	dynamic bool
}

func (o Options) analyze(spec reflect.Type, collect func(setterFunc, ...fragment)) {
	switch spec.Kind() {
	case reflect.Ptr:
		o.analyzePointer(spec, collect)
	case reflect.Struct:
		o.analyzeStruct(spec, collect)
	case reflect.Map:
		o.analyzeMap(spec, collect)
	case reflect.Slice:
		o.analyzeSlice(spec, collect)
	default:
		// do nothing
	}
}

func (o Options) analyzePointer(spec reflect.Type, collect func(setterFunc, ...fragment)) {
	o.analyze(spec.Elem(), func(set setterFunc, fragments ...fragment) {
		setter := func(target reflect.Value, values ...string) error {
			if target.Kind() != reflect.Ptr {
				return fmt.Errorf("invalid type: expected %s but got %s", reflect.Ptr, target.Kind())
			}
			if target.IsNil() {
				target.Set(reflect.New(target.Type().Elem()))
			}
			return set(target.Elem(), values...)
		}

		collect(setter, fragments...)
	})
}

func (o Options) analyzeStruct(spec reflect.Type, collect func(setterFunc, ...fragment)) {
	for i := 0; i < spec.NumField(); i++ {
		index := i
		field := spec.Field(index)

		if field.IsExported() {
			if isPrimitive(field.Type) {
				setter := func(target reflect.Value, values ...string) error {
					return setPrimitive(target.Field(index), values[0])
				}
				collect(setter, fragment{field.Name, false})
			} else {
				if isPrimitiveMap(field.Type) {
					setter := func(target reflect.Value, values ...string) error {
						return o.setPrimitiveMap(field.Type, target.Field(index), values[0])
					}
					collect(setter, fragment{field.Name, false})
				}

				if isPrimitiveSlice(field.Type) {
					setter := func(target reflect.Value, values ...string) error {
						return o.setPrimitiveSlice(field.Type, target.Field(index), values[0])
					}
					collect(setter, fragment{field.Name, false})
				}

				o.analyze(field.Type, func(set setterFunc, fragments ...fragment) {
					setter := func(target reflect.Value, values ...string) error {
						return set(target.Field(index), values...)
					}
					if field.Anonymous {
						collect(setter, fragments...)
					}
					collect(setter, append(fragments, fragment{field.Name, false})...)
				})
			}
		}
	}
}

func (o Options) analyzeMap(spec reflect.Type, collect func(setterFunc, ...fragment)) {
	keySpec := spec.Key()
	valueSpec := spec.Elem()
	if isPrimitive(keySpec) {
		_collect := func(set setterFunc, fragments ...fragment) {
			setter := func(target reflect.Value, values ...string) error {
				if target.Kind() != reflect.Map {
					return fmt.Errorf("invalid type: expected %s but got %s", reflect.Map, target.Kind())
				}

				keyElem := reflect.New(keySpec).Elem()
				if err := setPrimitive(keyElem, values[0]); err != nil {
					return err
				}

				valueElem := reflect.New(valueSpec).Elem()
				if err := set(valueElem, values[1:]...); err != nil {
					return err
				}

				if target.IsNil() {
					target.Set(reflect.MakeMap(target.Type()))
				}

				target.SetMapIndex(keyElem, valueElem)

				return nil
			}
			collect(setter, fragments...)
		}

		if isPrimitive(valueSpec) {
			setter := func(target reflect.Value, values ...string) error {
				return setPrimitive(target, values[0])
			}
			_collect(setter, fragment{o.Map.KeyPattern, true})
		} else {
			o.analyze(valueSpec, func(setter setterFunc, fragments ...fragment) {
				_collect(setter, append(fragments, fragment{o.Map.KeyPattern, true})...)
			})
		}
	}
}

func (o Options) analyzeSlice(spec reflect.Type, collect func(setterFunc, ...fragment)) {
	elementSpec := spec.Elem()
	_collect := func(set setterFunc, fragments ...fragment) {
		setter := func(target reflect.Value, values ...string) error {
			if target.Kind() != reflect.Slice {
				return fmt.Errorf("invalid type: expected %s but got %s", reflect.Slice, target.Kind())
			}

			index, err := strconv.Atoi(values[0])
			if err != nil {
				return err
			}

			length := index - o.Slice.FirstIndex + 1
			capacity := 16

			for capacity < length {
				capacity *= 2
			}

			if target.IsNil() {
				target.Set(reflect.MakeSlice(target.Type(), length, capacity))
			}

			if capacity > target.Cap() {
				slice := reflect.MakeSlice(target.Type(), target.Len(), capacity)
				reflect.Copy(slice, target)
				target.Set(slice)
			}

			element := target.Index(index - o.Slice.FirstIndex)

			if err := set(element, values[1:]...); err != nil {
				return err
			}

			return nil
		}
		collect(setter, fragments...)
	}

	if isPrimitive(elementSpec) {
		setter := func(target reflect.Value, values ...string) error {
			return setPrimitive(target, values[0])
		}
		_collect(setter, fragment{o.Slice.IndexPattern, true})
	} else {
		o.analyze(elementSpec, func(setter setterFunc, fragments ...fragment) {
			_collect(setter, append(fragments, fragment{o.Slice.IndexPattern, true})...)
		})
	}
}

func indirect(spec reflect.Type) reflect.Type {
	for spec.Kind() == reflect.Ptr {
		spec = spec.Elem()
	}

	return spec
}

func isPrimitive(spec reflect.Type) bool {
	spec = indirect(spec)

	return spec.Kind() == reflect.String ||
		spec.Kind() >= reflect.Bool && spec.Kind() <= reflect.Complex128
}

func isPrimitiveMap(spec reflect.Type) bool {
	spec = indirect(spec)

	return spec.Kind() == reflect.Map && isPrimitive(spec.Key()) && isPrimitive(spec.Elem())
}

func isPrimitiveSlice(spec reflect.Type) bool {
	spec = indirect(spec)

	return spec.Kind() == reflect.Slice && isPrimitive(spec.Elem())
}

func setPrimitive(target reflect.Value, value string) error {
	for target.Kind() == reflect.Ptr {
		if target.IsNil() {
			target.Set(reflect.New(target.Type().Elem()))
		}
		target = target.Elem()
	}

	switch target.Kind() {
	case reflect.Bool:
		boolValue, err := strconv.ParseBool(value)
		if err != nil {
			return err
		}
		target.SetBool(boolValue)
	case reflect.String:
		target.SetString(value)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		intValue, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return err
		}
		target.SetInt(intValue)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		uintValue, err := strconv.ParseUint(value, 10, 64)
		if err != nil {
			return err
		}
		target.SetUint(uintValue)
	case reflect.Float32, reflect.Float64:
		floatValue, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return err
		}
		target.SetFloat(floatValue)
	case reflect.Complex64, reflect.Complex128:
		complexValue, err := strconv.ParseComplex(value, 64)
		if err != nil {
			return err
		}
		target.SetComplex(complexValue)
	default:
		return fmt.Errorf("invalid type: %s", target.Kind())
	}
	return nil
}

func (o Options) setPrimitiveMap(spec reflect.Type, target reflect.Value, token string) error {
	pairs := strings.Split(token, o.Map.EntrySeparator)
	target.Set(reflect.MakeMap(spec))

	for _, pair := range pairs {
		tokens := strings.SplitN(pair, o.Map.KeyValueSeparator, 2)

		key := reflect.New(spec.Key()).Elem()
		if err := setPrimitive(key, strings.TrimSpace(tokens[0])); err != nil {
			return err
		}

		value := reflect.New(spec.Elem()).Elem()
		if err := setPrimitive(value, strings.TrimSpace(tokens[1])); err != nil {
			return err
		}

		target.SetMapIndex(key, value)
	}

	return nil
}

func (o Options) setPrimitiveSlice(spec reflect.Type, target reflect.Value, token string) error {
	values := strings.Split(token, o.Slice.ElementSeparator)
	target.Set(reflect.MakeSlice(spec, len(values), len(values)))

	for index, element := range values {
		if err := setPrimitive(target.Index(index), strings.TrimSpace(element)); err != nil {
			return err
		}
	}

	return nil
}
