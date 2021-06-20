.PHONY: runws
runws:
	cd backend/ws/scripts && make run

.PHONY: runmedia
runmedia:
	cd backend/media/scripts && make run

.PHONY: runfile
runfile:
	cd backend/file/scripts && make run

.PHONY: rundevice
rundevice:
	cd backend/device/scripts && make run

.PHONY:addws
addws:
	cd backend/ws/scripts && make build
	cd backend/ws/scripts && make addservice

.PHONY:addfile
addfile:
	cd backend/file/scripts && make build
	cd backend/file/scripts && make addservice

.PHONY:addmedia
addmedia:
	cd backend/media/scripts && make build
	cd backend/media/scripts && make addservice

.PHONY:adddevice
adddevice:
	cd backend/device/scripts && make build
	cd backend/device/scripts && make addservice

.PHONY:delws
delws:
	cd backend/ws/scripts && make delservice

.PHONY:delfile
delfile:
	cd backend/file/scripts && make delservice

.PHONY:delmedia
delmedia:
	cd backend/media/scripts && make delservice

.PHONY:deldevice
deldevice:
	cd backend/device/scripts && make delservice

.PHONY: test
test:
	cd backend/ws/internal/3_usecase && go test ./usecase
	cd backend/media/internal/3_usecase && go test ./usecase
	cd backend/file/internal/3_usecase && go test ./usecase
	cd backend/device/internal/3_usecase && go test ./usecase

.PHONY: compile
compile:
	protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    backend/commonpb/common.proto

.PHONY: resetvideo
resetvideo:
	echo "[]" > ./backend/file/storage/file/db.json
	rm ./backend/file/storage/file/video/*
	rm ./backend/file/storage/file/anime/*


