/*
Copyright 2021 CodeNotary, Inc. All rights reserved.

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

package sql

type SQLValueType = int

const (
	IntegerType SQLValueType = iota
	BooleanType
	StringType
	BLOBType
	TimestampType
)

type SQLStmt interface {
}

type CreateDatabaseStmt struct {
	db string
}

type UseDatabaseStmt struct {
	db string
}

type CreateTableStmt struct {
	table    string
	colsSpec []*ColSpec
}

type ColSpec struct {
	colName string
	colType SQLValueType
}

type CreateIndexStmt struct {
	table string
	col   string
}

type AddColumnStmt struct {
	table   string
	colSpec *ColSpec
}

type AlterColumnStmt struct {
	table   string
	colSpec *ColSpec
}

type InsertIntoStmt struct {
	table  string
	cols   []string
	values []Value
}

type Value interface {
}

type BooleanValue struct {
	value bool
}

type IntegerValue struct {
	value int
}

type StringValue struct {
	value string
}

type BLOBValue struct {
	value []byte
}
