import React from "react";
import { X, Trash2 } from "lucide-react";

interface NotificationItem {
  id: number;
  text: string;
  time: string;
}

interface NotificationProps {
  isOpen: boolean;
  onClose: () => void;
  notifications: NotificationItem[];
  onClearAll: () => void;
}

export default function Notification({
  isOpen,
  onClose,
  notifications,
  onClearAll,
}: NotificationProps): React.JSX.Element {
  return (
    <>
      {/* Semi-translucent Overlay backdrop blur */}
      <div
        onClick={onClose}
        className={`fixed inset-0 bg-stone-950/40 backdrop-blur-sm z-50 transition-opacity duration-300 ${
          isOpen ? "opacity-100 pointer-events-auto" : "opacity-0 pointer-events-none"
        }`}
      />

      {/* Sidebar Container Shell Panel matching 75% display parameters */}
      <div
        className={`fixed top-0 right-0 h-full w-3/4 max-w-lg bg-stone-50 dark:bg-stone-900 border-l border-stone-200 dark:border-stone-800 shadow-2xl z-50 flex flex-col justify-between transform transition-transform duration-300 ease-out ${
          isOpen ? "translate-x-0" : "translate-x-full"
        }`}
      >
        {/* Header Panel */}
        <div className="p-4 border-b border-stone-200 dark:border-stone-800 flex items-center justify-between">
          <div>
            <h2 className="text-xs font-bold uppercase tracking-wider text-stone-900 dark:text-stone-100">
              System Alerts
            </h2>
            <p className="text-[10px] text-stone-500">Live deployment log queues</p>
          </div>
          <button
            type="button"
            onClick={onClose}
            className="p-1 rounded-md text-stone-500 hover:bg-stone-200 dark:hover:bg-stone-800 transition-colors"
          >
            <X size={16} />
          </button>
        </div>

        {/* Core Content Body Scrolling Container Layout */}
        <div className="flex-1 p-4 overflow-y-auto space-y-2.5">
          {notifications.length === 0 ? (
            <div className="h-full flex flex-col items-center justify-center text-center opacity-60 py-12">
              <span className="text-lg">✨</span>
              <p className="text-[11px] text-stone-500 mt-1">Workspace completely clear.</p>
            </div>
          ) : (
            notifications.map((notif) => (
              <div
                key={notif.id}
                className="p-3 rounded-xl bg-white dark:bg-stone-950 border border-stone-200/50 dark:border-stone-800/60 shadow-sm"
              >
                <p className="text-[11px] text-stone-800 dark:text-stone-200 leading-normal">
                  {notif.text}
                </p>
                <span className="text-[9px] block mt-1.5 text-stone-400 font-medium">
                  {notif.time}
                </span>
              </div>
            ))
          )}
        </div>

        {/* Sticky Bottom Actions Button Bar Section Control Deck */}
        <div className="p-4 bg-stone-100/80 dark:bg-stone-900/80 backdrop-blur border-t border-stone-200 dark:border-stone-800">
          <button
            type="button"
            disabled={notifications.length === 0}
            onClick={onClearAll}
            className="w-full flex items-center justify-center gap-1.5 py-2 bg-stone-900 dark:bg-stone-100 text-stone-100 dark:text-stone-950 font-bold rounded-xl text-[10px] uppercase tracking-wider transition-all disabled:opacity-40 disabled:cursor-not-allowed hover:opacity-90"
          >
            <Trash2 size={12} />
            Clear All Notifications
          </button>
        </div>
      </div>
    </>
  );
}