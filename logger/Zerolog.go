//  Copyright DataWiseHQ/grule-rule-engine Authors
//
//  Licensed under the Apache License, Version 2.0 (the "License");
//  you may not use this file except in compliance with the License.
//  You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS,
//  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  See the License for the specific language governing permissions and
//  limitations under the License.

package logger

import (
	"fmt"
	"github.com/rs/zerolog"
)

var _ Logger = (*zeroLogger)(nil)

type zeroLogger struct {
	logger *zerolog.Logger
}

func NewZeroLog(logger *zerolog.Logger) LogEntry {
	l := zeroLogger{logger: logger}
	return l.WithFields(Fields{"lib": "grule-rule-engine"})
}

func (l *zeroLogger) Debug(args ...interface{}) {
	if l.logger.GetLevel() <= zerolog.DebugLevel {
		l.logger.Debug().Msg(fmt.Sprint(args...))
	}
}

func (l *zeroLogger) Info(args ...interface{}) {
	if l.logger.GetLevel() <= zerolog.InfoLevel {
		l.logger.Info().Msg(fmt.Sprint(args...))
	}
}

func (l *zeroLogger) Warn(args ...interface{}) {
	if l.logger.GetLevel() <= zerolog.WarnLevel {
		l.logger.Warn().Msg(fmt.Sprint(args...))
	}
}

func (l *zeroLogger) Error(args ...interface{}) {
	if l.logger.GetLevel() <= zerolog.ErrorLevel {
		l.logger.Error().Msg(fmt.Sprint(args...))
	}
}

func (l *zeroLogger) Panic(args ...interface{}) {
	if l.logger.GetLevel() <= zerolog.PanicLevel {
		l.logger.Panic().Msg(fmt.Sprint(args...))
	}
}

func (l *zeroLogger) Fatal(args ...interface{}) {
	if l.logger.GetLevel() <= zerolog.FatalLevel {
		l.logger.Fatal().Msg(fmt.Sprint(args...))
	}
}

func (l *zeroLogger) Debugf(template string, args ...interface{}) {
	if l.logger.GetLevel() <= zerolog.DebugLevel {
		l.logger.Debug().Msgf(template, args...)
	}
}

func (l *zeroLogger) Infof(template string, args ...interface{}) {
	if l.logger.GetLevel() <= zerolog.InfoLevel {
		l.logger.Info().Msgf(template, args...)
	}
}

func (l *zeroLogger) Warnf(template string, args ...interface{}) {
	if l.logger.GetLevel() <= zerolog.WarnLevel {
		l.logger.Warn().Msgf(template, args...)
	}
}

func (l *zeroLogger) Errorf(template string, args ...interface{}) {
	if l.logger.GetLevel() <= zerolog.ErrorLevel {
		l.logger.Error().Msgf(template, args...)
	}
}

func (l *zeroLogger) Panicf(template string, args ...interface{}) {
	if l.logger.GetLevel() <= zerolog.PanicLevel {
		l.logger.Panic().Msgf(template, args...)
	}
}

func (l *zeroLogger) Fatalf(template string, args ...interface{}) {
	if l.logger.GetLevel() <= zerolog.FatalLevel {
		l.logger.Fatal().Msgf(template, args...)
	}
}

func (l *zeroLogger) Trace(args ...interface{}) {
	if l.logger.GetLevel() <= zerolog.TraceLevel {
		l.logger.Trace().Msg(fmt.Sprint(args...))
	}
}

func (l *zeroLogger) Tracef(template string, args ...interface{}) {
	if l.logger.GetLevel() <= zerolog.TraceLevel {
		l.logger.Trace().Msgf(template, args...)
	}
}

func (l *zeroLogger) Print(args ...interface{}) {
	if l.logger.GetLevel() <= zerolog.InfoLevel {
		l.logger.Info().Msg(fmt.Sprint(args...))
	}
}

func (l *zeroLogger) Println(args ...interface{}) {
	if l.logger.GetLevel() <= zerolog.InfoLevel {
		l.logger.Info().Msg(fmt.Sprintln(args...))
	}
}

func (l *zeroLogger) Printf(format string, args ...interface{}) {
	if l.logger.GetLevel() <= zerolog.InfoLevel {
		l.logger.Info().Msgf(format, args...)
	}
}

func (l *zeroLogger) WithFields(fields Fields) LogEntry {
	context := l.logger.With()

	for k, v := range fields {
		context = context.Interface(k, v)
	}

	newLogger := context.Logger()

	return LogEntry{
		Logger: &zeroLogger{
			logger: &newLogger,
		},
		Level: convertZeroLogToInternalLevel(newLogger.GetLevel()),
	}
}
func convertZeroLogToInternalLevel(level zerolog.Level) Level {
	switch level {
	case zerolog.TraceLevel:

		return TraceLevel
	case zerolog.DebugLevel:

		return DebugLevel
	case zerolog.InfoLevel:

		return InfoLevel
	case zerolog.WarnLevel:

		return WarnLevel
	case zerolog.ErrorLevel:

		return ErrorLevel
	case zerolog.FatalLevel:

		return FatalLevel
	case zerolog.PanicLevel:
		return PanicLevel
	default:

		return DebugLevel
	}
}
