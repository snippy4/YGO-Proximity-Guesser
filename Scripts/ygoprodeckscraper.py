import requests
from bs4 import BeautifulSoup
import json
import time

# Define the range of IDs
start_id = 340000
end_id = 349000

# List to store deck data
decks = []

# Function to fetch and parse a single deck page
def fetch_deck(deck_id):
    url = f"https://ygoprodeck.com/deck/{deck_id}"
    try:
        response = requests.get(url, timeout=10)  # Timeout after 10 seconds
        if response.status_code == 200:
            soup = BeautifulSoup(response.text, 'html.parser')
            # Find all card elements
            cards = [card['href'].strip().split("=")[1] for card in soup.find_all(class_='ygodeckcard')]
            return {"deck_id": deck_id, "cards": cards}
        else:
            print(f"Deck {deck_id}: Page returned status code {response.status_code}")
            return None
    except requests.RequestException as e:
        print(f"Deck {deck_id}: Error fetching page - {e}")
        return None

# Loop through the range of IDs
for deck_id in range(start_id, end_id + 1):
    print(f"Processing deck {deck_id}...")
    deck_data = fetch_deck(deck_id)
    if deck_data and deck_data["cards"]:
        decks.append(deck_data)
        print(f"Deck {deck_id} saved with {len(deck_data['cards'])} cards.")
    time.sleep(0.01)  # Delay to avoid overwhelming the server

# Save the results to a JSON file
output_file = "ygopro_decks.json"
with open(output_file, 'w', encoding='utf-8') as file:
    json.dump(decks, file, ensure_ascii=False, indent=4)

print(f"Deck data saved to {output_file}. Total decks: {len(decks)}")
