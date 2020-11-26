module cmd/engine

require pkg/crawler v1.0.0
replace pkg/crawler => ../../pkg/crawler

require pkg/index v1.0.0
replace pkg/index => ../../pkg/index

require pkg/crawbot v1.0.0
replace pkg/crawbot => ../../pkg/crawbot

go 1.15
