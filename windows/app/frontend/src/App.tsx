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
import data from "./sections/data";

import { useSystemTheme } from "./utils/theme";
import Navbar from "./component/Navbar";
import Header from "./component/Header";

export default function App() {

  useSystemTheme();

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
    <div className="h-screen bg-white text-slate-900 dark:bg-slate-950 dark:text-white">
      <Suspense fallback={
        <div className="p-4">
          Loading...
        </div>
        }
      >
        <div className="bg-stone-100 dark:bg-stone-950 min-h-screen">
          {
            route == "agent" ? <></> : <Header title={data.name} />
          }
          <div className="p-2 pb-14 overflow-hidden">
            <Page />
          </div>
        </div>
        <Navbar />
      </Suspense>
    </div>
  );
}