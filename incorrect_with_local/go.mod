module github.com/chmeliik/gomod-multi/wrong_with_local

go 1.15

require github.com/chmeliik/some-module v0.0.0

replace github.com/chmeliik/some-module => ./src/chmeliik/some-module
