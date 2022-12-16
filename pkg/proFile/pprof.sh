#!/usr/bin/bash
go tool pprof -http=localhost:8001 /tmp/cpu.pprof
go tool pprof -http=localhost:8002 /tmp/mem.pprof
timeout 5s rm -rf \tmp\*.pprof

