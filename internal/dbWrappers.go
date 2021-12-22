package wrapper

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/tarantool/go-tarantool"
)

var dblayer = "db"

var Hits = prometheus.NewCounterVec(prometheus.CounterOpts{
	Name: "hits",
}, []string{"layer", "path"})

var Errors = prometheus.NewCounterVec(prometheus.CounterOpts{
	Name: "Errors",
}, []string{"status", "path"})

var Duration = prometheus.NewCounterVec(prometheus.CounterOpts{
	Name: "Duration",
}, []string{"status", "path"})

func MyInsert(tr *tarantool.Connection, path string, space interface{}, tuple interface{}) (resp *tarantool.Response, err error) {
	//TODO Metrics
	Hits.WithLabelValues(dblayer, path).Inc()
	result, err := tr.Insert(space, tuple)
	return result, err
}

func MyDelete(tr *tarantool.Connection, path string, space interface{}, index interface{}, key interface{}) (resp *tarantool.Response, err error) {
	//TODO Metrics
	Hits.WithLabelValues(dblayer, path).Inc()
	result, err := tr.Delete(space, index, key)
	return result, err
}

func MySelectTyped(tr *tarantool.Connection, path string, space interface{}, index interface{}, offset uint32, limit uint32, iterator uint32, key interface{}, result interface{}) (err error) {
	//TODO Metrics
	Hits.WithLabelValues(dblayer, path).Inc()
	err = tr.SelectTyped(space, index, offset, limit, iterator, key, result)
	return err
}

func MyReplace(tr *tarantool.Connection, path string, space interface{}, tuple interface{}) (resp *tarantool.Response, err error) {
	//TODO Metrics
	Hits.WithLabelValues(dblayer, path).Inc()
	result, err := tr.Replace(space, tuple)
	return result, err
}

func MyCall(tr *tarantool.Connection, path string, functionName string, args interface{}) (resp *tarantool.Response, err error) {
	//TODO Metrics
	Hits.WithLabelValues(dblayer, path).Inc()
	result, err := tr.Call(functionName, args)
	return result, err
}

func MyQuery(db *sqlx.DB, path string, query string, args ...interface{}) (*sqlx.Rows, error) {
	//TODO Metrics
	Hits.WithLabelValues(dblayer, path).Inc()
	rows, err := db.Queryx(query, args...)
	return rows, err
}

func MySelect(db *sqlx.DB, path string, query string, dest interface{}, args ...interface{}) error {
	//TODO Metrics
	Hits.WithLabelValues(dblayer, path).Inc()
	err := db.Select(dest, query, args...)
	return err
}

func MyGet(db *sqlx.DB, path string, query string, dest interface{}, args ...interface{}) error {
	//TODO Metrics
	Hits.WithLabelValues(dblayer, path).Inc()
	err := db.Get(dest, query, args...)
	return err
}

func MyExec(db *sqlx.DB, path string, query string, args ...interface{}) (sql.Result, error) {
	//TODO Metrics
	Hits.WithLabelValues(dblayer, path).Inc()
	result, err := db.Exec(query, args...)
	return result, err
}

func MyTxExec(tx *sqlx.Tx, path string, query string, args ...interface{}) (sql.Result, error) {
	//TODO Metrics
	Hits.WithLabelValues(dblayer, path).Inc()
	result, err := tx.Exec(query, args...)
	return result, err
}

func MyBegin(db *sqlx.DB, path string) (*sqlx.Tx, error) {
	//TODO Metrics
	Hits.WithLabelValues(dblayer, path).Inc()
	result, err := db.Beginx()
	return result, err
}

func MyRollBack(tx *sqlx.Tx, path string) error {
	//TODO Metrics
	Hits.WithLabelValues(dblayer, path).Inc()
	err := tx.Rollback()
	return err
}
func MyCommit(tx *sqlx.Tx, path string) error {
	//TODO Metrics
	Hits.WithLabelValues(dblayer, path).Inc()
	err := tx.Commit()
	return err
}
