module github.com/punqy/punqy

go 1.16

require (
	github.com/asaskevich/govalidator v0.0.0-20210307081110-f21760c49a8d // indirect
	github.com/go-ozzo/ozzo-validation/v4 v4.3.0
	github.com/golang-jwt/jwt v3.2.2+incompatible
	github.com/google/uuid v1.3.0
	github.com/jmoiron/sqlx v1.3.4
	github.com/joho/godotenv v1.4.0
	github.com/punqy/core v0.0.0-20220128212947-e010ef317451
	github.com/sethvargo/go-envconfig v0.4.0
	github.com/sirupsen/logrus v1.8.1
	github.com/slmder/migrate v0.4.0
	github.com/slmder/qbuilder v0.7.0
	github.com/spf13/cobra v1.2.1
	github.com/valyala/fasthttp v1.33.0 // indirect
)

// Work
//replace github.com/punqy/core => /home/sergey/Документы/Development/src/github.com/punqy/core

// Home
replace github.com/punqy/core => /home/sergey/GolangProjects/src/github.com/punqy/core
