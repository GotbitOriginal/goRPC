package logger

import (
	"bytes"
	"io"
	"strconv"
	"strings"

	"git.gotbit.io/gotbit/stroke"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func New(w io.Writer) stroke.Logger {
	const timeLayout = "2006-01-02 15:04:05"

	pe := zap.NewProductionEncoderConfig()
	pe.EncodeCaller = callerEncoder
	pe.EncodeTime = zapcore.TimeEncoderOfLayout(timeLayout)
	ce := zapcore.NewConsoleEncoder(pe)

	core := zapcore.NewCore(ce, zapcore.AddSync(w), zap.InfoLevel)

	l := zap.New(core, zap.AddCaller())
	s := l.Sugar()
	return s
}

func callerEncoder(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
	buffer := bytes.NewBuffer(make([]byte, len(caller.File)+1+5))
	buffer.WriteString(shortName(caller.Function))
	buffer.WriteRune('(')
	buffer.WriteString(caller.File)
	buffer.WriteRune(':')
	buffer.WriteString(strconv.Itoa(caller.Line))
	buffer.WriteRune(')')
	enc.AppendByteString(buffer.Bytes())
}

func shortName(qualifiedFunctionName string) string {
	name := qualifiedFunctionName

	if lastslash := strings.LastIndex(name, "/"); lastslash >= 0 {
		name = name[lastslash+1:]
	}
	if period := strings.Index(name, "."); period >= 0 {
		name = name[period+1:]
	}

	name = strings.Replace(name, "·", ".", -1)
	return name
}
