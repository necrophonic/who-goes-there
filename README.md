# Who Goes There

_Who Goes There_ is an application for managing membership of a github organisation and enforcing rules such as multi-factor authentication and naming policies

## Pre-requisites

To use this tool, you need a repository (preferably _private_ or _internal_) created on your organisation that will hold sponsorship details for your org members and will hold compliance issues as they're raised

## Required token scopes

The github access token for using this service requires the following scopes:

- `repo`
- `read:org`
- `read:packages`
- `read:repo_hook`
- `user`

## Required config

| Environment Var   | Cli Flag | Default | Description                                                                                |
| ----------------- | -------- | ------- | ------------------------------------------------------------------------------------------ |
| `API_TOKEN`       | `-t`     | _none_  | The API key to connect to github. Must have scopes as defined above                        |
| `GITHUB_ORG_NAME` | `-o`     | _none_  | The name of the Github organisation to scan (token above must have permission on this org) |
| `REPOSITORY`      | `-r`     | _none_  | The name of the repository to hold people detail and raise compliance issues on            |

## License

Copyright (c) 2019 Crown Copyright (Office for National Statistics)

Released under MIT license, see [LICENSE] for details.
