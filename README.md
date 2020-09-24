# Barong JWT
Tool to forge barong JWT to be used in any service intended to run in the Barong ecosystem.
It is very useful in developement mode to develop and test a microservice without the need of deploying any dependency of the stack.

## Build
```bash
go build .
```

## Example of API call to barong

```bash
# Create the secret key file if it doesn't exist
barong-jwt 

# Configure baron to use this secret file
export JWT_PRIVATE_KEY_PATH=config/rsa-key

# Run barong server
rails s
```

From an other console:
```
JWT=$(barong-jwt --uid IDBD12DEB15B)
curl -H "Authorization: Bearer ${JWT}" localhost:3000/api/v2/resource/users/me
```
