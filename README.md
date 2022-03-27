# Looped Headless Content Management System (Work in Progress)

[![API](https://github.com/looped-dev/cms/actions/workflows/api.yml/badge.svg)](https://github.com/looped-dev/cms/actions/workflows/api.yml)

> Early Work in Progress

Looped Headless CMS is a modern, lightweight, and opinionated headless CMS for
managing blogs, publications, and journals. This CMS is inspired by the Ghost
CMS but focuses on being a headless CMS only to enable you to build custom and
rich user experiences for your readers and users. At the same time, we handle
the nitty-gritty stuff like Content Management, Member Subscriptions, Payments,
and so on. This is all exposed using GraphQL technology, which is great for documentation
and easy to request the data you need in a single API call.

As a headless content management system (CMS), you bring your own frontend. This
may include:

- Web Apps - Use your framework of choice - React, Angular, Vue, Svelte,
  Marko, etc.
- Static Site Generators - Hugo, Astro, Jekyll, etc.
- Smartphones/Tables Apps - You can use looped-cms as the backend for your mobile app from
  iOS, Android and a whole host of cross-platform tools.
- Desktops Apps - Laptops and desktop computers are here to stay for the
  foreseeable and looped doesn't care what the frontend is made for.
- The sky is the limit - if it can integrate with an API, it can connect to our
  CMS.

## Features

This is still a work in progress; for the initial version, I am targetting the
following features:

- Content Management - Pages, Posts, Tags, etc.
- An Admin UI for managing content
- GraphQL API for connecting your frontend to the headless CMS
- Membership and Membership Subscription feature - allowing you to control access to content by members.
- Staff Management - Be able to give access to your staff to access the Admin UI and manage content with granular access control

## Contributing and Testing Early Versions

This is still an early work in progress and this will change in the future as I
intent to make easier to get started with looped, hopefully a CLI for
accomplishing various task.

### Running the CMS API (Backend)

> This is for development purpose, when I launch alpha version, I intend to have
> a full blown CLI to manage the CMS API.

You have two options:

#### Using Docker

- First, make sure you have docker installed
- Clone this repository - `git clone https://github.com/looped-dev/cms.git`
- Then at the root of the repository, run docker-compose up. This is going to
  build the API in a docker container and also get MongoDB container.
- Once the build process is done, you can open the graphql server playground:
  [http://localhost:8080](http://localhost:8080).

#### Using Golang

- First, make sure you have [Go](https://go.dev/doc/install) and [MongoDB](https://www.mongodb.com/docs/manual/installation/) Installed in your system.
- Clone this repository: `git clone https://github.com/looped-dev/cms.git`
- Copy the `looped.config.example.yaml` file and rename to `looped.config.yaml`
- Open the file you just renamed, and update the various configs appropriately -
  follow the comments in the file.
- You can now run the golang server: `go run ./api/server/server.go`. This will
  launch the server on port 8080.
- That's it 👏, you can explore graphql playground on your browser -
  [http://localhost:8080](http://localhost:8080).

### Running the CMS Frontend

The frontend is based on Angular and uses NX. Once you have cloned the
repository, you can run the frontend by running `nx server` which will serve the
webapp on port 4200. You will need to have [Angular CLI](https://angular.io/cli)
and [NX](https://nx.dev/getting-started/nx-setup#install-nx-cli) Installed globally.

## Questions, Suggestions and Ideas

If you have any questions or suggestions, feel free to start a discussion
[here](https://github.com/looped-dev/cms/discussions).
