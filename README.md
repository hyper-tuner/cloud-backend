<p align="center">
  <img src="/public/icon.png" alt="HyperTuner" width="100">
</p>

<h1 align="center">HyperTuner Cloud Backend</h1>

<div align="center">
  <p><a href="https://hypertuner.cloud"><strong>hypertuner.cloud</strong></a></p>
  <p><sub>The best way to share your tunes and logs.</sub></p>
</div>

<br/>

This is the backend for the [HyperTuner Cloud](https://github.com/hyper-tuner/hyper-tuner-cloud). It is based on a great [PocketBase](https://pocketbase.io) backend.

## This repository

- the source code of the **HyperTuner Cloud** backend that extends **PocketBase**.
- database schema
- configurations
- deployment scripts / **Docker** and docker compose files
- migration scripts and guides

## Setting up your instance

Use provider Docker files from the `/docker` directory or just grab the binary from the [Releases](https://github.com/hyper-tuner/cloud-backend/releases) page.

```bash
./cloud-backend serve
```

This will create `pb_data` directory where all the data will be stored (SQLite, uploaded files and metadata).

Now you can access the admin UI at: [https://your-instance.com/_/](https://your-instance.com/_/).

### Application name and URL

Located in admin UI: `Settings -> Application`.

- `Application name` - the name of your application
- `Application URL` - the URL of your **frontend** application

### Mail settings

Located in the admin UI: `Settings -> Mailer settings`.

- Verification email **ACTION URL**:

```bash
{APP_URL}/#/auth/email-verification/{TOKEN}
```

- Password reset **ACTION URL**:

```bash
{APP_URL}/#/auth/reset-password-confirmation/{TOKEN}
```

### OAuth2 settings

For every OAuth2 provider, you need to set the following redirect URLs:

```bash
https://{FRONTEND-URL}/?redirect=oauth&provider=google
https://{FRONTEND-URL}/?redirect=oauth&provider=github
https://{FRONTEND-URL}/?redirect=oauth&provider=facebook
```

### Loading schema

Copy/load `pb_schema.json` to `Settings -> Sync -> Import collections` in the admin UI.

## Building from source

```bash
go build
```

## Support this project

[GitHub Sponsors](https://github.com/sponsors/karniv00l)

## Discord server

[![HyperTuner Discord server](https://dcbadge.vercel.app/api/server/HdxznPUA)](https://discord.gg/HdxznPUA)

## License

[MIT](https://github.com/hyper-tuner/cloud-backend/blob/master/LICENSE)
