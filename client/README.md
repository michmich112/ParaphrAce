# Client
Implementation of the frontend client for ParaphrAce

## Get started

*Note that you will need to have [Node.js](https://nodejs.org) installed.*
Install the dependencies...

```bash
npm install
```

...then start [Rollup](https://rollupjs.org):

```bash
npm run dev
```

Navigate to [localhost:5000](http://localhost:5000). You should see your app running. Edit a component file in `src`, save it, and reload the page to see your changes.

By default, the server will only respond to requests from localhost. To allow connections from other computers, edit the `sirv` commands in package.json to include the option `--host 0.0.0.0`.


## Environment
To connect to the backend server you'll need to have the following environment variables set up:
| Variable             | Required | Description                                                                               |
|----------------------|----------|-------------------------------------------------------------------------------------------|
| `PARAPHRACE_API_URL` | Yes      | URL for the paraphrace api. Must be of the form `http://url.com/` with a `/` at the end.|

