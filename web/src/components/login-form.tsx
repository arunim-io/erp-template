import type { UserLoginMutation } from "$lib/gql/_generated/graphql";
import type { LoginFormValues } from "$src/lib/schemas/forms";
import { Button } from "$components/ui/button";
import { Form } from "$components/ui/form";
import { graphql } from "$lib/gql";
import { authStore } from "$lib/stores";
import { loginFormSchema } from "$src/lib/schemas/forms";
import { zodResolver } from "@hookform/resolvers/zod";
import { SiFacebook, SiGoogle } from "@icons-pack/react-simple-icons";
import { useMutation } from "@tanstack/react-query";
import { useLoaderData, useNavigate } from "@tanstack/react-router";
import { useForm } from "react-hook-form";
import { PasswordField, UsernameField } from "./form-fields";

export function LoginForm() {
  const navigate = useNavigate({ from: "/login" });
  const { loginMutationOptions } = useLoaderData({ from: "/login" });

  const { mutateAsync: loginUser } = useMutation<UserLoginMutation, Error, LoginFormValues>({
    ...loginMutationOptions,
    async onSuccess() {
      authStore.setState(state => ({
        ...state,
        isAuthenticated: true,
      }));

      await navigate({ to: "/" });
    },
  });
  const form = useForm<LoginFormValues>({
    resolver: zodResolver(loginFormSchema),
    defaultValues: { username: "", password: "" },
  });

  async function onSubmit(values: LoginFormValues) {
    await loginUser(values);
  }

  return (
    <div className="flex flex-col gap-6">
      <div className="flex flex-col items-center gap-2 text-center">
        <h1 className="text-2xl font-bold">Login to your account</h1>
        <p className="text-sm text-balance text-muted-foreground">Enter your email below to login to your account</p>
      </div>
      <Form {...form}>
        <form onSubmit={form.handleSubmit(onSubmit)} className="grid gap-6">
          <UsernameField control={form.control} />
          <PasswordField control={form.control} showForgotPasswordLink />
          <Button type="submit" className="w-full">
            Login
          </Button>
          <div className="relative text-center text-sm after:absolute after:inset-0 after:top-1/2 after:z-0 after:flex after:items-center after:border-t after:border-border">
            <span className="relative z-10 bg-background px-2 text-muted-foreground">Or continue with</span>
          </div>
          <div className="flex flex-col gap-3">
            <Button type="button" variant="outline" className="w-full">
              <SiGoogle className="mr-2 size-4" />
              Login with Google
            </Button>
            <Button type="button" variant="outline" className="w-full">
              <SiFacebook className="mr-2 size-4" />
              Login with Facebook
            </Button>
          </div>
        </form>
      </Form>
      <div className="flex flex-col gap-2 text-center text-sm">
        <p className="space-x-1">
          <span>
            Don&apos;t have an account?
          </span>
          <a href="#" className="underline underline-offset-4">
            Sign up
          </a>
        </p>
        <p className="space-x-1 text-xs text-muted-foreground">
          <span>
            By continuing, you agree to our
          </span>
          <a href="#" className="underline underline-offset-4 hover:text-primary">
            Terms of Service
          </a>
          <span>
            and
          </span>
          <a href="#" className="underline underline-offset-4 hover:text-primary">
            Privacy Policy
          </a>
        </p>
      </div>
    </div>
  );
}
