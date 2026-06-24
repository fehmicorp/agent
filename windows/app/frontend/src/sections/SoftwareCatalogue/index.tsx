import React, { useState, useMemo } from "react";
import { cardMain, spanText, mutedText } from "../../utils/colour";
import { software } from "../data/dummy"; // Adjust path according to your folder structure

export default function Software(): React.JSX.Element {
  const [searchQuery, setSearchQuery] = useState("");
  const [selectedCategory, setSelectedCategory] = useState<string>("All");
  const [deployingId, setDeployingId] = useState<number | null>(null);
  const [isRefreshing, setIsRefreshing] = useState(false);

  // Filter Pipeline Engine
  const filteredSoftware = useMemo(() => {
    return software.filter((item) => {
      const matchesSearch = 
        item.name.toLowerCase().includes(searchQuery.toLowerCase()) ||
        item.developer.toLowerCase().includes(searchQuery.toLowerCase()) ||
        item.description.toLowerCase().includes(searchQuery.toLowerCase());
      
      const matchesCategory = selectedCategory === "All" || item.category === selectedCategory;

      return matchesSearch && matchesCategory;
    });
  }, [searchQuery, selectedCategory]);

  const handleProvisioningRequest = (id: number) => {
    setDeployingId(id);
    setTimeout(() => {
      setDeployingId(null);
      alert(`Package configuration deployment workflow queued successfully for ID: ${id}`);
    }, 2500);
  };

  const handleRefreshRepository = () => {
    setIsRefreshing(true);
    setTimeout(() => {
      setIsRefreshing(false);
    }, 2000);
  };

  return (
    <div className="">
      {/* SECTION 1: HEADER & TELEMETRY STREAM ROW */}
      <div className={`${cardMain} rounded-2xl p-4 border border-stone-200/5 flex flex-col lg:flex-row lg:items-center justify-between gap-4`}>
        <div>
          <div className="flex items-center gap-2">
            <h1 className={`${spanText} font-bold text-base tracking-wide font-mono`}>
              Enterprise Software Catalogue
            </h1>
          </div>
          <p className={`${mutedText} text-[11px] font-mono mt-1`}>
            Centralized software repository provisioned securely by your administration console endpoint rule logic.
          </p>
        </div>
        <div className="flex flex-wrap items-center justify-end gap-2 font-mono text-[10px]">
          <span className="px-2 py-1 bg-stone-500/10 border border-stone-400/80 dark:border-stone-400/20 text-stone-700 dark:text-stone-400 rounded-lg whitespace-nowrap">
            Total Repo Packages: {software.length}
          </span>
          <span className="px-2 py-1 bg-emerald-500/15 border dark:border-emerald-500/20 border-emerald-500/50 text-emerald-800 dark:text-emerald-400 rounded-lg whitespace-nowrap">
            Available Updates: {software.filter(s => s.status === "Update Available").length}
          </span>
          
          <button
            onClick={handleRefreshRepository}
            disabled={isRefreshing}
            className={`px-3 py-1 bg-stone-800 text-white hover:bg-stone-700 active:scale-95 border border-stone-200/10 rounded-lg transition-all flex items-center gap-1.5 font-medium ${
              isRefreshing ? "dark:opacity-80 opacity-90 cursor-not-allowed" : ""
            }`}
          >
            <svg 
              className={`h-3 w-3 text-stone-400 ${isRefreshing ? "animate-spin text-blue-600 dark:text-blue-400" : ""}`} 
              fill="none" 
              viewBox="0 0 24 24" 
              stroke="currentColor" 
              strokeWidth={2.5}
            >
              <path 
                strokeLinecap="round" 
                strokeLinejoin="round" 
                d="M16.023 9.348h4.992v-.001M2.985 19.644v-4.992m0 0h4.992m-4.993 0l3.181 3.183a8.25 8.25 0 0013.803-3.7M4.031 9.865a8.25 8.25 0 0113.803-3.7l3.181 3.182m0-4.991v4.99" 
              />
            </svg>
            {isRefreshing ? "Syncing..." : "Update Repo"}
          </button>
        </div>
      </div>

      {/* SECTION 2: SEARCH FILTER CONTROLS TERMINAL (STICKY DECK) */}
      <div className="sticky top-[2.5rem] z-30 py-2.5 bg-stone-100 dark:bg-stone-950 shadow-md shadow-stone-300/40 dark:shadow-stone-950/80 grid grid-cols-6 gap-3 transition-colors duration-200">
        <div className="col-span-4 relative pt-5">
          <input
            type="text"
            placeholder="Query package index by namespace, target binary tag, or vendor description..."
            value={searchQuery}
            onChange={(e) => setSearchQuery(e.target.value)}
            className="w-full bg-stone-400/10 dark:bg-stone-950/40 text-stone-800 dark:text-stone-200 border border-stone-400/30 dark:border-stone-200/10 rounded-xl px-4 py-2.5 text-xs font-mono focus:outline-none focus:border-blue-500/50 transition-colors placeholder:text-stone-500"
          />
        </div>
        <div className="relative col-span-2 pt-5">
          <select
            value={selectedCategory}
            onChange={(e) => setSelectedCategory(e.target.value)}
            className="w-full bg-stone-100/80 dark:bg-stone-950/40 text-stone-700 dark:text-stone-300 border border-stone-400/30 dark:border-stone-200/10 rounded-md px-3 py-2.5 text-xs font-mono focus:outline-none focus:border-blue-500/50 transition-colors"
          >
            <option value="All">Select Categories</option>
            {Array.from(new Set(software.map((pkg) => pkg.category))).map((category) => (
              <option key={category} value={category}>
                {category}
              </option>
            ))}
          </select>
        </div>
      </div>

      {/* SECTION 3: CORE CATALOGUE DYNAMIC REGISTRY GRID */}
      {filteredSoftware.length === 0 ? (
        <div className={`${cardMain} rounded-2xl p-10 text-center border border-dashed dark:border-stone-800 border-stone-300`}>
          <p className="text-xs font-mono text-stone-500">
            ⚠ No software packages matched your search criteria index.
          </p>
        </div>
      ) : (
        <div className="grid grid-cols-1 xl:grid-cols-2 gap-4">
          {filteredSoftware.map((pkg) => {
            const isPackageDeploying = deployingId === pkg.id;

            return (
              <div 
                key={pkg.id} 
                className={`${cardMain} rounded-2xl p-3 border border-stone-200/5 flex flex-col justify-between hover:border-stone-400/20 dark:hover:border-stone-200/10 transition-all duration-200`}
              >
                <div>
                  <div className="flex items-start justify-between gap-2">
                    <div>
                      <h3 className="text-sm font-bold text-stone-900 dark:text-stone-100 font-mono tracking-wide">
                        {pkg.name}
                      </h3>
                      <p className="text-[11px] text-stone-500 font-mono mt-0.5">
                        {pkg.developer}
                      </p>
                    </div>

                    <span className={`text-[9px] uppercase font-mono px-2 py-0.5 rounded border font-semibold tracking-wider ${
                      pkg.status === "Installed"
                        ? "bg-emerald-500/10 text-emerald-600 dark:text-emerald-400 border-emerald-500/20"
                        : pkg.status === "Update Available"
                        ? "bg-amber-500/10 text-amber-600 dark:text-amber-400 border-amber-500/20 animate-pulse"
                        : "bg-stone-500/10 text-stone-500 border-stone-500/20"
                    }`}>
                      {pkg.status}
                    </span>
                  </div>

                  <p className={`${mutedText} text-[11px] font-sans leading-relaxed text-stone-600 dark:text-stone-400`}>
                    {pkg.description}
                  </p>
                </div>

                <div className="mt-2 pt-2 border-t border-stone-400/10 dark:border-stone-200/5 flex items-center justify-between gap-4 font-mono text-[11px]">
                  <div className="flex gap-2 text-stone-500">
                    <span className="bg-stone-500/10 dark:bg-stone-900/50 px-1.5 py-0.5 rounded border border-stone-400/40 dark:border-stone-200/5 text-[10px]">
                      Architecture: {pkg.architecture}
                    </span>
                    <span className="bg-stone-500/10 dark:bg-stone-900/50 px-1.5 py-0.5 rounded border border-stone-400/40 dark:border-stone-200/5 text-[10px]">
                      Version: v{pkg.version}
                    </span>
                    <span className="text-[10px] self-center">
                      Size: {pkg.size}
                    </span>
                  </div>

                  <button
                    disabled={pkg.status === "Installed" || isPackageDeploying}
                    onClick={() => handleProvisioningRequest(pkg.id)}
                    className={`px-3 py-1.5 rounded-xl text-[10px] font-medium transition-all active:scale-95 whitespace-nowrap ${
                      isPackageDeploying
                        ? "dark:bg-stone-800 text-amber-500 dark:border dark:border-stone-700/40 cursor-wait"
                        : pkg.status === "Installed"
                        ? "bg-stone-200/50 dark:bg-stone-900 text-stone-400 dark:text-stone-600 cursor-not-allowed border border-stone-400/10 dark:border-stone-800"
                        : pkg.status === "Update Available"
                        ? "bg-gradient-to-tr from-amber-500/20 to-orange-500/20 border border-amber-500/40 hover:border-amber-500 text-amber-700 dark:text-amber-400 hover:scale-105"
                        : "dark:bg-stone-800 bg-slate-900 text-stone-200 hover:bg-sky-950/90 dark:hover:bg-stone-800/40 hover:text-white border border-stone-200/10 hover:scale-105"
                    }`}
                  >
                    {isPackageDeploying ? (
                      <span className="flex items-center gap-1.5">
                        <span className="h-1.5 w-1.5 rounded-full bg-amber-500 animate-ping" />
                        Deploying...
                      </span>
                    ) : pkg.status === "Installed" ? (
                      "Active"
                    ) : pkg.status === "Update Available" ? (
                      "Update"
                    ) : (
                      "Provision"
                    )}
                  </button>
                </div>

              </div>
            );
          })}
        </div>
      )}
    </div>
  );
}