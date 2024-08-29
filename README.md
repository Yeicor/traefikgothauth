# Traefik Goth Auth

> **Warning**   
> This plugin should be considered as experimental. It is not yet tested in production environments.

Multi-provider authentication plugin for Traefik, thanks to [Goth](https://github.com/markbates/goth). Features:

- Only/any authenticated users can reach the next middleware/service.
- All available information of the user is published as headers. 
  - Use this to filter authorized accounts with other middlewares.
- If multiple configuration providers are configured, an initial selection screen is shown.
- Once logged in a cookie will avoid the need to contact the provider for a configurable amount of time.
- Configuration documentation is available [here](config.go).
- Available providers:

![providers.png](.github/providers.png)
