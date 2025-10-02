import { useState } from "react";
import { buttonStyle, buttonStyleHover } from "./Button.styles";

function Button({ children, onClick, disabled = false }) {
  const [buttonHover, setButtonHover] = useState(false);
  return (
    <button
      style={buttonHover ? buttonStyleHover : buttonStyle}
      onClick={onClick}
      disabled={disabled}
      onMouseEnter={() => setButtonHover(true)}
      onMouseLeave={() => setButtonHover(false)}
    >
      {children}
    </button>
  );
}

export default Button;
