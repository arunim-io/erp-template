import { authStore } from "$lib/stores";
import { graphql } from "$src/lib/gql";
import { queryOptions } from "@tanstack/react-query";
import { createFileRoute, Outlet, redirect } from "@tanstack/react-router";

const userAuthStatusQuery = graphql(`
  query UserAuthStatus {
    auth {
      status {
        username
        isActive
      }
    }
  }
`);

export const Route = createFileRoute("/_app")({
  component: RouteComponent,
  beforeLoad(ctx) {
    if (!authStore.getState().isAuthenticated) {
      throw redirect({ to: "/login", search: { redirect: ctx.location.href } });
    }
  },
  loader: ({ context: { graphqlClient } }) => ({
    authStatusQueryOptions: queryOptions({
      queryKey: ["auth", "user", "status"],
      retry: false,
      queryFn: () => graphqlClient.request(userAuthStatusQuery),
    }),
  }),
});

function RouteComponent() {
  return <Outlet />;
}
