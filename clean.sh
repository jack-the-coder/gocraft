#!/bin/sh

rm gocraft.db
go build && ./gocraft "$@"
