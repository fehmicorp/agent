import React from "react";
import Logo from "../assets/logo.svg";
import { headingText } from "../utils/colour";
import { dash } from "../sections/data/dummy";
import { Bell, List } from "lucide-react";

interface HeaderProps {
  title: string;
  isScanning?: boolean;
  onOpenNotifications: () => void;
  onOpenQueue: () => void; // Added interface prop for the queue trigger
}

export default function Header({ 
  title, 
  isScanning = false, 
  onOpenNotifications,
  onOpenQueue // Destructured hook parameter
}: HeaderProps): React.JSX.Element {
  let statusStyle = dash.status
    ? "shadow-green-500/20 animate-heartbeat-green"
    : "shadow-red-500/20 animate-heartbeat-red";
  if (isScanning) {
    statusStyle = "shadow-yellow-500/20 animate-heartbeat-yellow";
  }

  return (
    <div 
      className={`sticky top-0 z-40 p-2 border-stone-200/20 flex justify-between items-end shadow-lg bg-stone-100 dark:bg-stone-950 transition-colors duration-200 ${statusStyle}`}
    >
      <div className="flex items-end">
        <img src={Logo} alt="Logo" className="h-8 w-8 mr-1"/>
        <h1 className={headingText}>
          {title}
        </h1>
      </div>
      <div className="flex items-center space-x-1 mb-0.5">
        <button 
          type="button"
          onClick={onOpenNotifications}
          aria-label="Notifications"
          className="p-1.5 rounded-lg text-stone-500 dark:text-stone-400 hover:bg-stone-200 dark:hover:bg-stone-800 hover:text-stone-900 dark:hover:text-stone-100 transition-colors duration-200"
        >
          <Bell size={16} strokeWidth={2} />
        </button>
        
        <button 
          type="button"
          onClick={onOpenQueue} // Tied queue drawer toggle handler here
          aria-label="Queue"
          className="p-1.5 rounded-lg text-stone-500 dark:text-stone-400 hover:bg-stone-200 dark:hover:bg-stone-800 hover:text-stone-900 dark:hover:text-stone-100 transition-colors duration-200"
        >
          <List size={16} strokeWidth={2} />
        </button>
      </div>
    </div>
  );
}