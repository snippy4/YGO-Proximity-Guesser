import json
from collections import defaultdict
from itertools import combinations
import csv

# Load JSON data
def load_json(file_path):
    with open(file_path, 'r') as f:
        return json.load(f)

# Build the co-occurrence matrix
def build_co_occurrence_matrix(decks):
    co_occurrence = defaultdict(lambda: defaultdict(int))
    blacklist = ["10000000", "10000010", "10000020", "14315573", "55144522", "4206964", "69890967", "32491822", "6007213",
                 "71625222", "70781052", "62279055", "340521"]
    filtered = 0
    for deck in decks:
        skip = False
        cards = deck['cards']
        for card in blacklist:
            if card in cards:
                skip = True
        # Generate all pair combinations of cards in the deck
        if skip:
            filtered += 1
            continue
        for card1, card2 in combinations(list(set(cards)), 2):
            co_occurrence[card1][card2] += 1
            co_occurrence[card2][card1] += 1  # Since it's undirected
    print(f'filtered {filtered} decks')
    return co_occurrence

# Save the network as an edge list
def save_edge_list(co_occurrence, output_file):
    with open(output_file, 'w', newline='') as csvfile:
        writer = csv.writer(csvfile)
        writer.writerow(['Source', 'Target', 'Weight'])
        for card1, neighbors in co_occurrence.items():
            for card2, weight in neighbors.items():
                if card1 < card2:  # Avoid duplicates in undirected graph
                    writer.writerow([card1, card2, weight])

def main(json_file, output_file):
    data = load_json(json_file)
    
    # Check if data is a list or contains the "decks" key
    if isinstance(data, list):
        decks = data
    elif "decks" in data:
        decks = data["decks"]
    else:
        raise ValueError("Invalid JSON structure: expected a list or a dictionary with a 'decks' key")
    
    co_occurrence = build_co_occurrence_matrix(decks)
    save_edge_list(co_occurrence, output_file)
    print(f"Network saved to {output_file}")


# Example usage
# Replace 'input.json' and 'output.csv' with your actual file paths
main('ygopro_decks.json', 'output.csv')
