# frontend

This template should help get you started developing with Vue 3 in Vite.

## Environment Configuration

This project automatically handles API routing based on the environment:

- **Development** (`npm run dev`): API calls go to `http://localhost:8080`
- **Production** (`npm run build`): API calls use relative paths (same domain)

The routing is configured in `src/utils/routes.ts` using Vite's built-in environment flags:
- `import.meta.env.DEV` - true in development mode
- `import.meta.env.PROD` - true in production mode

### Development Server Setup

The Vite dev server runs on port 3000 and includes a proxy configuration to forward API calls to your Go backend on port 8080. This means you can run both servers simultaneously:

1. Start your Go backend: `go run main.go` (runs on port 8080)
2. Start the frontend dev server: `npm run dev` (runs on port 3000)

API calls from the frontend will be automatically proxied to the backend.

## Recommended IDE Setup

[VSCode](https://code.visualstudio.com/) + [Volar](https://marketplace.visualstudio.com/items?itemName=Vue.volar) (and disable Vetur).

## Type Support for `.vue` Imports in TS

TypeScript cannot handle type information for `.vue` imports by default, so we replace the `tsc` CLI with `vue-tsc` for type checking. In editors, we need [Volar](https://marketplace.visualstudio.com/items?itemName=Vue.volar) to make the TypeScript language service aware of `.vue` types.

## Customize configuration

See [Vite Configuration Reference](https://vite.dev/config/).

## Project Setup

```sh
npm install
```

### Compile and Hot-Reload for Development

```sh
npm run dev
```

### Type-Check, Compile and Minify for Production

```sh
npm run build
```

### Run Unit Tests with [Vitest](https://vitest.dev/)

```sh
npm run test:unit
```

### Run End-to-End Tests with [Playwright](https://playwright.dev)

```sh
# Install browsers for the first run
npx playwright install

# When testing on CI, must build the project first
npm run build

# Runs the end-to-end tests
npm run test:e2e
# Runs the tests only on Chromium
npm run test:e2e -- --project=chromium
# Runs the tests of a specific file
npm run test:e2e -- tests/example.spec.ts
# Runs the tests in debug mode
npm run test:e2e -- --debug
```

### Lint with [ESLint](https://eslint.org/)

```sh
npm run lint
```
