import ky from "ky";
import SuperJSON from "superjson";

export const client = ky.create({
  parseJson: SuperJSON.parse,
  stringifyJson: SuperJSON.stringify,
  prefixUrl: import.meta.env.VITE_API_URL,
  headers: {
    "Content-Type": "application/vnd.api+json",
    "Accept": "application/vnd.api+json",
  },
});
