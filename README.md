# Barong JWT
Tool to forge barong JWT to be used in any service intended to run in the Barong ecosystem.
It is very useful in developement mode to develop and test a microservice without the need of deploying any dependency of the stack.

## Parameters

| Param | Default         | Description |
|-------|-----------------|-------------|
| uid   | U487205863      | UID         |
| email | admin@barong.io | Email       |
| role  | admin           | Role        |
| level | 3               | Level       |

`JWT_PRIVATE_KEY_PATH` is used to specify JWT private key path, otherwise `config/rsa-key` would be created and used.

## Example of API call to barong

```bash
# Create the secret key file if it doesn't exist
barong-jwt

# Configure barong to use a specific secret file
export JWT_PRIVATE_KEY_PATH=config/rsa-key

# Run barong server
rails s
```

From another terminal tab:
```
JWT=$(barong-jwt --uid IDBD12DEB15B)
curl -H "Authorization: Bearer ${JWT}" localhost:3000/api/v2/resource/users/me
```

## License
Barong JWT is released under the terms of the [Apache License 2.0](LICENSE.md).

It's used in openware [cryptocurrency exchange software](https://www.openware.com) stack, and more!
