migrate_db:
	tern migrate --config ./migrations/tern.conf --migrations ./migrations

setup_local:
	cd docker && finch compose up

teardown_local:
	cd docker && finch compose down
