import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import { createRouter, RouterProvider } from "@tanstack/react-router";
import { GraphQLClient } from "graphql-request";
import { StrictMode } from "react";
import { createRoot } from "react-dom/client";
import { ThemeProvider } from "./components/theme";
import { routeTree } from "./routeTree.gen";
import "./index.css";

const queryClient = new QueryClient();

const graphqlClient = new GraphQLClient(import.meta.env.VITE_API_URL, { credentials: "include" });

const router = createRouter({
  routeTree,
  defaultPreloadStaleTime: 0,
  context: { queryClient, graphqlClient },
});

declare module "@tanstack/react-router" {
  interface Register {
    router: typeof router;
  }
}

const rootElement = document.getElementById("root");

if (!rootElement) {
  throw new Error("An element with id `root` must be specified in the body of the entrypoint template file.");
}
if (rootElement.innerHTML) {
  throw new Error("The element must be empty for React to fully render.");
}

createRoot(rootElement).render(
  <StrictMode>
    <ThemeProvider>
      <QueryClientProvider client={queryClient}>
        <RouterProvider router={router} />
      </QueryClientProvider>
    </ThemeProvider>
  </StrictMode>,
);
