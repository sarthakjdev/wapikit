import { defineConfig } from 'orval';

export default defineConfig({
  api: {
    input: '../spec.openapi.yaml',
    output: {
      packageJson: './package.json',
      mode: 'single',
      prettier: true,
      client: 'react-query',
      tsconfig: './tsconfig.json',
      target: './.generated.ts',
      override: {
        mutator: {
          path: './src/utils/api-client.ts',
          name: 'customInstance',
        },
      },
    },
    hooks: {
      afterAllFilesWrite: "prettier --write"
    }
  },
});