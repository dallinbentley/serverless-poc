.PHONY: build clean deploy gomodgen

build: gomodgen
	export GO111MODULE=on
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/hello hello/main.go
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/world world/main.go
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/stripe_create_business_account stripe_create_business_account/main.go
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/supabase_get_all_items supabase_get_all_items/main.go

clean:
	rm -rf ./bin ./vendor

deploy: clean build
	sls deploy --verbose --region us-west-2

gomodgen:
	chmod u+x gomod.sh
	./gomod.sh
