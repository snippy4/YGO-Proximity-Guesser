import {
  modalStyle,
  modalContentStyle,
  titleStyle,
  buttonStyle,
} from "./WinPopup.styles";
import Card from "../Game/Card";

function WinPopup({ guessCount, onClose, cardId, cardName }) {
  return (
    <div style={modalStyle}>
      <div style={modalContentStyle}>
        <h2 style={titleStyle}>Congratulations!</h2>
        <Card cardId={cardId} cardName={cardName} />
        <p>You guessed the card in {guessCount} tries!</p>
        <button style={buttonStyle} onClick={onClose}>
          Play Again
        </button>
      </div>
    </div>
  );
}

export default WinPopup;
