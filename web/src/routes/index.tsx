import { graphql } from "$lib/gql";
import { Button } from "$src/components/ui/button";
import { useQuery } from "@tanstack/react-query";
import { createFileRoute, Link, useRouteContext } from "@tanstack/react-router";

export const Route = createFileRoute("/")({
  component: RouteComponent,
});

const UserAuthStatusQuery = graphql(`
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
    queryFn: () => graphqlClient.request(UserAuthStatusQuery),
  });

  return isSuccess
    ? (
      <p>
        Logged In!
        <br />
        Username:
        {" "}
        {data?.status.username}
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
