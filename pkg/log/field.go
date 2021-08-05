package log

import (
	"fmt"
	"math"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc/status"
)

// Field is an alias for Field. Aliasing this type dramatically
// improves the navigability of this package's API documentation.
type Field = zap.Field

var (
	// Skip constructs a no-op field, which is often useful when handling invalid
	// inputs in other Field constructors.
	Skip = zap.Skip

	// Binary constructs a field that carries an opaque binary blob.
	//
	// Binary data is serialized in an encoding-appropriate format. For example,
	// zap's JSON encoder base64-encodes binary blobs. To log UTF-8 encoded text,
	// use ByteString.
	Binary = zap.Binary

	// Bool constructs a field that carries a bool.
	Bool = zap.Bool

	// ByteString constructs a field that carries UTF-8 encoded text as a []byte.
	// To log opaque binary blobs (which aren't necessarily valid UTF-8), use
	// Binary.
	ByteString = zap.ByteString

	// Complex128 constructs a field that carries a complex number. Unlike most
	// numeric fields, this costs an allocation (to convert the complex128 to
	// interface{}).
	Complex128 = zap.Complex128

	// Complex64 constructs a field that carries a complex number. Unlike most
	// numeric fields, this costs an allocation (to convert the complex64 to
	// interface{}).
	Complex64 = zap.Complex64

	// Float64 constructs a field that carries a float64. The way the
	// floating-point value is represented is encoder-dependent, so marshaling is
	// necessarily lazy.
	Float64 = zap.Float64

	// Float32 constructs a field that carries a float32. The way the
	// floating-point value is represented is encoder-dependent, so marshaling is
	// necessarily lazy.
	Float32 = zap.Float32

	// Int constructs a field with the given key and value.
	Int = zap.Int

	// Int64 constructs a field with the given key and value.
	Int64 = zap.Int64

	// Int32 constructs a field with the given key and value.
	Int32 = zap.Int32

	// Int16 constructs a field with the given key and value.
	Int16 = zap.Int16

	// Int8 constructs a field with the given key and value.
	Int8 = zap.Int8

	// String constructs a field with the given key and value.
	String = zap.String

	// Uint constructs a field with the given key and value.
	Uint = zap.Uint

	// Uint64 constructs a field with the given key and value.
	Uint64 = zap.Uint64

	// Uint32 constructs a field with the given key and value.
	Uint32 = zap.Uint32

	// Uint16 constructs a field with the given key and value.
	Uint16 = zap.Uint16

	// Uint8 constructs a field with the given key and value.
	Uint8 = zap.Uint8

	// Uintptr constructs a field with the given key and value.
	Uintptr = zap.Uintptr

	// Time constructs a Field with the given key and value. The encoder
	// controls how the time is serialized.
	Time = zap.Time

	// Duration constructs a field with the given key and value. The encoder
	// controls how the duration is serialized.
	Duration = zap.Duration

	// Any takes a key and an arbitrary value and chooses the best way to represent
	// them as a field, falling back to a reflection-based approach only if
	// necessary.
	//
	// Since byte/uint8 and rune/int32 are aliases, Any can't differentiate between
	// them. To minimize surprises, []byte values are treated as binary blobs, byte
	// values are treated as uint8, and runes are always treated as integers.
	Any = zap.Any

	// Error is shorthand for the common idiom NamedError("error", err).
	Error = zap.Error

	// NamedError constructs a field that lazily stores err.Error() under the
	// provided key. Errors which also implement fmt.Formatter (like those produced
	// by github.com/pkg/errors) will also have their verbose representation stored
	// under key+"Verbose". If passed a nil error, the field is a no-op.
	//
	// For the common case in which the key is simply "error", the Error function
	// is shorter and less repetitive.
	NamedError = zap.NamedError

	// Bools constructs a field that carries a slice of bools.
	Bools = zap.Bools

	// ByteStrings constructs a field that carries a slice of []byte, each of which
	// must be UTF-8 encoded text.
	ByteStrings = zap.ByteStrings

	// Complex128s constructs a field that carries a slice of complex numbers.
	Complex128s = zap.Complex128s

	// Complex64s constructs a field that carries a slice of complex numbers.
	Complex64s = zap.Complex64s

	// Durations constructs a field that carries a slice of time.Durations.
	Durations = zap.Durations

	// Float64s constructs a field that carries a slice of floats.
	Float64s = zap.Float64s

	// Float32s constructs a field that carries a slice of floats.
	Float32s = zap.Float32s

	// Ints constructs a field that carries a slice of integers.
	Ints = zap.Ints

	// Int64s constructs a field that carries a slice of integers.
	Int64s = zap.Int64s

	// Int32s constructs a field that carries a slice of integers.
	Int32s = zap.Int32s

	// Int16s constructs a field that carries a slice of integers.
	Int16s = zap.Int16s

	// Int8s constructs a field that carries a slice of integers.
	Int8s = zap.Int8s

	// Strings constructs a field that carries a slice of strings.
	Strings = zap.Strings

	// Times constructs a field that carries a slice of time.Times.
	Times = zap.Times

	// Uints constructs a field that carries a slice of unsigned integers.
	Uints = zap.Uints

	// Uint64s constructs a field that carries a slice of unsigned integers.
	Uint64s = zap.Uint64s

	// Uint32s constructs a field that carries a slice of unsigned integers.
	Uint32s = zap.Uint32s

	// Uint16s constructs a field that carries a slice of unsigned integers.
	Uint16s = zap.Uint16s

	// Uint8s constructs a field that carries a slice of unsigned integers.
	Uint8s = zap.Uint8s

	// Uintptrs constructs a field that carries a slice of pointer addresses.
	Uintptrs = zap.Uintptrs

	// Errors constructs a field that carries a slice of errors.
	Errors = zap.Errors
)

