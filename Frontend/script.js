// Generate a random number between 1 and 100
const randomNumber = Math.floor(Math.random() * 100) + 1;

// Track previous guesses
let previousGuesses = [];

function checkGuess() {
    // Get the user's guess
    const userGuess = parseInt(document.getElementById('guessInput').value);
    const resultMessage = document.getElementById('resultMessage');
    const previousGuessesDisplay = document.getElementById('previousGuesses');

    // Validate input
    if (isNaN(userGuess) || userGuess < 1 || userGuess > 100) {
        resultMessage.textContent = "Please enter a valid number between 1 and 100.";
        resultMessage.style.color = "red";
        return;
    }

    // Add the guess to the list of previous guesses
    previousGuesses.push(userGuess);

    // Build the feedback for previous guesses
    previousGuessesDisplay.textContent = `Previous guesses: ${previousGuesses.join(', ')}`;

    // Check the guess
    if (userGuess === randomNumber) {
        resultMessage.textContent = `🎉 Congratulations! You guessed it! The number was ${randomNumber}.`;
        resultMessage.style.color = "green";
        disableGame();
    } else if (userGuess < randomNumber) {
        resultMessage.textContent = "Too low! Try again.";
        resultMessage.style.color = "orange";
    } else {
        resultMessage.textContent = "Too high! Try again.";
        resultMessage.style.color = "orange";
    }
}

function disableGame() {
    document.getElementById('guessInput').disabled = true;
    document.querySelector('button').disabled = true;
    const resetButton = document.createElement('button');
    resetButton.textContent = 'Play Again';
    resetButton.style.marginTop = '20px';
    resetButton.onclick = () => window.location.reload();
    document.querySelector('.container').appendChild(resetButton);
}
