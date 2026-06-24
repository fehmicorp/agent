import React from "react";
import { cardMain, spanText } from "../../utils/colour";

interface AdminSettingItem {
  id: number;
  title: string;
  tag: string;
  value: boolean;
}

interface AdminLayoutProps {
  setIsAdmin: (val: boolean) => void;
  adminSettings: AdminSettingItem[];
  settingsState: Record<string, string>;
  onSettingChange: (tagIdKey: string, value: string) => void;
}

/* ========================================================
   1. ADMIN AUTHENTICATED COMPONENT LAYOUT
   ======================================================== */
export const AdminLayout: React.FC<AdminLayoutProps> = ({
  setIsAdmin,
  adminSettings,
  settingsState,
  onSettingChange,
}) => {
  return (
    <div className={`${cardMain} rounded-2xl p-4 border border-stone-500/20 bg-stone-200/50 dark:bg-stone-950/5 font-mono text-xs space-y-4 animate-fadeIn`}>
      {/* Header administrative toolbar banner */}
      <div className="flex justify-between items-center border-b border-stone-700/20 dark:border-stone-300/20 pb-2">
        <div>
          <div className="flex items-center gap-2">
            <h2 className={`${spanText} text-lg font-bold`}>Control Panel</h2>
          </div>
          <p className={`text-stone-700 dark:text-stone-400 text-[10px] mt-0.5`}>Authenticated Secure Session token active — Auto-expires in 30m.</p>
        </div>
        <button 
          onClick={() => setIsAdmin(false)} 
          className="px-3 py-1.5 bg-red-700 border border-stone-200/20 rounded-xl text-stone-200 hover:bg-red-600 transition-colors active:scale-95 text-[11px]"
        >
          Logout Session
        </button>
      </div>

      {/* Dynamic Radio Settings Interactive Responsive Grid Container */}
      <div className="grid grid-cols-2 gap-2">
        {adminSettings.map((item) => {
          const controlKey = `${item.tag}-${item.id}`;
          const currentSelection = settingsState[controlKey] || (item.value ? "on" : "off");
          return (
            <div key={controlKey} className={`flex justify-between items-center ${cardMain} p-4 rounded-xl`}>
              <p className={`${spanText} text-md truncate`}>{item.title}</p>
              <button
                  type="button"
                  onClick={() => onSettingChange(controlKey, currentSelection === "on" ? "off" : "on")}
                  className={`relative inline-flex h-5 w-10 shrink-0 cursor-pointer rounded-full border-2 border-transparent transition-colors duration-200 ease-in-out outline-none focus:ring-1 focus:ring-purple-500/50 ${
                  currentSelection === "on" ? "bg-green-600" : "bg-red-600"
                  }`}
              >
                  <span
                  className={`pointer-events-none inline-block h-4 w-4 transform rounded-full bg-stone-100 shadow ring-0 transition duration-200 ease-in-out ${
                      currentSelection === "on" ? "translate-x-5" : "translate-x-0"
                  }`}
                  />
              </button>
            </div>
          );
        })}
      </div>
    </div>
  );
};