import { createContext, use } from "react";

export type Theme = "dark" | "light" | "system";

interface ThemeProviderState {
  theme: Theme;
  setTheme: (theme: Theme) => void;
}

export const ThemeProviderContext = createContext<ThemeProviderState>({
  theme: "system",
  setTheme: () => null,
});

export function useTheme() {
  const context = use(ThemeProviderContext);

  if (context === undefined) {
    throw new Error("`useTheme` must be used within a `ThemeProvider`");
  }

  return context;
}

export { ThemeProvider } from "./provider";
export { ModeToggle as ThemeToggle } from "./toggle";
