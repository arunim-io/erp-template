import Antfu from "@antfu/eslint-config";
import TanstackQuery from "@tanstack/eslint-plugin-query";
import TanstackRouter from "@tanstack/eslint-plugin-router";

export default Antfu({
  ignores: ["**/*.gen.ts", "**/.venv/", "./backend/**/fixtures/*.json", "./web/src/lib/gql/"],
  stylistic: {
    quotes: "double",
    semi: true,
  },
  formatters: true,
  react: true,
  toml: false,
}, ...TanstackRouter.configs["flat/recommended"], ...TanstackQuery.configs["flat/recommended"]);
