create:
	cp dbconfig.yml.template dbconfig.yml
migrate:
	sql-migrate $(@F) -env="development" -config="dbconfig.yml"
