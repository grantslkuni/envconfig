package envconfig

import (
	"os"
	"testing"

	"github.com/c2fo/testify/assert"
)

func TestBoolField(t *testing.T) {
	type TestSpec struct {
		TestField bool
	}
	withEnv("TEST_FIELD", "true", func() {
		spec := TestSpec{}
		assert.NoError(t, Init(&spec))
		assert.Equal(t, true, spec.TestField)
	})
}

func TestBoolPointerField(t *testing.T) {
	type TestSpec struct {
		TestField *bool
	}
	withEnv("TEST_FIELD", "true", func() {
		spec := TestSpec{}
		assert.NoError(t, Init(&spec))
		assert.Equal(t, true, *spec.TestField)
	})
}

func TestStringField(t *testing.T) {
	type TestSpec struct {
		TestField string
	}
	withEnv("TEST_FIELD", "test", func() {
		spec := TestSpec{}
		assert.NoError(t, Init(&spec))
		assert.Equal(t, "test", spec.TestField)
	})
}

func TestStringPointerField(t *testing.T) {
	type TestSpec struct {
		TestField *string
	}
	withEnv("TEST_FIELD", "test", func() {
		spec := TestSpec{}
		assert.NoError(t, Init(&spec))
		assert.Equal(t, "test", *spec.TestField)
	})
}

func TestIntField(t *testing.T) {
	type TestSpec struct {
		TestField int
	}
	withEnv("TEST_FIELD", "123", func() {
		spec := TestSpec{}
		assert.NoError(t, Init(&spec))
		assert.Equal(t, 123, spec.TestField)
	})
}

func TestIntPointerField(t *testing.T) {
	type TestSpec struct {
		TestField *int
	}
	withEnv("TEST_FIELD", "123", func() {
		spec := TestSpec{}
		assert.NoError(t, Init(&spec))
		assert.Equal(t, 123, *spec.TestField)
	})
}

func TestInt8Field(t *testing.T) {
	type TestSpec struct {
		TestField int8
	}
	withEnv("TEST_FIELD", "123", func() {
		spec := TestSpec{}
		assert.NoError(t, Init(&spec))
		assert.Equal(t, 123, spec.TestField)
	})
}

func TestInt8PointerField(t *testing.T) {
	type TestSpec struct {
		TestField *int8
	}
	withEnv("TEST_FIELD", "123", func() {
		spec := TestSpec{}
		assert.NoError(t, Init(&spec))
		assert.Equal(t, 123, *spec.TestField)
	})
}

func TestInt16Field(t *testing.T) {
	type TestSpec struct {
		TestField int16
	}
	withEnv("TEST_FIELD", "123", func() {
		spec := TestSpec{}
		assert.NoError(t, Init(&spec))
		assert.Equal(t, 123, spec.TestField)
	})
}

func TestInt16PointerField(t *testing.T) {
	type TestSpec struct {
		TestField *int16
	}
	withEnv("TEST_FIELD", "123", func() {
		spec := TestSpec{}
		assert.NoError(t, Init(&spec))
		assert.Equal(t, 123, *spec.TestField)
	})
}

func TestInt32Field(t *testing.T) {
	type TestSpec struct {
		TestField int32
	}
	withEnv("TEST_FIELD", "123", func() {
		spec := TestSpec{}
		assert.NoError(t, Init(&spec))
		assert.Equal(t, 123, spec.TestField)
	})
}

func TestInt32PointerField(t *testing.T) {
	type TestSpec struct {
		TestField *int32
	}
	withEnv("TEST_FIELD", "123", func() {
		spec := TestSpec{}
		assert.NoError(t, Init(&spec))
		assert.Equal(t, 123, *spec.TestField)
	})
}

func TestInt64Field(t *testing.T) {
	type TestSpec struct {
		TestField int64
	}
	withEnv("TEST_FIELD", "123", func() {
		spec := TestSpec{}
		assert.NoError(t, Init(&spec))
		assert.Equal(t, 123, spec.TestField)
	})
}

func TestInt64PointerField(t *testing.T) {
	type TestSpec struct {
		TestField *int64
	}
	withEnv("TEST_FIELD", "123", func() {
		spec := TestSpec{}
		assert.NoError(t, Init(&spec))
		assert.Equal(t, 123, *spec.TestField)
	})
}

func TestUintField(t *testing.T) {
	type TestSpec struct {
		TestField uint
	}
	withEnv("TEST_FIELD", "123", func() {
		spec := TestSpec{}
		assert.NoError(t, Init(&spec))
		assert.Equal(t, 123, spec.TestField)
	})
}

func TestUintPointerField(t *testing.T) {
	type TestSpec struct {
		TestField *uint
	}
	withEnv("TEST_FIELD", "123", func() {
		spec := TestSpec{}
		assert.NoError(t, Init(&spec))
		assert.Equal(t, 123, *spec.TestField)
	})
}

func TestUint8Field(t *testing.T) {
	type TestSpec struct {
		TestField uint8
	}
	withEnv("TEST_FIELD", "123", func() {
		spec := TestSpec{}
		assert.NoError(t, Init(&spec))
		assert.Equal(t, 123, spec.TestField)
	})
}

