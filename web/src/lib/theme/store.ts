import { createContext, use } from "react";

export type Theme = "dark" | "light" | "system";

export interface ThemeState {
  theme: Theme;
  setTheme: (theme: Theme) => void;
}

export const ThemeContext = createContext<ThemeState>({
  theme: "system",
  setTheme: () => null,
});

export function useTheme() {
  const context = use(ThemeContext);

  if (context === undefined)
    throw new Error("useTheme must be used within a ThemeProvider");

  return context;
}
