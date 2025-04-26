import type { Theme, ThemeState } from "./store";
import { useEffect, useMemo, useState } from "react";
import { ThemeContext } from "./store";

interface Props {
  defaultTheme?: Theme;
  storageKey?: string;
}

export function ThemeProvider({
  children,
  defaultTheme = "system",
  storageKey = "vite-ui-theme",
  ...props
}: React.PropsWithChildren<Props>) {
  const [theme, setTheme] = useState<Theme>(
    () => (localStorage.getItem(storageKey) as Theme) || defaultTheme,
  );

  useEffect(() => {
    const root = window.document.documentElement;

    root.classList.remove("light", "dark");

    if (theme === "system") {
      const systemTheme = window.matchMedia("(prefers-color-scheme: dark)")
        .matches
        ? "dark"
        : "light";

      root.classList.add(systemTheme);
      return;
    }

    root.classList.add(theme);
  }, [theme]);

  const value = useMemo<ThemeState>(() => ({
    theme,
    setTheme(theme: Theme) {
      localStorage.setItem(storageKey, theme);
      setTheme(theme);
    },
  }), [storageKey, theme]);

  return (
    <ThemeContext {...props} value={value}>
      {children}
    </ThemeContext>
  );
}
