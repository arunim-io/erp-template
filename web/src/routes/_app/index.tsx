import { LogoutDialog } from "$components/auth/logout-dialog";
import { useQuery } from "@tanstack/react-query";
import { createFileRoute, useLoaderData } from "@tanstack/react-router";

export const Route = createFileRoute("/_app/")({
  component: RouteComponent,
});

function RouteComponent() {
  const { authStatusQueryOptions } = useLoaderData({ from: "/_app" });
  const { data } = useQuery(authStatusQueryOptions);

  return (
    <>
      <p>
        Logged In!
        <br />
        Username:
        {" "}
        {data?.auth.status.username}
      </p>
      <LogoutDialog />
    </>
  );
}
