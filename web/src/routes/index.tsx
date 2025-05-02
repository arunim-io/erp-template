import { graphql } from "$lib/gql";
import { Button } from "$src/components/ui/button";
import { useQuery } from "@tanstack/react-query";
import { createFileRoute, Link, useRouteContext } from "@tanstack/react-router";

export const Route = createFileRoute("/")({
  component: RouteComponent,
});

const UserAuthStatusQuery = graphql(/* graphql */`
  query UserAuthStatusQuery {
    status {
      username
      email
      firstName
      lastName
    }
  }
`);

function RouteComponent() {
  const { graphqlClient } = useRouteContext({ from: "/" });
  const { data, isSuccess } = useQuery({ queryKey: ["auth", "user", "status"], queryFn: () => graphqlClient.request(UserAuthStatusQuery) });

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
