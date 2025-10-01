import { buttonStyle } from "./Button.styles";

function Button({ children, onClick, disabled = false }) {
  return (
    <button style={buttonStyle} onClick={onClick} disabled={disabled}>
      {children}
    </button>
  );
}

export default Button;
