# invoice-dashboard

dashboard for managing your invoices(view, edit, delete, create)

### run locally(dev)
There are 2 options for running the backend locally

<dl>
  <dt>requirements</dt>
  <dd>you should have <a href="https://go.dev/doc/install">go</a> installed</dd>
  <dd>you should have <a href="https://stackoverflow.com/questions/32127524/how-to-install-and-use-make-in-windows">make</a> installed</dd>
  <dd>you should have <a href="https://www.docker.com/">docker</a> installed</dd>
</dl>

1. running the monolith application which uses [gin](https://gin-gonic.com/docs/) framework
- you need to run docker containers for local database and datastore (postgres + redis) for that you should run `make run-containers` command
- then build the application `make build-cmd` and run the `make run-cmd` binary

2. invoking serverless lambda functions locally

<dl>
  <dt>requirements</dt>
  <dd>you should have <a href="https://docs.aws.amazon.com/serverless-application-model/latest/developerguide/serverless-sam-cli-install.html">sam</a> installed</dd>
  <dd>also postgres&redis running and available on internet</dd>
  <dd>create and configure env-vars.json file in root of project as in the file dev-env.json</dd>
</dl>

- `make invoke-get-all` for invoking handler for getting all invoices
- `make invoke-get` for invoking handler for getting invoice by id (you need to update the id in event.json)
- `make invoke-create` for creating new invoice(you need to update body in event.json)
- `make invoke-put` for updating invoice by id(you need to update body and id in event.json)
- `make invoke-delete` for deleting invoice by id(you need to update id in event.json)
