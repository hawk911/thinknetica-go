module go-core/lesson_3/cmd/engine

require (
	golang.org/x/net v0.0.0-20201031054903-ff519b6c9102 // indirect
	pkg/crawler v1.0.0
)

replace pkg/crawler => ../../pkg/crawler

go 1.15
