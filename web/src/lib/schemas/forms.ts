import z from "zod";

export const loginFormSchema = z.object({
  username: z.string(),
  password: z.string().min(8, { message: "Password must be at least 8 characters" }),
});

export type LoginFormValues = z.infer<typeof loginFormSchema>;
