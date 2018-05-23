/*
Copyright [yyyy] [name of copyright owner]

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package status

import (
	"github.com/VividCortex/mysqlerr"
)

/*******************************************************************
 beego/orm 中定义的错误，通过 parseErrorCode 生成错误码
	 ErrTxHasBegan    = errors.New("<Ormer.Begin> transaction already begin")
	 ErrTxDone        = errors.New("<Ormer.Commit/Rollback> transaction not begin")
	 ErrMultiRows     = errors.New("<QuerySeter> return multi rows")
	 ErrNoRows        = errors.New("<QuerySeter> no row found")
	 ErrStmtClosed    = errors.New("<QuerySeter> stmt already closed")
	 ErrArgs          = errors.New("<Ormer> args error may be empty")   // returned by MultiInsert in beego
	 ErrNotImplement  = errors.New("have not implement")
 *******************************************************************/

const (
	// https://github.com/lib/pq/blob/master/error.go
	// https://github.com/VividCortex/mysqlerr/blob/93e3dc264d50c8ea98844075ef84224cacdc0c91/mysqlerr.go
	SQLSuccess                 uint32 = 1000 + iota // success
	SQLErrUnAuthorized                              // authorize failed, mysql 1045, postgres "28P01"/"28000"
	SQLErrSyntax                                    // syntax error, mysql 1064, postgres "42601"/"42000"
	SQLErrUnknownTable                              // unknown table, mysql 1146, postgres "42P01"
	SQLErrUnknownColumn                             // unknown column(field), mysql 1047, postgres "42703"
	SQLErrTooManyConnections                        // too many connections, mysql 1203,postgres "53300"
	SQLErrInsufficentResources                      // server resources not enough, mysql 1041, postgres "53000"
	SQLErrDiskFull                                  // server disk is full, mysql 1021, postgres "53100"
	SQLErrInternalErr                               // server internal error, mysql 1815,postgres "XX000"
	SQLErrDataTypeMismatch                          // data type not match, mysql 3064, postgres:"42804"
	SQLErrDuplicateEntry                            // duplicate on unique key

	// Errors defined by beego
	SQLErrNoRowFound         // no data rows found, when you want one, e.g. o.QueryRow(...).One(this), defined by beego
	SQLErrMultiRows          // too many rows found, when you want one only, e.g. o.QueryRow(...).One(this), defined by beego
	SQLErrInvalidMultiInsert // error defined by beego when you call InsertMulti(), but pass in zero rows of data

	// UnCategoried errors
	SQLErrUnCategoried // errors defined by mysql or postgres, but not included above

)

var (
	// mapping from mysql errors to customized errors
	// reference: https://github.com/VividCortex/mysqlerr
	mysqlErrorMapping = map[uint16]uint32{
		mysqlerr.ER_ACCESS_DENIED_ERROR:       SQLErrUnAuthorized,         // 1045
		mysqlerr.ER_PARSE_ERROR:               SQLErrSyntax,               // 1064
		mysqlerr.ER_NO_SUCH_TABLE:             SQLErrUnknownTable,         // 1146
		mysqlerr.ER_UNKNOWN_COM_ERROR:         SQLErrUnknownColumn,        // 1047
		mysqlerr.ER_TOO_MANY_USER_CONNECTIONS: SQLErrTooManyConnections,   // 1203
		mysqlerr.ER_OUT_OF_RESOURCES:          SQLErrInsufficentResources, // 1041
		mysqlerr.ER_DISK_FULL:                 SQLErrDiskFull,             // 1021
		mysqlerr.ER_INTERNAL_ERROR:            SQLErrInternalErr,          // 1815
		mysqlerr.ER_INCORRECT_TYPE:            SQLErrDataTypeMismatch,     // 3064
		mysqlerr.ER_DUP_ENTRY:                 SQLErrDuplicateEntry,       // 1062
	}
	// mapping from postgres errors to customized errors
	// reference: https://github.com/lib/pq/blob/master/error.go
	postgresErrorMapping = map[string]uint32{
		"invalid_authorization_specification": SQLErrUnAuthorized,
		"invalid_password":                    SQLErrUnAuthorized,
		"undefined_table":                     SQLErrUnknownTable,
		"undefined_column":                    SQLErrUnknownColumn,
		"too_many_connections":                SQLErrTooManyConnections,
		"insufficient_resources":              SQLErrInsufficentResources,
		"disk_full":                           SQLErrDiskFull,
		"internal_error":                      SQLErrInternalErr,
		"datatype_mismatch":                   SQLErrDataTypeMismatch,
		"duplicate_column":                    SQLErrDuplicateEntry,
	}
)
