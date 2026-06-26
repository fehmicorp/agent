import { anim_1, cardMain, spanText } from "./colour";

export function getStatTheme(value: number) {
  if (value <= 50) {
    return {
      text: "text-blue-500",
      shadow: "hover:shadow-blue-500/30",
    };
  }

  if (value <= 60) {
    return {
      text: "text-green-500",
      shadow: "hover:shadow-green-500/30",
    };
  }

  if (value <= 70) {
    return {
      text: "text-lime-500",
      shadow: "hover:shadow-lime-500/30",
    };
  }

  if (value <= 80) {
    return {
      text: "text-orange-500",
      shadow: "hover:shadow-orange-500/30",
    };
  }

  if (value <= 90) {
    return {
      text: "text-red-500",
      shadow: "hover:shadow-red-500/30",
    };
  }

  return {
    text: "text-red-900 dark:text-red-700",
    shadow: "hover:shadow-red-900/40",
  };
}


export function Card({
  title,
  value,
}: {
  title: string;
  value: number;
}) {
  const colour = getStatTheme(value);
  return (
    <div
      className={`
        ${cardMain}
        rounded-xl
        p-5
        ${anim_1}
        hover:shadow-lg
        ${colour.shadow}
      `}
    >
      <p className={spanText}>
        {title}
      </p>
      <p className={`text-xl font-semibold mt-2 ${colour.text}`}>
        {value}%
      </p>
    </div>
  );
}

export function Info({
  title,
  value,
}: {
  title: string;
  value: string;
}) {
  return (
    <div className="flex justify-between py-2 text-[12px] border-b border-stone-400/50 dark:border-stone-600/50">
      <span className="text-stone-600 dark:text-stone-400">
        {title}
      </span>
      <span className="text-stone-800 dark:text-stone-100/80">{value}</span>
    </div>
  );
}

export function Status({
  title,
  value
}: {
  title: string;
  value: boolean;
}) {
  return (
  <div className="flex justify-between py-2 text-[12px] border-b border-stone-400/50 dark:border-stone-600/50"> <span className="text-stone-600 dark:text-stone-400">{title}</span>
    <span
      className={
        value
          ? "dark:text-emerald-400 text-emerald-600"
          : "dark:text-red-400 text-red-600"
      }
    >
      {value ? "Enabled" : "Disabled"}
    </span>
  </div>
  );
}
