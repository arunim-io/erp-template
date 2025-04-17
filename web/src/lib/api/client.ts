import ky from "ky";
import SuperJSON from "superjson";

export const client = ky.create({
  parseJson: SuperJSON.parse,
  stringifyJson: SuperJSON.stringify,
  prefixUrl: import.meta.env.VITE_API_URL,
});
