import networkx as nx
import json

def compute_relative_strength(graph, importance_metric='pagerank'):
    if importance_metric == 'degree':
        importance = dict(graph.degree(weight='weight'))
    elif importance_metric == 'weighted_degree':
        importance = {node: sum(data['weight'] for _, _, data in graph.edges(node, data=True)) for node in graph}
    elif importance_metric == 'pagerank':
        importance = nx.pagerank(graph, weight='weight')
    else:
        raise ValueError("Unsupported importance metric")
    
    # adjust weights to account for staple cards
    relative_strengths = {}
    for u, v, data in graph.edges(data=True):
        weight = data.get('weight', 1.0)
        imp_u = importance[u]
        imp_v = importance[v]
        relative_strength = weight / ((imp_u + imp_v)**1.5)  
        relative_strengths[(u, v)] = relative_strength
    
    return relative_strengths

def normalize_weights(weights):
    values = list(weights.values())
    w_min = min(values)
    w_max = max(values)
    
    if w_max == w_min:
        return {key: 0.5 for key in weights}  
    
    normalized = {key: value / w_max for key, value in weights.items()}
    return normalized

G = nx.Graph()
with open("output.csv") as network:
    for edge in network:
        i = edge.split(",")
        G.add_edge(i[0], i[1], weight=int(i[2]))

relative_strengths = compute_relative_strength(G, importance_metric='pagerank')
relative_strengths_str = {str(k): v for k, v in relative_strengths.items()}
with open("proximity.json", "w") as f:
    f.write(json.dumps(relative_strengths_str))


