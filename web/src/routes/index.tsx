import { graphql } from "$lib/gql";
import { Button } from "$src/components/ui/button";
import { authStore } from "$src/lib/stores";
import { useQuery } from "@tanstack/react-query";
import { createFileRoute, Link, redirect, useRouteContext } from "@tanstack/react-router";

export const Route = createFileRoute("/")({
  component: RouteComponent,
  beforeLoad(ctx) {
    if (!authStore.getState().isAuthenticated) {
      throw redirect({ to: "/login", search: { redirect: ctx.location.href } });
    }
  },
});

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

function RouteComponent() {
  const { graphqlClient } = useRouteContext({ from: "/" });
  const { data, isSuccess } = useQuery({
    queryKey: ["auth", "user", "status"],
    retry: false,
    queryFn: () => graphqlClient.request(userAuthStatusQuery),
  });

  return isSuccess
    ? (
        <p>
          Logged In!
          <br />
          Username:
          {" "}
          {data?.auth.status.username}
        </p>
      )
    : (
        <>
          <h1>Not Logged In!</h1>
          <Link to="/login">
            <Button>Login</Button>
          </Link>
        </>
      );
}
