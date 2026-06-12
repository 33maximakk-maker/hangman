const hangmanPics = [
    `  +---+\n      |\n      |\n      |\n      |\n      |\n=========`,
    `  +---+\n  |   |\n  O   |\n      |\n      |\n      |\n=========`,
    `  +---+\n  |   |\n  O   |\n  |   |\n      |\n      |\n=========`,
    `  +---+\n  |   |\n  O   |\n /|   |\n      |\n      |\n=========`,
    `  +---+\n  |   |\n  O   |\n /|\\  |\n      |\n      |\n=========`,
    `  +---+\n  |   |\n  O   |\n /|\\  |\n /    |\n      |\n=========`,
    `  +---+\n  |   |\n  O   |\n /|\\  |\n / \\  |\n      |\n=========`
];

const words = {
    animals: ["python", "elephant", "giraffe", "tiger", "zebra", "kangaroo"],
    tech: ["computer", "keyboard", "monitor", "algorithm", "function", "variable"],
    programming: ["python", "javascript", "rust", "golang", "java", "typescript"]
};

let currentWord = "";
let guessed = new Set();
let attempts = 0;
const maxAttempts = hangmanPics.length - 1;

function getRandomWord() {
    const all = [...words.animals, ...words.tech, ...words.programming];
    return all[Math.floor(Math.random() * all.length)];
}

function updateDisplay() {
    document.getElementById("hangmanPic").innerText = hangmanPics[attempts];
    let display = "";
    for (let ch of currentWord) {
        display += (guessed.has(ch) ? ch : "_") + " ";
    }
    document.getElementById("wordDisplay").innerText = display.trim();
    document.getElementById("guessedLetters").innerText = Array.from(guessed).sort().join(", ");
    document.getElementById("attemptsLeft").innerText = maxAttempts - attempts;
}

function checkGameOver() {
    if (Array.from(currentWord).every(ch => guessed.has(ch))) {
        alert(`You win! The word was: ${currentWord}`);
        newGame();
    } else if (attempts >= maxAttempts) {
        alert(`You lost! The word was: ${currentWord}`);
        newGame();
    }
}

function guessLetter() {
    let input = document.getElementById("guessInput");
    let letter = input.value.toLowerCase();
    input.value = "";
    if (!letter || !letter.match(/[a-z]/)) {
        alert("Enter a single letter.");
        return;
    }
    if (guessed.has(letter)) {
        alert("Already guessed.");
        return;
    }
    guessed.add(letter);
    if (currentWord.includes(letter)) {
        // correct
    } else {
        attempts++;
    }
    updateDisplay();
    checkGameOver();
}

function newGame() {
    currentWord = getRandomWord();
    guessed.clear();
    attempts = 0;
    updateDisplay();
    document.getElementById("guessInput").focus();
}

document.getElementById("guessBtn").addEventListener("click", guessLetter);
document.getElementById("newGameBtn").addEventListener("click", newGame);
document.getElementById("guessInput").addEventListener("keypress", (e) => {
    if (e.key === "Enter") guessLetter();
});
newGame();
