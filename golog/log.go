// Copyright 2014 beego Author. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package golog

import (
	"github.com/sillydong/blog4go"
	"github.com/sillydong/goczd/gofile"
	"os"
	"path"
)

func InitGoLog(filename string, level int, maxdays int) {
	if !path.IsAbs(filename) {
		workingdir, _ := gofile.WorkDir()
		if workingdir != "" {
			filename = path.Join(workingdir, filename)
		}
	}
	diroflog := path.Dir(filename)
	os.MkdirAll(diroflog, 0777)

	blog4go.NewBaseFileWriter(filename, true)
	//兼容旧库loglevel 7
	if level > 5 {
		level = 1
	}
	blog4go.SetLevel(blog4go.Levels[level])
}

func InitGoMultiLog(filedir string, level int, maxdays int) {
	if !path.IsAbs(filedir) {
		workingdir, _ := gofile.WorkDir()
		if workingdir != "" {
			filedir = path.Join(workingdir, filedir)
		}
	}
	os.MkdirAll(filedir, 0777)

	blog4go.NewFileWriter(filedir, true)
	//兼容旧库loglevel 7
	if level > 5 {
		level = 1
	}
	blog4go.SetLevel(blog4go.Levels[level])
}

func Close() {
	blog4go.Close()
}

// Trace static function for Trace
func Trace(args ...interface{}) {
	blog4go.Trace(args...)
}

// Tracef static function for Tracef
func Tracef(format string, args ...interface{}) {
	blog4go.Tracef(format, args...)
}

// Debug static function for Debug
func Debug(args ...interface{}) {
	blog4go.Debug(args...)
}

// Debugf static function for Debugf
func Debugf(format string, args ...interface{}) {
	blog4go.Debugf(format, args...)
}

// Info static function for Info
func Info(args ...interface{}) {
	blog4go.Info(args...)
}

// Infof static function for Infof
func Infof(format string, args ...interface{}) {
	blog4go.Infof(format, args...)
}

// Warn static function for Warn
func Warn(args ...interface{}) {
	blog4go.Warn(args...)
}

// Warnf static function for Warnf
func Warnf(format string, args ...interface{}) {
	blog4go.Warnf(format, args...)
}

// Error static function for Error
func Error(args ...interface{}) {
	blog4go.Error(args...)
}

// Errorf static function for Errorf
func Errorf(format string, args ...interface{}) {
	blog4go.Errorf(format, args...)
}

// Critical static function for Critical
func Critical(args ...interface{}) {
	blog4go.Critical(args...)
}

// Criticalf static function for Criticalf
func Criticalf(format string, args ...interface{}) {
	blog4go.Criticalf(format, args...)
}
