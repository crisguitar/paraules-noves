module.exports = {
  root: true,
  env: {
    node: true
  },
  extends: ['eslint:recommended'],
  overrides: [
    {
      files: [
        '**/__tests__/*.{j,t}s?(x)',
        '**/*.spec.{j,t}s?(x)'
      ],
      env: {
        jest: true
      }
    }
  ]
}
