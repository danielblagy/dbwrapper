package dbwrapper

import "errors"

// ErrToSQLFail is returned when query builder fails to convert query to raw sql
var ErrToSQLFail = errors.New("can't convert to raw sql query")

// ErrNotSlice is returned when dest passed to Query.GetList is not a slice.
var ErrNotSlice = errors.New("GetList dest must be a slice")
