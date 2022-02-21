Connet to remote DB using command:
$ heroku pg:psql

App logic workflow (architecture):
HTTP request -> Handlers -> Service(busines logic) -> Repository (DB)
