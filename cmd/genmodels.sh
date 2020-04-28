#!/bin/bash
#usage: xorm reverse [-s] driverName datasourceName tmplPath [generatedPath] [tableFilterReg]
#
#according database's tables and columns to generate codes for Go, C++ and etc.
#
#    -s                Generated one go file for every table
#    driverName        Database driver name, now supported four: mysql mymysql sqlite3 postgres
#    datasourceName    Database connection uri, for detail infomation please visit driver's project page
#    tmplPath          Template dir for generated. the default templates dir has provide 1 template
#    generatedPath     This parameter is optional, if blank, the default value is models, then will
#                      generated all codes in models dir
#    tableFilterReg    Table name filter regexp

#xorm reverse mysql "root:123456@tcp(127.0.0.1:33060)/posts?charset=utf8mb4" api/xorm/goxorm api/models