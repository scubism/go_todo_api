Todo API
========

Set environments:

```
ENDPOINT="http://$HOST:$GO_TODO_API_PORT"
ITEM_ID="xxx"
# Here, the variables such as $HOST are defined in default.env in the central repository
```

Create an item:

```
curl -H "Accept:application/json" \
-H "Content-Type:application/json" \
-XPOST "$ENDPOINT/v1/todos" \
-d '{"title": "xxx"}'
```

List items:

```
curl $ENDPOINT/v1/todos
```

Get an item:

```
curl $ENDPOINT/v1/todos/$ITEM_ID
```

Update an item:

```
curl -H "Accept:application/json" \
-H "Content-Type:application/json" \
-XPUT "$ENDPOINT/v1/todos/$ITEM_ID" \
-d '{"title": "yyy"}'
```

Delete an item:

```
curl -H "Accept:application/json" \
-H "Content-Type:application/json" \
-XDELETE "$ENDPOINT/v1/todos/$ITEM_ID"
```

Move an item:

```
curl -H "Accept:application/json" \
-H "Content-Type:application/json" \
-XPOST "$ENDPOINT/v1/todos/$ITEM_ID/move" \
-d "{\"prior_sibling_id\": \"$PRIOR_SIBLING_ID\"}"
```
