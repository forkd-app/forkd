/** @type {import('eslint').Linter.Config} */
module.exports = {
  root: true, // Ensures ESLint doesn't search above this directory
  env: {
    node: true, // Enable Node.js globals
  },
  parserOptions: {
    ecmaVersion: "latest", // Support the latest ECMAScript syntax
    sourceType: "module", // Use ES modules
  },
  extends: [
    "eslint:recommended", // Base ESLint rules
    "plugin:prettier/recommended", // Integrates Prettier with ESLint
  ],
  ignorePatterns: ["node_modules/"], // Exclude unnecessary files
  rules: {
    "prettier/prettier": "error", // Prettier issues will be treated as errors
  },
  overrides: [
    {
      // Setup GraphQL Parser
      files: "*.{graphql,gql}",
      parser: "@graphql-eslint/eslint-plugin",
      plugins: ["@graphql-eslint"],
      rules: {
        "prettier/prettier": "error",
        "@graphql-eslint/naming-convention": [
          "error",
          {
            OperationDefinition: {
              style: "PascalCase",
              forbiddenPrefixes: ["Query", "Mutation", "Subscription", "Get"],
              forbiddenSuffixes: ["Query", "Mutation", "Subscription"],
            },
          },
        ],
      },
    },
    {
      // Setup processor for operations/fragments definitions on code-files
      files: "web/**/*.tsx",
      processor: "@graphql-eslint/graphql",
    },
    {
      // Setup recommended config for schema files
      files: "api/graph/schema/**/*.{graphql,gql}",
      extends: "plugin:@graphql-eslint/schema-recommended",
      rules: {
        // Override graphql-eslint rules for schema files
        "@graphql-eslint/require-description": "off",
        "@graphql-eslint/strict-id-in-types": [
          "error",
          {
            exceptions: {
              // This is for wrapper types that don't have a need for an id.
              types: [
                "PaginationInfo",
                "PaginatedRecipes",
                "PaginatedComments",
                "PaginatedRecipeRevisions",
                "UserQuery",
                "RecipeQuery",
              ],
            },
          },
        ],
      },
    },
  ],
}
