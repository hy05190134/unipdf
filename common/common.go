//
// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// This is a commercial product and requires a license to operate.
// A trial license can be obtained at https://unidoc.io
//
// DO NOT EDIT: generated by unitwist Go source code obfuscator.
//
// Use of this source code is governed by the UniDoc End User License Agreement
// terms that can be accessed at https://unidoc.io/eula/

// Package common contains common properties used by the subpackages.
package common ;import (_f "fmt";_ag "io";_cc "os";_c "path/filepath";_a "runtime";_e "time";);const _fcf =22;

// Warning does nothing for dummy logger.
func (DummyLogger )Warning (format string ,args ...interface{}){};

// Trace does nothing for dummy logger.
func (DummyLogger )Trace (format string ,args ...interface{}){};

// Trace logs trace message.
func (_aea WriterLogger )Trace (format string ,args ...interface{}){if _aea .LogLevel >=LogLevelTrace {_fbg :="\u005b\u0054\u0052\u0041\u0043\u0045\u005d\u0020";_aea .logToWriter (_aea .Output ,_fbg ,format ,args ...);};};const _fce =20;func (_cd ConsoleLogger )output (_fc _ag .Writer ,_gbc string ,_efa string ,_ea ...interface{}){_cb (_fc ,_gbc ,_efa ,_ea ...);};

// SetLogger sets 'logger' to be used by the unidoc unipdf library.
func SetLogger (logger Logger ){Log =logger };

// DummyLogger does nothing.
type DummyLogger struct{};

// LogLevel is the verbosity level for logging.
type LogLevel int ;func UtcTimeFormat (t _e .Time )string {return t .Format (_fgc )+"\u0020\u0055\u0054\u0043"};const _fgc ="\u0032\u0020\u004aan\u0075\u0061\u0072\u0079\u0020\u0032\u0030\u0030\u0036\u0020\u0061\u0074\u0020\u0031\u0035\u003a\u0030\u0034";

// Debug logs debug message.
func (_agd WriterLogger )Debug (format string ,args ...interface{}){if _agd .LogLevel >=LogLevelDebug {_ee :="\u005b\u0044\u0045\u0042\u0055\u0047\u005d\u0020";_agd .logToWriter (_agd .Output ,_ee ,format ,args ...);};};

// Trace logs trace message.
func (_gf ConsoleLogger )Trace (format string ,args ...interface{}){if _gf .LogLevel >=LogLevelTrace {_agf :="\u005b\u0054\u0052\u0041\u0043\u0045\u005d\u0020";_gf .output (_cc .Stdout ,_agf ,format ,args ...);};};

// Logger is the interface used for logging in the unipdf package.
type Logger interface{Error (_fa string ,_ed ...interface{});Warning (_ef string ,_g ...interface{});Notice (_cg string ,_bc ...interface{});Info (_ba string ,_be ...interface{});Debug (_gc string ,_gb ...interface{});Trace (_beg string ,_fae ...interface{});IsLogLevel (_d LogLevel )bool ;};

// IsLogLevel returns true if log level is greater or equal than `level`.
// Can be used to avoid resource intensive calls to loggers.
func (_fg ConsoleLogger )IsLogLevel (level LogLevel )bool {return _fg .LogLevel >=level };const _cac =2020;

// NewWriterLogger creates new 'writer' logger.
func NewWriterLogger (logLevel LogLevel ,writer _ag .Writer )*WriterLogger {_gcc :=WriterLogger {Output :writer ,LogLevel :logLevel };return &_gcc ;};var Log Logger =DummyLogger {};

// Error does nothing for dummy logger.
func (DummyLogger )Error (format string ,args ...interface{}){};

// IsLogLevel returns true from dummy logger.
func (DummyLogger )IsLogLevel (level LogLevel )bool {return true };

// Info does nothing for dummy logger.
func (DummyLogger )Info (format string ,args ...interface{}){};

// WriterLogger is the logger that writes data to the Output writer
type WriterLogger struct{LogLevel LogLevel ;Output _ag .Writer ;};const (LogLevelTrace LogLevel =5;LogLevelDebug LogLevel =4;LogLevelInfo LogLevel =3;LogLevelNotice LogLevel =2;LogLevelWarning LogLevel =1;LogLevelError LogLevel =0;);

// Info logs info message.
func (_fbb WriterLogger )Info (format string ,args ...interface{}){if _fbb .LogLevel >=LogLevelInfo {_fade :="\u005bI\u004e\u0046\u004f\u005d\u0020";_fbb .logToWriter (_fbb .Output ,_fade ,format ,args ...);};};

