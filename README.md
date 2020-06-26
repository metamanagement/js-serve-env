# js-serve-env

This package demonstrates how to inject environment variables passed to a docker image into a client-side javascript package that has already been built and minified.

## Usage

### Step 1

Expose the relevant environment variables to the relevant docker image.

### Step 2

Add a new environment variable, named `RUNTIME_VARS` that is a comma-separated list of the environment variable names that you want to pass to the running client-side javascript.

### Step 3

Make sure that your `index.html` file loads `/runtime_env_vars.js`. The environment variables will be made available through this file and will be of the format:

```
window.RUNTIME_<foo> = "<foo's value>"
```

Note that environment variable `<foo>`'s value will be escaped, so it may be necessary to parse or otherwise transform the variables in the existing javascript code.
