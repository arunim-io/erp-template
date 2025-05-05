import type { LoginFormValues } from "$src/lib/schemas/forms";
import { LoginForm } from "$components/login-form";
import { graphql } from "$src/lib/gql";
import { authStore } from "$src/lib/stores";
import { createFileRoute, useRouter, useSearch } from "@tanstack/react-router";
import { GalleryVerticalEnd } from "lucide-react";
import { useEffect } from "react";
import z from "zod";

const userLoginMutation = graphql(`
  mutation UserLogin($username: String!, $password: String!) {
    auth {
      login(username: $username, password: $password) {
      isActive
      }
    }
  }
`);

export const Route = createFileRoute("/_auth/login")({
  component: LoginPage,
  validateSearch: z.object({ redirect: z.string().optional() }),
  loader: ({ context: { graphqlClient } }) => ({
    loginMutationOptions: {
      mutationKey: ["auth", "user", "login"],
      mutationFn: (values: LoginFormValues) => graphqlClient.request(userLoginMutation, values),
    },
  }),
});

function LoginPage() {
  const { history } = useRouter();
  const { redirect } = useSearch({ from: "/_auth/login" });
  const { isAuthenticated } = authStore.getState();

  useEffect(() => {
    if (isAuthenticated && redirect) {
      history.push(redirect);
    }
  }, [history, isAuthenticated, redirect]);

  return (
    <div className="grid min-h-svh lg:grid-cols-2">
      <div className="flex flex-col gap-4 p-6 md:p-10">
        <div className="flex justify-center gap-2 md:justify-start">
          <a href="/" className="flex items-center gap-2 font-medium">
            <div className="flex h-6 w-6 items-center justify-center rounded-md bg-primary text-primary-foreground">
              <GalleryVerticalEnd className="size-4" />
            </div>
            Acme Inc.
          </a>
        </div>
        <div className="flex flex-1 items-center justify-center">
          <div className="w-full max-w-xs">
            <LoginForm />
          </div>
        </div>
      </div>
      <div className="relative hidden bg-muted lg:block">
        <img
          src="/placeholder.svg"
          alt="Image"
          className="absolute inset-0 h-full w-full object-cover dark:brightness-[0.2] dark:grayscale"
        />
      </div>
    </div>
  );
}