// Stringer safe nil stringer
func Stringer(key string, val fmt.Stringer) Field {
	if val == nil {
		return String(key, "nil")
	}
	return zap.Stringer(key, val)
}

// Fields map[string]interface{} to []Fields
func Fields(tags map[string]interface{}) (ret []Field) {
	for k, v := range tags {
		ret = append(ret, Any(k, v))
	}
	return ret
}

// FieldsToKV []Field to (k, v) []interface{}
func FieldsToKV(fields []Field) (ret []interface{}) {
	if len(fields) == 0 {
		return nil
	}

	for _, field := range fields {
		key, value := ExtractField(field)
		if value == nil {
			continue
		}

		ret = append(ret, key, value)
	}
	return ret
}

// ExtractField Field to string, interface{}
func ExtractField(f Field) (string, interface{}) {
	switch f.Type {
	case zapcore.ArrayMarshalerType,
		zapcore.ObjectMarshalerType,
		zapcore.ByteStringType,
		zapcore.Complex128Type,
		zapcore.Complex64Type,
		zapcore.ReflectType,
		zapcore.StringerType,
		zapcore.ErrorType,
		zapcore.BinaryType:
		return f.Key, f.Interface

	case zapcore.Int64Type:
		return f.Key, f.Integer

	case zapcore.Int32Type,
		zapcore.Int16Type,
		zapcore.Int8Type:
		return f.Key, int32(f.Integer)

	case zapcore.Uint64Type:
		return f.Key, uint64(f.Integer)

	case zapcore.Uint32Type,
		zapcore.Uint16Type,
		zapcore.Uint8Type:
		return f.Key, uint32(f.Integer)

	case zapcore.BoolType:
		return f.Key, f.Integer == 1

	case zapcore.DurationType:
		return f.Key, time.Duration(f.Integer)

	case zapcore.Float64Type:
		return f.Key, math.Float64frombits(uint64(f.Integer))

	case zapcore.Float32Type:
		return f.Key, math.Float32frombits(uint32(f.Integer))

	case zapcore.StringType:
		return f.Key, f.String

	case zapcore.TimeType:
		vtime := time.Unix(0, f.Integer)
		if f.Interface != nil {
			vtime = vtime.In(f.Interface.(*time.Location))
		}
		return f.Key, vtime

	case zapcore.UintptrType:
		return f.Key, uintptr(f.Integer)

	case zapcore.NamespaceType:
		return f.Key, nil

	default:
		return "", nil
	}
}

// ErrorDetails if error is *google.rpc.Status then also print details
func ErrorDetails(err error) Field {
	if err == nil {
		return Skip()
	}

	var details []interface{}
	if serr, ok := status.FromError(err); ok {
		details = serr.Details()
	}
	if len(details) == 0 {
		return Skip()
	}

	return Any("error.details", details)
}
