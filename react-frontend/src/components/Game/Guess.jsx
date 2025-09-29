import Card from "./Card";
import ClosnessBubble from "./ClosnessBubble";
import {
  guessContainerStyle,
  cardInfoStyle,
  cardNameStyle,
} from "./Guess.styles";

function Guess({ cardId, cardName, closenessText, closenessValue }) {
  return (
    <div style={guessContainerStyle}>
      <Card cardId={cardId} cardName={cardName} />
      <div style={cardInfoStyle}>
        <div style={cardNameStyle}>
          {cardName}: {closenessText}
        </div>
        <ClosnessBubble value={closenessValue} />
      </div>
    </div>
  );
}

export default Guess;
