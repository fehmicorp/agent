import React from "react";
import { navItems } from "../sections/navbar";

export default function Navbar() {

const current =
window.location.hash.replace("#", "") ||
"dashboard";

const navigate = (route: string) => {
window.location.hash = route;
};

return (
  <div className="fixed bottom-2 left-1/2 -translate-x-1/2 z-50">
    <div className="relative flex items-center gap-1 p-1 rounded-2xl bg-white/10 dark:bg-slate-950/10
      backdrop-blur-[0.5rem] backdrop-saturate-[150%] border border-white/10 dark:border-white/10
      shadow-[0_8px_32px_rgba(0,0,0,0.12)] overflow-visible transition-all px-2 duration-300 ease-in-out hover:shadow-blue-400/20 hover:shadow-md hover:scale-[1.05] hover:-translate-1">
      {navItems.map((item) => {
        const Icon = item.icon;
        const active = current === item.route;
          return (
            <button
              key={item.route}
              title={item.title}
              onClick={() =>
                navigate(item.route)
              }
              className={`relative flex items-center justify-center h-8 rounded-xl
                transition-all duration-200 ease-in-out
                ${
                  active
                    ? `bg-sky-100/10 text-stone-800 dark:text-stone-200 scale-120 hover:scale-[1.10] mx-2 hover:-translate-y-1`
                    : `text-stone-500 hover:text-stone-800 dark:hover:text-stone-200 hover:scale-[1.10] mx-1/2 hover:-translate-y-1 w-10`
                }
              `}
            >
              {active 
                ? <div className="flex items-center p-2 text-[12px] space-x-2"><Icon size={18} className="relative z-10 shrink-0"/></div>
                : <Icon size={20} className="relative z-10 shrink-0"/>
              }
            </button>
          );
      })}
    </div>
  </div>
);
}
