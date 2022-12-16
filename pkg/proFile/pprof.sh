#!/usr/bin/bash
go tool pprof -http=localhost:8001 /tmp/mem.pprof
go tool pprof -http=localhost:8002 /tmp/cpu.pprof
timeout 5s rm -rf \tmp\*.pprof

