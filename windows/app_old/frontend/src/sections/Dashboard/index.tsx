import React from "react";
import { anim_1, bgMain, cardMain, headingText, mutedText, spanText } from "../../utils/colour";
import Logo from "../../assets/logo.svg";
import App from "../data";
import { dash } from "../data/dummy";
import { Card, Info, Status } from "../../utils/func";

export default function Dashboard(): React.JSX.Element {
  return (
    <>
      <div className="grid grid-cols-4 gap-4 mb-6">
        {dash.stat.map((item, index) => (
          <Card
            key={item.title}
            title={item.title}
            value={item.value}
          />
        ))}
      </div>
      <div className="grid grid-cols-2 gap-6">
        <div className={`${cardMain} rounded-xl p-5 border dark:border-stone-500/20 border-stone-500/20`}>
          <h2 className={spanText}>
            Device Information
          </h2>
          {dash.dev.map((item) => (
            <Info
              title={item.title}
              value={item.value}
            />
          ))}
        </div>
        <div className={`${cardMain} rounded-xl p-5 border dark:border-stone-500/20 border-stone-500/20`}>
          <h2 className={spanText}>
            Security Status
          </h2>
          {dash.sec.map((item) => (
            <Status
              title={item.title}
              value={item.value}
            />
          ))}
        </div>
      </div>
    </>
  );
}
