import networkx as nx
import matplotlib.pyplot as plt
import csv

# Load the edge list from a CSV file
def load_edge_list(csv_file):
    edges = []
    with open(csv_file, 'r') as file:
        reader = csv.DictReader(file)
        for row in reader:
            source = row['Source']
            target = row['Target']
            weight = int(row['Weight'])
            edges.append((source, target, weight))
    return edges

# Build the full graph
def build_graph(edges):
    G = nx.Graph()
    for source, target, weight in edges:
        G.add_edge(source, target, weight=weight)
    return G

# Filter the graph to include only the selected node and its neighbors
def filter_graph(G, node_id):
    if node_id not in G:
        raise ValueError(f"Node '{node_id}' not found in the graph.")
    neighbors = list(G.neighbors(node_id))
    subgraph = G.subgraph([node_id] + neighbors)
    return subgraph

# Visualize the subgraph
def visualize_graph(G, node_id, output_file=None):
    # Position nodes using a spring layout
    pos = nx.spring_layout(G, seed=42)

    # Draw nodes
    nx.draw_networkx_nodes(G, pos, node_size=500, node_color="lightblue")

    # Highlight the central node
    nx.draw_networkx_nodes(G, pos, nodelist=[node_id], node_size=800, node_color="orange")

    # Draw edges
    nx.draw_networkx_edges(G, pos, width=2, alpha=0.6)

    # Draw labels
    nx.draw_networkx_labels(G, pos, font_size=10, font_color="black")

    # Add edge weight labels
    edge_labels = {(u, v): data['weight'] for u, v, data in G.edges(data=True)}
    nx.draw_networkx_edge_labels(G, pos, edge_labels=edge_labels)

    # Show or save the graph
    plt.title(f"Subgraph for Card ID: {node_id}")
    plt.axis("off")
    if output_file:
        plt.savefig(output_file, format="PNG")
        print(f"Graph saved as {output_file}")
    plt.show()

# Main function
def main(input_csv, node_id, output_file=None):
    edges = load_edge_list(input_csv)
    G = build_graph(edges)
    subgraph = filter_graph(G, node_id)
    visualize_graph(subgraph, node_id, output_file)

# Example usage
# Replace 'output.csv' with your edge list file
# Replace 'card1' with the ID you want to visualize
main('output.csv', '77152542', 'subgraph_card1.png')
