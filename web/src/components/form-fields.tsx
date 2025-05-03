import type { ComponentPropsWithoutRef } from "react";
import type { Control } from "react-hook-form";
import { Eye, EyeOff, Lock, Mail, User } from "lucide-react";
import { useState } from "react";
import { Button } from "./ui/button";
import { FormControl, FormField, FormItem, FormLabel, FormMessage } from "./ui/form";
import { Input } from "./ui/input";

interface FieldProps extends ComponentPropsWithoutRef<"input"> {
  control: Control<any>;
  label?: string;
}

export function UsernameField({
  control,
  name = "username",
  label = "Username",
  ...props
}: FieldProps) {
  return (
    <FormField
      control={control}
      name={name}
      render={({ field }) => (
        <FormItem className="grid gap-2">
          <FormLabel>{label}</FormLabel>
          <div className="relative">
            <User className="absolute top-1/2 left-3 size-4 -translate-y-1/2 text-muted-foreground" />
            <FormControl>
              <Input placeholder="johndoe" className="pl-10" {...props} {...field} />
            </FormControl>
          </div>
          <FormMessage />
        </FormItem>
      )}
    />
  );
}

export function EmailField({
  control,
  label = "Email Address",
  placeholder = "johndoe@gmail.com",
  name = "email",
  ...props
}: FieldProps) {
  return (
    <FormField
      control={control}
      name={name}
      render={({ field }) => (
        <FormItem className="grid gap-2">
          <FormLabel>{label}</FormLabel>
          <div className="relative">
            <Mail className="absolute top-1/2 left-3 h-4 w-4 -translate-y-1/2 text-muted-foreground" />
            <FormControl>
              <Input placeholder={placeholder} type="email" className="pl-10" {...props} {...field} />
            </FormControl>
          </div>
          <FormMessage />
        </FormItem>
      )}
    />
  );
}

interface PasswordFieldProps extends FieldProps {
  showForgotPasswordLink?: boolean;
}

export function PasswordField({
  control,
  name = "password",
  label = "Password",
  placeholder = "**********",
  showForgotPasswordLink = false,
  ...props
}: PasswordFieldProps) {
  const [show, setShow] = useState(false);

  const toggle = () => setShow(show => !show);
  const Icon = show ? EyeOff : Eye;

  return (
    <FormField
      control={control}
      name={name}
      render={({ field }) => (
        <FormItem className="grid gap-2">
          <div className="flex items-center">
            <FormLabel>{label}</FormLabel>
            {showForgotPasswordLink && (
              <a href="#" className="ml-auto text-sm underline-offset-4 hover:underline">
                Forgot your password?
              </a>
            )}
          </div>
          <div className="relative">
            <Lock className="absolute top-1/2 left-3 h-4 w-4 -translate-y-1/2 text-muted-foreground" />
            <FormControl>
              <Input type={show ? "text" : "password"} placeholder={placeholder} className="pl-10" {...props} {...field} />
            </FormControl>
            <Button
              type="button"
              variant="ghost"
              size="sm"
              className="absolute top-0 right-0 h-full px-3 py-2 hover:bg-transparent"
              onClick={toggle}
              aria-label={show ? "Hide password" : "Show password"}
            >
              <Icon className="h-4 w-4 text-muted-foreground" />
            </Button>
          </div>
          <FormMessage />
        </FormItem>
      )}
    />
  );
}