// IsLogLevel returns true if log level is greater or equal than `level`.
// Can be used to avoid resource intensive calls to loggers.
func (_ff WriterLogger )IsLogLevel (level LogLevel )bool {return _ff .LogLevel >=level };func _cb (_egf _ag .Writer ,_edd string ,_eee string ,_bag ...interface{}){_ ,_bec ,_gg ,_cbd :=_a .Caller (3);if !_cbd {_bec ="\u003f\u003f\u003f";_gg =0;}else {_bec =_c .Base (_bec );};_ge :=_f .Sprintf ("\u0025s\u0020\u0025\u0073\u003a\u0025\u0064 ",_edd ,_bec ,_gg )+_eee +"\u000a";_f .Fprintf (_egf ,_ge ,_bag ...);};

// Debug does nothing for dummy logger.
func (DummyLogger )Debug (format string ,args ...interface{}){};

// NewConsoleLogger creates new console logger.
func NewConsoleLogger (logLevel LogLevel )*ConsoleLogger {return &ConsoleLogger {LogLevel :logLevel }};func (_bf WriterLogger )logToWriter (_dd _ag .Writer ,_af string ,_age string ,_agfg ...interface{}){_cb (_dd ,_af ,_age ,_agfg );};const Version ="\u0033\u002e\u0031\u0032\u002e\u0030";const _bac =9;var ReleasedAt =_e .Date (_cac ,_bac ,_fce ,_fcf ,_gde ,0,0,_e .UTC );

// Warning logs warning message.
func (_agb ConsoleLogger )Warning (format string ,args ...interface{}){if _agb .LogLevel >=LogLevelWarning {_bg :="\u005b\u0057\u0041\u0052\u004e\u0049\u004e\u0047\u005d\u0020";_agb .output (_cc .Stdout ,_bg ,format ,args ...);};};

// Warning logs warning message.
func (_aa WriterLogger )Warning (format string ,args ...interface{}){if _aa .LogLevel >=LogLevelWarning {_bb :="\u005b\u0057\u0041\u0052\u004e\u0049\u004e\u0047\u005d\u0020";_aa .logToWriter (_aa .Output ,_bb ,format ,args ...);};};

// ConsoleLogger is a logger that writes logs to the 'os.Stdout'
type ConsoleLogger struct{LogLevel LogLevel ;};

// Debug logs debug message.
func (_ca ConsoleLogger )Debug (format string ,args ...interface{}){if _ca .LogLevel >=LogLevelDebug {_de :="\u005b\u0044\u0045\u0042\u0055\u0047\u005d\u0020";_ca .output (_cc .Stdout ,_de ,format ,args ...);};};

// Notice logs notice message.
func (_dc ConsoleLogger )Notice (format string ,args ...interface{}){if _dc .LogLevel >=LogLevelNotice {_fb :="\u005bN\u004f\u0054\u0049\u0043\u0045\u005d ";_dc .output (_cc .Stdout ,_fb ,format ,args ...);};};

// Error logs error message.
func (_gd ConsoleLogger )Error (format string ,args ...interface{}){if _gd .LogLevel >=LogLevelError {_ce :="\u005b\u0045\u0052\u0052\u004f\u0052\u005d\u0020";_gd .output (_cc .Stdout ,_ce ,format ,args ...);};};

// Notice does nothing for dummy logger.
func (DummyLogger )Notice (format string ,args ...interface{}){};const _gde =10;

// Info logs info message.
func (_fad ConsoleLogger )Info (format string ,args ...interface{}){if _fad .LogLevel >=LogLevelInfo {_bcc :="\u005bI\u004e\u0046\u004f\u005d\u0020";_fad .output (_cc .Stdout ,_bcc ,format ,args ...);};};

// Notice logs notice message.
func (_cf WriterLogger )Notice (format string ,args ...interface{}){if _cf .LogLevel >=LogLevelNotice {_eg :="\u005bN\u004f\u0054\u0049\u0043\u0045\u005d ";_cf .logToWriter (_cf .Output ,_eg ,format ,args ...);};};

// Error logs error message.
func (_cee WriterLogger )Error (format string ,args ...interface{}){if _cee .LogLevel >=LogLevelError {_ae :="\u005b\u0045\u0052\u0052\u004f\u0052\u005d\u0020";_cee .logToWriter (_cee .Output ,_ae ,format ,args ...);};};