func TestUint8PointerField(t *testing.T) {
	type TestSpec struct {
		TestField *uint8
	}
	withEnv("TEST_FIELD", "123", func() {
		spec := TestSpec{}
		assert.NoError(t, Init(&spec))
		assert.Equal(t, 123, *spec.TestField)
	})
}

func TestUint16Field(t *testing.T) {
	type TestSpec struct {
		TestField uint16
	}
	withEnv("TEST_FIELD", "123", func() {
		spec := TestSpec{}
		assert.NoError(t, Init(&spec))
		assert.Equal(t, 123, spec.TestField)
	})
}

func TestUint16PointerField(t *testing.T) {
	type TestSpec struct {
		TestField *uint16
	}
	withEnv("TEST_FIELD", "123", func() {
		spec := TestSpec{}
		assert.NoError(t, Init(&spec))
		assert.Equal(t, 123, *spec.TestField)
	})
}

func TestUint32Field(t *testing.T) {
	type TestSpec struct {
		TestField uint32
	}
	withEnv("TEST_FIELD", "123", func() {
		spec := TestSpec{}
		assert.NoError(t, Init(&spec))
		assert.Equal(t, 123, spec.TestField)
	})
}

func TestUint32PointerField(t *testing.T) {
	type TestSpec struct {
		TestField *uint32
	}
	withEnv("TEST_FIELD", "123", func() {
		spec := TestSpec{}
		assert.NoError(t, Init(&spec))
		assert.Equal(t, 123, *spec.TestField)
	})
}

func TestUint64Field(t *testing.T) {
	type TestSpec struct {
		TestField uint64
	}
	withEnv("TEST_FIELD", "123", func() {
		spec := TestSpec{}
		assert.NoError(t, Init(&spec))
		assert.Equal(t, 123, spec.TestField)
	})
}

func TestUint64PointerField(t *testing.T) {
	type TestSpec struct {
		TestField *uint64
	}
	withEnv("TEST_FIELD", "123", func() {
		spec := TestSpec{}
		assert.NoError(t, Init(&spec))
		assert.Equal(t, 123, *spec.TestField)
	})
}

func TestFloat32Field(t *testing.T) {
	type TestSpec struct {
		TestField float32
	}
	withEnv("TEST_FIELD", "12.3", func() {
		spec := TestSpec{}
		assert.NoError(t, Init(&spec))
		assert.Equal(t, 12.3, spec.TestField)
	})
}

func TestFloat32PointerField(t *testing.T) {
	type TestSpec struct {
		TestField *float32
	}
	withEnv("TEST_FIELD", "12.3", func() {
		spec := TestSpec{}
		assert.NoError(t, Init(&spec))
		assert.Equal(t, 12.3, *spec.TestField)
	})
}

func TestFloat64Field(t *testing.T) {
	type TestSpec struct {
		TestField float64
	}
	withEnv("TEST_FIELD", "12.3", func() {
		spec := TestSpec{}
		assert.NoError(t, Init(&spec))
		assert.Equal(t, 12.3, spec.TestField)
	})
}

func TestFloat64PointerField(t *testing.T) {
	type TestSpec struct {
		TestField *float64
	}
	withEnv("TEST_FIELD", "12.3", func() {
		spec := TestSpec{}
		assert.NoError(t, Init(&spec))
		assert.Equal(t, 12.3, *spec.TestField)
	})
}

func TestComplex64Field(t *testing.T) {
	type TestSpec struct {
		TestField complex64
	}
	withEnv("TEST_FIELD", "1+2i", func() {
		spec := TestSpec{}
		assert.NoError(t, Init(&spec))
		assert.Equal(t, complex(1, 2), spec.TestField)
	})
}

func TestComplex64PointerField(t *testing.T) {
	type TestSpec struct {
		TestField *complex64
	}

	withEnv("TEST_FIELD", "1+2i", func() {
		spec := TestSpec{}
		assert.NoError(t, Init(&spec))
		assert.Equal(t, complex(1, 2), *spec.TestField)
	})
}

func TestComplex128Field(t *testing.T) {
	type TestSpec struct {
		TestField complex128
	}

	withEnv("TEST_FIELD", "1+2i", func() {
		spec := TestSpec{}
		assert.NoError(t, Init(&spec))
		assert.Equal(t, complex(1, 2), spec.TestField)
	})
}

func TestComplex128PointerField(t *testing.T) {
	type TestSpec struct {
		TestField *complex128
	}

	withEnv("TEST_FIELD", "1+2i", func() {
		spec := TestSpec{}
		assert.NoError(t, Init(&spec))
		assert.Equal(t, complex(1, 2), *spec.TestField)
	})
}

func TestAnonymousStructField(t *testing.T) {
	type ChildSpec struct {
		TestField string
	}

	type ParentSpec struct {
		ChildSpec
	}

	testCases := []struct {
		name     string
		varKey   string
		varValue string
	}{
		{
			name:     "IncludingFieldName",
			varKey:   "CHILD_SPEC_TEST_FIELD",
			varValue: "test",
		},
		{
			name:     "ExcludingFieldName",
			varKey:   "TEST_FIELD",
			varValue: "test",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			withEnv(testCase.varKey, testCase.varValue, func() {
				spec := ParentSpec{}
				assert.NoError(t, Init(&spec))
				assert.Equal(t, testCase.varValue, spec.ChildSpec.TestField)
			})
		})
	}
}

