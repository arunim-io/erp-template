import antfu from "@antfu/eslint-config";

export default antfu({
  ignores: ["**/*.gen.ts", "**/.venv/"],
  stylistic: {
    quotes: "double",
    semi: true,
  },
  formatters: true,
  react: true,
  toml: false,
});
