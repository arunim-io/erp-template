import type { Theme } from ".";
import { useEffect, useMemo, useState } from "react";
import { ThemeProviderContext } from ".";

type ThemeProviderProps = React.PropsWithChildren<{
  children: React.ReactNode;
  defaultTheme?: Theme;
  storageKey?: string;
}>;

export function ThemeProvider({
  children,
  defaultTheme = "system",
  storageKey = "erp-web-theme",
  ...props
}: ThemeProviderProps) {
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

  const value = useMemo(() => ({
    theme,
    setTheme(theme: Theme) {
      localStorage.setItem(storageKey, theme);
      setTheme(theme);
    },
  }), [storageKey, theme]);

  return (
    <ThemeProviderContext {...props} value={value}>
      {children}
    </ThemeProviderContext>
  );
}
