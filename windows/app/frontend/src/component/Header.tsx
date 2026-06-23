import React from "react";
import Logo from "../assets/logo.svg";
import { headingText, mutedText } from "../utils/colour";
import { dash } from "../sections/data/dummy";

interface HeaderProps {
  title: string;
  isScanning?: boolean;
}

export default function Header({title,isScanning = false}:HeaderProps): React.JSX.Element {
  let statusStyle = dash.status
    ? "shadow-green-500/20 animate-heartbeat-green"
    : "shadow-red-500/20 animate-heartbeat-red";
  if (isScanning) {
    statusStyle = "shadow-yellow-500/20 animate-heartbeat-yellow";
  }
  return (
    <div className={`mb-4 border-b rounded-lg p-2 border-stone-200/20 shadow-lg ${statusStyle}`}>
      <div className="flex items-end">
        <img src={Logo} alt="Logo" className="h-8 w-8 mr-1"/>
        <h1 className={headingText}>
          {title}
        </h1>
      </div>
      {/* <p className={`${mutedText} text-sm letter-spacing-2`}>
        {dash.hostname} • {dash.status} • {dash.security}
      </p> */}
    </div>
  )
}