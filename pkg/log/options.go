package log

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/pflag"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"strings"
	"sync"
)

//@author: coachhe
//@create: 2022/9/2 13:21

var once sync.Once

const (
	flagLevel             = "log.level"
	flagDisableCaller     = "log.disable-caller"
	flagDisableStacktrace = "log.disable-stacktrace"
	flagFormat            = "log.sqlformat"
	flagEnableColor       = "log.enable-color"
	flagOutputPaths       = "log.output-paths"
	flagErrorOutputPaths  = "log.error-output-paths"
	flagDevelopment       = "log.development"
	flagName              = "log.name"

	ConsoleFormat = "console"
	JsonFormat    = "json"
)

/*
Options
日志选项
通过日志选项来配置日志的不同行为
参数：
	OutputPaths 输出到的位置,支持输出到多个文件，用逗号隔开
	Level 日志级别
    ErrorOutputPaths 错误日志输出路径，多个输出，用逗号隔开
	Formatter 输出格式，例如JSON或者Text
	DisableCaller 是否开启文件名和行号
	DisableStacktrace 是否在Panic及以上级别禁止打印堆栈信息
	EnableColor 是否开启颜色输出
	Name Logger的名字
*/
type Options struct {
	OutputPaths       []string
	ErrorOutputPaths  []string
	LogLevel          string
	Formatter         string
	Development       bool
	DisableCaller     bool
	DisableStacktrace bool
	EnableColor       bool
	Name              string
}

/*
NewOptions
利用默认设置创建一个Options
*/
func NewOptions() *Options {
	return &Options{
		LogLevel:          zap.InfoLevel.String(),
		DisableCaller:     false,
		DisableStacktrace: false,
		Formatter:         ConsoleFormat,
		EnableColor:       true,
		Development:       true,
		OutputPaths:       []string{"stdout"},
		ErrorOutputPaths:  []string{"stderr"},
	}
}

// 格式化输出
func (o *Options) String() string {
	data, _ := json.Marshal(o)
	return string(data)
}

// Build
// 构建一个全局zap logger
func (o *Options) Build() error {
	var zapLevel zapcore.Level
	err := zapLevel.UnmarshalText([]byte(o.LogLevel))
	// 默认为Info级别
	if err != nil {
		zapLevel = zapcore.InfoLevel
	}

	encodeLevel := zapcore.CapitalLevelEncoder
	if o.Formatter == ConsoleFormat && o.EnableColor {
		encodeLevel = zapcore.CapitalColorLevelEncoder
	}

	zc := &zap.Config{
		Level:             zap.NewAtomicLevelAt(zapLevel),
		Development:       o.Development,
		DisableCaller:     o.DisableCaller,
		DisableStacktrace: o.DisableStacktrace,
		Sampling: &zap.SamplingConfig{
			Initial:    100,
			Thereafter: 100,
		},
		Encoding: o.Formatter,
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:     "message",
			LevelKey:       "level",
			TimeKey:        "timestamp",
			NameKey:        "logger",
			CallerKey:      "caller",
			StacktraceKey:  "stacktrace",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    encodeLevel,
			EncodeTime:     timeEncoder,
			EncodeDuration: milliSecondsDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
			EncodeName:     zapcore.FullNameEncoder,
		},
		OutputPaths:      o.OutputPaths,
		ErrorOutputPaths: o.ErrorOutputPaths,
	}
	logger, err := zc.Build(zap.AddStacktrace(zapcore.PanicLevel))
	if err != nil {
		return err
	}
	zap.RedirectStdLog(logger.Named(o.Name))
	zap.ReplaceGlobals(logger)

	return nil
}

// AddFlags adds flags for log to the specified FlagSet object.
func (o *Options) AddFlags(fs *pflag.FlagSet) {
	fs.StringVar(&o.LogLevel, flagLevel, o.LogLevel, "Minimum log output `LEVEL`.")
	fs.BoolVar(&o.DisableCaller, flagDisableCaller, o.DisableCaller, "Disable output of caller information in the log.")
	fs.BoolVar(&o.DisableStacktrace, flagDisableStacktrace,
		o.DisableStacktrace, "Disable the log to record a stack trace for all messages at or above panic level.")
	fs.StringVar(&o.Formatter, flagFormat, o.Formatter, "Log output `FORMAT`, support plain or json sqlformat.")
	fs.BoolVar(&o.EnableColor, flagEnableColor, o.EnableColor, "Enable output ansi colors in plain sqlformat logs.")
	fs.StringSliceVar(&o.OutputPaths, flagOutputPaths, o.OutputPaths, "Output paths of log.")
	fs.StringSliceVar(&o.ErrorOutputPaths, flagErrorOutputPaths, o.ErrorOutputPaths, "Error output paths of log.")
	fs.BoolVar(
		&o.Development,
		flagDevelopment,
		o.Development,
		"Development puts the logger in development mode, which changes "+
			"the behavior of DPanicLevel and takes stacktraces more liberally.",
	)
	fs.StringVar(&o.Name, flagName, o.Name, "The name of the logger.")
}

// Validate validate the options fields.
func (o *Options) Validate() []error {
	var errs []error

	var zapLevel zapcore.Level
	if err := zapLevel.UnmarshalText([]byte(o.LogLevel)); err != nil {
		errs = append(errs, err)
	}

	format := strings.ToLower(o.Formatter)
	if format != ConsoleFormat && format != JsonFormat {
		errs = append(errs, fmt.Errorf("not a valid log sqlformat: %q", o.Formatter))
	}

	return errs
}
