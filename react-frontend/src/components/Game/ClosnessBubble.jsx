import { useState, useEffect } from "react";
import { bubbleStyle } from "./ClosnessBubble.styles";

function ClosnessBubble({ value }) {
  const [animatedValue, setAnimatedValue] = useState(0);

  const getBubbleColor = (val) => {
    if (val >= 0.8) return "#10b981";
    if (val >= 0.6) return "#f59e0b";
    if (val >= 0.4) return "#ef4444";
    if (val >= 0.2) return "#8b5cf6";
    return "#6b7280";
  };

  useEffect(() => {
    const normalizedValue = value > 1 ? 1 : value;
    const duration = 1000; // 1 second
    const steps = 50;
    const increment = normalizedValue / steps;
    const stepDuration = duration / steps;

    let currentStep = 0;
    const interval = setInterval(() => {
      currentStep++;
      setAnimatedValue(increment * currentStep);

      if (currentStep >= steps) {
        setAnimatedValue(normalizedValue);
        clearInterval(interval);
      }
    }, stepDuration);

    return () => clearInterval(interval);
  }, [value]);

  return (
    <div
      style={{
        ...bubbleStyle,
        background: `conic-gradient(${getBubbleColor(animatedValue)} ${
          animatedValue * 360
        }deg, #444 0deg)`,
      }}
    />
  );
}

export default ClosnessBubble;
