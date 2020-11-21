module cmd/engine

require pkg/crawler v1.0.0
replace pkg/crawler => ../../pkg/crawler

require pkg/invertindex v1.0.0
replace pkg/invertindex => ../../pkg/invertindex

require pkg/crawbot v1.0.0
replace pkg/crawbot => ../../pkg/crawbot

go 1.15
