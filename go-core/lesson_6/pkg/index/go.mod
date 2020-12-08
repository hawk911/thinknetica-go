module pkg/index

require pkg/model v1.0.0
replace pkg/model => ../../pkg/model

require pkg/crawler v1.0.0
replace pkg/crawler => ../../pkg/crawler

require pkg/storage/memory v1.0.0
replace pkg/storage/memory => ../../pkg/storage/memory

go 1.15
