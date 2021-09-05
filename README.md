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

If using Slack notifications, the following is required"

| Environment Var | Cli Flag | Default | Description                                                                             |
| --------------- | -------- | ------- | --------------------------------------------------------------------------------------- |
| `WEBHOOK_URL`   | n/a      | _none_  | Used to specify the webhook endpoint as given to you by slack when creating the webhook |

## Deploying

Uses the [serverless framework](https://serverless.com) to deploy to cloud platforms

To deploy the default stack (AWS):

```bash
# Export AWS credentials / set profile
# ...

# Export necessary environment vars
# ...

cd aws
make deploy
```

## Notification config

TODO

### Slack

TODO

#### Gopher codes

Some pre-spec'd avatars for your bot, from the awesome [Gopherize.me](gopherize.me)

- [Mechanic Gopher](resources/images/mechanic_gopher.png)

![Mechanic Gopher](resources/images/mechanic_gopher.png =250x250) (gopherize id: 6d2e08c34c63c0bc77160a263a0f98f4d60c41ea)

## License

Copyright (c) 2019 Crown Copyright (Office for National Statistics)

Released under MIT license, see [LICENSE] for details.
