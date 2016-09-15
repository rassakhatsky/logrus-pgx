// Package logrus_pgx provides ability to use Logrus with PGX
package logrus_pgx

import (
	"github.com/Sirupsen/logrus"
)

// pgxLogger type, used to extend standart logrus logger.
type PgxLogger logrus.Logger

//Print and format debug message using logrus.
func (w *PgxLogger) Debug(msg string, vars ...interface{}) {
	f := logrus.Fields{}
	for i := 0; i < len(vars)/2; i++ {
		f[vars[i*2].(string)] = vars[i*2+1]
	}
	(*logrus.Logger)(w).WithFields(f).Debug(msg)
}

//Print and format error message using logrus.
func (w *PgxLogger) Error(msg string, vars ...interface{}) {
	f := logrus.Fields{}
	for i := 0; i < len(vars)/2; i++ {
		f[vars[i*2].(string)] = vars[i*2+1]
	}
	(*logrus.Logger)(w).WithFields(f).Error(msg)
}

//Print and format info message using logrus.
func (w *PgxLogger) Info(msg string, vars ...interface{}) {
	f := logrus.Fields{}
	for i := 0; i < len(vars)/2; i++ {
		f[vars[i*2].(string)] = vars[i*2+1]
	}
	(*logrus.Logger)(w).WithFields(f).Info(msg)
}

//Print and format warning message using logrus.
func (w *PgxLogger) Warn(msg string, vars ...interface{}) {
	f := logrus.Fields{}
	for i := 0; i < len(vars)/2; i++ {
		f[vars[i*2].(string)] = vars[i*2+1]
	}
	(*logrus.Logger)(w).WithFields(f).Warn(msg)
}
