import React, { useState, useEffect, useRef } from "react";
import { cardMain, spanText, mutedText } from "../../utils/colour";
import App from "../data";
import { dash, upgradeAgent } from "../data/dummy";
import { AdminLayout } from "./admin";
import { useAppInfoStore } from "../../store/dashboard";

const adminSetting = [
  { id: 1, title: "Network Firewall", tag: "firewall", value: true },
  { id: 2, title: "Files Scanning", tag: "fscan", value: true },
  { id: 3, title: "Microsoft Windows Defender", tag: "mdefender", value: true },
  { id: 4, title: "BitLocker", tag: "firewall", value: true },
  { id: 5, title: "Verbose Telemetry Stream", tag: "log", value: true },
];

export default function Agent(): React.JSX.Element {
  const [isAdmin, setIsAdmin] = useState(false);
  const [password, setPassword] = useState("");
  const [isUpgrading, setIsUpgrading] = useState(false);
  const [error, setError] = useState(""); 
  const [isShaking, setIsShaking] = useState(false); // Triggers left-right error vibration

  // Focus reference pointer node target
  const inputRef = useRef<HTMLInputElement>(null);
  
  // Dynamic mapped setting pipeline dictionary replaces flat separate hooks
  const [settingsState, setSettingsState] = useState<Record<string, string>>(() => {
    const initialState: Record<string, string> = {};
    adminSetting.forEach((item) => {
      initialState[`${item.tag}-${item.id}`] = item.value ? "on" : "off";
    });
    return initialState;
  });
  
  // Track runtime state version dynamically
  const [currentVersion, setCurrentVersion] = useState({
    version: App.version,
    tag: App.tag
  });

  // Automatic Session Expiration Handler (Runs 30 minutes countdown)
  useEffect(() => {
    let logoutTimer: ReturnType<typeof setTimeout>;
    
    if (isAdmin) {
      logoutTimer = setTimeout(() => {
        setIsAdmin(false);
        alert("Administrative session expired (30-minute timeout security rule triggered).");
      }, 30 * 60 * 1000);
    }
    
    return () => clearTimeout(logoutTimer);
  }, [isAdmin]);

  // Handle 5-second automatic warning message clearance routines
  useEffect(() => {
    if (!error) return;

    const autoClearTimer = setTimeout(() => {
      setError("");
    }, 5000);

    return () => clearTimeout(autoClearTimer);
  }, [error]);

  const handleAdminLogin = (e: React.FormEvent) => {
    e.preventDefault();
    if (password === "admin") {
      setIsAdmin(true);
      setPassword("");
      setError("");
    } else {
      setPassword("");
      setError("Invalid Admin Authentication Credentials. Verification Rejected.");
      
      // Trigger side-to-side shake vibration sequence specifically on the input line component
      setIsShaking(true);
      setTimeout(() => setIsShaking(false), 500); // Matches animation timing duration

      // Return execution priority state to input field after string reset finishes
      setTimeout(() => {
        inputRef.current?.focus();
      }, 0);
    }
  };

  const handlePasswordChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setPassword(e.target.value);
    if (error) setError(""); 
  };

  const triggerUpgrade = () => {
    setIsUpgrading(true);
    setTimeout(() => {
      // Pull dynamic configuration from dummy upgrader target metadata
      setCurrentVersion({
        version: upgradeAgent.version,
        tag: upgradeAgent.tag
      });
      setIsUpgrading(false);
    }, 4000);
  };

  const handleSettingChange = (controlKey: string, value: string) => {
    setSettingsState((prev) => ({
      ...prev,
      [controlKey]: value,
    }));
  };

  const appinfo = useAppInfoStore((state) => state.appInfo);

  // Determine if the current version matches the latest upgrade target profile
  const isLatestVersion = currentVersion.version === upgradeAgent.version;

  /* ========================================================
     1. ADMIN AUTHENTICATED STATE: Clean screen control deck
     ======================================================== */
  if (isAdmin) {
    return (
      <AdminLayout 
        setIsAdmin={setIsAdmin}
        adminSettings={adminSetting}
        settingsState={settingsState}
        onSettingChange={handleSettingChange}
      />
    );
  }

  /* ========================================================
     2. STANDARD STATE: Default view workspace
     ======================================================== */
  return (
    <div className="space-y-4 pt-4 px-2">
      {/* 1. Top Row Dashboard Overview Grid */}
      <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
        
        {/* Agent Metadata Card */}
        <div className={`${cardMain} rounded-2xl p-3 border border-stone-200/5 relative overflow-hidden`}>
          <div className="flex justify-between items-start z-10 relative">
            <span className="text-[10px] uppercase font-mono tracking-wider dark:bg-sky-500/20 bg-sky-500/40 text-sky-800 dark:text-sky-300 px-2 py-0.5 rounded border border-sky-500/20 dark:border-sky-500/20">
              Core Endpoint Service
            </span>
            {dash.status ? (
              <div className="flex items-center gap-1.5 bg-emerald-500/20 border-emerald-500/60 dark:bg-emerald-500/10 px-2.5 py-1 rounded-full border dark:border-emerald-500/20">
                <span className="h-2 w-2 rounded-full bg-emerald-500 dark:bg-emerald-400 animate-pulse" />
                <span className="text-[10px] text-emerald-800 dark:text-emerald-400 font-bold uppercase font-mono">ONLINE</span>
              </div>
            ) : (
              <div className="flex items-center gap-1.5 bg-red-500/40 dark:bg-red-500/10 px-2.5 py-1 rounded-full border border-red-500/40 dark:border-red-500/20">
                <span className="h-2 w-2 rounded-full bg-red-600 dark:bg-red-400 animate-pulse" />
                <span className="text-[10px] text-red-800 dark:text-red-400 font-bold uppercase font-mono">OFFLINE</span>
              </div>
            )}
          </div>          
          <div className="grid grid-cols-[50%_25%_25%] items-end text-xs font-mono w-full my-1">
            <div className="flex flex-col justify-end min-w-0 pr-4">
              <h2 className={`text-stone-900 dark:text-stone-100 text-xl font-bold truncate`}>{appinfo.name}</h2>
              <p className={`${mutedText} text-xs mt-1 truncate`}>
                ID: {appinfo.deviceToken}
              </p>
            </div>
            <div className="min-w-0">
              <p className={`${spanText} text-sm mt-1 truncate`}>Service Version</p>
              <p className={`${mutedText} text-xs mt-0.5 truncate`}>
                {currentVersion.version}-{currentVersion.tag}
              </p>
            </div>
            <div className="min-w-0">
              <p className={`${spanText} text-sm mt-1 truncate`}>Last Handshake</p>
              <p className={`${mutedText} text-xs mt-0.5 truncate`}>Just Now</p>
            </div>
          </div>
        </div>

        {/* Upgrade Card View Engine */}
        <div className={`${cardMain} rounded-2xl p-3 flex flex-col justify-center h-full`}>
          <div className="flex items-center justify-between gap-4">
            <div>
              <h3 className={`${spanText} font-semibold text-xs`}>Agent Upgrade</h3>
              <p className={`${mutedText} font-mono text-[11px] mt-1`}>
                {isUpgrading 
                  ? "Staging binary updates..." 
                  : isLatestVersion 
                    ? "✓ System is up to date." 
                    : `⚙ Update '${upgradeAgent.version}-${upgradeAgent.tag}' ready.`}
              </p>
            </div>

            <div className="flex-shrink-0">
              {isUpgrading ? (
                <div className="flex items-center gap-2 font-mono text-[11px] text-amber-700 dark:text-amber-400">
                  <span className="h-2 w-2 rounded-full bg-amber-700 dark:bg-amber-400 animate-ping" />
                  <span>Processing...</span>
                </div>
              ) : (
                <button
                  disabled={isLatestVersion}
                  onClick={triggerUpgrade}
                  className={`px-4 py-[0.4rem] text-[11px] transition-all duration-300 ease-in-out hover:scale-105 font-medium rounded-xl font-mono whitespace-nowrap active:scale-95 ${
                    isLatestVersion
                      ? "bg-stone-800 text-stone-600 cursor-not-allowed border border-stone-700/20"
                      : "bg-gradient-to-tr from-sky-200 to-cyan-400 border border-blue-800/20 dark:border-sky-200 dark:from-sky-500 dark:to-slate-300 text-slate-950 tracking-[0.15rem] font-semibold hover:shadow-[0_0_10px_rgba(56,189,248,0.2)]"
                  }`}
                >
                  {isLatestVersion ? "Latest" : "Upgrade"}
                </button>
              )}
            </div>
          </div>
        </div>
      </div>

      {/* 3. Administrative Authorization Control Section (Standard Gate) */}
      <div className={`${cardMain} rounded-2xl px-3 py-4 border border-stone-200/5 mt-4 space-y-3`}>
        <div className="flex items-center gap-2">
          <h3 className={`${spanText} font-semibold text-sm`}>Administrative Terminal Gateway</h3>
        </div>
        
        <form onSubmit={handleAdminLogin} className="flex flex-row gap-3 max-w-xl">
          {/* The shake class wrapper is exclusively attached around the input context layout container node */}
          <div className={`relative flex-1 transition-transform duration-500 ${
            isShaking ? "animate-[shake_0.4s_ease-in-out]" : ""
          }`}>
            <input
              ref={inputRef}
              type="password"
              required
              placeholder="Provide master administration passphrase keys (type 'admin' to bypass)..."
              value={password}
              onChange={handlePasswordChange}
              className={`w-full bg-stone-400/30 dark:bg-stone-950/40 text-stone-800 dark:text-stone-200 border 
                rounded-xl px-4 py-2.5 text-xs font-mono focus:outline-none transition-colors 
                placeholder:text-stone-600 ${
                  error 
                    ? "border-red-500/50 focus:border-red-500 ring-1 ring-red-500/20" 
                    : "dark:border-stone-200/10 border-stone-400/30 dark:focus:border-blue-500/50 focus:border-blue-600/50"
                }`}
            />
          </div>
          <button
            type="submit"
            className="px-5 py-2.5 bg-stone-800 text-stone-200 dark:hover:bg-green-600/40 hover:bg-green-700/90 active:scale-95 text-xs font-medium font-mono rounded-xl border border-stone-200/10 transition-all shadow-md whitespace-nowrap"
          >
            Verify
          </button>
        </form>
        {error && (
          <p className="text-[11px] font-mono text-red-600 dark:text-red-400 font-semibold tracking-wide flex items-center gap-1.5 animate-fadeIn">
            <span className="h-1.5 w-1.5 rounded-full bg-red-500 animate-pulse" />
            {error}
          </p>
        )}
      </div>
    </div>
  );
}