.PHONY: schema

all: schema rpc client/types

schema: schema/schema.json schema/schema.ts schema/docs

schema/schema.json:
	./sutil build > schema/schema.json

schema/schema.ts:
	./sutil ts > schema/schema.ts

schema/docs: schema/schema.ts
	typedoc --out schema/docs schema

client/types: schema/lib/clientgen.go
	mkdir -p client/electrond
	go run schema/lib/clientgen.go > client/electrond/electrond.go

rpc:
	./sutil rpc | tsfmt --stdin > src/main/rpc.ts