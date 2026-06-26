import { useEffect } from "react";

export function useSystemTheme() {

  useEffect(() => {

    const media = window.matchMedia(
      "(prefers-color-scheme: dark)"
    );

    const updateTheme = () => {

      console.log(
        "Dark Mode:",
        media.matches
      );

      document.documentElement.classList.toggle(
        "dark",
        media.matches
      );

      document.documentElement.style.colorScheme =
        media.matches
          ? "dark"
          : "light";

      console.log(
        "HTML Classes:",
        document.documentElement.className
      );
    };

    updateTheme();

    media.addEventListener(
      "change",
      updateTheme
    );

    return () =>
      media.removeEventListener(
        "change",
        updateTheme
      );

  }, []);
}