func TestChildStructField(t *testing.T) {
	type ChildSpec struct {
		TestField string
	}
	type ParentSpec struct {
		ChildField ChildSpec
	}
	withEnv("CHILD_FIELD_TEST_FIELD", "test", func() {
		spec := ParentSpec{}
		assert.NoError(t, Init(&spec))
		assert.Equal(t, "test", spec.ChildField.TestField)
	})
}

func TestParseMapValuesFromString(t *testing.T) {
	type TestSpec struct {
		MapField map[string]string
	}
	withEnv("MAP_FIELD", "first:a, second: b", func() {
		spec := TestSpec{}
		assert.NoError(t, Init(&spec))
		assert.Len(t, spec.MapField, 2)
		assert.Equal(t, "a", spec.MapField["first"])
		assert.Equal(t, "b", spec.MapField["second"])
	})
}

func TestPrimitiveMapValue(t *testing.T) {
	type TestSpec struct {
		MapField map[string]string
	}
	withEnv("MAP_FIELD_MAP_KEY", "test", func() {
		spec := TestSpec{}
		assert.NoError(t, Init(&spec))
		assert.Equal(t, "test", spec.MapField["MAP_KEY"])
	})
}

func TestStructAsMapValue(t *testing.T) {
	type ChildSpec struct {
		TestField string
	}
	type ParentSpec struct {
		MapField map[string]ChildSpec
	}
	withEnv("MAP_FIELD_MAP_KEY_TEST_FIELD", "test", func() {
		spec := ParentSpec{}
		assert.NoError(t, Init(&spec))
		assert.Equal(t, "test", spec.MapField["MAP_KEY"].TestField)
	})
}

func TestSliceOfStringAsMapValue(t *testing.T) {
	type TestSpec struct {
		MapField map[string][]string
	}
	withEnv("MAP_FIELD_MAP_KEY_5", "test", func() {
		spec := TestSpec{}
		assert.NoError(t, Init(&spec))
		assert.Equal(t, "test", spec.MapField["MAP_KEY"][5])
	})
}

func TestSliceOfStructAsMapValue(t *testing.T) {
	type ChildSpec struct {
		TestField string
	}
	type ParentSpec struct {
		MapField map[string][]ChildSpec
	}
	withEnv("MAP_FIELD_MAP_KEY_5_TEST_FIELD", "test", func() {
		spec := ParentSpec{}
		assert.NoError(t, Init(&spec))
		assert.Equal(t, "test", spec.MapField["MAP_KEY"][5].TestField)
	})
}

func TestParseSliceElementsFromString(t *testing.T) {
	type TestSpec struct {
		SliceField []string
	}
	withEnv("SLICE_FIELD", "a,b, c", func() {
		spec := TestSpec{}
		assert.NoError(t, Init(&spec))
		assert.Len(t, spec.SliceField, 3)
		assert.Equal(t, "a", spec.SliceField[0])
		assert.Equal(t, "b", spec.SliceField[1])
		assert.Equal(t, "c", spec.SliceField[2])
	})
}

func TestPrimitiveSliceElement(t *testing.T) {
	type TestSpec struct {
		SliceField []string
	}
	withEnv("SLICE_FIELD_5", "test", func() {
		spec := TestSpec{}
		assert.NoError(t, Init(&spec))
		assert.Equal(t, "test", spec.SliceField[5])
	})
}

func TestStructAsSliceElement(t *testing.T) {
	type ChildSpec struct {
		TestField string
	}
	type ParentSpec struct {
		SliceField []ChildSpec
	}
	withEnv("SLICE_FIELD_5_TEST_FIELD", "test", func() {
		spec := ParentSpec{}
		assert.NoError(t, Init(&spec))
		assert.Equal(t, "test", spec.SliceField[5].TestField)
	})
}

func TestMapOfStringAsSliceElement(t *testing.T) {
	type TestSpec struct {
		SliceField []map[string]string
	}
	withEnv("SLICE_FIELD_5_MAP_KEY", "test", func() {
		spec := TestSpec{}
		assert.NoError(t, Init(&spec))
		assert.Equal(t, "test", spec.SliceField[5]["MAP_KEY"])
	})
}

func TestMapOfStructAsSliceElement(t *testing.T) {
	type ChildSpec struct {
		TestField string
	}
	type ParentSpec struct {
		SliceField []map[string]ChildSpec
	}
	withEnv("SLICE_FIELD_5_MAP_KEY_TEST_FIELD", "test", func() {
		spec := ParentSpec{}
		assert.NoError(t, Init(&spec))
		assert.Equal(t, "test", spec.SliceField[5]["MAP_KEY"].TestField)
	})
}

func withEnv(key, value string, test func()) {
	_ = os.Setenv(key, value)
	defer func() {
		_ = os.Unsetenv(key)
	}()
	test()
}
