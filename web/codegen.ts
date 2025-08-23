import type { CodegenConfig } from "@graphql-codegen/cli"

const config: CodegenConfig = {
  overwrite: true,
  schema: "http://localhost:8000/query",
  documents: "app/gql/**/*.graphql",
  generates: {
    "app/gql/forkd.g.ts": {
      plugins: [
        "typescript",
        "typescript-operations",
        "typescript-graphql-request",
      ],
    },
  },
}

export default config
