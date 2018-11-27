package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type Book_20181113_200224 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Book_20181113_200224{}
	m.Created = "20181113_200224"

	migration.Register("Book_20181113_200224", m)
}

// Run the migrations
func (m *Book_20181113_200224) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE book(`id` int(11) DEFAULT NULL,`name` varchar(128) NOT NULL)")
}

// Reverse the migrations
func (m *Book_20181113_200224) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE `book`")
}
