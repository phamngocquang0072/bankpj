go mod init github.com/phamngocquang0072/ginpj1
go get github.com/gin-gonic/gin

go install github.com/githubnemo/CompileDaemon@latest
go get github.com/githubnemo/CompileDaemon
go get golang.org/x/crypto/sha3@v0.28.0

go get -u gorm.io/gorm
go get -u gorm.io/driver/postgres
go get github.com/joho/godotenv

run program eachtime it build might add
CompileDaemon -command="./MyProgram -my-options"