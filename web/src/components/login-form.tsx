import { Button } from "$src/components/ui/button";
import { Input } from "$src/components/ui/input";
import { Label } from "$src/components/ui/label";
import { client } from "$src/lib/api";
import { zodResolver } from "@hookform/resolvers/zod";
import { SiFacebook, SiGoogle } from "@icons-pack/react-simple-icons";
import { useMutation } from "@tanstack/react-query";
import { useForm } from "react-hook-form";
import * as z from "zod";
import { Form, FormField } from "./ui/form";

const loginFormSchema = z.interface({
  username: z.string(),
  password: z.string(),
});

type LoginFormSchema = z.infer<typeof loginFormSchema>;

async function loginUser(payload: LoginFormSchema) {
  const response = await client.post("accounts/login/", {
    json: {
      data: {
        type: "auth",
        attributes: payload,
      },
    },
  });
}

export function LoginForm() {
  const mutation = useMutation({ mutationFn: loginUser });
  const form = useForm<LoginFormSchema>({
    resolver: zodResolver(loginFormSchema),
  });

  const onSubmit = form.handleSubmit(
    async values => mutation.mutateAsync(values),
  );

  return (
    <Form {...form}>
      <form onSubmit={onSubmit} className="flex flex-col gap-6">
        <div className="flex flex-col items-center gap-2 text-center">
          <h1 className="text-2xl font-bold">Login to your account</h1>
          <p className="text-sm text-balance text-muted-foreground">
            Enter your email below to login to your account
          </p>
        </div>
        <div className="grid gap-6">
          <FormField
            control={form.control}
            name="username"
            render={({ field }) => (
              <div className="grid gap-3">
                <Label htmlFor="email">Username</Label>
                <Input placeholder="johndoe" type="text" {...field} />
              </div>
            )}
          />
          <FormField
            control={form.control}
            name="password"
            render={({ field }) => (
              <div className="grid gap-3">
                <div className="flex items-center">
                  <Label htmlFor="password">Password</Label>
                  <a
                    href="#"
                    className="ml-auto text-sm underline-offset-4 hover:underline"
                  >
                    Forgot your password?
                  </a>
                </div>
                <Input placeholder="**********" type="password" {...field} />
              </div>
            )}
          />
          <Button type="submit" className="w-full">
            Login
          </Button>
          <div className="relative text-center text-sm after:absolute after:inset-0 after:top-1/2 after:z-0 after:flex after:items-center after:border-t after:border-border">
            <span className="relative z-10 bg-background px-2 text-muted-foreground">
              Or continue with
            </span>
          </div>
          <Button variant="outline" className="w-full">
            <SiGoogle />
            Login with Google
          </Button>
          <Button variant="outline" className="w-full">
            <SiFacebook />
            Login with Facebook
          </Button>
        </div>
        <div className="text-center text-sm">
          Don&apos;t have an account?
          {" "}
          <a href="#" className="underline underline-offset-4">
            Sign up
          </a>
        </div>
      </form>
    </Form>
  );
}
