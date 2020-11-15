module cmd/engine

require pkg/crawler v1.0.0
replace pkg/crawler => ../../pkg/crawler

require pkg/engine v1.0.0
replace pkg/engine => ../../pkg/engine

require pkg/crawbot v1.0.0
replace pkg/crawbot => ../../pkg/crawbot

go 1.15
