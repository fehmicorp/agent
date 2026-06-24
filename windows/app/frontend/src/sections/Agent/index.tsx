import React, { useState, useEffect } from "react";
import { cardMain, spanText, mutedText } from "../../utils/colour";
import App from "../data";
import { dash, upgradeAgent } from "../data/dummy";

export default function Agent(): React.JSX.Element {
  const [isAdmin, setIsAdmin] = useState(false);
  const [password, setPassword] = useState("");
  const [isUpgrading, setIsUpgrading] = useState(false);
  
  // Administrative radio control panel overrides states
  const [firewall, setFirewall] = useState("on");
  const [proxyMode, setProxyMode] = useState("off");
  const [liveLogs, setLiveLogs] = useState("on");
  
  // Track runtime state version dynamically
  const [currentVersion, setCurrentVersion] = useState({
    version: App.version,
    tag: App.tag
  });

  // Automatic Session Expiration Handler (Runs 30 minutes countdown)
  useEffect(() => {
    let logoutTimer: ReturnType<typeof setTimeout>; // Fixed: Replaced NodeJS.Timeout
    
    if (isAdmin) {
      logoutTimer = setTimeout(() => {
        setIsAdmin(false);
        alert("Administrative session expired (30-minute timeout security rule triggered).");
      }, 30 * 60 * 1000);
    }
    
    return () => clearTimeout(logoutTimer);
  }, [isAdmin]);

  const handleAdminLogin = (e: React.FormEvent) => {
    e.preventDefault();
    if (password === "admin") {
      setIsAdmin(true);
      setPassword("");
    } else {
      alert("Invalid Admin Authentication Credentials");
    }
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

  // Determine if the current version matches the latest upgrade target profile
  const isLatestVersion = currentVersion.version === upgradeAgent.version;

  /* ========================================================
     1. ADMIN AUTHENTICATED STATE: Clean screen control deck
     ======================================================== */
  if (isAdmin) {
    return (
      <div className={`${cardMain} rounded-2xl p-6 border border-purple-500/20 bg-purple-950/5 font-mono text-xs space-y-6 animate-fadeIn`}>
        {/* Header administrative toolbar banner */}
        <div className="flex justify-between items-center border-b border-stone-200/5 pb-4">
          <div>
            <div className="flex items-center gap-2">
              <span className="h-2 w-2 rounded-full bg-purple-500 shadow-[0_0_8px_rgba(168,85,247,0.5)]" />
              <h2 className={`${spanText} text-lg font-bold`}>Administrative Operations Control Panel</h2>
            </div>
            <p className={`${mutedText} text-[11px] mt-0.5`}>Authenticated Secure Session token active — Auto-expires in 30m.</p>
          </div>
          <button 
            onClick={() => setIsAdmin(false)} 
            className="px-3 py-1.5 bg-stone-900 border border-stone-200/10 rounded-xl text-stone-400 hover:text-stone-200 transition-colors active:scale-95 text-[11px]"
          >
            Logout Session
          </button>
        </div>

        {/* Radio Settings Core Interactive Row Grid */}
        <div className="grid grid-cols-1 md:grid-cols-3 gap-4">
          
          {/* Radio Group 1: Firewall */}
          <div className="bg-stone-950/40 p-4 rounded-xl border border-stone-200/5 space-y-3">
            <p className="text-stone-400 font-medium">Network Firewall</p>
            <div className="flex items-center gap-4">
              <label className="flex items-center gap-2 cursor-pointer text-stone-300 select-none">
                <input 
                  type="radio" 
                  name="firewall" 
                  value="on" 
                  checked={firewall === "on"}
                  onChange={(e) => setFirewall(e.target.value)}
                  className="accent-purple-500 bg-stone-900 border-stone-700 w-3.5 h-3.5"
                />
                <span>Active (On)</span>
              </label>
              <label className="flex items-center gap-2 cursor-pointer text-stone-500 select-none">
                <input 
                  type="radio" 
                  name="firewall" 
                  value="off" 
                  checked={firewall === "off"}
                  onChange={(e) => setFirewall(e.target.value)}
                  className="accent-purple-500 bg-stone-900 border-stone-700 w-3.5 h-3.5"
                />
                <span>Disabled (Off)</span>
              </label>
            </div>
          </div>

          {/* Radio Group 2: Reverse Proxy Routing */}
          <div className="bg-stone-950/40 p-4 rounded-xl border border-stone-200/5 space-y-3">
            <p className="text-stone-400 font-medium">Reverse Proxy Routing</p>
            <div className="flex items-center gap-4">
              <label className="flex items-center gap-2 cursor-pointer text-stone-300 select-none">
                <input 
                  type="radio" 
                  name="proxy" 
                  value="on" 
                  checked={proxyMode === "on"}
                  onChange={(e) => setProxyMode(e.target.value)}
                  className="accent-purple-500 bg-stone-900 border-stone-700 w-3.5 h-3.5"
                />
                <span>Enable Proxy</span>
              </label>
              <label className="flex items-center gap-2 cursor-pointer text-stone-500 select-none">
                <input 
                  type="radio" 
                  name="proxy" 
                  value="off" 
                  checked={proxyMode === "off"}
                  onChange={(e) => setProxyMode(e.target.value)}
                  className="accent-purple-500 bg-stone-900 border-stone-700 w-3.5 h-3.5"
                />
                <span>Bypass Route</span>
              </label>
            </div>
          </div>

          {/* Radio Group 3: Telemetry Stream */}
          <div className="bg-stone-950/40 p-4 rounded-xl border border-stone-200/5 space-y-3">
            <p className="text-stone-400 font-medium">Verbose Telemetry Stream</p>
            <div className="flex items-center gap-4">
              <label className="flex items-center gap-2 cursor-pointer text-stone-300 select-none">
                <input 
                  type="radio" 
                  name="liveLogs" 
                  value="on" 
                  checked={liveLogs === "on"}
                  onChange={(e) => setLiveLogs(e.target.value)}
                  className="accent-purple-500 bg-stone-900 border-stone-700 w-3.5 h-3.5"
                />
                <span>Stream Logs</span>
              </label>
              <label className="flex items-center gap-2 cursor-pointer text-stone-500 select-none">
                <input 
                  type="radio" 
                  name="liveLogs" 
                  value="off" 
                  checked={liveLogs === "off"}
                  onChange={(e) => setLiveLogs(e.target.value)}
                  className="accent-purple-500 bg-stone-900 border-stone-700 w-3.5 h-3.5"
                />
                <span>Mute Terminal</span>
              </label>
            </div>
          </div>

        </div>

        {/* Supplementary system control overrides panels */}
        <div className="grid grid-cols-1 sm:grid-cols-2 gap-3 pt-2">
          <div className="bg-stone-900/30 p-3 rounded-xl border border-stone-200/5 flex justify-between items-center">
            <span className="text-stone-500">System Cache File Handles</span>
            <button type="button" className="text-red-400 hover:text-red-300 font-semibold transition-colors">
              Purge Local Cache
            </button>
          </div>
          <div className="bg-stone-900/30 p-3 rounded-xl border border-stone-200/5 flex justify-between items-center">
            <span className="text-stone-500">Core Process Threads</span>
            <button type="button" className="text-amber-400 hover:text-amber-300 font-semibold transition-colors">
              Restart Agent Service
            </button>
          </div>
        </div>
      </div>
    );
  }

  /* ========================================================
     2. STANDARD STATE: Default view workspace
     ======================================================== */
  return (
    <div className="space-y-4">
      {/* 1. Top Row Dashboard Overview Grid */}
      <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
        
        {/* Agent Metadata Card */}
        <div className={`${cardMain} rounded-2xl p-3 border border-stone-200/5 relative overflow-hidden`}>
          <div className="flex justify-between items-start z-10 relative">
            <span className="text-[10px] uppercase font-mono tracking-wider bg-sky-500/10 text-sky-400 px-2 py-0.5 rounded border border-sky-500/20">
              Core Endpoint Service
            </span>
            <div className="flex items-center gap-1.5 bg-emerald-500/10 px-2.5 py-1 rounded-full border border-emerald-500/20">
              <span className="h-2 w-2 rounded-full bg-emerald-400 animate-pulse" />
              <span className="text-[10px] text-emerald-400 font-bold font-mono">ONLINE</span>
            </div>
          </div>
          
          <div className="grid grid-cols-[40%_30%_30%] items-end text-xs font-mono w-full mt-4">
            {/* First Column: 40% Width */}
            <div className="flex flex-col justify-end min-w-0 pr-4">
              <h2 className={`${spanText} text-xl font-bold truncate`}>{App.name}</h2>
              <p className={`${mutedText} text-xs mt-1 truncate`}>
                ID: {dash.agentId}
              </p>
            </div>
            
            {/* Second Column: 30% Width */}
            <div className="min-w-0">
              <p className="text-stone-500">Service Version</p>
              <p className="text-stone-200 font-medium text-sm mt-0.5 truncate">
                {currentVersion.version}-{currentVersion.tag}
              </p>
            </div>
            
            {/* Third Column: 30% Width */}
            <div className="min-w-0">
              <p className="text-stone-500">Last Handshake</p>
              <p className="text-stone-200 font-medium text-sm mt-0.5 truncate">Just Now</p>
            </div>
          </div>
        </div>

        {/* 2. Upgrade Action Engine Base */}
        <div className={`${cardMain} rounded-2xl p-3 border border-stone-200/5 flex flex-col justify-center h-full`}>
          <div className="flex items-center justify-between gap-4">
            <div>
              <h3 className={`${spanText} font-semibold text-xs`}>Agent Upgrade</h3>
              <p className="text-[11px] font-mono text-stone-500 mt-0.5">
                {isUpgrading 
                  ? "Staging binary updates..." 
                  : isLatestVersion 
                    ? "✓ System is up to date." 
                    : `⚙ Update '${upgradeAgent.version}-${upgradeAgent.tag}' ready.`}
              </p>
            </div>

            <div className="flex-shrink-0">
              {isUpgrading ? (
                <div className="flex items-center gap-2 font-mono text-[11px] text-amber-400">
                  <span className="h-2 w-2 rounded-full bg-amber-400 animate-ping" />
                  <span>Processing...</span>
                </div>
              ) : (
                <button
                  disabled={isLatestVersion}
                  onClick={triggerUpgrade}
                  className={`px-3 py-1.5 text-[11px] font-medium rounded-xl transition-all font-mono whitespace-nowrap active:scale-95 ${
                    isLatestVersion
                      ? "bg-stone-800 text-stone-600 cursor-not-allowed border border-stone-700/20"
                      : "bg-gradient-to-r from-sky-500 to-cyan-400 text-slate-950 font-bold hover:shadow-[0_0_10px_rgba(56,189,248,0.2)]"
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
      <div className={`${cardMain} rounded-2xl p-5 border border-stone-200/5 mt-4`}>
        <div className="flex items-center gap-2 mb-3">
          <span className="h-2 w-2 rounded-full bg-stone-600" />
          <h3 className={`${spanText} font-semibold text-sm`}>Administrative Terminal Gateway</h3>
        </div>

        <form onSubmit={handleAdminLogin} className="flex flex-col sm:flex-row gap-3 max-w-xl">
          <div className="relative flex-1">
            <input
              type="password"
              required
              placeholder="Provide master administration passphrase keys (type 'admin' to bypass)..."
              value={password}
              onChange={(e) => setPassword(e.target.value)}
              className="w-full bg-stone-950/40 text-stone-200 border border-stone-200/10 rounded-xl px-4 py-2.5 text-xs font-mono focus:outline-none focus:border-purple-500/40 transition-colors placeholder:text-stone-600"
            />
          </div>
          <button
            type="submit"
            className="px-5 py-2.5 bg-stone-800 text-stone-200 hover:bg-stone-700 active:scale-95 text-xs font-medium font-mono rounded-xl border border-stone-200/10 transition-all shadow-md whitespace-nowrap"
          >
            Verify Credentials
          </button>
        </form>
      </div>
    </div>
  );
}