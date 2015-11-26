Todo API
========

Set environments:

```
LOCALIP=$(ip addr | grep "global enp" | awk '{print $2}' | sed 's/\/.*$//')
ENDPOINT="http://$LOCALIP:8001"
ITEM_ID="xxx"
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
