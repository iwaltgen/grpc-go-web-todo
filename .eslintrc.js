module.exports = {
  root: true,
  parser: '@typescript-eslint/parser',
  parserOptions: {
    ecmaVersion: 2017,
    project: join(__dirname, "./tsconfig.json"),
    sourceType: "module"
  },
  plugins: ['@typescript-eslint', 'import', 'svelte3', 'rxjs', 'prettier'],
  extends: [
    'eslint:recommended',
    'plugin:import/recommended',
    'plugin:@typescript-eslint/eslint-recommended',
    'plugin:@typescript-eslint/recommended',
    'plugin:prettier/recommended',
    'prettier/@typescript-eslint',
  ],
  overrides: [
    {
      files: ['**/*.svelte'],
      processor: 'svelte3/svelte3',
    },
  ],
  rules: {
    "rxjs/no-async-subscribe": "error",
    "rxjs/no-ignored-observable": "error",
    "rxjs/no-ignored-subscription": "error",
    "rxjs/no-nested-subscribe": "error",
    "rxjs/no-unbound-methods": "error",
    "rxjs/throw-error": "error"
  }
};
