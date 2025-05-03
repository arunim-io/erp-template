import { Button } from "$components/ui/button";
import { Form } from "$components/ui/form";
import { zodResolver } from "@hookform/resolvers/zod";
import { SiFacebook, SiGoogle } from "@icons-pack/react-simple-icons";
import { useForm } from "react-hook-form";
import { z } from "zod";
import { EmailField, PasswordField } from "./form-fields";

const loginFormSchema = z.object({
  email: z.string().email({ message: "Please enter a valid email address" }),
  password: z.string().min(8, { message: "Password must be at least 8 characters" }),
});

type LoginFormValues = z.infer<typeof loginFormSchema>;

export function LoginForm() {
  const form = useForm<LoginFormValues>({
    resolver: zodResolver(loginFormSchema),
  });

  function onSubmit(values: LoginFormValues) {
    // eslint-disable-next-line no-console
    console.log(values);
  }

  return (
    <div className="flex flex-col gap-6">
      <div className="flex flex-col items-center gap-2 text-center">
        <h1 className="text-2xl font-bold">Login to your account</h1>
        <p className="text-sm text-balance text-muted-foreground">Enter your email below to login to your account</p>
      </div>
      <Form {...form}>
        <form onSubmit={form.handleSubmit(onSubmit)} className="grid gap-6">
          <EmailField control={form.control} />
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
