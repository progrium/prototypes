
default: run

run:
	node ./dev.js

schema/schema.ts:
	./sutil ts > schema/schema.ts

schema/docs: schema/schema.ts
	typedoc --out schema/docs schema