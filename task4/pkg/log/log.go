package log

import (
	"os"
	"path/filepath"
	"task4/internal/config"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.Logger

func Init() error {
	cf := config.GetConfig()
	err := InitLogger("default", cf.Log.Path, cf.Log.Level, 10, 10)
	return err
}

// InitLogger 初始化日志（在项目启动时调用，如 main.go）
// 参数说明：
//   env: 环境（dev=开发环境，prod=生产环境）
//   logPath: 日志文件保存路径（如 "logs/app.log"）
//   level: 日志级别（debug/info/warn/error/dpanic/panic/fatal）
//   maxAge: 日志文件最大保存时间（天）
//   rotationTime: 日志文件轮转周期（小时）

func InitLogger(env, logPath, level string, maxAge, rotationTime int) error {
	// 1. 解析日志级别（字符串转 zapcore.Level）
	logLevel, err := parseLogLevel(level)
	if err != nil {
		return err
	}
	// 2. 配置日志编码器（输出格式）
	encoder := getEncoder(env)

	// 3. 配置日志输出目标（控制台/文件）
	var cores []zapcore.Core

	switch env {
	case "dev":
		// 开发环境：仅控制台输出（彩色、简洁格式）
		consoleCore := zapcore.NewCore(encoder, zapcore.Lock(os.Stdout), logLevel)
		cores = append(cores, consoleCore)
	case "prod":
		// 生产环境：仅文件输出（JSON 格式 + 轮转切割）
		fileWriter, err := getFileWriter(logPath, maxAge, rotationTime)
		if err != nil {
			return err
		}
		fileCore := zapcore.NewCore(encoder, zapcore.Lock(fileWriter), logLevel)
		cores = append(cores, fileCore)
	default:
		// 默认：控制台 + 文件双输出
		consoleCore := zapcore.NewCore(encoder, zapcore.Lock(os.Stdout), logLevel)
		fileWriter, err := getFileWriter(logPath, maxAge, rotationTime)
		if err != nil {
			return err
		}
		fileCore := zapcore.NewCore(encoder, zapcore.Lock(fileWriter), logLevel)
		cores = append(cores, consoleCore, fileCore)
	}

	// 4. 创建日志实例（添加调用者信息、堆栈追踪）
	Logger = zap.New(
		zapcore.NewTee(cores...),              // 多输出目标聚合
		zap.AddCaller(),                       // 输出日志调用的文件名和行号（如 main.go:23）
		zap.AddCallerSkip(1),                  // 跳过当前工具函数，显示真实调用位置
		zap.AddStacktrace(zapcore.ErrorLevel), // 仅在 Error 及以上级别输出堆栈信息
	)

	// 测试日志输出
	Logger.Info("日志初始化成功",
		zap.String("env", env),
		zap.String("log_path", logPath),
		zap.String("log_level", level),
	)

	return nil
}

// getFileWriter 配置文件输出（日志轮转 + 过期清理）
func getFileWriter(logPath string, maxAge, rotationTime int) (zapcore.WriteSyncer, error) {
	// 确保日志目录存在（如 logs/ 不存在则创建）
	logDir := filepath.Dir(logPath)
	if err := os.MkdirAll(logDir, 0755); err != nil {
		return nil, err
	}

	// 配置日志轮转规则
	writer, err := rotatelogs.New(
		logPath+".%Y%m%d%H", // 日志文件名格式（如 app.log.2023100112，按小时轮转）
		rotatelogs.WithMaxAge(time.Duration(maxAge)*24*time.Hour),          // 日志最大保存时间
		rotatelogs.WithRotationTime(time.Duration(rotationTime)*time.Hour), // 轮转周期
		rotatelogs.WithRotationCount(0),                                    // 不限制轮转文件数量（由 maxAge 控制）
	)
	if err != nil {
		return nil, err
	}

	// 返回 zap 兼容的 WriteSyncer
	return zapcore.AddSync(writer), nil
}

// getEncoder 配置日志编码器（根据环境返回不同格式）
func getEncoder(env string) zapcore.Encoder {
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",                         // 日志时间字段名
		LevelKey:       "level",                        // 日志级别字段名
		NameKey:        "logger",                       // 日志器名称字段名
		CallerKey:      "caller",                       // 调用者信息字段名
		MessageKey:     "msg",                          // 日志消息字段名
		StacktraceKey:  "stacktrace",                   // 堆栈信息字段名
		LineEnding:     zapcore.DefaultLineEnding,      // 行结束符
		EncodeDuration: zapcore.SecondsDurationEncoder, // 耗时编码（秒级）
		EncodeCaller:   zapcore.ShortCallerEncoder,     // 调用者信息格式（简短路径，如 pkg/utils/logger.go:45）
	}

	switch env {
	case "dev":
		// 开发环境：控制台彩色输出 + 人类可读格式
		encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder // 彩色级别（如 INFO 为蓝色，ERROR 为红色）
		encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder        // 时间格式：2023-10-01T12:00:00Z07:00
		return zapcore.NewConsoleEncoder(encoderConfig)              // 控制台编码器
	case "prod":
		// 生产环境：JSON 格式（便于日志收集工具解析，如 ELK）
		encoderConfig.EncodeLevel = zapcore.LowercaseLevelEncoder // 小写级别（如 info、error）
		encoderConfig.EncodeTime = zapcore.EpochTimeEncoder       // 时间格式：时间戳（秒级）
		return zapcore.NewJSONEncoder(encoderConfig)              // JSON 编码器
	default:
		// 默认：JSON 格式
		encoderConfig.EncodeLevel = zapcore.LowercaseLevelEncoder
		encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
		return zapcore.NewJSONEncoder(encoderConfig)
	}
}

// parseLogLevel 字符串日志级别转 zapcore.Level
func parseLogLevel(level string) (zapcore.Level, error) {
	switch level {
	case "debug":
		return zapcore.DebugLevel, nil
	case "info":
		return zapcore.InfoLevel, nil
	case "warn":
		return zapcore.WarnLevel, nil
	case "error":
		return zapcore.ErrorLevel, nil
	case "dpanic":
		return zapcore.DPanicLevel, nil
	case "panic":
		return zapcore.PanicLevel, nil
	case "fatal":
		return zapcore.FatalLevel, nil
	default:
		// 默认日志级别：info
		return zapcore.InfoLevel, nil
	}
}
