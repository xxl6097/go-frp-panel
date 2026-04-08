# rolldown-string

[![npm version][npm-version-src]][npm-version-href]
[![npm downloads][npm-downloads-src]][npm-downloads-href]
[![Unit Test][unit-test-src]][unit-test-href]

A compatibility layer for [magic-string](https://github.com/Rich-Harris/magic-string) to work with Rolldown and other bundlers.

- In Rolldown, [native `magic-string`](https://rolldown.rs/in-depth/native-magic-string) is used to optimize performance.
- If native support is unavailable, it gracefully falls back to the JavaScript implementation of `magic-string`.

Recommended for use with [unplugin](https://github.com/unjs/unplugin).

## Install

```bash
npm i rolldown-string
```

## Usage

### `withMagicString`

Higher-order function to create a `transform` hook with `magic-string` support.

```ts
import { withMagicString } from 'rolldown-string'

const plugin = {
  transform: withMagicString((s, id) => {
    // your transformations...
    s.replace('42', '43')
  }),
}
```

### `rolldownString` / `generateTransform`

More flexible way to use `rolldown-string`.

```ts
import { generateTransform, rolldownString } from 'rolldown-string'

const yourPlugin = {
  transform(code, id, meta) {
    const s = rolldownString(code, id, meta)

    // your transformations...
    s.replace('42', '43')

    return generateTransform(s, id)
  },
}
```

## Sponsors

<p align="center">
  <a href="https://cdn.jsdelivr.net/gh/sxzz/sponsors/sponsors.svg">
    <img src='https://cdn.jsdelivr.net/gh/sxzz/sponsors/sponsors.svg'/>
  </a>
</p>

## License

[MIT](./LICENSE) License © 2025-PRESENT [Kevin Deng](https://github.com/sxzz)

<!-- Badges -->

[npm-version-src]: https://img.shields.io/npm/v/rolldown-string.svg
[npm-version-href]: https://npmjs.com/package/rolldown-string
[npm-downloads-src]: https://img.shields.io/npm/dm/rolldown-string
[npm-downloads-href]: https://www.npmcharts.com/compare/rolldown-string?interval=30
[unit-test-src]: https://github.com/sxzz/rolldown-string/actions/workflows/unit-test.yml/badge.svg
[unit-test-href]: https://github.com/sxzz/rolldown-string/actions/workflows/unit-test.yml
