# Atomic LTI

[![Client Build](https://github.com/atomicjolt/atomiclti/actions/workflows/node.js.yml/badge.svg)](https://github.com/atomicjolt/atomiclti/actions/workflows/node.js.yml)
[![Server Build](https://github.com/atomicjolt/atomiclti/actions/workflows/go.yml/badge.svg)](https://github.com/atomicjolt/atomiclti/actions/workflows/go.yml)

This is a LTI stack written in Go that handles a lot of the boilerplate so that you can get to writing logic for your app.

*TODO*:
- [ ] Deep Linking
- [ ] Oauth2 token dance
- [ ] Canvas API integration
- [ ] Deploy process

## Setup

To get set up, follow these steps in order.

### 0: Dependencies
If you are using `asdf`, just run `asdf install` after the clone step in the environment setup. Otherwise, just make sure you have PostgreSQL, NodeJS, and the GoLang dev environment set up for this project. You will also need `nginx`, but that is covered in the following install steps.

*Note*: These instructions assume you are either on MacOS or a Debian-based Linux distribution.

### 1: Environment setup
* Clone this repo down
* Create the development and test database with `createdb atomiclti_dev` and `createdb atomiclti_test`
* Copy `server_config.example.json` to `server_config.json` and update values inside appropriately
  * These can be left as the default, unless you would like to change it:
    * `db_user`
    * `db_host`
    * `database`
    * `server_port`
  * `client_id`: You will need to set up a new LTI key for your dev environment, so see the instructions for that
  * `auth0_client_id`: This can be left alone, or set to something you choose. It just identifies your client when doing the Open ID Connect process to initiate an LTI Advantage Launch with Canvas.
  * `autho0_client_secret`: For production, this of course needs to be kept secret, but for dev you could just use the default value. If you want, you can generate a new random secret and replace the existing one.
* Copy `.env.example` to `.env`. Note that nothing sources `.env` automatically yet, so this is more or less a no-op. The lack of an `LTI_ENV` environment variable being set makes the app default to the `development` environment.
  * The three supported environments are `development`, `test`, and `production`

### 2: SSL/Web Stack Setup:
* Run the `nginx` setup script in `./bin/setup` (you will need `nginx` configured with SSL certificates for `atomicjolt.xyz`)
  * If you haven't done this yet, the setup script is in our shared Dropbox folder `$DROPBOX/aj-dev/atomicjolt/local-development` (`setup` is for MacOS, and `setup-linux` is for Debian based systems, but you could copy and modify it for other systems)
* Install client dependencies first with `cd client` to be in the right directory, then `yarn` to install dependencies
* Create a client build with `yarn build` in the client directory
* Run the server with `go run server.go` in the project root (to view the project, you will need to install it as an LTI tool)

### 3: LTI Tool Install
* This tool is installed via LTI Advantage. This means that you will need to generate a new LTI key.
  * A note: LTI keys normally identify an OAUTHv2 relationship with an LMS for the establishment of durable API tokens. For whatever reason, LTI keys can also represent a LTI tool install. This was a decision that the Canvas devs made and stuck with.
* Head to the Canvas admin panel for the main account and navigate to the "Developer Keys" section on the left sidebar.
* Add a new developer key and fill out the application settings
  * **Key Name**: The name of the app, I usually choose "Atomic Insight - My Name - Environment"
  * **Owner Email**: Your AJ email
  * **Redirect URIs**: https://atomicinsight.atomicjolt.xyz/lti_launches
  * **Notes**: Leave blank
  * **Method**: Manual Entry
  * **Title**: Atomic Insight
  * **Description**: Atomic Insight - Analytics for Canvas
  * **Target Link URI**: https://atomicinsight.atomicjolt.xyz/lti_launches
  * **OpenID Connect Initiation Url**: https://atomicinsight.atomicjolt.xyz/oidc_init
  * **JWK Method**: Public JWK (you have to manually enter it for dev, since Canvas usually can't hit your dev instance from WAN unless you are using `ngrok`, or some other form of NAT punching)
  * **Public JWK**: With the Atomic Insight GoLang server running, navigate to https://atomicinsight.atomicjolt.xyz/jwks and grab just one of the JWKs in the array (one JSON object inside of the `keys` array). Paste that JSON object into the Public JWK text box for your LTI key.
  * **LTI Advantage Services**: Enable *Can retrieve user data associated with the context the tool is installed in.*
  * **Additional Settings**: Set the scope to *Public*
  * **Placements**: Delete all of the default placements and add *Course Navigation*
    * **Target Link URI**: https://atomicinsight.atomicjolt.xyz/lti_launches
    * **Icon URL**: https://atomicinsight.atomicjolt.xyz/logo.png
    * **Text**: Atomic Insight
    * **Selection Height**: 600
    * **Selection Width**: 800
* Save your LTI key
* Enable your LTI key (turn the *State* toggle in the key list from *Off* to *On*)
* Copy the Client ID from your key (the top number above the *Show Key* button) and put that into your `server_config.json` for both the `development` and `test` environments.
* Head to the course you will be testing in and install a new tool via Client ID by pasting this client ID into the install form.

### 4: Database setup:
* Run any migrations with `./bin/migrate migrate`
* Run any seeds with `./bin/seed` (this is not idempotent, running it multiple times will create duplicate data)

### 5: Running
* Make sure you are running `go run server.go` in the project root and have run `yarn build` in the client folder
  - To enable hot reload for the client, run `yarn start` in the client directory, and make sure to set the environment variable: `LTI_ENV=development`
  - If you have `entr` installed, you can run `find . -name "*.go" | entr -cr go run server.go` or similar to enable hot reload for the server
* Navigate to your test course and open up *Atomic LTI* on the sidebar, and the tool should launch

## Development

### Graphql
This project used [gqlgen](https://github.com/99designs/gqlgen) to generate code for a graphql server.
Documentation can be found at https://gqlgen.com/.

#### Configuration
The locations for generated code, mappings between graphql types and go types, and other configuration values are in `gqlgen.yml`

#### Schema
When the graphql schema is changed, the generated code needs to be updated with `go generate ./...`
