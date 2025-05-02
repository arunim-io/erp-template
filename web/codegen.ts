import type { CodegenConfig } from "@graphql-codegen/cli";
import process from "node:process";

const config: CodegenConfig = {
  overwrite: true,
  schema: process.env.VITE_API_URL,
  documents: ["src/**/*.ts", "src/**/*.tsx"],
  ignoreNoDocuments: true,
  generates: {
    "src/lib/gql/": {
      preset: "client",
      presetConfig: {
        persistedDocuments: true,
      },
      config: {
        documentMode: "string",
        immutableTypes: true,
        useTypeImports: true,
        extractAllFieldsToTypes: true,
        fetcher: "graphql-request",
      },
    },
  },
};

export default config;
