package wrapper

import (
	"context"
	"database/sql"
	"encoding/json"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/tarantool/go-tarantool"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

var dblayer = "db"
var method = "Db"

func NewLogger() *MyLogger {
	rawJSON := []byte(`{
	"level": "debug",
	"encoding": "json",
	"outputPaths": ["stdout", "/tmp/logs"],
	"errorOutputPaths": ["stderr"],
	"initialFields": {"foo": "bar"},
	"encoderConfig": {
	  "messageKey": "message",
	  "levelKey": "level",
	  "levelEncoder": "lowercase"
	}
  }`)

	var cfg zap.Config
	if err := json.Unmarshal(rawJSON, &cfg); err != nil {
		panic(err)
	}
	logger, err := cfg.Build()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()
	log := NewMyLogger(logger)
	return log
}

type MyLogger struct {
	Logger *zap.Logger
}

func NewMyLogger(logger *zap.Logger) *MyLogger {
	return &MyLogger{logger}
}

var Hits = prometheus.NewCounterVec(prometheus.CounterOpts{
	Name: "Hits",
}, []string{"layer", "path", "method"})

var Errors = prometheus.NewCounterVec(prometheus.CounterOpts{
	Name: "Errors",
}, []string{"layer", "status", "path"})

var Duration = prometheus.NewHistogramVec(prometheus.HistogramOpts{
	Name: "Duratin",
}, []string{"layer", "path"})

func (m *MyLogger) Err(path, method, err string, passed time.Duration) {
	m.Logger.Error("request=",
		zap.String("path", path),
		zap.String("method", method),
		zap.String("error", err),
		zap.Duration("latency", passed),
	)
}

func (m *MyLogger) Inf(path, method string, passed time.Duration) {
	m.Logger.Info("request=",
		zap.String("path", path),
		zap.String("method", method),
		zap.Duration("latency", passed),
	)

}
func (m *MyLogger) MyInsert(tr *tarantool.Connection, path string, space interface{}, tuple interface{}) (resp *tarantool.Response, err error) {
	start := time.Now()
	Hits.WithLabelValues(dblayer, path, method).Inc()
	result, err := tr.Insert(space, tuple)
	passed := time.Since(start)
	Duration.WithLabelValues(dblayer, path).Observe(float64(passed.Milliseconds()))
	m.Inf(path, method, passed)
	if err != nil {
		Errors.WithLabelValues(dblayer, err.Error(), path).Inc()
		m.Err(path, method, err.Error(), passed)
	}
	return result, err
}

func (m *MyLogger) MyDelete(tr *tarantool.Connection, path string, space interface{}, index interface{}, key interface{}) (resp *tarantool.Response, err error) {
	start := time.Now()
	Hits.WithLabelValues(dblayer, path, method).Inc()
	result, err := tr.Delete(space, index, key)
	passed := time.Since(start)
	Duration.WithLabelValues(dblayer, path).Observe(float64(passed.Milliseconds()))
	m.Inf(path, method, passed)
	if err != nil {
		Errors.WithLabelValues(dblayer, err.Error(), path).Inc()
		m.Err(path, method, err.Error(), passed)
	}
	return result, err
}

func (m *MyLogger) MySelectTyped(tr *tarantool.Connection, path string, space interface{}, index interface{}, offset uint32, limit uint32, iterator uint32, key interface{}, result interface{}) (err error) {
	start := time.Now()
	Hits.WithLabelValues(dblayer, path, method).Inc()
	err = tr.SelectTyped(space, index, offset, limit, iterator, key, result)
	passed := time.Since(start)
	Duration.WithLabelValues(dblayer, path).Observe(float64(passed.Milliseconds()))
	m.Inf(path, method, passed)
	if err != nil {
		Errors.WithLabelValues(dblayer, err.Error(), path).Inc()
		m.Err(path, method, err.Error(), passed)
	}
	return err
}

func (m *MyLogger) MyReplace(tr *tarantool.Connection, path string, space interface{}, tuple interface{}) (resp *tarantool.Response, err error) {
	start := time.Now()
	Hits.WithLabelValues(dblayer, path, method).Inc()
	result, err := tr.Replace(space, tuple)
	passed := time.Since(start)
	Duration.WithLabelValues(dblayer, path).Observe(float64(passed.Milliseconds()))
	m.Inf(path, method, passed)
	if err != nil {
		Errors.WithLabelValues(dblayer, err.Error(), path).Inc()
		m.Err(path, method, err.Error(), passed)
	}
	return result, err
}

func (m *MyLogger) MyCall(tr *tarantool.Connection, path string, functionName string, args interface{}) (resp *tarantool.Response, err error) {
	start := time.Now()
	Hits.WithLabelValues(dblayer, path, method).Inc()
	result, err := tr.Call(functionName, args)
	passed := time.Since(start)
	Duration.WithLabelValues(dblayer, path).Observe(float64(passed.Milliseconds()))
	m.Inf(path, method, passed)
	if err != nil {
		Errors.WithLabelValues(dblayer, err.Error(), path).Inc()
		m.Err(path, method, err.Error(), passed)
	}
	return result, err
}

func (m *MyLogger) MyQuery(db *sqlx.DB, path string, query string, args ...interface{}) (*sqlx.Rows, error) {
	start := time.Now()
	Hits.WithLabelValues(dblayer, path, method).Inc()
	rows, err := db.Queryx(query, args...)
	passed := time.Since(start)
	Duration.WithLabelValues(dblayer, path).Observe(float64(passed.Milliseconds()))
	m.Inf(path, method, passed)
	if err != nil {
		Errors.WithLabelValues(dblayer, err.Error(), path).Inc()
		m.Err(path, method, err.Error(), passed)
	}
	return rows, err
}

func (m *MyLogger) MySelect(db *sqlx.DB, path string, query string, dest interface{}, args ...interface{}) error {
	start := time.Now()
	Hits.WithLabelValues(dblayer, path, method).Inc()
	err := db.Select(dest, query, args...)
	passed := time.Since(start)
	Duration.WithLabelValues(dblayer, path).Observe(float64(passed.Milliseconds()))
	m.Inf(path, method, passed)
	if err != nil {
		Errors.WithLabelValues(dblayer, err.Error(), path).Inc()
		m.Err(path, method, err.Error(), passed)
	}
	return err
}

func (m *MyLogger) MyGet(db *sqlx.DB, path string, query string, dest interface{}, args ...interface{}) error {
	start := time.Now()
	Hits.WithLabelValues(dblayer, path, method).Inc()
	err := db.Get(dest, query, args...)
	passed := time.Since(start)
	Duration.WithLabelValues(dblayer, path).Observe(float64(passed.Milliseconds()))
	m.Inf(path, method, passed)
	if err != nil {
		Errors.WithLabelValues(dblayer, err.Error(), path).Inc()
		m.Err(path, method, err.Error(), passed)
	}
	return err
}

func (m *MyLogger) MyExec(db *sqlx.DB, path string, query string, args ...interface{}) (sql.Result, error) {
	start := time.Now()
	Hits.WithLabelValues(dblayer, path, method).Inc()
	result, err := db.Exec(query, args...)
	passed := time.Since(start)
	Duration.WithLabelValues(dblayer, path).Observe(float64(passed.Milliseconds()))
	m.Inf(path, method, passed)
	if err != nil {
		Errors.WithLabelValues(dblayer, err.Error(), path).Inc()
		m.Err(path, method, err.Error(), passed)
	}
	return result, err
}

func (m *MyLogger) MyTxExec(tx *sqlx.Tx, path string, query string, args ...interface{}) (sql.Result, error) {
	start := time.Now()
	Hits.WithLabelValues(dblayer, path, method).Inc()
	result, err := tx.Exec(query, args...)
	passed := time.Since(start)
	Duration.WithLabelValues(dblayer, path).Observe(float64(passed.Milliseconds()))
	m.Inf(path, method, passed)
	if err != nil {
		Errors.WithLabelValues(dblayer, err.Error(), path).Inc()
		m.Err(path, method, err.Error(), passed)
	}
	return result, err
}

func (m *MyLogger) MyBegin(db *sqlx.DB, path string) (*sqlx.Tx, error) {
	start := time.Now()
	Hits.WithLabelValues(dblayer, path, method).Inc()
	result, err := db.Beginx()
	passed := time.Since(start)
	Duration.WithLabelValues(dblayer, path).Observe(float64(passed.Milliseconds()))
	m.Inf(path, method, passed)
	if err != nil {
		Errors.WithLabelValues(dblayer, err.Error(), path).Inc()
		m.Err(path, method, err.Error(), passed)
	}
	return result, err
}

func (m *MyLogger) MyRollBack(tx *sqlx.Tx, path string) error {
	start := time.Now()
	Hits.WithLabelValues(dblayer, path, method).Inc()
	err := tx.Rollback()
	passed := time.Since(start)
	Duration.WithLabelValues(dblayer, path).Observe(float64(passed.Milliseconds()))
	m.Inf(path, method, passed)
	if err != nil {
		Errors.WithLabelValues(dblayer, err.Error(), path).Inc()
		m.Err(path, method, err.Error(), passed)
	}
	return err
}
func (m *MyLogger) MyCommit(tx *sqlx.Tx, path string) error {
	start := time.Now()
	Hits.WithLabelValues(dblayer, path, method).Inc()
	err := tx.Commit()
	passed := time.Since(start)
	Duration.WithLabelValues(dblayer, path).Observe(float64(passed.Milliseconds()))
	m.Inf(path, method, passed)
	if err != nil {
		Errors.WithLabelValues(dblayer, err.Error(), path).Inc()
		m.Err(path, method, err.Error(), passed)
	}
	return err
}

func (m *MyLogger) MetricsInterceptor(
	ctx context.Context,
	method string,
	req interface{},
	reply interface{},
	cc *grpc.ClientConn,
	invoker grpc.UnaryInvoker,
	opts ...grpc.CallOption,
) error {
	start := time.Now()
	err := invoker(ctx, method, req, reply, cc, opts...)
	duration := time.Since(start)
	Hits.WithLabelValues("grpc", cc.Target(), method).Inc()
	md, ok := metadata.FromIncomingContext(ctx)
	id := md["x-request-id"]
	if !ok {
		m.Logger.Error("no ID")
	}
	m.Logger.Info("request=",
		zap.Any("Id", id),
		zap.String("method", method),
		zap.Duration("latency", duration),
	)
	Duration.WithLabelValues("grpc", cc.Target()).Observe(float64(duration.Milliseconds()))
	if err != nil {
		Errors.WithLabelValues("grpc", err.Error(), cc.Target()).Inc()
		m.Logger.Error("request=",
			zap.Any("Id", id),
			zap.String("method", method),
			zap.String("error", err.Error()),
			zap.Duration("latency", duration),
		)
	}
	return err
}
