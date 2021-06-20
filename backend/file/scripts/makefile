appName := fileApp
useCaseDir := ../internal/3_usecase/usecase
mockDir := $(useCaseDir)/mock_repository

# ----------------------------
.PHONY: createmock
createmock:
	if [ ! -d $(mockDir) ]; then \
		mkdir $(mockDir); \
	fi
	cd $(useCaseDir); \
	mockgen -source repository.go -destination mock_repository/mock_repository.go

# ----------------------------
# github.com/securego/gosec/v2/cmd/gosec
# .PHONY: gosec
# gosec:
# 	gosec -exclude=G303 ../...

# ----------------------------
.PHONY: run
run:
	go run ../cmd/*.go

# ----------------------------
.PHONY: build
build:
	go build -o ../cmd/$(appName) ../cmd/

# ----------------------------
.PHONY: addservice
addservice:
	sudo cp -f ../init/$(appName).service /etc/systemd/system
	sudo systemctl daemon-reload
	sudo systemctl enable $(appName).service
	sudo systemctl start $(appName).service
	sudo systemctl status $(appName).service

# ----------------------------
.PHONY: delservice
delservice:
	sudo systemctl stop $(appName).service
	sudo systemctl disable $(appName).service
	sudo rm /etc/systemd/system/$(appName).service
	sudo systemctl daemon-reload
	sudo systemctl status $(appName).service

# ----------------------------
.PHONY: stopservice
stopservice:
	sudo systemctl stop $(appName).service

# ----------------------------
.PHONY: startservice
startservice:
	sudo systemctl start $(appName).service

# ----------------------------
.PHONY: statusservice
statusservice:
	sudo systemctl status $(appName).service

