/**
 * This is intended to be a basic starting point for linting in your app.
 * It relies on recommended configs out of the box for simplicity, but you can
 * and should modify this configuration to best suit your team's needs.
 */

/** @type {import('eslint').Linter.Config} */
module.exports = {
  root: false, // Inherit from the root-level configuration
  extends: ["../.eslintrc.cjs"], // Extend the root configuration
  parserOptions: {
    ecmaVersion: "latest", // Support the latest ECMAScript syntax
    sourceType: "module", // Use ES modules
    ecmaFeatures: {
      jsx: true, // Enable JSX
    },
  },
  env: {
    browser: true, // Enable browser globals for React
    es6: true, // Enable ES6 globals
  },
  overrides: [
    // React-specific configuration
    {
      files: ["**/*.{js,jsx,ts,tsx}"],
      plugins: ["react", "jsx-a11y", "react-hooks"],
      extends: [
        "plugin:react/recommended",
        "plugin:react/jsx-runtime",
        "plugin:react-hooks/recommended",
        "plugin:jsx-a11y/recommended",
      ],
      settings: {
        react: {
          version: "detect", // Automatically detect React version
        },
        "import/resolver": {
          typescript: {
            alwaysTryTypes: true, // Resolve TypeScript imports
            project: "./tsconfig.json", // Point to the tsconfig file in /web
          },
        },
      },
    },

    // TypeScript-specific configuration
    {
      files: ["**/*.{ts,tsx}"],
      parser: "@typescript-eslint/parser",
      plugins: ["@typescript-eslint", "import"],
      settings: {
        "import/resolver": {
          typescript: {
            alwaysTryTypes: true,
            project: "./tsconfig.json", // Point to the correct tsconfig file
          },
        },
      },
      extends: [
        "plugin:@typescript-eslint/recommended",
        "plugin:import/recommended",
        "plugin:import/typescript",
      ],
    },
  ],
}
