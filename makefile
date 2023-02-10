GO2BAN_BINARY=go2ban
GO2BAN_CONF=deploy/go2ban.conf
GO2BAN_SERVICE=deploy/go2ban.service
GO2BAN_LOGDIR=/var/log/go2ban
GO2BAN_CONFDIR=/etc/go2ban

build:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o $(GO2BAN_BINARY)

install: build
	@if [ $(EUID) -ne 0 ]; then \
		echo "You must run this with superuser privileges. Try \"sudo make install\"" 2>&1; \
		exit 1; \
	fi
	systemctl stop go2ban || true
	mkdir -p $(GO2BAN_LOGDIR) $(GO2BAN_CONFDIR)
	install -m 644 $(GO2BAN_CONF) $(GO2BAN_CONFDIR)
	install -m 755 $(GO2BAN_BINARY) /usr/local/bin
	install -m 644 $(GO2BAN_SERVICE) /etc/systemd/system
	systemctl daemon-reload

clean:
	rm $(GO2BAN_BINARY)

lint:
	golangci-lint run $(go list ./... | grep -v example)

test:
	sudo go test $(go list ./... | grep -v example)

.PHONY: build install clean lint test
.DEFAULT_GOAL := build