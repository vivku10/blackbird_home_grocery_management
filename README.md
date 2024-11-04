# Blackbird Home Grocery Managmenent

## Setup
#### Step 1 - Use Blackbird AI Bot to create API spec
```
https://blackbird.a8r.io/
```

#### Step 2 - Use Blackbird CLI for API, Mock, Code and Deploy
```
https://www.getambassador.io/docs/blackbird/latest/guides/api/quickstart
```

#### Step 3 - Make code changes and update deployment
```
blackbird login
blackbird deployment update grocery-management --dockerfile Dockerfile --context .
```

## Endpoint definitions and commands

### Endpoint 1 - Adds a new grocery item to the refrigerator
```
curl --request POST \
  --url https://default-blackbird-viveks-organization-865b3-0.blackbird-relay.a8r.io/grocery-management/items \
  --header 'Accept: application/json' \
  --header 'Content-Type: application/json' \
  --data '{
  "name": "Paneer",
  "category": "Dairy",
  "quantity": "1 pound",
  "expiration_date": "2024-10-31T00:00:00Z"
}'
```

### Endpoint 2 - Searches for grocery items based on name
```
curl --request GET \
  --url 'https://default-blackbird-viveks-organization-865b3-0.blackbird-relay.a8r.io/grocery-management/items/search?query=Burger' \
  --header 'Accept: application/json'
```

### Endpoint 3 - Retrieves a specific grocery item by its ID
```
curl --request GET \
  --url https://default-blackbird-viveks-organization-865b3-0.blackbird-relay.a8r.io/grocery-management/items/3 \
  --header 'Accept: application/json'
```

### Endpoint 4 - Deletes a grocery item from the refrigerator
```
curl --request DELETE \
  --url https://default-blackbird-viveks-organization-865b3-0.blackbird-relay.a8r.io/grocery-management/items/22
```

### Endpoint 5 - Updates the details of an existing grocery item
```
curl --request PUT \
  --url https://default-blackbird-viveks-organization-865b3-0.blackbird-relay.a8r.io/grocery-management/items/3 \
  --header 'Accept: application/json' \
  --header 'Content-Type: application/json' \
  --data '{
  "id": 3,
  "name": "Coke",
  "category": "Beverages",
  "quantity": "2 litre",
  "expiration_date": "2025-08-21T00:00:00Z"
}'
```

### Endpoint 6 - Retrieves a list of all expired grocery items in the refrigerator
```
curl --request GET \
  --url https://default-blackbird-viveks-organization-865b3-0.blackbird-relay.a8r.io/grocery-management/items/expired \
  --header 'Accept: application/json'
```

### Endpoint 7 - Retrieves a list of all grocery items in the refrigerator along with their quantities and expiration dates
```
curl --request GET \
  --url https://default-blackbird-viveks-organization-865b3-0.blackbird-relay.a8r.io/grocery-management/items \
  --header 'Accept: application/json'
```
