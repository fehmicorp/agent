import React from "react";
import Logo from "../assets/logo.svg";
import { headingText, mutedText } from "../utils/colour";
import { dash } from "../sections/data/dummy";

interface HeaderProps {
  title: string;
  isScanning?: boolean;
}

export default function Header({ title, isScanning = false }: HeaderProps): React.JSX.Element {
  let statusStyle = dash.status
    ? "shadow-green-500/20 animate-heartbeat-green"
    : "shadow-red-500/20 animate-heartbeat-red";
  if (isScanning) {
    statusStyle = "shadow-yellow-500/20 animate-heartbeat-yellow";
  }

  return (
    <div 
      className={`sticky top-0 z-40 mb-4 p-2 border-stone-200/20 shadow-lg bg-stone-100 dark:bg-stone-950 transition-colors duration-200 ${statusStyle}`}
    >
      <div className="flex items-end">
        <img src={Logo} alt="Logo" className="h-8 w-8 mr-1"/>
        <h1 className={headingText}>
          {title}
        </h1>
      </div>
    </div>
  );
}