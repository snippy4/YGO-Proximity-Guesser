import { useState } from "react";
import SearchBar from "../Game/SearchBar";
import Guess from "../Game/Guess";
import {
  containerStyle,
  titleStyle,
  guessCounterStyle,
  guessListStyle,
} from "./HomeBody.styles";
import Button from "./Button";
import WinPopup from "./WinPopup";

function HomeBody() {
  const [guesses, setGuesses] = useState([]);
  const [guessCount, setGuessCount] = useState(0);
  const [showWinScreen, setShowWinScreen] = useState(false);

  const handleCardSelect = async (card) => {
    try {
      const response = await fetch(
        `http://localhost:8080/select?q=${encodeURIComponent(card.name)}`
      );
      const result = await response.json();

      const newGuess = {
        cardId: card.id,
        cardName: card.name,
        closenessValue:
          result.value === "Correct" ? 21 : parseFloat(result.value) || 0,
      };

      setGuesses((prev) =>
        [newGuess, ...prev].sort((a, b) => b.closenessValue - a.closenessValue)
      );
      setGuessCount((prev) => prev + 1);
      if (result.value === "Correct") {
        setShowWinScreen(true);
      }
    } catch (error) {
      console.error("Guess error:", error);
    }
  };

  const getClosenessText = (value) => {
    const parsedValue = parseFloat(value);

    if (value === "Correct") return "Correct!";
    if (isNaN(parsedValue)) return "No connection :(";
    if (parsedValue === 21) return "Correct!";
    if (parsedValue === 20) return "The Closest Card!";
    if (parsedValue >= 1) return `${20 - parsedValue} card(s) away!`;
    if (parsedValue >= 0.8) return "Flaming hot";
    if (parsedValue >= 0.6) return "Warm";
    if (parsedValue >= 0.4) return "Tepid";
    if (parsedValue >= 0.2) return "Cold";
    return "At least someone played these cards together?";
  };

  const handleGetHint = async () => {
    try {
      const hintId = guesses.length > 0 ? guesses[0].cardId : "23434538";
      const response = await fetch(`http://localhost:8080/hint?q=${hintId}`);
      const hint = await response.json();

      if (hint) {
        // Find the card data for the hint
        const searchResponse = await fetch(
          `http://localhost:8080/search?q=${encodeURIComponent(hint)}`
        );
        let searchData = await searchResponse.json();

        if (Array.isArray(searchData) && typeof searchData[0] === "string") {
          searchData = JSON.parse(searchData[0]);
        }

        const hintCard = searchData.find((card) => card.name === hint);
        if (hintCard) {
          handleCardSelect(hintCard);
        }
      }
    } catch (error) {
      console.error("Hint error:", error);
    }
  };

  const handleGiveUp = async () => {
    try {
      const response = await fetch(`http://localhost:8080/answer`);
      const answer = await response.json();
      console.log(answer);
      handleCardSelect(JSON.parse(answer)[0]);
    } catch (error) {
      console.error("Give up error:", error);
    }
  };

  return (
    <>
      <div style={containerStyle}>
        <h1 style={titleStyle}>Yu-Gi-Oh Proximity Game</h1>
        <SearchBar onSelect={handleCardSelect} />
        <div style={{ textAlign: "center", margin: "20px 0" }}>
          <Button onClick={handleGetHint}>Get Hint</Button>
          <Button onClick={handleGiveUp}>Give Up</Button>
        </div>

        <div style={guessCounterStyle}>
          You have guessed {guessCount} times.
        </div>
        <div style={guessListStyle}>
          {guesses.map((guess) => (
            <Guess
              key={guess.cardId}
              cardId={guess.cardId}
              cardName={guess.cardName}
              closenessText={getClosenessText(guess.closenessValue)}
              closenessValue={guess.closenessValue}
            />
          ))}
        </div>
      </div>
      {showWinScreen && (
        <WinPopup
          guessCount={guessCount}
          onClose={() => setShowWinScreen(false)}
          cardId={guesses[0].cardId}
          cardName={guesses[0].cardName}
        />
      )}
    </>
  );
}

export default HomeBody;
