FROM scubism/go_api_base:latest

# === Set app specific settings ===

ENV IMPORT_PATH="todo_center/go_todo_api"

WORKDIR $GOPATH/src/$IMPORT_PATH

COPY . $GOPATH/src/$IMPORT_PATH
RUN godep restore

VOLUME $GOPATH/src/$IMPORT_PATH

EXPOSE 3000
ENTRYPOINT ["./docker-entrypoint.sh"]
