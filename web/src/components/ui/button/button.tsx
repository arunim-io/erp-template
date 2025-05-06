import type { ButtonVariantProps } from "./variants";
import { cn } from "$lib/utils";
import { Slot } from "@radix-ui/react-slot";
import { buttonVariants } from "./variants";

export function Button({
  className,
  variant,
  size,
  asChild = false,
  ...props
}: React.ComponentProps<"button"> & ButtonVariantProps & { asChild?: boolean }) {
  const Comp = asChild ? Slot : "button";

  return (
    <Comp
      data-slot="button"
      className={cn(buttonVariants({ variant, size, className }))}
      {...props}
    />
  );
}
