import { useState } from "react";
import {
  searchBarStyle,
  suggestionsContainerStyle,
  suggestionItemStyle,
  suggestionImageStyle,
} from "./SearchBar.styles";

function SearchBar({ onSearch, onSelect }) {
  const [input, setInput] = useState("");
  const [suggestions, setSuggestions] = useState([]);

  const handleInputChange = async (e) => {
    const value = e.target.value;
    setInput(value);

    if (value.length > 0) {
      try {
        const response = await fetch(
          `http://localhost:8080/search?q=${encodeURIComponent(value)}`
        );
        let data = await response.json();

        // Handle if response is wrapped in array
        if (Array.isArray(data) && typeof data[0] === "string") {
          data = JSON.parse(data[0]);
        }

        setSuggestions(data.slice(0, 5));
      } catch (error) {
        console.error("Search error:", error);
        setSuggestions([]);
      }
    } else {
      setSuggestions([]);
    }
  };

  const handleSelect = (suggestion) => {
    setInput(suggestion.name);
    setSuggestions([]);
    onSelect(suggestion);
  };

  return (
    <div>
      <input
        type="text"
        value={input}
        onChange={handleInputChange}
        placeholder="Search for a Yu-Gi-Oh card..."
        style={searchBarStyle}
      />
      {suggestions.length > 0 && (
        <div style={suggestionsContainerStyle}>
          {suggestions.map((suggestion, index) => (
            <div
              key={index}
              style={suggestionItemStyle}
              onClick={() => handleSelect(suggestion)}
            >
              <img
                src={`https://card.yugioh-api.com/${suggestion.id}/image`}
                alt={suggestion.name}
                style={suggestionImageStyle}
              />

              <span>{suggestion.name}</span>
            </div>
          ))}
        </div>
      )}
    </div>
  );
}

export default SearchBar;
