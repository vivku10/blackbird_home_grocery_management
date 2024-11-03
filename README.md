# api_hack

# home-grocery-management
```
blackbird login

blackbird deployment update grocery-management --dockerfile Dockerfile --context .
```

# home-grocery-management API commands
## Adds a new grocery item to the refrigerator
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

## Searches for grocery items based on name
curl --request GET \
  --url 'https://default-blackbird-viveks-organization-865b3-0.blackbird-relay.a8r.io/grocery-management/items/search?query=Burger' \
  --header 'Accept: application/json'

## Retrieves a specific grocery item by its ID
curl --request GET \
  --url https://default-blackbird-viveks-organization-865b3-0.blackbird-relay.a8r.io/grocery-management/items/3 \
  --header 'Accept: application/json'

## Deletes a grocery item from the refrigerator
curl --request DELETE \
  --url https://default-blackbird-viveks-organization-865b3-0.blackbird-relay.a8r.io/grocery-management/items/22

## Updates the details of an existing grocery item
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

## Retrieves a list of all expired grocery items in the refrigerator
curl --request GET \
  --url https://default-blackbird-viveks-organization-865b3-0.blackbird-relay.a8r.io/grocery-management/items/expired \
  --header 'Accept: application/json'

## Retrieves a list of all grocery items in the refrigerator along with their quantities and expiration dates
curl --request GET \
  --url https://default-blackbird-viveks-organization-865b3-0.blackbird-relay.a8r.io/grocery-management/items \
  --header 'Accept: application/json'
