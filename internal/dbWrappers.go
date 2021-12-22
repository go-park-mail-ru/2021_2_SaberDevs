package wrapper

import (
	"database/sql"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/tarantool/go-tarantool"
)

var dblayer = "db"

var Hits = prometheus.NewCounterVec(prometheus.CounterOpts{
	Name: "Hits",
}, []string{"layer", "path", "method"})

var Errors = prometheus.NewCounterVec(prometheus.CounterOpts{
	Name: "Errors",
}, []string{"layer", "status", "path"})

var Duration = prometheus.NewHistogramVec(prometheus.HistogramOpts{
	Name: "Duratin",
}, []string{"layer", "path"})

func MyInsert(tr *tarantool.Connection, path string, space interface{}, tuple interface{}) (resp *tarantool.Response, err error) {
	start := time.Now()
	Hits.WithLabelValues(dblayer, path).Inc()
	result, err := tr.Insert(space, tuple)
	time := time.Since(start)
	Duration.WithLabelValues(dblayer, path).Observe(float64(time.Milliseconds()))
	if err != nil {
		Errors.WithLabelValues(dblayer, err.Error()[:15], path).Inc()
	}
	return result, err
}

func MyDelete(tr *tarantool.Connection, path string, space interface{}, index interface{}, key interface{}) (resp *tarantool.Response, err error) {
	start := time.Now()
	Hits.WithLabelValues(dblayer, path).Inc()
	result, err := tr.Delete(space, index, key)
	time := time.Since(start)
	Duration.WithLabelValues(dblayer, path).Observe(float64(time.Milliseconds()))
	if err != nil {
		Errors.WithLabelValues(dblayer, err.Error()[:15], path).Inc()
	}
	return result, err
}

func MySelectTyped(tr *tarantool.Connection, path string, space interface{}, index interface{}, offset uint32, limit uint32, iterator uint32, key interface{}, result interface{}) (err error) {
	start := time.Now()
	Hits.WithLabelValues(dblayer, path).Inc()
	err = tr.SelectTyped(space, index, offset, limit, iterator, key, result)
	time := time.Since(start)
	Duration.WithLabelValues(dblayer, path).Observe(float64(time.Milliseconds()))
	if err != nil {
		Errors.WithLabelValues(dblayer, err.Error()[:15], path).Inc()
	}
	return err
}

func MyReplace(tr *tarantool.Connection, path string, space interface{}, tuple interface{}) (resp *tarantool.Response, err error) {
	start := time.Now()
	Hits.WithLabelValues(dblayer, path).Inc()
	result, err := tr.Replace(space, tuple)
	time := time.Since(start)
	Duration.WithLabelValues(dblayer, path).Observe(float64(time.Milliseconds()))
	if err != nil {
		Errors.WithLabelValues(dblayer, err.Error()[:15], path).Inc()
	}
	return result, err
}

func MyCall(tr *tarantool.Connection, path string, functionName string, args interface{}) (resp *tarantool.Response, err error) {
	start := time.Now()
	Hits.WithLabelValues(dblayer, path).Inc()
	result, err := tr.Call(functionName, args)
	time := time.Since(start)
	Duration.WithLabelValues(dblayer, path).Observe(float64(time.Milliseconds()))
	if err != nil {
		Errors.WithLabelValues(dblayer, err.Error()[:15], path).Inc()
	}
	return result, err
}

func MyQuery(db *sqlx.DB, path string, query string, args ...interface{}) (*sqlx.Rows, error) {
	start := time.Now()
	Hits.WithLabelValues(dblayer, path).Inc()
	rows, err := db.Queryx(query, args...)
	time := time.Since(start)
	Duration.WithLabelValues(dblayer, path).Observe(float64(time.Milliseconds()))
	if err != nil {
		Errors.WithLabelValues(dblayer, err.Error()[:15], path).Inc()
	}
	return rows, err
}

func MySelect(db *sqlx.DB, path string, query string, dest interface{}, args ...interface{}) error {
	start := time.Now()
	Hits.WithLabelValues(dblayer, path).Inc()
	err := db.Select(dest, query, args...)
	time := time.Since(start)
	Duration.WithLabelValues(dblayer, path).Observe(float64(time.Milliseconds()))
	if err != nil {
		Errors.WithLabelValues(dblayer, err.Error()[:15], path).Inc()
	}
	return err
}

func MyGet(db *sqlx.DB, path string, query string, dest interface{}, args ...interface{}) error {
	start := time.Now()
	Hits.WithLabelValues(dblayer, path).Inc()
	err := db.Get(dest, query, args...)
	time := time.Since(start)
	Duration.WithLabelValues(dblayer, path).Observe(float64(time.Milliseconds()))
	if err != nil {
		Errors.WithLabelValues(dblayer, err.Error()[:15], path).Inc()
	}
	return err
}

func MyExec(db *sqlx.DB, path string, query string, args ...interface{}) (sql.Result, error) {
	start := time.Now()
	Hits.WithLabelValues(dblayer, path).Inc()
	result, err := db.Exec(query, args...)
	time := time.Since(start)
	Duration.WithLabelValues(dblayer, path).Observe(float64(time.Milliseconds()))
	if err != nil {
		Errors.WithLabelValues(dblayer, err.Error()[:15], path).Inc()
	}
	return result, err
}

func MyTxExec(tx *sqlx.Tx, path string, query string, args ...interface{}) (sql.Result, error) {
	start := time.Now()
	Hits.WithLabelValues(dblayer, path).Inc()
	result, err := tx.Exec(query, args...)
	time := time.Since(start)
	Duration.WithLabelValues(dblayer, path).Observe(float64(time.Milliseconds()))
	if err != nil {
		Errors.WithLabelValues(dblayer, err.Error()[:15], path).Inc()
	}
	return result, err
}

func MyBegin(db *sqlx.DB, path string) (*sqlx.Tx, error) {
	start := time.Now()
	Hits.WithLabelValues(dblayer, path).Inc()
	result, err := db.Beginx()
	time := time.Since(start)
	Duration.WithLabelValues(dblayer, path).Observe(float64(time.Milliseconds()))
	if err != nil {
		Errors.WithLabelValues(dblayer, err.Error()[:15], path).Inc()
	}
	return result, err
}

func MyRollBack(tx *sqlx.Tx, path string) error {
	start := time.Now()
	Hits.WithLabelValues(dblayer, path).Inc()
	err := tx.Rollback()
	time := time.Since(start)
	Duration.WithLabelValues(dblayer, path).Observe(float64(time.Milliseconds()))
	if err != nil {
		Errors.WithLabelValues(dblayer, err.Error()[:15], path).Inc()
	}
	return err
}
func MyCommit(tx *sqlx.Tx, path string) error {
	start := time.Now()
	Hits.WithLabelValues(dblayer, path).Inc()
	err := tx.Commit()
	time := time.Since(start)
	Duration.WithLabelValues(dblayer, path).Observe(float64(time.Milliseconds()))
	if err != nil {
		Errors.WithLabelValues(dblayer, err.Error()[:15], path).Inc()
	}
	return err
}
