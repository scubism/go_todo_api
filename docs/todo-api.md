Todo API
========

Set environments:

```
# Load config environment
eval $(./init.sh | grep config_)

ENDPOINT="https://$config_host:$config_todo_api_gateway_port"
ITEM_ID="xxx"
# Here, the variables such as $HOST are defined in default.env in the central repository
```

Create an item:

```
curl -k -H "Accept:application/json" \
-H "Content-Type:application/json" \
-XPOST "$ENDPOINT/v1/todos" \
-d '{"title": "xxx"}'
```

List items:

```
curl -k $ENDPOINT/v1/todos
```

Get an item:

```
curl -k $ENDPOINT/v1/todos/$ITEM_ID
```

Update an item:

```
curl -k -H "Accept:application/json" \
-H "Content-Type:application/json" \
-XPUT "$ENDPOINT/v1/todos/$ITEM_ID" \
-d '{"title": "yyy"}'
```

Delete an item:

```
curl -k -H "Accept:application/json" \
-H "Content-Type:application/json" \
-XDELETE "$ENDPOINT/v1/todos/$ITEM_ID"
```

Move an item:

```
curl -k -H "Accept:application/json" \
-H "Content-Type:application/json" \
-XPOST "$ENDPOINT/v1/todos/$ITEM_ID/move" \
-d "{\"prior_sibling_id\": \"$PRIOR_SIBLING_ID\"}"
```
