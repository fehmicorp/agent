import React, {
  Suspense,
  useEffect,
  useState,
} from "react";

import {
  routes,
  AppRoute,
  getCurrentRoute,
} from "./sections/registry";

export default function App() {

  const [route, setRoute] =
    useState<AppRoute>(
      getCurrentRoute()
    );

  useEffect(() => {

    const handler = () => {
      setRoute(
        getCurrentRoute()
      );
    };

    window.addEventListener(
      "hashchange",
      handler
    );

    handler();

    return () => {
      window.removeEventListener(
        "hashchange",
        handler
      );
    };

  }, []);

  const Page = routes[route];

  return (
    <div className="h-screen bg-slate-950 text-white">

      <Suspense
        fallback={
          <div className="p-4">
            Loading...
          </div>
        }
      >
        <Page />
      </Suspense>

    </div>
  );
}