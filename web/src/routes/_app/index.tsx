import { Button } from "$components/ui/button";
import { useQuery } from "@tanstack/react-query";
import { createFileRoute, Link, useLoaderData } from "@tanstack/react-router";

export const Route = createFileRoute("/_app/")({
  component: RouteComponent,
});

function RouteComponent() {
  const { authStatusQueryOptions } = useLoaderData({ from: "/_app" });
  const { data, isSuccess } = useQuery(authStatusQueryOptions);

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
