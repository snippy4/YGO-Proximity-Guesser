import { bubbleStyle } from "./ClosnessBubble.styles";

function ClosnessBubble({ value }) {
  const getBubbleColor = (val) => {
    if (val >= 0.8) return "#00ff00";
    if (val >= 0.6) return "#51ad00";
    if (val >= 0.4) return "#7c8200";
    if (val >= 0.2) return "#a35b00";
    return "#ff0000";
  };

  // Handle whole numbers as 100%
  const normalizedValue = value > 1 ? 1 : value;

  return (
    <div
      style={{
        ...bubbleStyle,
        background: `conic-gradient(${getBubbleColor(normalizedValue)} ${
          normalizedValue * 360
        }deg, #444 0deg)`,
      }}
    />
  );
}

export default ClosnessBubble;
