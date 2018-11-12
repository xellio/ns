## ns - Network Scanner
ns is a wrapper for nmap, parsing and storing the (nmap) output in a mysql database.

[![go report card](https://goreportcard.com/badge/github.com/xellio/ns "go report card")](https://goreportcard.com/report/github.com/xellio/ns)
[![Build Status](https://travis-ci.org/xellio/ns.svg?branch=master)](https://travis-ci.org/xellio/ns)
[![MIT license](http://img.shields.io/badge/license-MIT-brightgreen.svg)](http://opensource.org/licenses/MIT)
[![GoDoc](https://godoc.org/github.com/xellio/ns?status.svg)](https://godoc.org/github.com/xellio/ns)

### Setup
1. create a database for storing the results
```
CREATE DATABASE `ns`;
```
2. add your database login/password (and database name) to config.yml file
```
database:
  login: 'DB_LOGIN'
  password: 'DB_PASSWORD'
  host: '127.0.0.1'
  port: '3306'
  name: 'ns'
  ...
```
3. build it
```
make
```
if not using upx run
```
make build
```

### Usage
ns uses nmap for scanning. In theory, all nmap arguments/parameters should work.